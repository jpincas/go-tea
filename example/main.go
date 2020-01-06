package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"

	"github.com/CloudyKit/jet"
	gt "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/example/tagselector"
)

const (
	animationBackgroundSize = 500
	animationBallSize       = 20
)

type Model struct {
	gt.BaseModel
	TemplateName string
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
		crashMessages,
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

	m.TemplateName = template
}

// Rendering

var templates *jet.Set

func (m Model) Render(w io.Writer) error {
	t, err := templates.GetTemplate(m.TemplateName)
	if err != nil {
		return err
	}

	vars := make(jet.VarMap)
	return t.Execute(w, vars, m)
}

func (m Model) RenderError(w io.Writer, errorToRender error) {
	t, _ := templates.GetTemplate("error")
	vars := make(jet.VarMap)
	vars.Set("Msg", errorToRender.Error())
	t.Execute(w, vars, m)
}

func main() {
	templates = jet.NewHTMLSet("templates")
	for funcName, f := range gt.TemplateFuncs {
		templates.AddGlobal(funcName, f)
	}

	templates.AddGlobal("prettyPrint", func(s interface{}) string {
		res, _ := json.MarshalIndent(s, "<br />", "<span style='margin-left:15px'></span>")
		return string(res)
	})

	app := gt.NewApp(
		gt.DefaultAppConfig,
		&Model{},
		nil,
	)

	log.Printf("Starting application server on %v\n", app.Config.Port)
	http.ListenAndServe(fmt.Sprintf(":%v", app.Config.Port), app.Router)
}
