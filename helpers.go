package gotea

import "encoding/json"

func mergeMaps(msgMaps ...MessageMap) MessageMap {
	startMap := MessageMap{}

	for _, thisMap := range msgMaps {
		for k, v := range thisMap {
			startMap[k] = v
		}
	}

	return startMap
}

func WithNextMsg(state State, msg string, args json.RawMessage) (State, *Message, error) {
	return state, &Message{
		Message:   msg,
		Arguments: args,
	}, nil
}

func WithNoMsg(state State) (State, *Message, error) {
	return state, nil, nil
}
