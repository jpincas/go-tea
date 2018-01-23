package gotea

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func init() {
	// parse app specific templates
	// therefore, your app templates should go in the folder 'templates'
	// in the root directory of your app
	App.Templates = template.Must(template.New("main").ParseGlob("templates/*.html"))

	//set basic config
	App.Config = Config{
		AppPort: 8080,
	}
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

	session, err := newSession(conn)
	if err != nil {
		renderError(conn, err)
	}

	// defer closing the connection and removing it from the session
	defer conn.Close()
	defer session.remove()

	for {

		_, message, err := conn.ReadMessage()
		if err != nil {
			renderError(conn, err)
			break
		}

		var msg Msg
		err = json.Unmarshal(message, &msg)
		if err != nil {
			renderError(conn, err)
		}

		if err := msg.Process(session); err != nil {
			renderError(conn, err)
		}

	}
}
