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
	PixelCanvas  PixelCanvas
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
		PixelCanvas: NewPixelCanvas(),
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
	model.Register("/pixelcanvas", func(s gt.State) []byte {
		return model.PixelCanvas.render().Bytes()
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
		pixelCanvasMessages,
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
			h.Title(a.Attrs(), h.Text(fmt.Sprintf("Gotea | %s", m.TemplateName))),
			h.Script(a.Attrs(a.Src("https://cdn.tailwindcss.com"))),
			h.Link(a.Attrs(a.Rel("preconnect"), a.Href("https://fonts.googleapis.com"))),
			h.Link(a.Attrs(a.Rel("preconnect"), a.Href("https://fonts.gstatic.com"), a.Custom("crossorigin", ""))),
			h.Link(a.Attrs(a.Rel("stylesheet"), a.Href("https://fonts.googleapis.com/css2?family=DM+Serif+Display&family=JetBrains+Mono:wght@400;600&family=Outfit:wght@300;400;500;600;700&display=swap"))),
			h.Link(a.Attrs(a.Rel("icon"), a.Type("image/svg+xml"), a.Href("data:image/svg+xml,<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'><text y='.9em' font-size='90'>🍵</text></svg>"))),
			h.Style(a.Attrs(), h.Text(customStyles())),
		),
		h.Body(
			a.Attrs(a.Class("min-h-screen bg-stone-100"), a.Custom("style", "font-family: 'Outfit', sans-serif;")),
			// Decorative background pattern
			h.Div(a.Attrs(a.Class("fixed inset-0 -z-10 overflow-hidden")),
				h.Div(a.Attrs(a.Class("absolute -top-40 -right-40 w-96 h-96 bg-amber-200 rounded-full mix-blend-multiply filter blur-3xl opacity-30 animate-blob"))),
				h.Div(a.Attrs(a.Class("absolute top-40 -left-20 w-72 h-72 bg-emerald-200 rounded-full mix-blend-multiply filter blur-3xl opacity-30 animate-blob animation-delay-2000"))),
				h.Div(a.Attrs(a.Class("absolute -bottom-20 left-1/2 w-80 h-80 bg-rose-200 rounded-full mix-blend-multiply filter blur-3xl opacity-30 animate-blob animation-delay-4000"))),
			),
			// Header/Nav
			h.Header(
				a.Attrs(a.Class("sticky top-0 z-50 backdrop-blur-md bg-stone-100/80 border-b-2 border-stone-900")),
				h.Div(
					a.Attrs(a.Class("max-w-7xl mx-auto px-6 py-4")),
					h.Div(
						a.Attrs(a.Class("flex items-center justify-between")),
						// Logo
						h.A(
							a.Attrs(a.Href("/"), a.Class("group flex items-center gap-3")),
							h.Span(a.Attrs(a.Class("text-4xl transition-transform group-hover:rotate-12")), h.Text("🍵")),
							h.Div(a.Attrs(),
								h.Span(a.Attrs(a.Class("block text-2xl font-bold tracking-tight text-stone-900"), a.Custom("style", "font-family: 'DM Serif Display', serif;")), h.Text("Gotea")),
								h.Span(a.Attrs(a.Class("block text-xs font-medium text-stone-500 tracking-widest uppercase")), h.Text("Elm Architecture for Go")),
							),
						),
						// Nav pills
						h.Nav(
							a.Attrs(a.Class("hidden lg:flex items-center gap-1")),
							navPill("/", "Home", m.Router.Route == "/"),
							navPill("/memory", "Memory", m.Router.Route == "/memory"),
							navPill("/form", "Forms", m.Router.Route == "/form"),
							navPill("/components", "Components", m.Router.Route == "/components"),
							navPill("/routing", "Routing", m.Router.Route == "/routing"),
							navPill("/animation", "Animation", m.Router.Route == "/animation"),
							navPill("/chat", "Chat", m.Router.Route == "/chat"),
							navPill("/blocktrader", "Blocktrader", m.Router.Route == "/blocktrader"),
							navPill("/pixelcanvas", "Pixels", m.Router.Route == "/pixelcanvas"),
						),
					),
				),
			),
			// Main content
			h.Main(
				a.Attrs(a.Class("max-w-7xl mx-auto px-6 py-12")),
				h.Div(
					a.Attrs(a.Class("bg-white border-2 border-stone-900 rounded-2xl shadow-brutal p-8 md:p-12")),
					h.Div(a.Attrs(a.Id("view")), h.UnsafeRaw(string(m.RenderRoute(m)))),
				),
			),
			// Footer
			h.Div(
				a.Attrs(a.Class("mt-auto py-8"), a.Custom("role", "contentinfo")),
				h.Div(
					a.Attrs(a.Class("max-w-7xl mx-auto px-6")),
					h.Div(
						a.Attrs(a.Class("flex flex-col md:flex-row items-center justify-between gap-4 text-stone-500")),
						h.Div(
							a.Attrs(a.Class("flex items-center gap-2 text-sm")),
							h.Span(a.Attrs(a.Class("font-mono text-xs bg-stone-200 px-2 py-1 rounded")), h.Text(m.sessionID.String()[:8]+"...")),
							h.Span(a.Attrs(a.Class("text-stone-300")), h.Text("•")),
							h.Span(a.Attrs(), h.Text(m.Router.Route)),
						),
						h.Div(
							a.Attrs(a.Class("text-sm")),
							h.Text("Built with "),
							h.Span(a.Attrs(a.Class("font-semibold text-stone-700")), h.Text("Gotea")),
							h.Text(" — Server-driven UI with The Elm Architecture"),
						),
					),
				),
			),
			h.Script(a.Attrs(a.Src("static/main.js"))),
		),
	)

	return el.Bytes()
}

