package gotea

import (
	"encoding/json"
)

type Routable interface {
	SetRoute(string)
	GetRoute() string
}

type Router struct {
	Route string
}

func (r *Router) SetRoute(newRoute string) {
	r.Route = newRoute
}

func (r Router) GetRoute() string {
	return r.Route
}

func changeRoute(args json.RawMessage, s State) (State, *Message, error) {
	var newRoute string
	if err := json.Unmarshal(args, &newRoute); err != nil {
		return s, nil, err
	}

	s.SetRoute(newRoute)
	return s, nil, nil
}

func MatchRoute(currentRoute, toMatch string) bool {
	if currentRoute == toMatch {
		return true
	}

	return false
}
