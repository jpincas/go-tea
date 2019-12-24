package gotea

import (
	"encoding/json"
	"net/url"
)

type Routable interface {
	SetNewRoute(string, string)
	GetRoute() string
	GetTemplate() string
	RouteParam(string) string
}

type Router struct {
	Route        string
	TemplateName string
}

var routingMessages = MessageMap{
	"CHANGE_ROUTE": changeRoute,
}

func (r *Router) SetNewRoute(route, template string) {
	r.Route = route
	r.TemplateName = template
}

func changeRoute(args json.RawMessage, state State) Response {
	var newRoute string
	if err := json.Unmarshal(args, &newRoute); err != nil {
		return RespondWithError(err)
	}

	setRoute(state, newRoute)
	return Respond()
}

func setRoute(state State, newRoute string) {
	// Parse the new route
	u, _ := url.Parse(newRoute)

	// Before returning, fire the route update hook.
	// This is provided by the application and can implement any
	// custom logic required
	newTemplate := state.Mux(u)

	// Set the new route and template on the router
	state.SetNewRoute(newRoute, newTemplate)
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
	return r.Route
}

func (r Router) GetTemplate() string {
	return r.TemplateName
}