func customStyles() string {
	return `
		.shadow-brutal { box-shadow: 6px 6px 0px 0px #1c1917; }
		.shadow-brutal-sm { box-shadow: 3px 3px 0px 0px #1c1917; }
		.shadow-brutal-hover:hover { box-shadow: 8px 8px 0px 0px #1c1917; transform: translate(-2px, -2px); }
		@keyframes blob { 0%, 100% { transform: translate(0, 0) scale(1); } 33% { transform: translate(30px, -50px) scale(1.1); } 66% { transform: translate(-20px, 20px) scale(0.9); } }
		.animate-blob { animation: blob 7s infinite; }
		.animation-delay-2000 { animation-delay: 2s; }
		.animation-delay-4000 { animation-delay: 4s; }
		.font-display { font-family: 'DM Serif Display', serif; }
		.font-mono { font-family: 'JetBrains Mono', monospace; }
		input, select, textarea { border: 2px solid #1c1917 !important; }
		input:focus, select:focus, textarea:focus { outline: none !important; box-shadow: 4px 4px 0px 0px #1c1917 !important; }
	`
}

func navPill(href, text string, active bool) h.Element {
	if active {
		return h.A(
			a.Attrs(a.Href(href), a.Class("px-4 py-2 text-sm font-semibold bg-stone-900 text-white rounded-full")),
			h.Text(text),
		)
	}
	return h.A(
		a.Attrs(a.Href(href), a.Class("px-4 py-2 text-sm font-medium text-stone-600 hover:text-stone-900 hover:bg-stone-200 rounded-full transition-colors")),
		h.Text(text),
	)
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
	case "/pixelcanvas":
		return m.PixelCanvas.render()
	default:
		return renderHome(m.Counter)
	}
}

