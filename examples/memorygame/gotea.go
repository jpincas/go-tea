////////////////////////////////////////////////////////////////
//               GOTEA FRAMEWORK / RUNTIME                    //
////////////////////////////////////////////////////////////////

// - this is the code that represents the core gotea 'runtime'
// - ideally it would be in a separate package and imported
// - at the moment I can't quite work out a sane way to do that
// - so for now, this file needs to be included with each gotea app
// - you DO NOT touch this code when building your app
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

// APPLICATION

// Application is the top-level representation of your application
type Application struct {
	Templates *template.Template
	Messages  map[string]func(map[string]interface{}, *Session)
	Sessions  SessionStore
}

// App instantiates a new application and makes it globally available
var App Application

// init:
// - parses templates in '/templates', so thats where you should put your templates!
// - initialises the session store map
func init() {
	// parse templates
	App.Templates = template.Must(template.New("main").ParseGlob("templates/*.html"))

	// initialise the session store
	App.Sessions = SessionStore{}
}

// SESSION

// Session is a combination of a websocket connection and some client state
// - the state is a Model, and is defined specifically by your app
type Session struct {
	ID    uuid.UUID
	Conn  *websocket.Conn
	State Model
}

// SessionStore stores sessions by ID (currently UUID)
type SessionStore map[uuid.UUID]*Session

// newSession creates a new session
// - assigns the specified websocket connection to the session
// - sets the intial stats by calling the app specific initialState()
// - saves the sessions
func newSession(conn *websocket.Conn) (*Session, error) {
	// new session ID
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return nil, err
	}

	// create the session and save it
	session := Session{
		ID:    u2,
		Conn:  conn,
		State: initialState(),
	}
	session.save()

	return &session, nil

}

// save saves a session in the map
func (session *Session) save() {
	App.Sessions[session.ID] = session
}

// removeConnection sets the websocket connection associated to a session as nil and saves it
func (session *Session) removeConnection() {
	session.Conn = nil
	session.save()
}

// render takes the session state, runs it through the main view template
// and renders it to the socket connection on the session
func (session *Session) render() {
	tpl := bytes.Buffer{}
	App.Templates.ExecuteTemplate(&tpl, "view.html", session.State)
	if session.Conn == nil {
		// Oops! There's no socket connection
		log.Println("Could not render, no socket to render to")
	}
	session.Conn.WriteMessage(1, tpl.Bytes())
}

// HANDLER

var upgrader = websocket.Upgrader{}

// handler is the function called when a client connects:
// - it is basically the core of the runtime
// - upgrades the connection to a websocket
// - creates a new session and stores it
// - renders the initial view
// - waits for a message from the client
// - sends the messages for processing
// - waits again
func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	// TODO: error handling
	session, _ := newSession(conn)

	// defer closing the connection and removing it from the session
	defer conn.Close()
	defer session.removeConnection()

	session.render()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		var msg Msg
		err = json.Unmarshal(message, &msg)
		if err != nil {
			fmt.Println("unmarshalling error:", err)
		}

		msg.Process(session)

	}
}

// MESSAGES

// Msg repesents is the data envelope for a message
// - the Message is a simple string
// - if any data needs to be attached, it is in the form of a map of interfaces(anything)
type Msg struct {
	Message string
	Data    map[string]interface{}
}

// Process a message
// - lookup the message in the App-level messages map
// - if it is not found, return an error
func (msg Msg) Process(session *Session) error {
	// check that the message exists, return an error if not
	funcToExecute, found := App.Messages[msg.Message]
	if !found {
		return fmt.Errorf("Could not process message %s: message does not exist", msg.Message)
	}

	// execute the function attached to the message
	funcToExecute(msg.Data, session)

	// rerender the session
	session.render()

	return nil
}

// MAIN

// main starts the server
func main() {

	fs := http.FileServer(http.Dir("../../dist"))

	http.HandleFunc("/server", handler)
	http.Handle("/", fs)

	log.Println("Staring server...")
	http.ListenAndServe(":8080", nil)
}
