package gotea

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/CloudyKit/jet"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

//go:generate hardcode

// The idea behind gotea is simple, but can take a bit of getting your head around.
// The code is structured more like a story than regular code -
// read along from top to bottom to understand how it works

// The concepts of 'State' and 'Session' are fundamental to gotea.

// State is private to each user (session) and is what is rendered
// by the runtime on each update.
// It can essentially be anything  - you define it as a struct in your application code.
// Conventionally, you'd call it 'Model', but you don't have to!
// You get a router by embedding the go-tea Router in your model
type State interface {
	InitState() State
	Routable
}

// Session is just a holder for a websocket connection (or more accurately, a pointer to one),
// and a lump of state (see above).  When a client connects, a new session is opened with a pointer
// to the websocker connection and the initial state (more on that later)
type Session struct {
	Conn  *websocket.Conn
	State State
}

// SessionStore tracks a list of active sessions.  When a session is opened,
// it gets added to the SessionStore.
// When a session is closed, it gets removed.
type SessionStore []*Session

// newSession creates a new session, as outlined above.
// It assigns the specified websocket connection to the session,
// sets the initial state by calling a special function that your application needs
// to provide (more on that later),
// adds the session to the session store so we can keep track of it.
func (app *Application) newSession(conn *websocket.Conn, path string) (*Session, error) {
	session := Session{
		State: app.newState(path),
	}
	session.Conn = conn
	session.add(app)

	return &session, nil
}

// add simply updates the list of active sessions, as explained above.
func (session *Session) add(app *Application) {
	app.Sessions = append(app.Sessions, session)
}

// remove, as you could expect, just removes the session from the session store.
func (session *Session) remove(app *Application) {
	for i, stored := range app.Sessions {
		if stored == session {
			// safe delete NOT preserving order
			// https://github.com/golang/go/wiki/SliceTricks
			app.Sessions[i] = app.Sessions[len(app.Sessions)-1]
			app.Sessions[len(app.Sessions)-1] = nil
			app.Sessions = app.Sessions[:len(app.Sessions)-1]
		}
	}
}

// Here's where the magic starts to happen

// render takes the session state, runs it through the main view template
// and renders it to the socket connection on the session -
// this is what the client sees in their browser.
// The JS part of gotea takes this HTML and patches it efficiently onto
// the existing DOM, so the browser only updates what has actually changed
func (session *Session) render(app *Application, errorToRender error) {
	if session.Conn == nil {
		// Oops! There's no socket connection
		log.Println("Could not render, no socket to render to")
		return
	}

	w, err := session.Conn.NextWriter(1)

	if err != nil {
		log.Printf("Error opening websocket connection writer: %v\n", err)
		return
	}

	if errorToRender == nil {
		app.render(w, session.State)
	} else {
		app.renderError(w, session.State, errorToRender)
	}

	if err := w.Close(); err != nil {
		log.Printf("Error closing websocket connection writer: %v\n", err)
		return
	}
}

// But on its own, the above would be quite boring,
// it wouldn't allow for any interactivity.
// Fortunately, interactivity is baked right into the gotea runime with 'Messages'

// Message is a data structure that is triggered in JS in the browser,
// and send through the open websocket connection to the gotea runtime.
// It's quite simple and consists of just two pieces of information:
// 1 - the name of the message (a string)
// 2 - some optional accompanying data (JSON)
type Message struct {
	Message   string          `json:"message"`
	Arguments json.RawMessage `json:"args"`
}

// gotea is all about state.
// Messages arriving at the runtime are handled by

// MessageHandler functions which can do absolutely anything
// as long as they return a new state.
// Typically they would be used to make some sort of mutation.
// They can optioally send back another 'Message' to the runtime,
// which will be processed in turn,
// effectively setting off a chain
type MessageHandler func(json.RawMessage, State) (State, *Message, error)

// Your application, then, will define a set of message handling functions
// that will be called in response to incoming messages.

// MessageMap holds a record of MessageHandler functions keyed against message.
// This enables the runtime to look up the correct function to execute for each message received.
type MessageMap map[string]MessageHandler

// Process does the actual work of dealing with an incoming message.
// It checks to make sure a message handling function is assigned to that message,
// raising an error if not.
// Assuming a message handling function is found, it is executed,
// resulting in a new state.  This new state is set on the session,
// and any further messages are sent for processing in the same way (recursively).
// Finally, a render of the new state takes place, sending new HTML down the websocket to the client,
// and starting the cycle again.
func (message Message) Process(session *Session, app *Application) error {
	// Try system messages first.
	// At the moment, just the router, but could expand
	systemMessages := routingMessages
	funcToExecute, found := systemMessages[message.Message]

	// TODO: We might want to check both maps here and raise
	// some sort of log message if there is a clash of names

	if !found {
		// Care to overwrite the funcToExecute variable above
		funcToExecute, found = app.Messages[message.Message]
		if !found {
			return fmt.Errorf("Could not process message %s: message does not exist", message.Message)
		}
	}

	newState, nextMessage, err := funcToExecute(message.Arguments, session.State)

	if err != nil {
		return err
	}

	session.State = newState
	session.render(app, nil)

	// TODO: new thread?
	if nextMessage != nil {
		nextMessage.Process(session, app)
	}

	return nil
}

