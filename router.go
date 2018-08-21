package gotea

import (
	"encoding/json"
	"net/url"
	"strings"
)

type Routable interface {
	SetRoute(string)
	BaseRoute() string
	RouteTemplate(string) string
	RouteParam(string) string
}

type Router struct {
	Route string
}

func (r *Router) SetRoute(newRoute string) {
	r.Route = newRoute
}

func (r Router) BaseRoute() string {
	trimmed := strings.Trim(r.Route, "/")
	base := strings.Split(trimmed, "/")[0]

	paramsIndex := strings.Index(base, "?")
	if paramsIndex == -1 {
		return base
	}

	return base[0:paramsIndex]
}

func (r Router) RouteTemplate(extension string) string {
	return r.BaseRoute() + "." + extension
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
