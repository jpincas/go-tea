package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"

	gotea "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/components/tagselector"
)

func renderView(state gotea.State) []byte {
	template := state.RouteTemplate("html")

	buf := bytes.Buffer{}
	err := Templates.ExecuteTemplate(&buf, template, state)
	if err != nil {
		return []byte(fmt.Sprintf("Executing template %s. Error: %v", template, err))
	}

	return buf.Bytes()
}

var Templates *template.Template

func parseTemplates() {
	auxFuncs := template.FuncMap{
		"prettyPrint": func(s interface{}) template.HTML {
			res, _ := json.MarshalIndent(s, "<br />", "<span style='margin-left:15px'></span>")
			return template.HTML(string(res))
		},
	}

	funcMap := gotea.SquashFuncMaps(
		auxFuncs,
		gotea.TemplateHelpers,
	)

	Templates = template.Must(template.New("main").Funcs(funcMap).ParseGlob("templates/*.html"))

	template.Must(Templates.ParseGlob("../components/*/*.html"))
}

type Model struct {
	*gotea.Router
	MemoryGame   MemoryGame
	NameSelector tagselector.Model
	TeamSelector tagselector.Model
	Form         Form
}

func main() {
	// Register the function that returns a new session
	gotea.App.NewSession = func() gotea.Session {
		return gotea.Session{
			State: Model{
				Router: &gotea.Router{
					Route: "/home",
				},
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
			},
		}
	}

	// Register all the message maps
	gotea.App.Messages.
		MergeMap(memoryGameMessages).
		MergeMap(formMessages).
		MergeMap(nameSelector.UniqueMsgMap(nameSelectorMessages)).
		MergeMap(teamSelector.UniqueMsgMap(teamSelectorMessages))

	// Parse templates
	parseTemplates()

	// Register the main render function
	gotea.App.RenderView = renderView

	// Start the app
	fmt.Println("starting server")
	gotea.App.Start("dist", 8080)
}