// How do we wire this whole thing up?
// We'll need a server to serve a single endpoint.
// The handler for that endpoint will do the work
// of establishing the websocket connection and running the infinite
// loop until disconnection.

// upgrader prepares the upgrader for websocket connections
var upgrader = websocket.Upgrader{
	// TODO: this needs to be configured
	CheckOrigin: func(r *http.Request) bool { return true },
}

// websocketHandler is the handler function called when a client connects.
// It is basically the core of the runtime.  Here's what it does:
// - upgrades the connection to a websocket
// - creates a new session and adds it to the list of active sessions
// - waits for a message from the client
// - sends the messages for processing
// - waits again
func (app *Application) websocketHandler(w http.ResponseWriter, r *http.Request) {
	// upgrade the connections
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	// The session needs to know which route it is starting from,
	// else the first template render will fail.
	// We can't just use the path from the URL, since the websocket
	// connection is always through /server.
	// Therefore, the JS adds a ?whence=route parameter to /server
	// when making the connection, so we get the starting route from there
	r.ParseForm()
	startingRoute := r.URL.Query().Get("whence")

	// create a new session
	// this will use the state seeder to create a default state
	session, err := app.newSession(conn, startingRoute)
	if err != nil {
		session.render(app, err)
	}

	// defer closing the connection and removing it from the session
	defer conn.Close()
	defer session.remove(app)

	// the main runtime loop
	for {
		// read the incoming message
		var message Message
		err := conn.ReadJSON(&message)
		if err != nil {
			session.render(app, err)
			break
		}

		// and send for processing
		if err := message.Process(session, app); err != nil {
			session.render(app, err)
		}

	}
}

// OK, that's pretty much it.
// All that's left now is to bring all this together and start the server

type AppConfig struct {
	Port            int
	HomeTemplate    string
	StaticDirectory string
}

var DefaultAppConfig = AppConfig{
	Port:            8080,
	HomeTemplate:    "home",
	StaticDirectory: "static",
}

// Application is the holder for all the bits and pieces go-tea needs
type Application struct {
	Config AppConfig

	// MessageMap is the global map of messages -> message handling functions
	Messages MessageMap

	// Sessions is a list of the current active sessions
	Sessions SessionStore

	// Templates for rending, provided by the appication
	Templates *jet.Set

	Router *chi.Mux

	Model State
}

// render is the main render function for the whole app
// It specifies how to render state.
// It will look for a template whose name matches the route, e.g.
// /myroute -> myroute.html
// /myroute/subroute -> myroute_subroute.html

func write(w io.Writer, t *jet.Template, state State, vars jet.VarMap) error {
	return t.Execute(w, vars, state)
}

func (app Application) renderError(w io.Writer, state State, errorToRender error) {
	// If no 'error' template has been specified, just write the error
	t, err := app.Templates.GetTemplate("error")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	vars := make(jet.VarMap)
	vars.Set("Msg", errorToRender.Error())

	write(w, t, state, vars)
}

func (app Application) render(w io.Writer, state State) {
	templateName := state.RouteTemplate(app.Config.HomeTemplate)

	t, err := app.Templates.GetTemplate(templateName)
	if err != nil {
		app.renderError(w, state, err)
		return
	}

	vars := make(jet.VarMap)
	if err = write(w, t, state, vars); err != nil {
		app.renderError(w, state, err)
		return
	}
}

// And finally you are ready to start gotea.

// NewApp is used by the calling application to set up a new gotea app
func NewApp(config AppConfig, model State, msgMaps ...MessageMap) *Application {
	app := Application{
		Config:   config,
		Model:    model,
		Sessions: SessionStore{},
		// Combine the built-in messages with the application level messages
		Messages:  mergeMaps(msgMaps...),
		Templates: parseTemplates(),
	}

	app.initRouter()
	return &app
}

func (app *Application) initRouter() {
	router := chi.NewRouter()

	// Attach the w ebsocket handler at /server,
	router.Get("/server", app.websocketHandler)

	// Serve the gotea JS
	router.Get("/gotea.js", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		w.Write([]byte(goteaJS))
	})

	// Attach the static file server if required
	if app.Config.StaticDirectory != "" {
		fileServer(router, "/static", http.Dir(app.Config.StaticDirectory))
	}

	// For all other routes, serve index.html
	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		state := app.newState(r.URL.Path)
		app.render(w, state)
	})

	app.Router = router
}

func (app Application) newState(path string) State {
	state := app.Model.InitState()
	state.SetRoute(path)
	return state
}

// Start creates the router, and serves it!
func (app *Application) Start() {
	fmt.Printf("Starting application server on %v\n", app.Config.Port)
	http.ListenAndServe(fmt.Sprintf(":%v", app.Config.Port), app.Router)
}

// That's all the important stuff.
// There are just a few other odds and ends.

// If your application needs to rerender all active sessions
// because of some change in global state, there's an easy function for that

// Broadcast re-renders every active session
func (app *Application) Broadcast() {
	for _, session := range app.Sessions {
		if session.Conn != nil {
			session.render(app, nil)
		}
	}
}
