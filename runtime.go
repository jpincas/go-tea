package gotea

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/CloudyKit/jet"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/websocket"
)

//go:generate hardcode

// The concepts of 'State' and 'Session' are fundamental to gotea.

// State is private to each user (session) and is what is rendered by the runtime on each update.
// It can essentially be anything  - you just define it as a struct in your application code.
// By Elm convention it would be called 'Model', but that's up to you.
type State interface {
	// USER DEFINED METHODS

	// Init must be defined by the user and describes the 'blank' state from which a session starts.
	// It gets passed a pointer to the original http request which might help to set some starting parameters,
	// but can most often be ignored
	Init(r *http.Request) State

	// Update is defined by the user and defines the MessageMap that is used to modify state.
	// on each loop of the runtime
	Update() MessageMap

	// PROVIDED METHODS

	// SetOriginal request records the original http request on the state for later use if required
	SetOriginalRequest(*http.Request)

	// Routable provides all the routing functions
	Routable
}

// BaseModel should be embedded in the client model
type BaseModel struct {
	Router
	OriginalRequest *http.Request
}

// SetOriginalRequest records the original http request on the model
func (b *BaseModel) SetOriginalRequest(request *http.Request) {
	b.OriginalRequest = request
}

// Session is just a holder for a websocket connection (or more accurately, a pointer to one),
// and a lump of state (see above).  When a client connects, a new session is opened with a pointer
// to the websocker connection and the initial state (more on that later)
type Session struct {
	Closed bool
	Conn   *websocket.Conn
	State  State
}

// SessionStore tracks a list of active sessions.  When a session is opened,
// it gets added to the SessionStore. When a session is closed, it gets removed.
type SessionStore []*Session

// newSession creates a new session, as outlined above.
// It assigns the specified websocket connection to the session,
// sets the initial state by calling a special function that your application needs
// to provide (more on that later),
// and adds the session to the session store so we can keep track of it.
func (app *Application) newSession(r *http.Request, conn *websocket.Conn, path string) *Session {
	newState := app.newState(r, path)

	session := Session{
		State: newState,
	}
	session.Conn = conn
	session.add(app)

	return &session
}

// add updates the list of active sessions
func (session *Session) add(app *Application) {
	app.Sessions = append(app.Sessions, session)
}

// close does the work of closing a session, including recording on the model that it is about to close,
// actually closing the connection, and removing it from the list.
func (session *Session) close(app *Application) {
	log.Println("Closing session")
	session.Closed = true
	session.Conn.Close()

	// remove, as you could expect, just removes the session from the session store.
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

// render takes the session state, runs it through the main view template
// and renders it to the socket connection on the session -
// this is what the client sees in their browser.
// The JS part of gotea takes this HTML and patches it efficiently onto
// the existing DOM, so the browser only updates what has actually changed
func (session *Session) render(app *Application, errorToRender error) {
	// There is no point trying to render a seesion if the message is CLOSE
	// because logically it will fail
	if websocket.IsCloseError(errorToRender, websocket.CloseGoingAway) {
		return
	}

	if session.Conn == nil {
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

// Message is a data structure that is triggered in JS in the browser,
// and send through the open websocket connection to the gotea runtime.
// It's quite simple and consists of just two pieces of information:
// 1 - the name of the message (a string)
// 2 - some optional accompanying data (JSON)
type Message struct {
	Message   string          `json:"message"`
	Arguments json.RawMessage `json:"args"`
}

// Response is returned by MessageHandler functions.  The most important part of the
// response is the new state, but they can optionally return another message to be
// processed (after an optional delay)
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

// Process does the actual work of dealing with an incoming message.
// It checks to make sure a message handling function is assigned to that message, raising an error if not.
// Assuming a message handling function is found, it is executed,
// resulting in a new state.  This new state is set on the session,
// and any further messages are sent for processing in the same way (recursively).
// Finally, a render of the new state takes place, sending new HTML down the websocket to the client,
// and starting the cycle again.
func (message Message) Process(session *Session, app *Application) error {
	// Since messages can trigger themselves, they can potentially set off an infinite loop,
	// which would not be interrupted by the connection closing.  So here we check that the connection is open
	// before processing the message.
	if session.Closed {
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
		funcToExecute, found = app.Messages[message.Message]
		if !found {
			return fmt.Errorf("Could not process message %s: message does not exist", message.Message)
		}
	}

	response := funcToExecute(message.Arguments, session.State)
	if response.Error != nil {
		return response.Error
	}

	// session.State = response.State
	session.render(app, nil)

	if response.NextMsg != nil {
		if response.Delay > 0 {
			time.Sleep(response.Delay * time.Millisecond)
		}

		response.NextMsg.Process(session, app)
	}

	return nil
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
	upgrader := websocket.Upgrader{
		CheckOrigin: app.Config.CheckOrigin,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("error upgrading connection:", err)
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
	session := app.newSession(r, conn, startingRoute)
	session.render(app, err)

	// defer closing the connection and removing it from the session
	defer session.close(app)

	// the main runtime loop
	for {
		// read the incoming message
		var message Message
		if err := conn.ReadJSON(&message); err != nil {
			session.render(app, err)
			break
		}

		// process inside a go routine so as to not block the runtime loop
		go func() {
			if err := message.Process(session, app); err != nil {
				session.render(app, err)
			}
		}()
	}
}

// OK, that's pretty much it.
// All that's left now is to bring all this together and start the server

type AppConfig struct {
	Port                                int
	TemplatesDirectory, StaticDirectory string
	CheckOrigin                         func(r *http.Request) bool
}

var DefaultAppConfig = AppConfig{
	Port:               8080,
	TemplatesDirectory: "templates",
	StaticDirectory:    "static",
	CheckOrigin:        func(_ *http.Request) bool { return true },
}

// Application is the holder for all the bits and pieces go-tea needs
type Application struct {
	Config AppConfig

	// Sessions is a list of the current active sessions
	Sessions SessionStore

	// Templates for rending, provided by the appication
	Templates *jet.Set

	Router *chi.Mux

	Model State

	Messages MessageMap
}

// render is the main render function for the whole app
// It specifies how to render state.

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

	t, err := app.Templates.GetTemplate(state.GetTemplate())
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
func NewApp(config AppConfig, model State, router *chi.Mux) *Application {
	app := Application{
		Config:    config,
		Model:     model,
		Messages:  model.Update(),
		Sessions:  SessionStore{},
		Templates: parseTemplates(config.TemplatesDirectory),
	}

	// If user has not passed in a preexisting router, use a new one
	if router == nil {
		router = chi.NewRouter()
		router.Use(middleware.Logger)
	}

	app.initRouter(router)
	return &app
}

func (app *Application) initRouter(router *chi.Mux) {
	router.Route("/", func(router chi.Router) {
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
			state := app.newState(r, r.URL.Path)
			app.render(w, state)
		})
	})

	app.Router = router
}

// newState bootstraps a new state model according to the init()
// provided by the calling app.  We also record the original http request
// and perform a 'route set' which will run any route dependent logic
// as well as set the starting template.
func (app Application) newState(r *http.Request, path string) State {
	state := app.Model.Init(r)
	state.SetOriginalRequest(r)
	changeRoute(state, path)
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
