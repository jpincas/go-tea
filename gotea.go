package gotea

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// APPLICATION

var App Application = Application{
	Sessions:      SessionStore{},
	Messages:      MessageMap{},
	ErrorTemplate: template.Must(template.New("error").Parse(errorTemplate)),
}

type Application struct {
	ErrorTemplate *template.Template
	Messages      MessageMap
	Sessions      SessionStore
	NewSession    func() Session
	RenderView    func(State) []byte
}

// Start runs the application server
func (app Application) Start(distDirectory string, port int) {
	if app.NewSession == nil {
		log.Fatalln("ERROR: No session state seeder function specificied.  Exiting...")
	}

	if app.RenderView == nil {
		log.Fatalln("Error: No main view render function set. Exiting...")
	}

	fs := http.FileServer(http.Dir(distDirectory))
	http.HandleFunc("/server", handler)
	http.Handle("/", fs)
	log.Println("Staring gotea app server...")
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

// Broadcast re-renders every active session
func (app Application) Broadcast() {
	for _, session := range app.Sessions {
		if session.Conn != nil {
			session.render()
		}
	}
}

// WEBSOCKET HANDLER

// prepare the upgrader for websocket connections
var upgrader = websocket.Upgrader{}

// handler is the function called when a client connects:
// - it is basically the core of the runtime
// - upgrades the connection to a websocket
// - creates a new session and adds it to the list of active sessions
// - waits for a message from the client
// - sends the messages for processing
// - waits again
func handler(w http.ResponseWriter, r *http.Request) {
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

// SESSIONS

// State is the per session model that is rerendered on every cycle
// - it can essentially be anything
// - conventionally you would define it a as Model in your app
type State interface{}

// Session is a combination of a websocket connection and some client state
// - the state is a Model, and is defined specifically by your app
type Session struct {
	Conn  *websocket.Conn
	State State
}

// SessionStore tracks a list of active sessions
type SessionStore []*Session

// newSession creates a new session
// - assigns the specified websocket connection to the session
// - sets the intial stats by calling the app specific initialState()
// - saves the sessions
func newSession(conn *websocket.Conn) (*Session, error) {
	// create the session from seed state, add the connection
	// and add the session to the active sessions list
	session := App.NewSession()
	session.Conn = conn
	session.add()

	// render the initial state straight away
	session.render()

	return &session, nil
}

// add updates the list of active sessions
func (session *Session) add() {
	App.Sessions = append(App.Sessions, session)
}

// remove a session from the active session list
// when its connection expires
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

// render takes the session state, runs it through the main view template
// and renders it to the socket connection on the session
func (session *Session) render() {
	if session.Conn == nil {
		// Oops! There's no socket connection
		log.Println("Could not render, no socket to render to")
		return
	}

	// render the state to the websocket connection
	session.Conn.WriteMessage(1, App.RenderView(session.State))
}

// MESSAGES

// MessageArguments are the parameters used in messages
// - they can essentially be anything
// - they are serialised and deserialised to JSON
type MessageArguments interface{}

// Msg repesents is the data envelope for a message
// - the actual function is not de/serialised
// - instead the string representing the functions name is de/serialised
type Message struct {
	Message   string           `json:"message"`
	Arguments MessageArguments `json:"args"`
}

type MessageHandler func(MessageArguments, *Session) (State, *Message)
type MessageMap map[string]MessageHandler

// Process a messages
// - lookup the message in the App-level messages map
// - if it is not found, return an error
func (message Message) Process(session *Session) error {
	// check that the message exists, return an error if not
	funcToExecute, found := App.Messages[message.Message]
	if !found {
		return fmt.Errorf("Could not process message %s: message does not exist", message.Message)
	}

	// execute the function attached to the message
	// supplying the tag as argument
	newState, nextMessage := funcToExecute(message.Arguments, session)

	// set new state and render
	session.State = newState
	session.render()

	// if there is another message o process, do it now
	// Question: new thread?
	if nextMessage != nil {
		nextMessage.Process(session)
	}

	return nil
}

// ERROR

var errorTemplate = `
<h1>Whoops!</h1>
<h2>There was a gotea runtime error</h2>
<hr />
<p>{{ .ErrorMessage }}</p>
`

// renderError renders a friendly error message in the browser
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

// FUTURE ROUTING CODE - LEAVE FOR NOW

// TODO:
// USe this code for routing
// http.HandleFunc("/server", handler)
// http.HandleFunc("/", renderIndex)
// log.Println("Staring gotea app server...")
// http.ListenAndServe(fmt.Sprintf(":%v", app.Config.AppPort), nil)

// func renderIndex(w http.ResponseWriter, r *http.Request) {

// 	templateData := struct {
// 		AppTitle string
// 	}{
// 		"gotea App",
// 	}

// 	App.IndexTemplate.Execute(w, templateData)

// }
