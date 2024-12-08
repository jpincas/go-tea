package gotea

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/olahol/melody"
)

const (
	melodyStateKey = "state"
)

// ROUTING

// Router is embedded by the application model to provide routing functionality
type Router struct {
	Route string
}

func (r *Router) SetNewRoute(route string) {
	r.Route = route
}

func (r Router) RouteParam(param string) string {
	rel, err := url.Parse(r.Route)
	if err != nil {
		return ""
	}

	return rel.Query().Get(param)
}

func (r Router) GetRoute() string {
	return r.Route
}

// Routable will be fulfilled by the applicaiton model by embedding the Router
// and defining the OnRouteChange function
type Routable interface {
	// These methods are fulfilled just by embedding the Router struct
	SetNewRoute(string)
	GetRoute() string
	RouteParam(string) string

	// OnRouteChange must be defined by the user.  It is a routing function that determines the template to use as well as any logic to perform based on the route.
	OnRouteChange(string)
}

// Messages relating to routing that will be merged into the main message map
var routingMessages = MessageMap{
	"CHANGE_ROUTE": changeRouteMsgHandler,
}

// changeRouteMsgHandler is the built in message handler which is fired when a
// navigation event is detected
func changeRouteMsgHandler(message Message, state State) Response {
	changeRoute(state, message.Arguments.(string))
	return Respond()
}

// changeRoute is fired both by the route change message handler and on establishment
// of a new state blob.  It fires the app-provided routing logic and sets the new route /// on the model.
func changeRoute(state State, newRoute string) {
	state.OnRouteChange(newRoute)
	state.SetNewRoute(newRoute)
}

// STATE

// State is attached to each session and is what is rendered by the Gotea runtime on each update.
// It can essentially anything  - you just define it as a struct in your application code.
// By Elm convention it would be called 'Model', but that's up to you.
type State interface {
	Routable

	// Init must be defined by the user and describes the 'blank' state from which a session starts.
	Init(uuid.UUID) State

	// Update is defined by the user and returns the list of messages that is used to modify state
	Update() MessageMap

	// Render and RenderError are the view functions provided by the user to render out the state
	Render() []byte
	RenderError(error) []byte
}

// onConnect is the Melody handler that is called when a new session is established
// It is responsible for setting up the initial state of the session, including routing
func onConnect(model State) func(s *melody.Session) {
	return func(s *melody.Session) {
		// We need to get the session id from the cookie
		cookie, err := s.Request.Cookie("session_id")
		if err != nil {
			log.Printf("Error getting session ID from cookie: %v", err)
			return
		}

		// Set the session ID on the state
		state := model.Init(uuid.MustParse(cookie.Value))

		// We can't just use the path from the URL, since the websocket
		// connection is always through /server.
		// Therefore, the JS adds a ?whence=route parameter to /server
		// when making the connection, so we get the starting route from there
		s.Request.ParseForm()
		startingRoute := s.Request.URL.Query().Get("whence")
		changeRoute(state, startingRoute)

		s.Set(melodyStateKey, state)
	}
}

// MESSAGE HANDLING

// Message is a data structure that is triggered in JS in the browser,
// and sent through the open websocket connection to the Gotea runtime.
// It's quite simple and consists of just two pieces of information:
// 1 - the name of the message (a string)
// 2 - some optional accompanying data (JSON) (can be nil)
// 3 - an optional identifier (a string) - useful for reusing the same message handler for multiple messages
// 4 - an optional blockRerender flag (bool) - useful for messages that don't require a rerender
type Message struct {
	Message string `json:"message"`

	Arguments     any    `json:"args"`
	Identifier    string `json:"identifier"`
	BlockRerender bool   `json:"blockRerender"`
}

// Some helpers for decoding messages

func (m Message) toJson() string {
	jsonData, _ := json.Marshal(m)
	return string(jsonData)
}

func (m Message) MustDecodeArgs(target any) {
	// First marshall the arguments to JSON
	jsonData, _ := json.Marshal(m.Arguments)
	// Then unmarshall to the target
	json.Unmarshal(jsonData, target)
}

func (m Message) ArgsToString() string {
	s, _ := m.Arguments.(string)
	return s
}

func (m Message) ArgsToFloat() float64 {
	f, _ := m.Arguments.(float64)
	return f
}

func (m Message) ArgsToInt() int {
	// JSON alwayss marshalls numbers to float64
	f, _ := m.Arguments.(float64)
	return int(f)
}

// Response is returned by MessageHandler functions.  The most important part of the
// response is the new state, but they can optionally return another message to be
// processed after an optional delay.
type Response struct {
	NextMsg *Message
	Delay   time.Duration
	Error   error
}

// Here are a bunch of helper functions to create Responses

// Respond is the basic message response when no error has ocurred and no subsequent messages are required.
func Respond() Response {
	return Response{
		NextMsg: nil,
		Delay:   0,
		Error:   nil,
	}
}

// Respond with error responds with an error message
func RespondWithError(err error) Response {
	return Response{
		NextMsg: nil,
		Delay:   0,
		Error:   err,
	}
}

