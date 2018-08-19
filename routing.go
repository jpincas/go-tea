package gotea

import (
	"encoding/json"
)

type Routable interface {
	SetRoute(string) State
	GetRoute() string
}

func changeRoute(args json.RawMessage, s State) (State, *Message, error) {
	var newRoute string
	if err := json.Unmarshal(args, &newRoute); err != nil {
		return s, nil, err
	}

	newState := s.SetRoute(newRoute)
	return newState, nil, nil
}
