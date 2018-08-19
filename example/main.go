package main

import (
	"bytes"
	"html/template"

	gotea "github.com/jpincas/go-tea"
)

func renderView(state gotea.State) []byte {
	buf := bytes.Buffer{}
	Templates.ExecuteTemplate(&buf, "main.html", state)
	return buf.Bytes()
}

var Templates *template.Template

func parseTemplates() {

	Templates = template.Must(template.New("main").Funcs(gotea.TemplateHelpers).ParseGlob("templates/*.html"))
}

type Model struct {
	Route      string
	MemoryGame MemoryGame
}

func (m Model) SetRoute(newRoute string) gotea.State {
	m.Route = newRoute
	return m
}

func (m Model) GetRoute() string {
	return m.Route
}

func main() {
	// Register the function that returns a new session
	gotea.App.NewSession = func() gotea.Session {
		return gotea.Session{
			State: Model{
				Route: "/home",
				MemoryGame: MemoryGame{
					Deck:              NewDeck(4),
					TurnsTaken:        0,
					LastAttemptedCard: 5, //hack
					Score:             0,
				},
			},
		}
	}

	// gotea.App.Messages
	gotea.App.Messages.
		MergeMap(memoryGameMessages)

	// Parse templates
	parseTemplates()

	// Register the main render function
	gotea.App.RenderView = renderView

	// Start the app
	gotea.App.Start("dist", 8080)
}
