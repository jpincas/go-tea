package gotea

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/olahol/melody"
)

// BaseModel is embedded in the application model struct to provide client-side routing and more
type BaseModel struct {
	Router
	OriginalRequest *http.Request
}

// State is attached to each 'session' and is what is rendered by the Gotea runtime on each update.
// It can essentially anything  - you just define it as a struct in your application code.
// By Elm convention it would be called 'Model', but that's up to you.
type State interface {
	// Init must be defined by the user and describes the 'blank' state from which a session starts.
	// It gets passed a pointer to the originating http request which might help to set some starting parameters, but can most often be ignored.
	Init(*http.Request) State

	// Update is defined by the user and returns the list of messages that is used to modify state
	// on each loop of the runtime.
	Update() MessageMap

	// Render is the view function provided by the user to render out the state
	Render() []byte

	// Routable provides all the routing functions
	Routable
}

// render uses the app-provided render function to send HTML through the socket -
// this is what the client sees in their browser.
// The JS part of gotea takes this HTML and patches it efficiently onto
// the existing DOM, so the browser only updates what has actually changed
func renderState(s *melody.Session) {
	st, _ := s.Get("state")
	state := st.(State)
	s.Write(state.Render())
}

// Message is a data structure that is triggered in JS in the browser,
// and sent through the open websocket connection to the Gotea runtime.
// It's quite simple and consists of just two pieces of information:
// 1 - the name of the message (a string)
// 2 - some optional accompanying data (JSON) (can be nil)
type Message struct {
	Message   string          `json:"message"`
	Arguments json.RawMessage `json:"args"`
}

// Response is returned by MessageHandler functions.  The most important part of the
// response is the new state, but they can optionally return another message to be
// processed after an optional delay.
type Response struct {
	NextMsg *Message
	Delay   time.Duration
	Error   error
}

// MessageHandler functions can do absolutely anything as long as they return a new state.
// Typically they would be used to make some sort of mutation to the state.
type MessageHandler func(json.RawMessage, State) Response

// MessageMap holds a record of MessageHandler functions keyed against message.
// This enables the runtime to look up the correct function to execute for each message received.
// The client application must provide this when bootstrapping the app.
type MessageMap map[string]MessageHandler

// process does the actual work of dealing with an incoming message.
// It checks to make sure a message handling function is assigned to that message, raising an error if not.
// Assuming a message handling function is found, it is executed,
// resulting in a new state.  This new state is set on the session,
// and any further messages are sent for processing in the same way (recursively).
// Finally, a render of the new state takes place, sending new HTML down the websocket to the client,
// and starting the cycle again.
func (message Message) process(s *melody.Session) error {
	// Since messages can trigger themselves, they can potentially set off an infinite loop,
	// which would not be interrupted by the connection closing.
	// So here we check that the connection is open before processing the message.
	if s.IsClosed() {
		return fmt.Errorf("Could not process message %s: connection has been closed", message.Message)
	}

	log.Println("Processing message: ", message.Message)

	st, _ := s.Get("state")
	state := st.(State)

	// Try system messages first.
	// At the moment, just the router, but could expand
	systemMessages := routingMessages
	funcToExecute, found := systemMessages[message.Message]

	// TODO: We might want to check both maps here and raise
	// some sort of log message if there is a clash of names
	if !found {
		// Care to overwrite the funcToExecute variable above
		funcToExecute, found = state.Update()[message.Message]
		if !found {
			return fmt.Errorf("Could not process message %s: message does not exist", message.Message)
		}
	}

	response := funcToExecute(message.Arguments, state)
	if response.Error != nil {
		return response.Error
	}

	renderState(s)

	if response.NextMsg != nil {
		go func() {
			if response.Delay > 0 {
				time.Sleep(response.Delay * time.Millisecond)
			}

			response.NextMsg.process(s)
		}()
	}

	return nil
}

// websocketHandler is the handler function called when a client connects.
// It is basically the core of the runtime.  Here's what it does:
// - upgrades the connection to a websocket
// - creates a new session and adds it to the list of active sessions
// - waits for a message from the client
// - sends the messages for processing
// - waits again, etc etc
func (app Application) onConnect(s *melody.Session) {
	// The session needs to know which route it is starting from,
	// else the first template render will fail.
	// We can't just use the path from the URL, since the websocket
	// connection is always through /server.
	// Therefore, the JS adds a ?whence=route parameter to /server
	// when making the connection, so we get the starting route from there
	s.Request.ParseForm()
	startingRoute := s.Request.URL.Query().Get("whence")
	state := app.newState(s.Request)

	// run any app-specific logic to when the session connects
	// Im not sure exactly at which point to fire this yet.
	// state.OnConnect(s.)

	// we need to run the client supplied route change logic before the first render
	/// but since it might be fail, we wrap it in error checking so that it can't crash
	// the runtime
	changeRoute(state, startingRoute)

	// Originally, I didn't actually need to render the session now. Why?
	// Because the intial HTML render (before the websocket was established) was enough
	// to perfectly render the initial state.  Rendering it again now would be redundant.
	// Therefore we wait for some interaction before doing our first render through the websocket.
	// However, I added the OnConnect() method (see above) which can essentailly do anything once
	// the client connects, so there state might therefore change, so we probably DO need this render.
	// session.render(nil)

	s.Set("state", state)
}

func handleMessage(s *melody.Session, msg []byte) {
	var message Message
	if err := json.Unmarshal(msg, &message); err != nil {
		log.Printf("Error unmarshalling JSON message: %v", err)
		return
	}

	message.process(s)
}

// AppConfig specifies the configuration for the Gotea app
type AppConfig struct {
	// Port is the port to start the app on when calling Start
	Port int

	// StaticDirectory (optional) will set up a convenient static file server on the specified directory
	// at /static.  If left blank, because you have a more complex setup with your own static file server,
	// then no static file server will be initiated.
	StaticDirectory string
}

// DefaultAppConfig provides a set of sane defaults for a Gotea app
var DefaultAppConfig = AppConfig{
	Port:            8080,
	StaticDirectory: "static",
}

// Application is the holder for all the bits and pieces go-tea needs
type Application struct {
	*melody.Melody

	Config AppConfig

	// The client provided state model
	Model State
}

// NewApp is used by the calling application to set up a new gotea app
func NewApp(config AppConfig, model State) *Application {
	app := Application{
		Melody: melody.New(),
		Config: config,
		Model:  model,
	}

	app.Melody.HandleConnect(app.onConnect)
	app.Melody.HandleMessage(handleMessage)

	return &app
}

// newState bootstraps a new state model according to the init() provided by the calling app.
// We also perform a 'route set' which will run any route dependent logic
// as well as set the starting template.
func (app Application) newState(r *http.Request) State {
	state := app.Model.Init(r)
	return state
}

// Start creates the router, and serves it!
func (app *Application) Start() {
	http.HandleFunc("/server", func(w http.ResponseWriter, r *http.Request) {
		app.Melody.HandleRequest(w, r)
	})

	if app.Config.StaticDirectory != "" {
		fs := http.FileServer(http.Dir(app.Config.StaticDirectory))
		http.Handle("/static/", http.StripPrefix("/static/", fs))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		state := app.newState(r)
		changeRoute(state, r.URL.Path)
		w.Write(state.Render())
	})

	log.Printf("Starting application server on %v\n", app.Config.Port)
	http.ListenAndServe(fmt.Sprintf(":%v", app.Config.Port), nil)
}
