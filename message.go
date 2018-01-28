package gotea

import (
	"encoding/json"
	"fmt"
)

func init() {
	// initialise the message map
	App.Messages = map[string]func(MessageArguments, *Session) (State, *Message){}

}

// MessageArguments are the parameters used in messages
// - they can essentially be anything
// - they are serialised and deserialised to JSON
type MessageArguments interface{}

// Msg repesents is the data envelope for a message
// - the actual function is not de/serialised
// - instead the string representing the functions name is de/serialised
type Message struct {
	Arguments MessageArguments `json:"arguments"`
	FuncCode  string           `json:"func"`
}

// Process a messages
// - lookup the message in the App-level messages map
// - if it is not found, return an error
func (message Message) Process(session *Session) error {
	// check that the message exists, return an error if not
	funcToExecute, found := App.Messages[message.FuncCode]
	if !found {
		return fmt.Errorf("Could not process message %s: message does not exist", message.FuncCode)
	}

	// execute the function attached to the message
	// supplying the tag as argument
	newState, nextMessage := funcToExecute(message.Arguments, session)

	// set new state and render
	session.State = newState
	session.render()

	// if there is another message o process, do it now
	// Question: new thread?
	if nextMessage != nil {
		nextMessage.Process(session)
	}

	return nil
}

// formatMessageAsHtmlAttrs does the work of serialising a message
// this is assigned the function 'Msg' in the tempalte function map
// the idea is to use a message generator to create a message and pipe it
// to this function for serialisation
func Msg(message Message) string {
	b, _ := json.Marshal(message)
	return fmt.Sprintf("%s", b)
}
