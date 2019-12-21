package gotea

import (
	"encoding/json"
	"net/url"
	"strings"
)

type Routable interface {
	SetRoute(string)
	GetRoute() string
	RouteTemplate(string) string
	RouteParam(string) string
	FireUpdateHook(State) State
}

type Router struct {
	Route      string
	UpdateHook func(State) State
}

func (r Router) FireUpdateHook(s State) State {
	if r.UpdateHook != nil {
		return r.UpdateHook(s)
	}

	return s
}

var routingMessages = MessageMap{
	"CHANGE_ROUTE": changeRoute,
}

func (r *Router) SetRoute(newRoute string) {
	r.Route = newRoute
}

func (r Router) GetRoute() string {

	// Trim any / at the start/end
	s := strings.Trim(r.Route, "/")

	// Remove any params
	paramsStart := strings.Index(s, "?")
	if paramsStart != -1 {
		s = s[0:paramsStart]
	}

	return s
}

// RouteTemplate is a helper for associating a template file to a route
// It replaces all intermediate slashes with an underscore, so
// /baseroute/subroute -> baseroute_subroute.html
func (r Router) RouteTemplate(homeTemplate string) string {
	routeTemplate := strings.Replace(r.GetRoute(), "/", "_", -1)
	if routeTemplate == "" {
		return homeTemplate
	}

	return routeTemplate
}

func (r Router) RouteParam(param string) string {
	rel, err := url.Parse(r.Route)
	if err != nil {
		return ""
	}

	return rel.Query().Get(param)
}

func changeRoute(args json.RawMessage, s State) (State, *Message, error) {
	var newRoute string
	if err := json.Unmarshal(args, &newRoute); err != nil {
		return s, nil, err
	}

	// Set the new route on the router
	s.SetRoute(newRoute)

	// Before returning, fire the route update hook.
	// This is provided by the application and can implement any
	// custom logic required
	return s.FireUpdateHook(s), nil, nil
}
