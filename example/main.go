package main

import (
	"bytes"
	"html/template"

	gotea "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/components/tagselector"
)

func renderView(state gotea.State) []byte {
	buf := bytes.Buffer{}
	Templates.ExecuteTemplate(&buf, "main.html", state)
	return buf.Bytes()
}

var Templates *template.Template

func parseTemplates() {
	Templates = template.Must(template.New("main").Funcs(gotea.TemplateHelpers).ParseGlob("templates/*.html"))

	template.Must(Templates.ParseGlob("../components/*/*.html"))
}

type Model struct {
	Route       string
	MemoryGame  MemoryGame
	TagSelector tagselector.Model
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
				TagSelector: tagselector.Model{
					AvailableTags: []string{"tag1", "tag2", "tag3"},
				},
			},
		}
	}

	gotea.App.Messages.
		MergeMap(memoryGameMessages).
		MergeMap(tagselector.Messages)

	// Parse templates
	parseTemplates()

	// Register the main render function
	gotea.App.RenderView = renderView

	// Start the app
	gotea.App.Start("dist", 8080)
}
