package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"

	gt "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/example/tagselector"
)

const (
	animationBackgroundSize = 500
	animationBallSize       = 20
)

type Model struct {
	gt.BaseModel
	MemoryGame   MemoryGame
	NameSelector tagselector.Model
	TeamSelector tagselector.Model
	Form         Form
	RouteData    int
	Animation    Animation
	Counter      int
}

func model(s gt.State) *Model {
	return s.(*Model)
}

func (m Model) Init(_ *http.Request) gt.State {
	return &Model{
		BaseModel: gt.BaseModel{},
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
	}
}

func (m Model) Update() gt.MessageMap {
	return gt.MergeMaps(
		memoryGameMessages,
		formMessages,
		nameSelector.UniqueMsgMap(nameSelectorMessages),
		teamSelector.UniqueMsgMap(teamSelectorMessages),
		animationMessages,
		counterMessages,
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

	m.SetNewTemplate(template)
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
