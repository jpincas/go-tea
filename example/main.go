package main

import (
	"encoding/json"
	"fmt"
	"html/template"

	gotea "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/components/tagselector"
)

type Model struct {
	*gotea.Router
	MemoryGame   MemoryGame
	NameSelector tagselector.Model
	TeamSelector tagselector.Model
	Form         Form
}

func initModel() gotea.Session {
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

func main() {
	app := gotea.NewApp(
		initModel,
		memoryGameMessages,
		formMessages,
		nameSelector.UniqueMsgMap(nameSelectorMessages),
		teamSelector.UniqueMsgMap(teamSelectorMessages),
	)

	// Start the app
	fmt.Println("starting server")

	// Define a custom template func map
	templateFuncs := template.FuncMap{
		"prettyPrint": func(s interface{}) template.HTML {
			res, _ := json.MarshalIndent(s, "<br />", "<span style='margin-left:15px'></span>")
			return template.HTML(string(res))
		},
	}

	app.Start("dist", 8080, templateFuncs, "templates/*.html", "../components/*/*.html")
}
