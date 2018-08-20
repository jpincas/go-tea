package gotea

type ComponentID string

func (c ComponentID) UniqueMsg(msg string) string {
	return string(c) + "_" + msg
}

func (c ComponentID) UniqueMsgMap(messagesWithHandlers MessageMap) MessageMap {
	msgMap := MessageMap{}

	for message, handler := range messagesWithHandlers {
		msgMap[c.UniqueMsg(message)] = handler
	}

	return msgMap
}
