package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"strconv"

	"github.com/google/uuid"
	gt "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/example/blocktrader"
	"github.com/jpincas/go-tea/example/tagselector"
	a "github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
)

type Model struct {
	gt.Router
	sessionID uuid.UUID

	TemplateName string
	MemoryGame   MemoryGame
	NameSelector tagselector.Model
	TeamSelector tagselector.Model
	Form         Form
	RouteData    int
	Animation    Animation
	Counter      int
	Chat         Chat
	Blocktrader  blocktrader.Model
}

func (m *Model) GetBlocktrader() *blocktrader.Model {
	return &m.Blocktrader
}

func (m *Model) Serialize() ([]byte, error) {
	// Create a serializable version of state
	snapshot := struct {
		Counter    int `json:"counter"`
		RouteData  int `json:"route_data"`
		MemoryGame struct {
			Score      int        `json:"score"`
			TurnsTaken int        `json:"turns_taken"`
			Difficulty Difficulty `json:"difficulty"`
		} `json:"memory_game"`
		Blocktrader blocktrader.Model `json:"blocktrader"`
	}{
		Counter:   m.Counter,
		RouteData: m.RouteData,
		Blocktrader: m.Blocktrader,
	}

	snapshot.MemoryGame.Score = m.MemoryGame.Score
	snapshot.MemoryGame.TurnsTaken = m.MemoryGame.TurnsTaken
	snapshot.MemoryGame.Difficulty = m.MemoryGame.Difficulty

	return json.Marshal(snapshot)
}

func (m *Model) Deserialize(data []byte) error {
	var snapshot struct {
		Counter    int `json:"counter"`
		RouteData  int `json:"route_data"`
		MemoryGame struct {
			Score      int        `json:"score"`
			TurnsTaken int        `json:"turns_taken"`
			Difficulty Difficulty `json:"difficulty"`
		} `json:"memory_game"`
		Blocktrader blocktrader.Model `json:"blocktrader"`
	}

	if err := json.Unmarshal(data, &snapshot); err != nil {
		return err
	}

	// Restore state
	m.Counter = snapshot.Counter
	m.RouteData = snapshot.RouteData
	m.MemoryGame.Score = snapshot.MemoryGame.Score
	m.MemoryGame.TurnsTaken = snapshot.MemoryGame.TurnsTaken
	m.MemoryGame.Difficulty = snapshot.MemoryGame.Difficulty
	m.Blocktrader = snapshot.Blocktrader

	// Ensure Blocktrader config is valid (handle case where state was saved before Blocktrader was added)
	if m.Blocktrader.Config.BoardSize == 0 {
		m.Blocktrader.Config = blocktrader.DefaultConfig
	}

	// Ensure Blocktrader state is valid
	if len(m.Blocktrader.State.Rows) == 0 || len(m.Blocktrader.State.Cols) == 0 {
		m.Blocktrader.State = blocktrader.NewState(m.Blocktrader.Config)
	}

	return nil
}

func model(s gt.State) *Model {
	return s.(*Model)
}

func (m *Model) Init(sid uuid.UUID) gt.State {
	l, err := LoadLeaderboard()
	if err != nil {
		fmt.Printf("Error loading leaderboard: %v\n", err)
		l = NewLeaderboard()
	}

	model := &Model{
		sessionID: sid,
		MemoryGame: MemoryGame{
			Deck:              NewDeck(Medium.Pairs()),
			TurnsTaken:        0,
			LastAttemptedCard: 5, //hack
			Score:             0,
			Difficulty:        Medium,
			Leaderboard:       l,
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
			Username: usernames[sid],
			Messages: &messages,
		},
		Blocktrader: blocktrader.NewModel(),
	}

	// Register Routes
	model.Register("/", func(s gt.State) []byte {
		return renderHome(model.Counter).Bytes()
	})
	model.Register("/memory", func(s gt.State) []byte {
		return model.MemoryGame.render().Bytes()
	})
	model.Register("/form", func(s gt.State) []byte {
		return model.Form.render().Bytes()
	})
	model.Register("/components", func(s gt.State) []byte {
		return renderComponents(model.NameSelector, model.TeamSelector).Bytes()
	})
	model.Register("/routing", func(s gt.State) []byte {
		return model.renderRouting().Bytes()
	})
	model.Register("/animation", func(s gt.State) []byte {
		return model.Animation.render().Bytes()
	})
	model.Register("/chat", func(s gt.State) []byte {
		return model.Chat.render().Bytes()
	})
	model.Register("/blocktrader", func(s gt.State) []byte {
		return model.Blocktrader.Render().Bytes()
	})

	return model
}