// Placeholder functions for individual template rendering
func renderHome(counter int) h.Element {
	return h.Div(
		a.Attrs(a.Class("space-y-10")),
		// Hero section
		h.Div(
			a.Attrs(a.Class("text-center space-y-4 pb-8")),
			h.H1(
				a.Attrs(a.Class("text-5xl md:text-6xl font-bold text-stone-900 tracking-tight"), a.Custom("style", "font-family: 'DM Serif Display', serif;")),
				h.Text("Welcome to Gotea"),
			),
			h.P(
				a.Attrs(a.Class("text-xl text-stone-600 max-w-2xl mx-auto leading-relaxed")),
				h.Text("The Elm Architecture for Go. Build reactive, server-rendered SPAs with state on the server and WebSocket-powered updates."),
			),
			// Feature badges
			h.Div(
				a.Attrs(a.Class("flex flex-wrap justify-center gap-3 pt-4")),
				featureBadge("🔄", "Real-time Updates", "emerald"),
				featureBadge("🖥️", "Server-Side State", "amber"),
				featureBadge("⚡", "WebSocket Powered", "rose"),
				featureBadge("🎯", "Zero JavaScript*", "sky"),
			),
		),
		// Counter demo
		h.Div(
			a.Attrs(a.Class("bg-gradient-to-br from-stone-50 to-stone-100 border-2 border-stone-900 rounded-2xl p-8 shadow-brutal-sm")),
			h.Div(
				a.Attrs(a.Class("flex flex-col md:flex-row md:items-center md:justify-between gap-6")),
				h.Div(
					a.Attrs(a.Class("space-y-2")),
					h.H2(
						a.Attrs(a.Class("text-2xl font-bold text-stone-900")),
						h.Text("The Classic Counter"),
					),
					h.P(
						a.Attrs(a.Class("text-stone-600")),
						h.Text("Every click sends a message to the server, which updates the state and re-renders."),
					),
				),
				renderCounter(counter),
			),
		),
		renderExplanatoryNote(
			"How this works",
			`
			<p class="mb-3">This is the classic "Counter" example, the "Hello World" of the Elm Architecture.</p>
			<ul class="list-disc pl-5 space-y-2">
				<li><strong class="text-stone-900">State:</strong> A simple integer <code class="bg-stone-200 px-1.5 py-0.5 rounded text-xs font-mono">Counter</code> in the Model.</li>
				<li><strong class="text-stone-900">Messages:</strong> <code class="bg-stone-200 px-1.5 py-0.5 rounded text-xs font-mono">INCREMENT</code> and <code class="bg-stone-200 px-1.5 py-0.5 rounded text-xs font-mono">DECREMENT</code> messages are sent when buttons are clicked.</li>
				<li><strong class="text-stone-900">Update:</strong> The update function modifies the Counter based on the message.</li>
				<li><strong class="text-stone-900">Render:</strong> The view function re-renders the current count and buttons.</li>
			</ul>
			`,
		),
		// Demo grid
		h.Div(
			a.Attrs(a.Class("space-y-6")),
			h.H3(a.Attrs(a.Class("text-lg font-semibold text-stone-900")), h.Text("Explore the demos →")),
			h.Div(
				a.Attrs(a.Class("grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4")),
				demoCard("/memory", "🎴", "Memory Game", "Card matching with delayed messages and leaderboards"),
				demoCard("/form", "📝", "Forms", "Two-way binding and real-time form state"),
				demoCard("/chat", "💬", "Chat Room", "Real-time broadcasting to all connected clients"),
				demoCard("/animation", "🏀", "Animation", "Server-driven 30fps animations"),
				demoCard("/pixelcanvas", "🎨", "Pixel Canvas", "Collaborative drawing with shared state"),
				demoCard("/blocktrader", "📈", "Blocktrader", "Complex game with persistence"),
			),
		),
		// Divider
		h.Div(a.Attrs(a.Class("border-t-2 border-dashed border-stone-300 my-4"))),
		// Error handling demo
		h.Div(
			a.Attrs(a.Class("bg-rose-50 border-2 border-rose-300 rounded-xl p-6")),
			h.H3(a.Attrs(a.Class("text-lg font-semibold text-rose-900 mb-2")), h.Text("🚨 Error Handling Demo")),
			h.P(a.Attrs(a.Class("text-rose-700 text-sm mb-4")), h.Text("Click to trigger a panic and see how Gotea handles server-side errors gracefully.")),
			renderCrash(),
		),
	)
}

func featureBadge(icon, label, color string) h.Element {
	colorClasses := map[string]string{
		"emerald": "bg-emerald-100 text-emerald-800 border-emerald-300",
		"amber":   "bg-amber-100 text-amber-800 border-amber-300",
		"rose":    "bg-rose-100 text-rose-800 border-rose-300",
		"sky":     "bg-sky-100 text-sky-800 border-sky-300",
	}
	return h.Span(
		a.Attrs(a.Class(fmt.Sprintf("inline-flex items-center gap-1.5 px-3 py-1.5 rounded-full text-sm font-medium border %s", colorClasses[color]))),
		h.Span(a.Attrs(), h.Text(icon)),
		h.Text(label),
	)
}

