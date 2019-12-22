package gotea

import (
	"encoding/json"
	"time"
)

func mergeMaps(msgMaps ...MessageMap) MessageMap {
	startMap := MessageMap{}

	for _, thisMap := range msgMaps {
		for k, v := range thisMap {
			startMap[k] = v
		}
	}

	return startMap
}

// Message Response Helpers

func Respond(state State) Response {
	return Response{
		State:   state,
		NextMsg: nil,
		Delay:   0,
		Error:   nil,
	}
}

func RespondWithError(state State, err error) Response {
	return Response{
		State:   state,
		NextMsg: nil,
		Delay:   0,
		Error:   err,
	}
}

func RespondWithNextMsg(state State, msg string, args json.RawMessage) Response {
	return Response{
		State: state,
		NextMsg: &Message{
			Message:   msg,
			Arguments: args,
		},
		Delay: 0,
		Error: nil,
	}
}

func RespondWithDelayedNextMsg(state State, msg string, args json.RawMessage, delay time.Duration) Response {
	return Response{
		State: state,
		NextMsg: &Message{
			Message:   msg,
			Arguments: args,
		},
		Delay: delay,
		Error: nil,
	}
}