func (m *Model) Update() gt.MessageMap {
	return gt.MergeMaps(
		memoryGameMessages,
		formMessages,
		nameSelector.UniqueMsgMap(nameSelectorMessages),
		teamSelector.UniqueMsgMap(teamSelectorMessages),
		animationMessages,
		counterMessages,
		crashMessages,
		chatMessages,
		blocktrader.Messages,
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
	fmt.Printf("Route changed to %s\n", path)
}

// Rendering

func (m *Model) Render() []byte {
	el := h.Html(
		a.Attrs(a.Lang("en")),
		h.Head(
			a.Attrs(),
			h.Meta(a.Attrs(a.Charset("UTF-8"))),
			h.Meta(a.Attrs(a.Name("viewport"), a.Content("width=device-width, initial-scale=1.0"))),
			h.Meta(a.Attrs(a.HttpEquiv("X-UA-Compatible"), a.Content("ie=edge"))),
			h.Title(a.Attrs(), h.Text(m.TemplateName)),
			h.Script(a.Attrs(a.Src("https://cdn.tailwindcss.com"))),
			h.Link(a.Attrs(a.Rel("icon"), a.Type("image/png"), a.Href("data:image/png;base64,iVBORw0KGgo="))),
		),
		h.Body(
			a.Attrs(a.Class("bg-gray-100 min-h-screen font-sans text-gray-800")),
			h.Nav(
				a.Attrs(a.Class("bg-white shadow-md")),
				h.Div(
					a.Attrs(a.Class("max-w-7xl mx-auto px-4 sm:px-6 lg:px-8")),
					h.Div(
						a.Attrs(a.Class("flex justify-between h-16")),
						h.Div(
							a.Attrs(a.Class("flex")),
							h.Div(
								a.Attrs(a.Class("flex-shrink-0 flex items-center")),
								h.Span(a.Attrs(a.Class("font-bold text-xl text-indigo-600")), h.Text("🍵 Gotea")),
							),
							h.Div(
								a.Attrs(a.Class("hidden sm:ml-6 sm:flex sm:space-x-8")),
								navLink("/", "Home", m.Router.Route == "/"),
								navLink("/memory", "Memory Game", m.Router.Route == "/memory"),
								navLink("/form", "Form", m.Router.Route == "/form"),
								navLink("/components", "Components", m.Router.Route == "/components"),
								navLink("/routing", "Routing", m.Router.Route == "/routing"),
								navLink("/animation", "Animation", m.Router.Route == "/animation"),
								navLink("/chat", "Chat", m.Router.Route == "/chat"),
								navLink("/blocktrader", "Blocktrader", m.Router.Route == "/blocktrader"),
							),
						),
					),
				),
			),
			h.Main(
				a.Attrs(a.Class("max-w-7xl mx-auto py-6 sm:px-6 lg:px-8")),
				h.Div(
					a.Attrs(a.Class("px-4 py-6 sm:px-0")),
					h.Div(
						a.Attrs(a.Class("border-4 border-dashed border-gray-200 rounded-lg p-4 bg-white")),
						h.Div(a.Attrs(a.Id("view")), h.UnsafeRaw(string(m.RenderRoute(m)))),
					),
				),
			),
			h.Div(
				a.Attrs(a.Class("bg-white border-t border-gray-200 mt-auto")),
				h.Div(
					a.Attrs(a.Class("max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8")),
					h.P(a.Attrs(a.Class("text-center text-gray-500 text-sm")), 
						h.Text(fmt.Sprintf("Session ID: %s | Current Route: %s", m.sessionID.String(), m.Router.Route)),
					),
				),
			),
			h.Script(a.Attrs(a.Src("static/main.js"))),
		),
	)

	return el.Bytes()
}

func navLink(href, text string, active bool) h.Element {
	classes := "inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
	if active {
		classes += " border-indigo-500 text-gray-900"
	} else {
		classes += " border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700"
	}
	return h.A(a.Attrs(a.Href(href), a.Class(classes)), h.Text(text))
}

func (m *Model) RenderError(err error) []byte {
	return []byte(fmt.Sprintf("An error occurred: %s", err.Error()))
}

func (m *Model) RenderView() h.Element {
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
	case "/blocktrader":
		return m.Blocktrader.Render()
	default:
		return renderHome(m.Counter)
	}
}

// Placeholder functions for individual template rendering
func renderHome(counter int) h.Element {
	return h.Div(
		a.Attrs(a.Class("space-y-6")),
		renderExplanatoryNote(
			"How this works",
			`
			<p class="mb-2">This is the classic "Counter" example, the "Hello World" of the Elm Architecture.</p>
			<ul class="list-disc pl-5 space-y-1">
				<li><strong>State:</strong> A simple integer <code>Counter</code> in the <code>Model</code>.</li>
				<li><strong>Messages:</strong> <code>INCREMENT</code> and <code>DECREMENT</code> messages are sent when buttons are clicked.</li>
				<li><strong>Update:</strong> The update function modifies the <code>Counter</code> based on the message.</li>
				<li><strong>Render:</strong> The view function renders the current count and buttons.</li>
			</ul>
			`,
		),
		h.H2(a.Attrs(a.Class("text-2xl font-bold text-gray-900")), h.Text("Gotea Demo Site")),
		h.P(a.Attrs(a.Class("text-gray-600")), h.Text("Choose an example from the menu to get started, or have a play with the classic 'counter' example below:")),
		renderCounter(counter),
		h.Hr(a.Attrs(a.Class("my-8 border-gray-300"))),
		renderCrash(),
	)
}

