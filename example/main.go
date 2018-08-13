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

func concat(s1, s2 string) string {
	return s1 + s2
}

func parseTemplates() {
	funcMap := template.FuncMap{
		"concat": concat,
	}

	Templates = template.Must(template.New("main").Funcs(funcMap).ParseGlob("templates/*.html"))
}

func main() {
	// Register the function that returns a new session
	gotea.App.NewSession = func() gotea.Session {
		return gotea.Session{
			State: Model{
				Deck:              NewDeck(4),
				TurnsTaken:        0,
				LastAttemptedCard: 5, //hack
				Score:             0,
			},
		}
	}

	// gotea.App.Messages
	gotea.App.Messages = gotea.MessageMap{
		"FLIP_CARD":      FlipCard,
		"REMOVE_MATCHES": RemoveMatches,
		"FLIP_ALL_BACK":  FlipAllBack,
	}

	// Parse templates
	parseTemplates()

	// Register the main render function
	gotea.App.RenderView = renderView

	// Start the app
	gotea.App.Start("dist", 8080)
}
