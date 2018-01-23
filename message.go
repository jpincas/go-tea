package gotea

import (
	"encoding/json"
	"fmt"
	"html/template"
	"runtime"
	"strings"
)

func init() {
	// initialise the message map
	App.Messages = map[string]func(MessageArguments, *Session){}

	// provide global template funcs
	App.TemplateFuncs = template.FuncMap{
		"Msg": formatMessageAsHtmlAttrs,
	}
}

// MessageArguments are the parameters used in messages
// - they can essentially be anything
// - they are serialised and deserialised to JSON
type MessageArguments interface{}

// Msg repesents is the data envelope for a message
// - the actual function is not de/serialised
// - instead the string representing the functions name is de/serialised
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

// formatMessageAsHtmlAttrs does the work of serialising a message
// this is assigned the function 'Msg' in the tempalte function map
// the idea is to use a message generator to create a message and pipe it
// to this function for serialisation
func formatMessageAsHtmlAttrs(message Message) string {
	b, _ := json.Marshal(message)
	return fmt.Sprintf("%s", b)
}

// RegisterMessages registers message generators at specific app level
// - the implenting app calls RegisterMessages with all its message generators as parameters
// - this function goes through them one by one and adds them to the main message map
// - and the tempalte function map so they can be accessed from templates
// - to do this, it actually calls them with nil interface argument
// - and uses the message they return to populate the maps
func RegisterMessages(funcs ...func(MessageArguments) Message) {
	for _, f := range funcs {
		App.Messages[f(nil).FuncCode] = f(nil).Func
		App.TemplateFuncs[f(nil).FuncCode] = f
	}
}

// NewMsg is a helper function to be used in app-level message generators
// - its main function is to auto fill the FuncCode field
// - by using the runtime package to find out the function name
// - this removes the need for 'users' to manually set the FuncCode
// - which eliminates duplicate naming bugs and similar
func NewMsg(f func(MessageArguments, *Session), args MessageArguments) Message {

	fpcs := make([]uintptr, 1)
	runtime.Callers(2, fpcs)
	fun := runtime.FuncForPC(fpcs[0] - 1)
	funcCode := strings.TrimPrefix(fun.Name(), "main.")

	return Message{
		Func:      f,
		FuncCode:  funcCode,
		Arguments: args,
	}

}
