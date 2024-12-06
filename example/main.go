package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	gt "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/example/tagselector"
	a "github.com/jpincas/htmlfunc/attributes"
	css "github.com/jpincas/htmlfunc/css"
	h "github.com/jpincas/htmlfunc/html"
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

func (m Model) Render(w io.Writer) error {
	el := h.Html(
		a.Attrs(a.Lang("en")),
		h.Head(
			a.Attrs(),
			h.Meta(a.Attrs(a.Charset("UTF-8"))),
			h.Meta(a.Attrs(a.Name("viewport"), a.Content("width=device-width, initial-scale=1.0"))),
			h.Meta(a.Attrs(a.HttpEquiv("X-UA-Compatible"), a.Content("ie=edge"))),
			h.Title(a.Attrs(), h.Text(m.TemplateName)),
			h.Link(a.Attrs(a.Rel("stylesheet"), a.Href("static/styles.css"))),
			h.Link(a.Attrs(a.Rel("icon"), a.Type("image/png"), a.Href("data:image/png;base64,iVBORw0KGgo="))),
		),
		h.Body(
			a.Attrs(),
			h.Div(
				a.Attrs(a.Class("navbar")),
				h.Ul(
					a.Attrs(),
					h.Li(a.Attrs(), h.A(a.Attrs(a.Href("/")), h.Text("Home"))),
					h.Li(a.Attrs(), h.A(a.Attrs(a.Href("/memory")), h.Text("Memory Game"))),
					h.Li(a.Attrs(), h.A(a.Attrs(a.Href("/form")), h.Text("Form"))),
					h.Li(a.Attrs(), h.A(a.Attrs(a.Href("/components")), h.Text("Components"))),
					h.Li(a.Attrs(), h.A(a.Attrs(a.Href("/routing")), h.Text("Routing"))),
					h.Li(a.Attrs(), h.A(a.Attrs(a.Href("/animation")), h.Text("Animation"))),
					h.Li(a.Attrs(), h.Text(fmt.Sprintf("Current Route: %s", m.BaseModel.Router.Route))),
				),
			),
			h.Div(a.Attrs(a.Id("view")), m.RenderView()),
			h.Script(a.Attrs(a.Src("static/main.js"))),
		),
	)

	return el.Write(w)
}

func (m Model) RenderError(w io.Writer, errorToRender error) {
	el := h.Html(
		a.Attrs(a.Lang("en")),
		h.Head(
			a.Attrs(),
			h.Meta(a.Attrs(a.Charset("UTF-8"))),
			h.Meta(a.Attrs(a.Name("viewport"), a.Content("width=device-width, initial-scale=1.0"))),
			h.Meta(a.Attrs(a.HttpEquiv("X-UA-Compatible"), a.Content("ie=edge"))),
			h.Title(a.Attrs(), h.Text("Error")),
			h.Link(a.Attrs(a.Rel("stylesheet"), a.Href("static/styles.css"))),
		),
		h.Body(
			a.Attrs(),
			h.Div(
				a.Attrs(a.Class("error-container")),
				h.H1(a.Attrs(), h.Text("An Error Occurred")),
				h.P(a.Attrs(), h.Text(errorToRender.Error())),
				h.A(a.Attrs(a.Href("/")), h.Text("Go back to Home")),
			),
		),
	)

	el.WriteDoc(w)
}

func (m Model) RenderView() h.Element {
	switch m.TemplateName {
	case "/home":
		return m.renderHome()
	case "/memory":
		return m.renderMemoryGame()
	case "/form":
		return m.renderForm()
	case "/components":
		return m.renderComponents()
	case "/routing":
		return m.renderRouting()
	case "/animation":
		return m.renderAnimation()
	default:
		return m.renderHome()
	}
}

// Placeholder functions for individual template rendering
func (m Model) renderHome() h.Element {
	return h.Div(
		a.Attrs(),
		h.H2(a.Attrs(), h.Text("Gotea Demo Site")),
		h.P(a.Attrs(), h.Text("Choose an example from the menu to get started, or have a play with the classic 'counter' example below:")),
		m.renderCounter(m.Counter),
		h.Hr(a.Attrs()),
		m.renderCrash(),
	)
}

func (m Model) renderCounter(counter int) h.Element {
	return h.Div(
		a.Attrs(a.Id("counter")),
		h.Button(
			a.Attrs(a.Class("counter-increment"), a.OnClick(gt.SendMessage("INCREMENT_COUNTER", -1))),
			h.Text("Down"),
		),
		h.Div(
			a.Attrs(a.Class("counter-readout")),
			h.Text(fmt.Sprintf("%d", counter)),
		),
		h.Button(
			a.Attrs(a.Class("counter-increment"), a.OnClick(gt.SendMessage("INCREMENT_COUNTER", 1))),
			h.Text("Up"),
		),
	)
}

