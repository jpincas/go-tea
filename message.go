package gotea

import (
	"encoding/json"
	"fmt"
)

func init() {
	// initialise the message map
	// App.Messages = map[string]func(MessageArguments, *Session){}
}

type MessageArguments interface{}

// Msg repesents is the data envelope for a message
// - the Message is a simple string
// - if any data needs to be attached, it is in the form of a map of interfaces(anything)
type Message struct {
	Func      func(MessageArguments, *Session) `json:"-"`
	Arguments MessageArguments                 `json:"arguments"`
	FuncCode  string                           `json:"func"`
}

// Process a message
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
	funcToExecute(message.Arguments, session)

	// rerender
	session.render()

	return nil
}

func formatMessageAsHtmlAttrs(message Message) string {

	b, _ := json.Marshal(message)
	return fmt.Sprintf("%s", b)
}

func RegisterMessages(funcs ...func(MessageArguments) Message) {
	for _, f := range funcs {
		App.Messages[f(nil).FuncCode] = f(nil).Func
		App.TemplateFuncs[f(nil).FuncCode] = f
	}
}
