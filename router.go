package gotea

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type Routable interface {
	SetRoute(string)
	GetRoute() string
	RouteTemplate(string) string
	RouteParam(string) string
}

type Router struct {
	Route string
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
func (r Router) RouteTemplate(extension string) string {
	underscoredRoute := strings.Replace(r.GetRoute(), "/", "_", -1)
	return fmt.Sprintf("%s.%s", underscoredRoute, extension)
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

	s.SetRoute(newRoute)
	return s, nil, nil
}
