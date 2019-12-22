package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	gt "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/example/tagselector"
)

type Model struct {
	*gt.Router
	MemoryGame   MemoryGame
	NameSelector tagselector.Model
	TeamSelector tagselector.Model
	Form         Form
}

func (m Model) Init(_ *http.Request) gt.State {
	return Model{
		Router: &gt.Router{},
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

func (m Model) RouteTemplate() string {
	currentRoute := m.GetRoute()
	routeTemplate := strings.Replace(currentRoute, "/", "_", -1)

	if routeTemplate == "" {
		return "home"
	}

	return routeTemplate
}

func (m Model) RouteUpdateHook() gt.State {
	fmt.Println("Hey! Changing route here")
	return m
}

func main() {
	app := gt.NewApp(
		gt.DefaultAppConfig,
		Model{},
		nil,
	)

	app.Templates.AddGlobal("prettyPrint", func(s interface{}) string {
		res, _ := json.MarshalIndent(s, "<br />", "<span style='margin-left:15px'></span>")
		return string(res)
	})

	app.Start()
}
