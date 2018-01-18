package main

import "github.com/gorilla/websocket"

func processMessage(msg MsgPayload, conn *websocket.Conn) {

	if funcToExecute, ok := Messages[msg.Message]; ok {
		funcToExecute(msg.Data, conn)
	}

	// rerender the view
	renderView(conn)

}
