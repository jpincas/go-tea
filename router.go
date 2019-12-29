package gotea

import (
	"encoding/json"
	"errors"
	"net/url"
)

type Routable interface {
	SetNewRoute(string)
	GetRoute() string
	GetTemplate() (string, error)
	RouteParam(string) string

	// OnRouteChange must be defined by the user.  It is a routing function that determines the template to use as well as any logic to perform based on the route.
	OnRouteChange(string)
}

type Router struct {
	Route        string
	TemplateName string
}

func (r *Router) SetNewRoute(route string) {
	r.Route = route
}

func (r *Router) SetNewTemplate(templateName string) {
	r.TemplateName = templateName
}

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

func (r Router) GetTemplate() (string, error) {
	if r.TemplateName == "" {
		return "", errors.New("No template has been set")
	}

	return r.TemplateName, nil
}

var routingMessages = MessageMap{
	"CHANGE_ROUTE": changeRouteMsgHandler,
}

// changeRouteMsgHandler is the built in message handler which is fired when a
// navigation event is detected
func changeRouteMsgHandler(args json.RawMessage, state State) Response {
	var newRoute string
	if err := json.Unmarshal(args, &newRoute); err != nil {
		return RespondWithError(err)
	}

	changeRoute(state, newRoute)
	return Respond()
}

// changeRoute is fired both by the route change message handler and on establishment
// of a new state blob.  It fires the app-provided routing logic and sets the new route /// on the model.
func changeRoute(state State, newRoute string) {
	state.OnRouteChange(newRoute)
	state.SetNewRoute(newRoute)
}
