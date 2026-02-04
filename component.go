package gotea

type ComponentID string

func (c ComponentID) UniqueMsg(msg string) string {
	return string(c) + "_" + msg
}

func (c ComponentID) UniqueID(id string) string {
	return string(c) + "-" + id
}

func (c ComponentID) UniqueMsgMap(messagesWithHandlers MessageMap) MessageMap {
	msgMap := MessageMap{}

	for message, handler := range messagesWithHandlers {
		msgMap[c.UniqueMsg(message)] = handler
	}

	return msgMap
}

// String returns the ComponentID as a string
func (c ComponentID) String() string {
	return string(c)
}

// SendMessage creates a JS call for a component message with ComponentID automatically set
func (c ComponentID) SendMessage(msg string, args any) string {
	return SendMessage(Message{
		Message:     c.UniqueMsg(msg),
		Arguments:   args,
		ComponentID: string(c),
	})
}

// SendMessageNoArgs creates a JS call for a component message with no arguments
func (c ComponentID) SendMessageNoArgs(msg string) string {
	return SendMessage(Message{
		Message:     c.UniqueMsg(msg),
		ComponentID: string(c),
	})
}

// SendMessageWithValueFromInput creates a JS call that reads value from an input
func (c ComponentID) SendMessageWithValueFromInput(msg string, inputID string) string {
	return SendMessageWithValueFromInput(
		Message{
			Message:     c.UniqueMsg(msg),
			ComponentID: string(c),
		},
		inputID,
	)
}

// UpdateForm creates a JS call to serialize a form and send as message args
func (c ComponentID) UpdateForm(msg string, formID string) string {
	return UpdateFormState(
		Message{
			Message:     c.UniqueMsg(msg),
			ComponentID: string(c),
		},
		formID,
	)
}
