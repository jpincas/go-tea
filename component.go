package gotea

type Component struct {
	UniqueID string
}

func (c Component) UniqueMsg(msg string) string {
	return c.UniqueID + "_" + msg
}

func (c Component) MessageMap(messagesWithHandlers map[string]MessageHandler) MessageMap {
	msgMap := MessageMap{}

	for message, handler := range messagesWithHandlers {
		msgMap[c.UniqueMsg(message)] = handler
	}

	return msgMap
}
