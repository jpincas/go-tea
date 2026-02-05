package main

import (
	"fmt"

	"github.com/google/uuid"
	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/go-tea/attributes"
	"github.com/jpincas/go-tea/css"
	h "github.com/jpincas/go-tea/html"
)

// Model is the application state. It MUST embed gt.Router for routing to work.
type Model struct {
	gt.Router
	sessionID uuid.UUID

	// Feature state
	Counter  int
	Form     FormState
	Selector TagSelector
}

// model is a helper for type assertion in handlers
func model(s gt.State) *Model {
	return s.(*Model)
}

// Init creates a new session state. Called once per WebSocket connection.
func (m *Model) Init(sid uuid.UUID) gt.State {
	model := &Model{
		sessionID: sid,
		Counter:   0,
		Form: FormState{
			Options: []string{"Option A", "Option B", "Option C"},
		},
		Selector: NewTagSelector("my-selector", []string{
			"Go", "JavaScript", "Python", "Rust", "TypeScript", "Ruby",
		}),
	}

	// Register routes - each path maps to a render function
	model.Register("/", func(s gt.State) []byte {
		return renderHome(model).Bytes()
	})
	model.Register("/form", func(s gt.State) []byte {
		return renderFormPage(model).Bytes()
	})
	model.Register("/component", func(s gt.State) []byte {
		return renderComponentPage(model).Bytes()
	})

	return model
}

// Update returns all message handlers. Use gt.MergeMaps for multiple sources.
func (m *Model) Update() gt.MessageMap {
	return gt.MergeMaps(
		counterMessages,
		formMessages,
		m.Selector.UniqueMsgMap(tagSelectorMessages), // Namespaced component messages
	)
}

// OnRouteChange is called whenever the route changes. Use for route-specific setup.
func (m *Model) OnRouteChange(path string) {
	fmt.Printf("Route changed to: %s\n", path)
}

// Render returns the full HTML page. Called after every state mutation.
func (m *Model) Render() []byte {
	return h.Html(a.Attrs(a.Lang("en")),
		h.Head(a.Attrs(),
			h.Meta(a.Attrs(a.Charset("UTF-8"))),
			h.Meta(a.Attrs(
				a.Name("viewport"),
				a.Content("width=device-width, initial-scale=1.0"))),
			h.Title(a.Attrs(), h.Text("Gotea Starter")),
			// Tailwind via CDN for simplicity
			h.Script(a.Attrs(a.Src("https://cdn.tailwindcss.com"))),
			h.Link(a.Attrs(
				a.Rel("stylesheet"),
				a.Href("static/styles.css")))),
		h.Body(a.Attrs(
			a.Class("min-h-screen bg-gray-100"),
			a.Style(css.FontFamily("system-ui, sans-serif"))),
			// Navigation
			renderNav(m.Router.Route),
			// Main content - renders the current route
			h.Main(a.Attrs(a.Class("max-w-4xl mx-auto p-6")),
				h.Div(a.Attrs(a.Id("view")),
					h.UnsafeRaw(string(m.RenderRoute(m))))),
			// Gotea client script
			h.Script(a.Attrs(a.Src("static/main.js"))))).Bytes()
}

// RenderError is called when a handler returns an error or panics.
func (m *Model) RenderError(err error) []byte {
	return h.Div(a.Attrs(a.Class("p-4 bg-red-100 text-red-800 rounded")),
		h.Text(fmt.Sprintf("Error: %s", err.Error()))).Bytes()
}

// renderNav creates the navigation bar
func renderNav(currentRoute string) h.Element {
	navLink := func(href, text string) h.Element {
		active := currentRoute == href
		classes := "px-4 py-2 rounded transition-colors"
		if active {
			classes += " bg-blue-600 text-white"
		} else {
			classes += " text-gray-700 hover:bg-gray-200"
		}
		return h.A(a.Attrs(a.Href(href), a.Class(classes)), h.Text(text))
	}

	return h.Nav(a.Attrs(a.Class("bg-white shadow-sm border-b")),
		h.Div(a.Attrs(a.Class("max-w-4xl mx-auto px-6 py-4 flex items-center gap-4")),
			h.Span(a.Attrs(a.Class("text-xl font-bold text-gray-900")), h.Text("Gotea Starter")),
			h.Div(a.Attrs(a.Class("flex gap-2")),
				navLink("/", "Home"),
				navLink("/form", "Form"),
				navLink("/component", "Component"))))
}

// Application instance - accessible for broadcasting if needed
var app = gt.NewApp(&Model{})

func main() {
	fmt.Println("Starting Gotea app at http://localhost:8080")
	app.Start(8080, "static")
}
