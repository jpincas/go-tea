package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type MsgPayload struct {
	Message string
	Data    map[string]interface{}
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	WSCX = conn
	defer conn.Close()

	// initial view render
	renderView()

	for {

		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		var msg MsgPayload
		err = json.Unmarshal(message, &msg)
		if err != nil {
			fmt.Println("unmarshalling error:", err)
		}

		processMessage(msg)

	}
}

func renderView() {
	tpl := bytes.Buffer{}
	Templates.ExecuteTemplate(&tpl, "view.html", State)

	WSCX.WriteMessage(1, tpl.Bytes())
}
