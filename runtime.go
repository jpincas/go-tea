package gotea

import (
	"compress/flate"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/gorilla/websocket"
)

// BaseModel is embedded in the application model struct to provide client-side routing and more
type BaseModel struct {
	Router
	OriginalRequest *http.Request
}

// RenderError is a default error renderer that can be overwritten by the app
func (b BaseModel) RenderError(w io.Writer, err error) {
	msg := fmt.Sprintf("Sorry! An error has ocurred.  Please try again or contact the site admin for help.\n\nError details: %s", err)
	w.Write([]byte(msg))
}

func (b BaseModel) OnConnect(_ SessionID) {}

// State is attached to each 'session' and is what is rendered by the Gotea runtime on each update.
// It can essentially anything  - you just define it as a struct in your application code.
// By Elm convention it would be called 'Model', but that's up to you.
type State interface {
	// Init must be defined by the user and describes the 'blank' state from which a session starts.
	// It gets passed a pointer to the originating http request which might help to set some starting parameters, but can most often be ignored.
	Init(*http.Request) State

	OnConnect(SessionID)

	// Update is defined by the user and returns the list of messages that is used to modify state
	// on each loop of the runtime.
	Update() MessageMap

	// Render is the view function provided by the user to render out the state
	Render(io.Writer) error

	// RenderError is an optional error renderer than can be provided by the app
	RenderError(io.Writer, error)

	// Routable provides all the routing functions
	Routable
}

// sessionStore tracks a list of active sessions.  When a session is opened,
// it gets added to the sessionStore. When a session is closed, it gets removed.
type sessionStore map[SessionID]*session

type SessionID = uuid.UUID

// session is a holder for a websocket connection (or more accurately, a pointer to one),
// and a lump of state (see above).  When a client connects, a new session is opened with a pointer
// to the websocker connection and the initial state (more on that later)
type session struct {
	id      uuid.UUID
	session bool
	conn    *websocket.Conn
	state   State
}

// changeRoute is a wrapper around the route change hook which is unsafe code provided by the calling app.
// In case it crashes, it won't bring down the Gotea runtime
func (session *session) changeRoute(path string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Could not run initial route change hook")
		}
	}()

	changeRoute(session.state, path)
}

// add adds the session to the list of active sessions
func (session *session) add(app *Application) {
	app.Sessions[session.id] = session
}

// close does the work of closing a session:
// actually closing the connection, and removing it from the list.
func (session *session) close(app *Application) {
	log.Printf("Closing session %s", session.id)
	session.session = true
	if err := session.conn.Close(); err != nil {
		log.Printf("Error closing WebSocket connection: %v", err)
	}

	// remove from session store
	delete(app.Sessions, session.id)
}

