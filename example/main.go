package main

import (
	"encoding/json"

	gotea "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/example/tagselector"
)

type Model struct {
	*gotea.Router
	MemoryGame   MemoryGame
	NameSelector tagselector.Model
	TeamSelector tagselector.Model
	Form         Form
}

func initModel() gotea.State {
	return Model{
		Router: &gotea.Router{},
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
	}
}

func main() {
	app := gotea.NewApp(
		gotea.DefaultAppConfig,
		initModel,
		memoryGameMessages,
		formMessages,
		nameSelector.UniqueMsgMap(nameSelectorMessages),
		teamSelector.UniqueMsgMap(teamSelectorMessages),
	)

	app.Templates.AddGlobal("prettyPrint", func(s interface{}) string {
		res, _ := json.MarshalIndent(s, "<br />", "<span style='margin-left:15px'></span>")
		return string(res)
	})

	app.Start()
}
