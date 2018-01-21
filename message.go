package gotea

import "fmt"

func init() {
	// initialise the message map
	App.Messages = map[string]func(MsgTag, *Session){}
}

// Msg repesents is the data envelope for a message
// - the Message is a simple string
// - if any data needs to be attached, it is in the form of a map of interfaces(anything)
type Msg struct {
	Message string
	Tag     MsgTag
}

// MsgTag is a map of arguments to be supplied to the function
// that corresponds to the message
type MsgTag map[string]interface{}

// Process a message
// - lookup the message in the App-level messages map
// - if it is not found, return an error
func (msg Msg) Process(session *Session) error {
	// check that the message exists, return an error if not
	funcToExecute, found := App.Messages[msg.Message]
	if !found {
		return fmt.Errorf("Could not process message %s: message does not exist", msg.Message)
	}

	// execute the function attached to the message
	// supplying the tag as argument
	funcToExecute(msg.Tag, session)

	// rerender
	session.render()

	return nil
}