func (m Model) renderCrash() h.Element {
	return h.Div(
		a.Attrs(a.Id("crash")),
		h.P(a.Attrs(), h.Text("In order to stop errors in client-app code killing the Gotea runtime, crash protection is included!")),
		h.Button(
			a.Attrs(a.Class("crash-button"), a.OnClick(gt.SendMessageNoArgs("CRASH_ME"))),
			h.Text("Crash!"),
		),
	)
}

func (m Model) renderMemoryGame() h.Element {
	return h.Div(
		a.Attrs(),
		h.H2(a.Attrs(), h.Text(fmt.Sprintf("Turns Taken: %d", m.MemoryGame.TurnsTaken))),
		h.H2(a.Attrs(), h.Text(fmt.Sprintf("Pairs Found: %d", m.MemoryGame.Score))),
		h.H2(a.Attrs(), h.Text(fmt.Sprintf("Best Score: %d", m.MemoryGame.BestScore))),
		m.renderDeck(m.MemoryGame.Deck),
		m.renderGameStatus(),
	)
}

func (m Model) renderGameStatus() h.Element {
	if m.MemoryGame.HasWon() {
		return m.renderGameWon()
	}
	return m.renderGameOngoing()
}

func (m Model) renderGameWon() h.Element {
	return h.Div(
		a.Attrs(),
		h.H2(a.Attrs(), h.Text("Well Done! You Won!")),
		h.Button(
			a.Attrs(a.OnClick(gt.SendMessageNoArgs("RESTART_GAME"))),
			h.Text("New Game"),
		),
	)
}

func (m Model) renderGameOngoing() h.Element {
	return h.Div(a.Attrs(), h.Text("Keep going!"))
}

func (m Model) renderDeck(deck []Card) h.Element {
	return h.Div(
		a.Attrs(a.Id("deck")),
		func() []h.Element {
			var elements []h.Element
			for index, card := range deck {
				elements = append(elements, m.renderCardContainer(card, index))
			}
			return elements
		}()...,
	)
}

func (m Model) renderCardContainer(card Card, index int) h.Element {
	return h.Div(
		a.Attrs(a.Class("card-container")),
		m.renderCard(card),
		m.renderFlipButton(card, index),
	)
}

func (m Model) renderFlipButton(card Card, index int) h.Element {
	if !card.Matched {
		return h.Button(
			a.Attrs(a.Class("flipcard"), a.OnClick(gt.SendMessage("FLIP_CARD", index))),
			h.Text("Flip"),
		)
	}
	return h.Nothing(a.Attrs())
}

func (m Model) renderCard(card Card) h.Element {
	return h.Div(
		a.Attrs(
			a.Class(m.getCardClass(card)),
		),
		h.Span(
			a.Attrs(a.Class("value")),
			m.getCardValue(card),
		),
	)
}

func (m Model) getCardClass(card Card) string {
	classes := []string{"card"}
	if card.Flipped {
		classes = append(classes, "faceup")
	} else {
		classes = append(classes, "facedown")
	}
	if card.Matched {
		classes = append(classes, "matched")
	} else {
		classes = append(classes, "unmatched")
	}

	// Join with a space
	return strings.Join(classes, " ")
}

func (m Model) getCardValue(card Card) h.Element {
	if card.Flipped {
		return h.Text(fmt.Sprintf("%d", card.Value))
	}
	return h.Nothing(a.Attrs())
}