// render uses the app-provided render function to send HTML through the socket -
// this is what the client sees in their browser.
// The JS part of gotea takes this HTML and patches it efficiently onto
// the existing DOM, so the browser only updates what has actually changed
func (session *session) render(errorToRender error) {
	// There is no point trying to render a seesion if the message is CLOSE
	// because logically it will fail
	// TODO: better close checking required here
	if websocket.IsCloseError(errorToRender, websocket.CloseGoingAway) {
		return
	}

	if session.conn == nil {
		log.Println("Could not render, no socket to render to")
		return
	}

	w, err := session.conn.NextWriter(1)
	if err != nil {
		log.Printf("Error opening websocket connection writer: %v\n", err)
		return
	}

	// Now we actually start the business of trying to render to the socket.
	if errorToRender != nil {
		session.state.RenderError(w, errorToRender)
		return
	}

	if renderAttemptErr := session.state.Render(w); renderAttemptErr != nil {
		session.state.RenderError(w, renderAttemptErr)
	}

	if err := w.Close(); err != nil {
		log.Printf("Error closing websocket connection writer: %v\n", err)
		return
	}
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
func (message Message) process(session *session, app *Application) error {
	// Since messages can trigger themselves, they can potentially set off an infinite loop,
	// which would not be interrupted by the connection closing.
	// So here we check that the connection is open before processing the message.
	if session.session {
		return fmt.Errorf("Could not process message %s: connection has been closed", message.Message)
	}

	// Try system messages first.
	// At the moment, just the router, but could expand
	systemMessages := routingMessages
	funcToExecute, found := systemMessages[message.Message]

	// TODO: We might want to check both maps here and raise
	// some sort of log message if there is a clash of names
	if !found {
		// Care to overwrite the funcToExecute variable above
		funcToExecute, found = session.state.Update()[message.Message]
		if !found {
			return fmt.Errorf("Could not process message %s: message does not exist", message.Message)
		}
	}

	response := funcToExecute(message.Arguments, session.state)
	if response.Error != nil {
		return response.Error
	}

	session.render(nil)

	if response.NextMsg != nil {
		if response.Delay > 0 {
			time.Sleep(response.Delay * time.Millisecond)
		}

		response.NextMsg.process(session, app)
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
func (app *Application) websocketHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to upgrade connection to WebSocket")

	// upgrade the connection without permessage-deflate compression
	upgrader := websocket.Upgrader{
		EnableCompression: true,
		CheckOrigin:       app.Config.CheckOrigin,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		BaseModel{}.RenderError(w, errors.New("error upgrading connection"))
		return
	}

	conn.EnableWriteCompression(true)
	conn.SetCompressionLevel(flate.BestSpeed)

	log.Println("WebSocket connection upgraded successfully")

	// The session needs to know which route it is starting from,
	// else the first template render will fail.
	// We can't just use the path from the URL, since the websocket
	// connection is always through /server.
	// Therefore, the JS adds a ?whence=route parameter to /server
	// when making the connection, so we get the starting route from there
	r.ParseForm()
	startingRoute := r.URL.Query().Get("whence")

	// Session ID
	id := uuid.NewV4()

	// create a new session based on the current connection
	// this will use the state seeder to create a default state
	session := &session{
		id:    id,
		conn:  conn,
		state: app.newState(r),
	}

	// run any app-specific logic to when the session connects
	// Im not sure exactly at which point to fire this yet.
	session.state.OnConnect(session.id)

	// we need to run the client supplied route change logic before the first render
	/// but since it might be fail, we wrap it in error checking so that it can't crash
	// the runtime
	session.changeRoute(startingRoute)
	session.add(app)

	// defer closing the connection and removing it from the session
	defer session.close(app)

	// Originally, I didn't actually need to render the session now. Why?
	// Because the intial HTML render (before the websocket was established) was enough
	// to perfectly render the initial state.  Rendering it again now would be redundant.
	// Therefore we wait for some interaction before doing our first render through the websocket.
	// However, I added the OnConnect() method (see above) which can essentailly do anything once
	// the client connects, so there state might therefore change, so we probably DO need this render.
	session.render(nil)

	// main runtime loop
	for {
		// read the incoming message
		var message Message
		if err := conn.ReadJSON(&message); err != nil {
			log.Printf("Error reading JSON message: %v", err)
			session.render(err)
			break
		}

		// process inside a go routine so as to not block the runtime loop
		go func() {
			// Client handlers could potentially cause all manner of havoc,
			// with the potential to shutdown the Gotea runtime.
			// Before processing any messages, which are essentially arbitrary client code,
			// we'll defer a recoverer which will render a nice error and save the runtime.
			defer func() {
				if r := recover(); r != nil {
					session.render(fmt.Errorf("I caught a panic while processing the '%s' message: %s", message.Message, r))
				}
			}()

			if err := message.process(session, app); err != nil {
				session.render(err)
			}

		}()
	}
}

// AppConfig specifies the configuration for the Gotea app
type AppConfig struct {
	// Port is the port to start the app on when calling Start
	Port int

	// TemplatesDirectory specifies where to find the Jet templates
	// TemplatesDirectory string

	// StaticDirectory (optional) will set up a convenient static file server on the specified directory
	// at /static.  If left blank, because you have a more complex setup with your own static file server,
	// then no static file server will be initiated.
	StaticDirectory string

	// CheckOrigin is the CORS checking function for websocket upgrade
	CheckOrigin func(r *http.Request) bool
}

// DefaultAppConfig provides a set of sane defaults for a Gotea app
var DefaultAppConfig = AppConfig{
	Port:            8080,
	StaticDirectory: "static",
	CheckOrigin:     func(_ *http.Request) bool { return true },
}

// Application is the holder for all the bits and pieces go-tea needs
type Application struct {
	Config AppConfig

	// Sessions is a list of the current active sessions
	Sessions sessionStore

	// The client provided state model
	Model State
}

// NewApp is used by the calling application to set up a new gotea app
func NewApp(config AppConfig, model State) *Application {
	app := Application{
		Config:   config,
		Model:    model,
		Sessions: sessionStore{},
	}

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
	http.HandleFunc("/server", app.websocketHandler)

	if app.Config.StaticDirectory != "" {
		fs := http.FileServer(http.Dir(app.Config.StaticDirectory))
		http.Handle("/static/", http.StripPrefix("/static/", fs))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		state := app.newState(r)

		defer func() {
			if rec := recover(); rec != nil {
				msg := fmt.Errorf("Gotea crashed while initialising state: %s", rec)
				state.RenderError(w, msg)
			}
		}()

		changeRoute(state, r.URL.Path)
		state.Render(w)
	})

	log.Printf("Starting application server on %v\n", app.Config.Port)
	http.ListenAndServe(fmt.Sprintf(":%v", app.Config.Port), nil)
}
