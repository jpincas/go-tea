package gotea

import (
	"bytes"
	"log"

	"github.com/gorilla/websocket"
)

func init() {
	// initialise the session store
	App.Sessions = SessionStore{}
}

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
	tpl := bytes.Buffer{}
	App.Templates.ExecuteTemplate(&tpl, "view.html", session.State)
	session.Conn.WriteMessage(1, tpl.Bytes())
}