func (m Model) renderForm() h.Element {
	return h.Div(
		a.Attrs(),
		h.H1(a.Attrs(), h.Text("Form Example")),
		h.Div(
			a.Attrs(a.Class("row")),
			h.Form(
				a.Attrs(a.Class("equalchild"), a.Id("my-form")),
				h.H2(a.Attrs(), h.Text("Text Input")),
				h.Input(
					a.Attrs(
						a.Class("input"),
						a.Type("text"),
						a.Placeholder("Some simple text input"),
						a.Value(m.Form.TextInput),
						a.Name("textInput"),
						a.OnKeyUp(gt.SubmitForm("FORM_UPDATE", "my-form")),
					),
				),
				h.H2(a.Attrs(), h.Text("Simple Select")),
				h.Select(
					a.Attrs(a.Name("selectInput"), a.OnChange(gt.SubmitForm("FORM_UPDATE", "my-form"))),
					func() []h.Element {
						var options []h.Element
						for _, option := range m.Form.Options {
							options = append(options, h.Option(
								a.Attrs(
									a.Value(option),
									func() a.Attribute {
										if m.Form.SelectInput == option {
											return a.Selected(true)
										}
										return a.Selected(false)
									}(),
								),
								h.Text(option),
							))
						}
						return options
					}()...,
				),
				h.H2(a.Attrs(), h.Text("Multiple Select")),
				h.Select(
					a.Attrs(a.Name("MultipleTextInput"), a.OnChange(gt.SubmitForm("FORM_UPDATE", "my-form")), a.Size(4), a.Multiple(true)),
					h.Option(a.Attrs(a.Value("first"), m.isSelected("first", m.Form.MultipleTextInput)), h.Text("first")),
					h.Option(a.Attrs(a.Value("second"), m.isSelected("second", m.Form.MultipleTextInput)), h.Text("second")),
					h.Option(a.Attrs(a.Value("third"), m.isSelected("third", m.Form.MultipleTextInput)), h.Text("third")),
					h.Option(a.Attrs(a.Value("fourth"), m.isSelected("fourth", m.Form.MultipleTextInput)), h.Text("fourth")),
				),
				h.H2(a.Attrs(), h.Text("Text Area")),
				h.TextArea(
					a.Attrs(a.OnKeyUp(gt.SubmitForm("FORM_UPDATE", "my-form")), a.Name("TextboxInput"), a.Rows(10), a.Cols(30)),
					h.Text(m.Form.TextboxInput),
				),
				h.H2(a.Attrs(), h.Text("Radio")),
				h.Input(
					a.Attrs(a.Type("radio"), a.Name("RadioTextInput"), a.Value("male"), a.OnChange(gt.SubmitForm("FORM_UPDATE", "my-form")), m.isChecked("male", m.Form.RadioTextInput)),
				),
				h.Text("Male"),
				h.Input(
					a.Attrs(a.Type("radio"), a.Name("RadioTextInput"), a.Value("female"), a.OnChange(gt.SubmitForm("FORM_UPDATE", "my-form")), m.isChecked("female", m.Form.RadioTextInput)),
				),
				h.Text("Female"),
				h.H2(a.Attrs(), h.Text("Checkbox")),
				h.Input(
					a.Attrs(a.Type("checkbox"), a.Name("CheckboxInput"), a.OnChange(gt.SubmitForm("FORM_UPDATE", "my-form")), m.isCheckedBool(m.Form.CheckboxInput)),
				),
				h.Text("True?"),
			),
			h.Div(
				a.Attrs(a.Class("equalchild")),
				h.H2(a.Attrs(), h.Text("State")),
				m.Form.renderValues(),
			),
		),
	)
}

func (m Model) isSelected(value string, selectedValues []string) a.Attribute {
	for _, v := range selectedValues {
		if v == value {
			return a.Selected(true)
		}
	}
	return a.Selected(false)
}

func (m Model) isChecked(value string, selectedValue string) a.Attribute {
	if value == selectedValue {
		return a.Checked(true)
	}
	return a.Checked(false)
}

func (m Model) isCheckedBool(checked bool) a.Attribute {
	if checked {
		return a.Checked(true)
	}
	return a.Checked(false)
}

func (m Model) renderComponents() h.Element {
	return h.Div(
		a.Attrs(),
		h.H2(a.Attrs(), h.Text("Components Demo")),
		h.P(a.Attrs(), h.Text("Two instantiations of the 'tag selector' component, running side-by-side")),
		h.Div(
			a.Attrs(),
			h.Div(
				a.Attrs(),
				h.H3(a.Attrs(), h.Text("Select a Name")),
				m.renderTagSelector(m.NameSelector),
			),
			h.Div(
				a.Attrs(),
				h.H3(a.Attrs(), h.Text("Select a Team")),
				m.renderTagSelector(m.TeamSelector),
			),
		),
	)
}

func (m Model) renderRouting() h.Element {
	return h.Div(
		a.Attrs(),
		h.H2(a.Attrs(), h.Text("Routing")),
		h.H3(a.Attrs(), h.Text(fmt.Sprintf("Current Route: %s", m.GetRoute()))),
		h.P(a.Attrs(), h.Text("The main navigation bar above is an example of GoTea routing.")),
		h.P(a.Attrs(), h.Text("Clicking on any of the links triggers a \"CHANGE_ROUTE\" message to be sent, with the new route as an argument. The app-provided routing function then decides whether to change the view template and performs any route-conditional logic.")),
		h.H3(a.Attrs(), h.Text("Here's an internal link:")),
		h.A(a.Attrs(a.Href("/memory")), h.Text("Memory Game")),
		h.H3(a.Attrs(), h.Text("Here's an external link:")),
		h.A(a.Attrs(a.Class("external"), a.Target("_blank"), a.Href("https://duckduckgo.com")), h.Text("Search DuckDuckGo")),
		h.H3(a.Attrs(), h.Text("Route Parameters")),
		h.Ul(
			a.Attrs(),
			h.Li(a.Attrs(), h.A(a.Attrs(a.Href("/routing?myparam=1")), h.Text("myparam = 1"))),
			h.Li(a.Attrs(), h.A(a.Attrs(a.Href("/routing?myparam=2")), h.Text("myparam = 2"))),
			h.Li(a.Attrs(), h.A(a.Attrs(a.Href("/routing?myparam=3")), h.Text("myparam = 3"))),
		),
		h.H4(a.Attrs(), h.Text(fmt.Sprintf("Value of 'myparam': %s", m.RouteParam("myparam")))),
		h.H3(a.Attrs(), h.Text(fmt.Sprintf("A random number generated before rendering by the routing hook (a simple representation of route-specific data loading): %d", m.RouteData))),
	)
}

