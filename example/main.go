package main

import (
	"encoding/json"
	"net/http"
	"net/url"

	gt "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/example/tagselector"
)

type Model struct {
	*gt.BaseModel
	MemoryGame   MemoryGame
	NameSelector tagselector.Model
	TeamSelector tagselector.Model
	Form         Form
}

func (m Model) Init(_ *http.Request) gt.State {
	return &Model{
		BaseModel: &gt.BaseModel{},
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

func (m Model) Update() gt.MessageMap {
	return gt.MergeMaps(
		memoryGameMessages,
		formMessages,
		nameSelector.UniqueMsgMap(nameSelectorMessages),
		teamSelector.UniqueMsgMap(teamSelectorMessages),
	)
}

func (m *Model) Mux(path *url.URL) string {
	// Ridiculously simle routing model -
	// the name of the template is the name of the path
	if path.Path == "/" {
		return "home"
	}

	return path.Path
}

func main() {
	app := gt.NewApp(
		gt.DefaultAppConfig,
		&Model{},
		nil,
	)

	app.Templates.AddGlobal("prettyPrint", func(s interface{}) string {
		res, _ := json.MarshalIndent(s, "<br />", "<span style='margin-left:15px'></span>")
		return string(res)
	})

	app.Start()
}