func (m *Model) renderRouting() h.Element {
	return h.Div(
		a.Attrs(a.Class("space-y-6")),
		renderExplanatoryNote(
			"Routing in GoTea",
			`
			<p class="mb-2">GoTea provides a simple but flexible routing system.</p>
			<ul class="list-disc pl-5 space-y-1">
				<li><strong>Navigation:</strong> Links trigger a <code>CHANGE_ROUTE</code> message.</li>
				<li><strong>Router:</strong> The <code>Router</code> struct (embedded in <code>Model</code>) handles the current route state.</li>
				<li><strong>Hooks:</strong> The <code>OnRouteChange</code> hook allows you to execute logic (like data fetching) when the route changes.</li>
				<li><strong>Rendering:</strong> The main <code>Render</code> function switches on <code>m.Router.Route</code> (or a derived template name) to decide what to render.</li>
			</ul>
			`,
		),
		h.H2(a.Attrs(a.Class("text-2xl font-bold text-gray-900")), h.Text("Routing")),
		h.Div(
			a.Attrs(a.Class("bg-blue-50 p-4 rounded-md border border-blue-200")),
			h.H3(a.Attrs(a.Class("text-lg font-medium text-blue-900")), h.Text(fmt.Sprintf("Current Route: %s", m.GetRoute()))),
		),
		h.P(a.Attrs(a.Class("text-gray-600")), h.Text("The main navigation bar above is an example of GoTea routing.")),
		h.P(a.Attrs(a.Class("text-gray-600")), h.Text("Clicking on any of the links triggers a \"CHANGE_ROUTE\" message to be sent, with the new route as an argument. The app-provided routing function then decides whether to change the view template and performs any route-conditional logic.")),
		
		h.Div(
			a.Attrs(a.Class("grid grid-cols-1 md:grid-cols-2 gap-6")),
			h.Div(
				a.Attrs(a.Class("bg-white p-6 rounded-lg shadow-sm border border-gray-200")),
				h.H3(a.Attrs(a.Class("text-lg font-medium text-gray-900 mb-4")), h.Text("Internal Link")),
				h.A(a.Attrs(a.Class("text-indigo-600 hover:text-indigo-900 underline"), a.Href("/memory")), h.Text("Go to Memory Game")),
			),
			h.Div(
				a.Attrs(a.Class("bg-white p-6 rounded-lg shadow-sm border border-gray-200")),
				h.H3(a.Attrs(a.Class("text-lg font-medium text-gray-900 mb-4")), h.Text("External Link")),
				h.A(a.Attrs(a.Class("text-indigo-600 hover:text-indigo-900 underline flex items-center"), a.Target("_blank"), a.Href("https://duckduckgo.com")), 
					h.Text("Search DuckDuckGo"),
					h.Span(a.Attrs(a.Class("ml-1 text-sm")), h.Text("↗")),
				),
			),
		),

		h.Div(
			a.Attrs(a.Class("bg-white p-6 rounded-lg shadow-sm border border-gray-200")),
			h.H3(a.Attrs(a.Class("text-lg font-medium text-gray-900 mb-4")), h.Text("Route Parameters")),
			h.Ul(
				a.Attrs(a.Class("space-y-2 mb-6")),
				h.Li(a.Attrs(), h.A(a.Attrs(a.Class("text-indigo-600 hover:text-indigo-900"), a.Href("/routing?myparam=1")), h.Text("myparam = 1"))),
				h.Li(a.Attrs(), h.A(a.Attrs(a.Class("text-indigo-600 hover:text-indigo-900"), a.Href("/routing?myparam=2")), h.Text("myparam = 2"))),
				h.Li(a.Attrs(), h.A(a.Attrs(a.Class("text-indigo-600 hover:text-indigo-900"), a.Href("/routing?myparam=3")), h.Text("myparam = 3"))),
			),
			h.Div(
				a.Attrs(a.Class("bg-gray-50 p-4 rounded-md")),
				h.H4(a.Attrs(a.Class("font-medium text-gray-700")), h.Text(fmt.Sprintf("Value of 'myparam': %s", m.RouteParam("myparam")))),
			),
		),

		h.Div(
			a.Attrs(a.Class("bg-yellow-50 p-4 rounded-md border border-yellow-200")),
			h.H3(a.Attrs(a.Class("text-sm text-yellow-800")), h.Text(fmt.Sprintf("A random number generated before rendering by the routing hook (a simple representation of route-specific data loading): %d", m.RouteData))),
		),
	)
}

func renderExplanatoryNote(title, content string) h.Element {
	return h.Details(
		a.Attrs(a.Class("bg-blue-50 border border-blue-200 rounded-md mb-6")),
		h.Summary(
			a.Attrs(a.Class("cursor-pointer p-4 font-bold text-blue-900 hover:bg-blue-100 rounded-t-md outline-none")),
			h.Text(fmt.Sprintf("ℹ️ %s", title)),
		),
		h.Div(
			a.Attrs(a.Class("p-4 pt-0 text-blue-800 text-sm leading-relaxed")),
			h.UnsafeRaw(content),
		),
	)
}

var app = gt.NewApp(&Model{})

func main() {
	app.Start(8080, "static")
}