func (m Model) renderAnimation() h.Element {
	return h.Div(
		a.Attrs(),
		h.H2(a.Attrs(), h.Text("A Server-Driven Animation")),
		h.P(a.Attrs(), h.Text("A 30fps bouncing ball animation being driven entirely by the server. You'd probably never want to do this, but is fun to know that you can! Clicking 'start' fires off a never-ending sequence of messages with a 33ms delay between each. Each message sets the x and y coordinates of the ball and the scene is rerendered by the Gotea runtime.")),
		h.Button(
			a.Attrs(a.OnClick(gt.SendMessageNoArgs("START_ANIMATION"))),
			h.Text("Start Animation"),
		),
		h.Button(
			a.Attrs(a.OnClick(gt.SendMessageNoArgs("STOP_ANIMATION"))),
			h.Text("Stop Animation"),
		),
		h.Button(
			a.Attrs(a.OnClick(gt.SendMessageNoArgs("RESET_ANIMATION"))),
			h.Text("Reset Animation"),
		),
		h.P(a.Attrs(), h.Text(fmt.Sprintf("Coordinates: (%d, %d)", m.Animation.X, m.Animation.Y))),
		h.Div(
			a.Attrs(a.Id("animation-background")),
			h.Div(
				a.Attrs(a.Id("animation-ball"), a.Style(
					css.Transform(fmt.Sprintf("translate(%dpx, %dpx)", m.Animation.TranslateX, m.Animation.TranslateY)),
				)),
			),
		),
	)
}

func (m Model) renderTagSelector(selector tagselector.Model) h.Element {
	msgSelectTag := selector.UniqueMsg("TAG.SELECT")
	msgUpdateSearchInput := selector.UniqueMsg("SEARCHINPUT.UPDATE")
	msgRemoveTag := selector.UniqueMsg("TAG.REMOVE")
	searchInputID := selector.UniqueID("search-input")

	return h.Div(
		a.Attrs(a.Class("tagselector tagselector-container")),
		h.Div(
			a.Attrs(a.Class("tagselector-suggestedtags")),
			h.Input(
				a.Attrs(
					a.Id(searchInputID),
					a.Class("input"),
					a.Type("text"),
					a.Placeholder("Start typing to see tags"),
					a.Value(selector.SearchInput),
					a.OnKeyUp(gt.SendMessageWithInputValue(msgUpdateSearchInput, searchInputID)),
				),
			),
			h.Ul(
				a.Attrs(a.Class("tagselector-tagslist tagselector-suggestedtagslist")),
				func() []h.Element {
					var elements []h.Element
					for _, tag := range selector.SuggestedTags {
						elements = append(elements, h.Li(
							a.Attrs(a.Class("tagselector-tag tagselector-suggestedtag"), a.OnClick(gt.SendMessage(msgSelectTag, tag))),
							h.Text(tag),
						))
					}
					return elements
				}()...,
			),
			func() h.Element {
				if selector.ShowNoMatchMessage() {
					return h.P(a.Attrs(), h.Text(selector.NoMatchMessage))
				}
				return h.Nothing(a.Attrs())
			}(),
		),
		h.Div(
			a.Attrs(a.Class("tagselector-selectedtags")),
			h.H4(a.Attrs(a.Class("tagselector-selectedtagstitle")), h.Text("Selected Tags:")),
			h.Ul(
				a.Attrs(a.Class("tagselector-tagslist tagselector-selectedtagslist")),
				func() []h.Element {
					var elements []h.Element
					for _, tag := range selector.SelectedTags {
						elements = append(elements, h.Li(
							a.Attrs(a.Class("tagselector-tag tagselector-selectedtag"), a.OnClick(gt.SendMessage(msgRemoveTag, tag))),
							h.Text(tag),
						))
					}
					return elements
				}()...,
			),
		),
	)
}

func main() {
	app := gt.NewApp(
		gt.DefaultAppConfig,
		&Model{},
		nil,
	)

	log.Printf("Starting application server on %v\n", app.Config.Port)
	http.ListenAndServe(fmt.Sprintf(":%v", app.Config.Port), app.Router)
}