// RespondWithNextMessage responds and queues up another message with 0 delay
func RespondWithNextMsg(message Message) Response {
	return RespondWithDelayedNextMsg(message, 0)
}

// RespondWithNextMessage responds and queues up another message with a delay of N milliseconds
func RespondWithDelayedNextMsg(message Message, delay time.Duration) Response {
	return Response{
		NextMsg: &message,
		Delay:   delay,
		Error:   nil,
	}
}

// MessageHandler functions are the functions that are called when a message is received.
// Typically they would be used to make some sort of mutation to the state.
// They can also return a new message to be processed, and optionally a delay.
type MessageHandler func(Message, State) Response

// MessageMap holds a record of MessageHandler functions keyed against message.
// This enables the runtime to look up the correct function to execute for each message received.
// The client application must provide this when bootstrapping the app.
type MessageMap map[string]MessageHandler

// MergeMaps is a helper that combines several message maps into one
// This is useful for splitting up the message handling functions into separate files
func MergeMaps(msgMaps ...MessageMap) MessageMap {
	startMap := MessageMap{}

	for _, thisMap := range msgMaps {
		for k, v := range thisMap {
			startMap[k] = v
		}
	}

	return startMap
}

// handleMessage is the Melody handler that is called when a websocket message is received
// In gotea, all it does is retrieve the state from the session, and then pass the message processor
func handleMessage(s *melody.Session, msg []byte) {
	st, _ := s.Get(melodyStateKey)
	state := st.(State)

	var message Message
	if err := json.Unmarshal(msg, &message); err != nil {
		s.Write(state.RenderError(err))
		return
	}

	if err := message.process(s, state); err != nil {
		s.Write(state.RenderError(err))
		return
	}
}

// process does the actual work of dealing with an incoming message.
// It checks to make sure a message handling function is assigned to that message, raising an error if not.
// Assuming a message handling function is found, it is executed and the new state is rendered
// Any further messages are sent for processing in the same way (recursively).
func (message Message) process(s *melody.Session, state State) error {
	// Since messages can trigger themselves, they can potentially set off an infinite loop,
	// which would not be interrupted by the connection closing.
	// So here we check that the connection is open before processing the message.
	if s.IsClosed() {
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
		funcToExecute, found = state.Update()[message.Message]
		if !found {
			return fmt.Errorf("Could not process message %s: message does not exist", message.Message)
		}
	}

	// Execute the message handler function and get the response
	response := funcToExecute(message, state)
	if response.Error != nil {
		return response.Error
	}

	// Now we can render the new state
	if !message.BlockRerender {
		s.Write(state.Render())
	}

	// If there is a next message, we process it
	// Note: this must happen in a go routine to unblock this session from receiving further messages
	if response.NextMsg != nil {
		go func() {
			if response.Delay > 0 {
				time.Sleep(response.Delay * time.Millisecond)
			}

			response.NextMsg.process(s, state)
		}()
	}

	return nil
}

// APPLICATION

// Application is the holder for
// - the Melody instance
// - the template for creating new state
type Application struct {
	*melody.Melody
	Model State
}

// NewApp is used by the calling application to set up a new gotea app
// - sets up a new Melody instance
// - and attach the connection and message handlers
func NewApp(model State) *Application {
	melody := melody.New()
	melody.HandleConnect(onConnect(model))
	melody.HandleMessage(handleMessage)

	return &Application{
		Melody: melody,
		Model:  model,
	}
}

// Starts the application on a specified port
// - serves the websocket connection endpoint
// - serves static files from the specified directory
// - initial render for all other routes
func (app *Application) Start(port int, staticDirectory string) {
	http.HandleFunc("/server", func(w http.ResponseWriter, r *http.Request) {
		app.Melody.HandleRequest(w, r)
	})

	staticDirectoryWithBothSlashes := fmt.Sprintf("/%v/", staticDirectory)
	fs := http.FileServer(http.Dir(staticDirectory))
	http.Handle(staticDirectoryWithBothSlashes, http.StripPrefix(staticDirectoryWithBothSlashes, fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Check for session cookie
		cookie, err := r.Cookie("session_id")
		var sessionID string
		if err != nil || cookie.Value == "" {
			// Create a new session ID if not present
			sessionID = uuid.New().String()
			http.SetCookie(w, &http.Cookie{
				Name:    "session_id",
				Value:   sessionID,
				Expires: time.Now().Add(24 * time.Hour),
			})
			log.Printf("New session created with ID: %s", sessionID)
		} else {
			sessionID = cookie.Value
			log.Printf("Existing session found with ID: %s", sessionID)
		}

		state := app.Model.Init(uuid.MustParse(sessionID))
		changeRoute(state, r.URL.Path)
		w.Write(state.Render())
	})

	log.Printf("Starting application server on %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

// Broadcast rerenders all sessions
func (app *Application) Broadcast() {
	sessions, _ := app.Melody.Sessions()
	for _, s := range sessions {
		st, _ := s.Get("state")
		state := st.(State)
		s.Write(state.Render())
	}
}
