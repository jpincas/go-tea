package gotea

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func init() {

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

		var message Message

		err := conn.ReadJSON(&message)
		if err != nil {
			renderError(conn, err)
			break
		}

		if err := message.Process(session); err != nil {
			renderError(conn, err)
		}

	}
}
