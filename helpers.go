package gotea

import (
	"encoding/json"
	"time"
)

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

func Respond() Response {
	return Response{
		NextMsg: nil,
		Delay:   0,
		Error:   nil,
	}
}

func RespondWithError(err error) Response {
	return Response{
		NextMsg: nil,
		Delay:   0,
		Error:   err,
	}
}

func RespondWithNextMsg(msg string, args json.RawMessage) Response {
	return Response{
		NextMsg: &Message{
			Message:   msg,
			Arguments: args,
		},
		Delay: 0,
		Error: nil,
	}
}

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
