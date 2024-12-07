package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"

	gt "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/example/tagselector"
	a "github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
)

type Model struct {
	gt.BaseModel
	TemplateName string
	MemoryGame   MemoryGame
	NameSelector tagselector.Model
	TeamSelector tagselector.Model
	Form         Form
	RouteData    int
	Animation    Animation
	Counter      int
	Chat         Chat
}

func model(s gt.State) *Model {
	return s.(*Model)
}

func (m Model) Init(_ *http.Request) gt.State {
	return &Model{
		BaseModel: gt.BaseModel{},
		MemoryGame: MemoryGame{
			Deck:              NewDeck(4),
			TurnsTaken:        0,
			LastAttemptedCard: 5, //hack
			Score:             0,
		},
		NameSelector: nameSelector,
		TeamSelector: teamSelector,
		Form: Form{
			Options: []string{"option 1", "option 2", "option 3"},
		},
		RouteData: 0,
		Animation: Animation{
			X:              50,
			Y:              50,
			XDirection:     true,
			YDirection:     true,
			BackgroundSize: animationBackgroundSize,
			BallSize:       animationBallSize,
			TranslateX:     translate(50, animationBackgroundSize, animationBallSize),
			TranslateY:     translate(50, animationBackgroundSize, animationBallSize),
			IncrementX:     2,
			IncrementY:     1,
		},
		Counter: 0,
		Chat: Chat{
			Username: randomChatUsername(),
			Messages: &messages,
		},
	}
}

func (m Model) Update() gt.MessageMap {
	return gt.MergeMaps(
		memoryGameMessages,
		formMessages,
		nameSelector.UniqueMsgMap(nameSelectorMessages),
		teamSelector.UniqueMsgMap(teamSelectorMessages),
		animationMessages,
		counterMessages,
		crashMessages,
		chatMessages,
	)
}

func (m *Model) OnRouteChange(path string) {
	// Ridiculously simle routing model -
	// the name of the template is the name of the path
	p, _ := url.Parse(path)

	template := p.Path
	if template == "/" {
		template = "home"
	}

	// Generate some 'route data'
	param := m.RouteParam("myparam")
	n, _ := strconv.Atoi(param)
	m.RouteData = rand.Intn((100 * n) + 100)

	m.TemplateName = template
}

// Rendering

func (m Model) Render() []byte {
	el := h.Html(
		a.Attrs(a.Lang("en")),
		h.Head(
			a.Attrs(),
			h.Meta(a.Attrs(a.Charset("UTF-8"))),
			h.Meta(a.Attrs(a.Name("viewport"), a.Content("width=device-width, initial-scale=1.0"))),
			h.Meta(a.Attrs(a.HttpEquiv("X-UA-Compatible"), a.Content("ie=edge"))),
			h.Title(a.Attrs(), h.Text(m.TemplateName)),
			h.Link(a.Attrs(a.Rel("stylesheet"), a.Href("static/styles.css"))),
			h.Link(a.Attrs(a.Rel("icon"), a.Type("image/png"), a.Href("data:image/png;base64,iVBORw0KGgo="))),
		),
		h.Body(
			a.Attrs(),
			h.Div(
				a.Attrs(a.Class("navbar")),
				h.Ul(
					a.Attrs(),
					h.Li(a.Attrs(), h.A(a.Attrs(a.Href("/")), h.Text("Home"))),
					h.Li(a.Attrs(), h.A(a.Attrs(a.Href("/memory")), h.Text("Memory Game"))),
					h.Li(a.Attrs(), h.A(a.Attrs(a.Href("/form")), h.Text("Form"))),
					h.Li(a.Attrs(), h.A(a.Attrs(a.Href("/components")), h.Text("Components"))),
					h.Li(a.Attrs(), h.A(a.Attrs(a.Href("/routing")), h.Text("Routing"))),
					h.Li(a.Attrs(), h.A(a.Attrs(a.Href("/animation")), h.Text("Animation"))),
					h.Li(a.Attrs(), h.A(a.Attrs(a.Href("/chat")), h.Text("Chat"))),
					h.Li(a.Attrs(), h.Text(fmt.Sprintf("Current Route: %s", m.BaseModel.Router.Route))),
				),
			),
			h.Div(a.Attrs(a.Id("view")), m.RenderView()),
			h.Script(a.Attrs(a.Src("static/main.js"))),
		),
	)

	return el.Bytes()
}

func (m Model) RenderView() h.Element {
	switch m.TemplateName {
	case "/home":
		return renderHome(m.Counter)
	case "/memory":
		return m.MemoryGame.render()
	case "/form":
		return m.Form.render()
	case "/components":
		return renderComponents(m.NameSelector, m.TeamSelector)
	case "/routing":
		return m.renderRouting()
	case "/animation":
		return m.Animation.render()
	case "/chat":
		return m.Chat.render()
	default:
		return renderHome(m.Counter)
	}
}

// Placeholder functions for individual template rendering
func renderHome(counter int) h.Element {
	return h.Div(
		a.Attrs(),
		h.H2(a.Attrs(), h.Text("Gotea Demo Site")),
		h.P(a.Attrs(), h.Text("Choose an example from the menu to get started, or have a play with the classic 'counter' example below:")),
		renderCounter(counter),
		h.Hr(a.Attrs()),
		renderCrash(),
	)
}

func (m Model) renderRouting() h.Element {
	return h.Div(
		a.Attrs(),
		h.H2(a.Attrs(), h.Text("Routing")),
		h.H3(a.Attrs(), h.Text(fmt.Sprintf("Current Route: %s", m.GetRoute()))),
		h.P(a.Attrs(), h.Text("The main navigation bar above is an example of GoTea routing.")),
		h.P(a.Attrs(), h.Text("Clicking on any of the links triggers a \"CHANGE_ROUTE\" message to be sent, with the new route as an argument. The app-provided routing function then decides whether to change the view template and performs any route-conditional logic.")),
		h.H3(a.Attrs(), h.Text("Here's an internal link:")),
		h.A(a.Attrs(a.Href("/memory")), h.Text("Memory Game")),
		h.H3(a.Attrs(), h.Text("Here's an external link:")),
		h.A(a.Attrs(a.Class("external"), a.Target("_blank"), a.Href("https://duckduckgo.com")), h.Text("Search DuckDuckGo")),
		h.H3(a.Attrs(), h.Text("Route Parameters")),
		h.Ul(
			a.Attrs(),
			h.Li(a.Attrs(), h.A(a.Attrs(a.Href("/routing?myparam=1")), h.Text("myparam = 1"))),
			h.Li(a.Attrs(), h.A(a.Attrs(a.Href("/routing?myparam=2")), h.Text("myparam = 2"))),
			h.Li(a.Attrs(), h.A(a.Attrs(a.Href("/routing?myparam=3")), h.Text("myparam = 3"))),
		),
		h.H4(a.Attrs(), h.Text(fmt.Sprintf("Value of 'myparam': %s", m.RouteParam("myparam")))),
		h.H3(a.Attrs(), h.Text(fmt.Sprintf("A random number generated before rendering by the routing hook (a simple representation of route-specific data loading): %d", m.RouteData))),
	)
}

var app = gt.NewApp(
	gt.DefaultAppConfig,
	&Model{},
)

func main() {
	app.Start()
}
