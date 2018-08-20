package gotea

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

// The idea behind gotea is simple, but can take a bit of getting your head around.
// The code is structured more like a story than regular code -
// read along from top to bottom to understand how it works

// The concepts of 'State' and 'Session' are fundamental to gotea.

// State is private to each user (session) and is what is rendered
// by the runtime on each update.
// It can essentially be anything  - you define it as a struct in your application code.
// Conventionally, you'd call it 'Model', but you don't have to!
// The only requirement is that the state you define should be able to set and get a 'route' - predictably used for routing.
// Typically, you do that by including a 'route' field on your state and then define 'SetState()' and 'GetState()' on your model.  This will cause your app to respond correctly to in-built routing messages.
type State Routable

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
// adds the session to the session store so we can keep track of it,
// and finally, does an initial render (again, more on that coming up)
func newSession(conn *websocket.Conn) (*Session, error) {
	session := App.NewSession()
	session.Conn = conn
	session.add()

	session.render()

	return &session, nil
}

// add simply updates the list of active sessions, as explained above.
func (session *Session) add() {
	App.Sessions = append(App.Sessions, session)
}

// remove, as you could expect, just removes the session from the session store.
func (session *Session) remove() {
	for i, stored := range App.Sessions {
		if stored == session {
			// safe delete NOT preserving order
			// https://github.com/golang/go/wiki/SliceTricks
			App.Sessions[i] = App.Sessions[len(App.Sessions)-1]
			App.Sessions[len(App.Sessions)-1] = nil
			App.Sessions = App.Sessions[:len(App.Sessions)-1]
		}
	}
}

// Here's where the magic starts to happen

// render takes the session state, runs it through the main view template
// and renders it to the socket connection on the session -
// this is what the client sees in their browser.
// The JS part of gotea takes this HTML and patches it efficiently onto
// the existing DOM, so the browser only updates what has actually changed
func (session *Session) render() {
	if session.Conn == nil {
		// Oops! There's no socket connection
		log.Println("Could not render, no socket to render to")
		return
	}

	session.Conn.WriteMessage(1, App.RenderView(session.State))
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
func (message Message) Process(session *Session) error {
	funcToExecute, found := App.Messages[message.Message]
	if !found {
		return fmt.Errorf("Could not process message %s: message does not exist", message.Message)
	}

	newState, nextMessage, err := funcToExecute(message.Arguments, session.State)

	if err != nil {
		return err
	}

	session.State = newState
	session.render()

	// TODO: new thread?
	if nextMessage != nil {
		nextMessage.Process(session)
	}

	return nil
}

// How do we wire this whole thing up?
// We'll need a server to serve a single endpoint.
// The handler for that endpoint will do the work
// of establishing the websocket connection and running the infinite
// loop until disconnection.

// upgrader prepares the upgrader for websocket connections
var upgrader = websocket.Upgrader{}

// websocketHandler is the handler function called when a client connects.
// It is basically the core of the runtime.  Here's what it does:
// - upgrades the connection to a websocket
// - creates a new session and adds it to the list of active sessions
// - waits for a message from the client
// - sends the messages for processing
// - waits again
func websocketHandler(w http.ResponseWriter, r *http.Request) {
	// upgrade the connections
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	// create a new session
	// this will use the state seeder to create a default state
	session, err := newSession(conn)
	if err != nil {
		renderError(conn, err)
	}

	// defer closing the connection and removing it from the session
	defer conn.Close()
	defer session.remove()

	// the main runtime loop
	for {

		// read the incoming message
		var message Message
		err := conn.ReadJSON(&message)
		if err != nil {
			renderError(conn, err)
			break
		}

		// and send for processing
		if err := message.Process(session); err != nil {
			renderError(conn, err)
		}

	}
}

// OK, that's pretty much it.
// All that's left now is to bring all this together and start the server

// Application is the holder for all the bits and pieces go-tea needs
type Application struct {
	// ErrorTemplate is the system template for rendering all errors
	ErrorTemplate *template.Template

	// MessageMap is the global map of messages -> message handling functions
	Messages MessageMap

	// Sessions is a list of the current active sessions
	Sessions SessionStore

	// NewSession is the function assigned by our application
	// to initialise a new sesson with a default initial state.
	NewSession func() Session

	// RenderView is the main render function
	// that turns state into HTML to be sent to the client
	RenderView func(State) []byte
}

// App kicks everything off, holding the global application state
var App = Application{
	Sessions: SessionStore{},
	// Rather than starting with a completely blank maessage map,
	// we start with some built in go-tea messages.
	// Therefore, when defining your application message map, you will ADD
	// it to this existing map, rather than overwrite it.
	// gotea provides the ''MergeMap' method on MessageMap for that!
	Messages: MessageMap{
		"CHANGE_ROUTE": changeRoute,
	},
	ErrorTemplate: template.Must(template.New("error").Parse(errorTemplate)),
}

// And finally you are ready to start gotea.

// Start creates the router, and serves it!
func (app Application) Start(distDirectory string, port int) {
	if app.NewSession == nil {
		log.Fatalln("ERROR: No session state seeder function specificied.  Exiting...")
	}

	if app.RenderView == nil {
		log.Fatalln("Error: No main view render function set. Exiting...")
	}

	router := chi.NewRouter()
	// Attach the websocket handler at /server
	router.Get("/server", websocketHandler)

	// Attach the static file serer at /dist
	fileServer(router, "/dist", http.Dir(distDirectory))

	// For all other routes, serve index.html
	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, distDirectory+"/index.html")
	})

	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}

// That's all the important stuff.
// There are just a few other odds and ends.

// Error handling at various points in the gotea runtime is handled by a simple
// error rendering procedure.

// errorTemplate is the HTML used to render errors
var errorTemplate = `
<h1>Whoops!</h1>
<h2>There was a gotea runtime error</h2>
<hr />
<p>{{ .ErrorMessage }}</p>
`

// renderError renders the above error template
func renderError(conn *websocket.Conn, err error) {
	if conn == nil {
		// Oops! There's no socket connection
		log.Println("Could not render, no socket to render to")
		return
	}
	tpl := bytes.Buffer{}

	templateData := struct {
		ErrorMessage string
	}{
		err.Error(),
	}

	App.ErrorTemplate.Execute(&tpl, templateData)
	conn.WriteMessage(1, tpl.Bytes())
}

// That's all there is to error handling.

// If your application needs to rerender all active sessions
// because of some change in global state, there's an easy function for that

// Broadcast re-renders every active session
func (app Application) Broadcast() {
	for _, session := range app.Sessions {
		if session.Conn != nil {
			session.render()
		}
	}
}
