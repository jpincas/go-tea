package main

func processMessage(msg MsgPayload) {

	if funcToExecute, ok := Messages[msg.Message]; ok {
		funcToExecute(msg.Data)
	}

	// rerender the view
	renderView()

}
