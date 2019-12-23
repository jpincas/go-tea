package gotea

import (
	"encoding/json"
	"net/url"
	"strings"
)

type Routable interface {
	SetRoute(string)
	RouteParam(string) string
	RouteTemplate() string
	RouteUpdateHook() State
}

type Router struct {
	Route string
}

var routingMessages = MessageMap{
	"CHANGE_ROUTE": changeRoute,
}

func (r *Router) SetRoute(newRoute string) {
	r.Route = newRoute
}

func changeRoute(args json.RawMessage, s State) Response {
	var newRoute string
	if err := json.Unmarshal(args, &newRoute); err != nil {
		return RespondWithError(s, err)
	}

	// Set the new route on the router
	s.SetRoute(newRoute)

	// Before returning, fire the route update hook.
	// This is provided by the application and can implement any
	// custom logic required
	return Respond(s.RouteUpdateHook())
}

//  Route Helpers
func (r Router) RouteParam(param string) string {
	rel, err := url.Parse(r.Route)
	if err != nil {
		return ""
	}

	return rel.Query().Get(param)
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