func demoCard(href, icon, title, description string) h.Element {
	return h.A(
		a.Attrs(
			a.Href(href),
			a.Class("group block bg-white border-2 border-stone-900 rounded-xl p-5 transition-all hover:shadow-brutal hover:-translate-x-1 hover:-translate-y-1"),
		),
		h.Div(
			a.Attrs(a.Class("flex items-start gap-4")),
			h.Span(a.Attrs(a.Class("text-3xl")), h.Text(icon)),
			h.Div(
				a.Attrs(),
				h.H4(a.Attrs(a.Class("font-semibold text-stone-900 group-hover:text-emerald-700 transition-colors")), h.Text(title)),
				h.P(a.Attrs(a.Class("text-sm text-stone-500 mt-1")), h.Text(description)),
			),
		),
	)
}

func (m *Model) renderRouting() h.Element {
	return h.Div(
		a.Attrs(a.Class("space-y-8")),
		// Header
		h.Div(
			a.Attrs(a.Class("text-center space-y-2")),
			h.H1(a.Attrs(a.Class("text-4xl font-bold text-stone-900"), a.Custom("style", "font-family: 'DM Serif Display', serif;")), h.Text("🧭 Client-Side Routing")),
			h.P(a.Attrs(a.Class("text-stone-600")), h.Text("SPA-style navigation without page reloads, powered by WebSocket messages.")),
		),

		// Current route display
		h.Div(
			a.Attrs(a.Class("bg-gradient-to-r from-sky-50 to-blue-50 p-5 rounded-xl border-2 border-stone-900 shadow-brutal-sm")),
			h.Div(
				a.Attrs(a.Class("flex items-center justify-between")),
				h.Div(
					a.Attrs(),
					h.Span(a.Attrs(a.Class("text-sm font-medium text-stone-500")), h.Text("Current Route")),
					h.Div(a.Attrs(a.Class("text-2xl font-bold text-stone-900"), a.Custom("style", "font-family: 'JetBrains Mono', monospace;")), h.Text(m.GetRoute())),
				),
				h.Div(
					a.Attrs(a.Class("text-4xl")),
					h.Text("📍"),
				),
			),
		),

		// Link examples
		h.Div(
			a.Attrs(a.Class("grid grid-cols-1 md:grid-cols-2 gap-4")),
			// Internal link
			h.Div(
				a.Attrs(a.Class("bg-white p-6 rounded-xl border-2 border-stone-900 shadow-brutal-sm")),
				h.H3(a.Attrs(a.Class("text-sm font-semibold text-stone-500 uppercase tracking-wide mb-3")), h.Text("Internal Link")),
				h.A(
					a.Attrs(
						a.Class("inline-flex items-center gap-2 text-emerald-700 hover:text-emerald-900 font-semibold"),
						a.Href("/memory"),
					),
					h.Span(a.Attrs(), h.Text("→")),
					h.Text("Go to Memory Game"),
				),
				h.P(a.Attrs(a.Class("text-sm text-stone-500 mt-2")), h.Text("Handled by Gotea (no page reload)")),
			),
			// External link
			h.Div(
				a.Attrs(a.Class("bg-white p-6 rounded-xl border-2 border-stone-900 shadow-brutal-sm")),
				h.H3(a.Attrs(a.Class("text-sm font-semibold text-stone-500 uppercase tracking-wide mb-3")), h.Text("External Link")),
				h.A(
					a.Attrs(
						a.Class("inline-flex items-center gap-2 text-blue-700 hover:text-blue-900 font-semibold"),
						a.Target("_blank"),
						a.Href("https://duckduckgo.com"),
					),
					h.Text("Search DuckDuckGo"),
					h.Span(a.Attrs(), h.Text("↗")),
				),
				h.P(a.Attrs(a.Class("text-sm text-stone-500 mt-2")), h.Text("Normal browser navigation")),
			),
		),

		// Route parameters
		h.Div(
			a.Attrs(a.Class("bg-white p-6 rounded-xl border-2 border-stone-900 shadow-brutal-sm")),
			h.H3(a.Attrs(a.Class("text-lg font-bold text-stone-900 mb-4")), h.Text("Query Parameters")),
			h.Div(
				a.Attrs(a.Class("flex flex-wrap gap-2 mb-4")),
				routeParamLink("/routing?myparam=1", "myparam=1"),
				routeParamLink("/routing?myparam=2", "myparam=2"),
				routeParamLink("/routing?myparam=3", "myparam=3"),
			),
			h.Div(
				a.Attrs(a.Class("bg-stone-100 p-4 rounded-lg flex items-center gap-3")),
				h.Span(a.Attrs(a.Class("text-sm text-stone-500")), h.Text("myparam =")),
				h.Span(a.Attrs(a.Class("text-lg font-bold text-stone-900"), a.Custom("style", "font-family: 'JetBrains Mono', monospace;")), h.Text(func() string {
					val := m.RouteParam("myparam")
					if val == "" {
						return "(empty)"
					}
					return val
				}())),
			),
		),

		// Route data
		h.Div(
			a.Attrs(a.Class("bg-gradient-to-r from-amber-50 to-yellow-50 p-5 rounded-xl border-2 border-amber-300")),
			h.Div(
				a.Attrs(a.Class("flex items-start gap-3")),
				h.Span(a.Attrs(a.Class("text-2xl")), h.Text("🎲")),
				h.Div(
					a.Attrs(),
					h.H4(a.Attrs(a.Class("font-semibold text-amber-900")), h.Text("Route-Specific Data")),
					h.P(a.Attrs(a.Class("text-sm text-amber-800")), h.Text(fmt.Sprintf("Random number generated in OnRouteChange hook: %d", m.RouteData))),
				),
			),
		),

		// Explanatory note
		renderExplanatoryNote(
			"Routing in GoTea",
			`
			<p class="mb-3">GoTea provides a simple but flexible routing system.</p>
			<ul class="list-disc pl-5 space-y-2">
				<li><strong class="text-stone-900">Navigation:</strong> Links trigger a <code class="bg-stone-200 px-1.5 py-0.5 rounded text-xs font-mono">CHANGE_ROUTE</code> message.</li>
				<li><strong class="text-stone-900">Router:</strong> The <code class="bg-stone-200 px-1.5 py-0.5 rounded text-xs font-mono">Router</code> struct (embedded in Model) handles the current route state.</li>
				<li><strong class="text-stone-900">Hooks:</strong> The <code class="bg-stone-200 px-1.5 py-0.5 rounded text-xs font-mono">OnRouteChange</code> hook allows you to execute logic when the route changes.</li>
				<li><strong class="text-stone-900">Rendering:</strong> The main Render function switches on the route to decide what to render.</li>
			</ul>
			`,
		),
	)
}

func routeParamLink(href, label string) h.Element {
	return h.A(
		a.Attrs(
			a.Href(href),
			a.Class("px-4 py-2 bg-stone-900 text-white font-mono text-sm rounded-lg hover:bg-stone-700 transition-colors"),
		),
		h.Text(label),
	)
}

func renderExplanatoryNote(title, content string) h.Element {
	return h.Details(
		a.Attrs(a.Class("bg-amber-50 border-2 border-stone-900 rounded-xl mb-8 shadow-brutal-sm overflow-hidden")),
		h.Summary(
			a.Attrs(a.Class("cursor-pointer px-6 py-4 font-semibold text-stone-900 hover:bg-amber-100 transition-colors flex items-center gap-2")),
			h.Span(a.Attrs(a.Class("text-lg")), h.Text("📖")),
			h.Text(title),
		),
		h.Div(
			a.Attrs(a.Class("px-6 pb-5 text-stone-700 text-sm leading-relaxed border-t-2 border-stone-200 pt-4")),
			h.UnsafeRaw(content),
		),
	)
}

var app = gt.NewApp(&Model{})

func main() {
	app.Start(8080, "static")
}
