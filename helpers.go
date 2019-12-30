package gotea

import (
	"encoding/json"
	"time"
)

// MessageMap creation helpers

// MergeMaps combines several message maps into one
func MergeMaps(msgMaps ...MessageMap) MessageMap {
	startMap := MessageMap{}

	for _, thisMap := range msgMaps {
		for k, v := range thisMap {
			startMap[k] = v
		}
	}

	return startMap
}

// Message Response Helpers

// Respond is the basic message response when no error has ocurred and no subsequent messages are required.
func Respond() Response {
	return Response{
		NextMsg: nil,
		Delay:   0,
		Error:   nil,
	}
}

// Respond with error responds with an error message
func RespondWithError(err error) Response {
	return Response{
		NextMsg: nil,
		Delay:   0,
		Error:   err,
	}
}

// RespondWithNextMessage responds and queues up another message with 0 delay
func RespondWithNextMsg(msg string, args json.RawMessage) Response {
	return RespondWithDelayedNextMsg(msg, args, 0)
}

// RespondWithNextMessage responds and queues up another message with a delay of N milliseconds
func RespondWithDelayedNextMsg(msg string, args json.RawMessage, delay time.Duration) Response {
	return Response{
		NextMsg: &Message{
			Message:   msg,
			Arguments: args,
		},
		Delay: delay,
		Error: nil,
	}
}

// Application level helpers

// Broadcast re-renders every active session
func (app *Application) Broadcast() {
	for _, session := range app.Sessions {
		if session.conn != nil {
			session.render(app, nil)
		}
	}
}
