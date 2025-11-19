package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
)

var memoryGameMessages gt.MessageMap = gt.MessageMap{
	"FLIP_CARD":         FlipCard,
	"REMOVE_MATCHES":    RemoveMatches,
	"FLIP_ALL_BACK":     FlipAllBack,
	"RESTART_GAME":      RestartGame,
	"CHANGE_DIFFICULTY": ChangeDifficulty,
	"SUBMIT_INITIALS":   SubmitInitials,
}

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
)

func (d Difficulty) String() string {
	return [...]string{"Easy", "Medium", "Hard"}[d]
}

func (d Difficulty) Pairs() int {
	return [...]int{6, 8, 12}[d]
}

type MemoryGame struct {
	Deck              Deck
	LastAttemptedCard int
	TurnsTaken        int
	Score             int
	BestScore         int
	Difficulty        Difficulty
	Leaderboard       *Leaderboard
	AskingForInitials bool
}

func (m *MemoryGame) takeTurn() {
	m.TurnsTaken++
}

func (m *MemoryGame) incrementScore() {
	m.Score++
}

func (m *MemoryGame) resetScores() {
	m.TurnsTaken = 0
	m.Score = 0
}

func (m MemoryGame) HasWon() bool {
	for _, card := range m.Deck {
		if !card.Matched {
			return false
		}
	}
	return true
}

type Deck []Card

type Card struct {
	Value   int
	Flipped bool
	Matched bool
}

func FlipCard(m gt.Message, s gt.State) gt.Response {
	state := model(s)
	cardToFlip := m.ArgsToInt()

	state.MemoryGame.Deck.flipCard(cardToFlip)

	if state.MemoryGame.Deck.numberFlippedCards() == 2 {
		state.MemoryGame.takeTurn()

		if state.MemoryGame.Deck.hasFoundMatch() {
			state.MemoryGame.incrementScore()
			return gt.RespondWithDelayedNextMsg(gt.Message{Message: "REMOVE_MATCHES"}, 1000)
		}

		return gt.RespondWithDelayedNextMsg(gt.Message{Message: "FLIP_ALL_BACK"}, 1000)
	}

	return gt.Respond()
}

func ChangeDifficulty(m gt.Message, s gt.State) gt.Response {
	state := model(s)
	newDifficulty := Difficulty(m.ArgsToInt())
	state.MemoryGame.Difficulty = newDifficulty
	state.MemoryGame.Deck = NewDeck(newDifficulty.Pairs())
	state.MemoryGame.resetScores()
	return gt.Respond()
}

func RestartGame(_ gt.Message, s gt.State) gt.Response {
	state := model(s)
	state.MemoryGame.Deck = NewDeck(state.MemoryGame.Difficulty.Pairs())
	state.MemoryGame.resetScores()
	state.MemoryGame.AskingForInitials = false
	return gt.Respond()
}

func FlipAllBack(_ gt.Message, s gt.State) gt.Response {
	state := model(s)
	state.MemoryGame.Deck.flipAllBack()
	return gt.Respond()
}

func SubmitInitials(m gt.Message, s gt.State) gt.Response {
	state := model(s)
	initials := m.ArgsToString()
	
	// Basic validation
	if len(initials) > 3 {
		initials = initials[:3]
	}
	initials = strings.ToUpper(initials)

	score := HighScore{
		Initials: initials,
		Score:    state.MemoryGame.TurnsTaken,
		Date:     time.Now(),
	}

	state.MemoryGame.Leaderboard.AddScore(state.MemoryGame.Difficulty, score)
	state.MemoryGame.Leaderboard.Save()
	state.MemoryGame.AskingForInitials = false

	return gt.Respond()
}

func RemoveMatches(_ gt.Message, s gt.State) gt.Response {
	state := model(s)
	state.MemoryGame.Deck.removeMatches()

	if state.MemoryGame.HasWon() {
		if state.MemoryGame.BestScore == 0 || state.MemoryGame.TurnsTaken < state.MemoryGame.BestScore {
			state.MemoryGame.BestScore = state.MemoryGame.TurnsTaken
		}

		if state.MemoryGame.Leaderboard.IsHighScore(state.MemoryGame.Difficulty, state.MemoryGame.TurnsTaken) {
			state.MemoryGame.AskingForInitials = true
		}
	}

	return gt.Respond()
}

func NewDeck(n int) (deck Deck) {
	cardValues := append([]int{}, append(rand.Perm(n), rand.Perm(n)...)...)
	for _, v := range cardValues {
		deck = append(deck, Card{v, false, false})
	}
	return
}

func (deck Deck) flipAllBack() {
	for i := range deck {
		deck[i].Flipped = false
	}
}

func (deck Deck) flipCard(i int) {
	deck[i].Flipped = !deck[i].Flipped
}

func (deck Deck) onlyFlipped() (flippedCards Deck) {
	// filter to only flipped cards
	for _, card := range deck {
		if card.Flipped {
			flippedCards = append(flippedCards, card)
		}
	}
	return
}

func (deck Deck) numberFlippedCards() int {
	return len(deck.onlyFlipped())
}

func (deck Deck) hasFoundMatch() bool {
	// if there are exactly 2 cards flipped
	// and their valus match
	flippedCards := deck.onlyFlipped()
	if len(flippedCards) == 2 {
		if flippedCards[0].Value == flippedCards[1].Value {
			return true
		}
	}
	return false
}

func (deck Deck) removeMatches() {
	for i, card := range deck {
		if card.Flipped {
			deck[i].Matched = true
			deck[i].Flipped = false
		}
	}
}

func (deck Deck) reset() {
	for i := range deck {
		deck[i].Matched = false
		deck[i].Flipped = false
	}
}

// Render
func (game MemoryGame) render() h.Element {
	return h.Div(
		a.Attrs(a.Class("space-y-6")),
		h.Div(
			a.Attrs(a.Class("flex justify-center space-x-4")),
			renderDifficultyButton(Easy, game.Difficulty),
			renderDifficultyButton(Medium, game.Difficulty),
			renderDifficultyButton(Hard, game.Difficulty),
		),
		h.Div(
			a.Attrs(a.Class("grid grid-cols-3 gap-4 text-center")),
			h.Div(
				a.Attrs(a.Class("bg-white p-4 rounded-lg shadow")),
				h.H3(a.Attrs(a.Class("text-sm font-medium text-gray-500")), h.Text("Turns Taken")),
				h.P(a.Attrs(a.Class("mt-1 text-3xl font-semibold text-gray-900")), h.Text(fmt.Sprintf("%d", game.TurnsTaken))),
			),
			h.Div(
				a.Attrs(a.Class("bg-white p-4 rounded-lg shadow")),
				h.H3(a.Attrs(a.Class("text-sm font-medium text-gray-500")), h.Text("Pairs Found")),
				h.P(a.Attrs(a.Class("mt-1 text-3xl font-semibold text-gray-900")), h.Text(fmt.Sprintf("%d", game.Score))),
			),
			h.Div(
				a.Attrs(a.Class("bg-white p-4 rounded-lg shadow")),
				h.H3(a.Attrs(a.Class("text-sm font-medium text-gray-500")), h.Text("Best Score")),
				h.P(a.Attrs(a.Class("mt-1 text-3xl font-semibold text-gray-900")), h.Text(fmt.Sprintf("%d", game.BestScore))),
			),
		),
		game.Deck.render(game.Difficulty),
		game.renderStatus(),
		game.renderLeaderboard(),
	)
}

func (game MemoryGame) renderLeaderboard() h.Element {
	scores := game.Leaderboard.GetScores(game.Difficulty)
	if len(scores) == 0 {
		return h.Nothing(a.Attrs())
	}

	return h.Div(
		a.Attrs(a.Class("mt-8 bg-white p-6 rounded-lg shadow")),
		h.H3(a.Attrs(a.Class("text-lg font-medium text-gray-900 mb-4 text-center")), h.Text(fmt.Sprintf("%s Leaderboard", game.Difficulty))),
		h.Div(
			a.Attrs(a.Class("flex flex-col")),
			h.Div(
				a.Attrs(a.Class("-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8")),
				h.Div(
					a.Attrs(a.Class("py-2 align-middle inline-block min-w-full sm:px-6 lg:px-8")),
					h.Div(
						a.Attrs(a.Class("shadow overflow-hidden border-b border-gray-200 sm:rounded-lg")),
						h.Div(
							a.Attrs(a.Class("min-w-full divide-y divide-gray-200")),
							h.Div(
								a.Attrs(a.Class("bg-gray-50 grid grid-cols-4 gap-4 px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider")),
								h.Div(a.Attrs(), h.Text("Rank")),
								h.Div(a.Attrs(), h.Text("Initials")),
								h.Div(a.Attrs(), h.Text("Score")),
								h.Div(a.Attrs(), h.Text("Date")),
							),
							h.Div(
								a.Attrs(a.Class("bg-white divide-y divide-gray-200")),
								func() []h.Element {
									var rows []h.Element
									for i, score := range scores {
										rows = append(rows, h.Div(
											a.Attrs(a.Class("grid grid-cols-4 gap-4 px-6 py-4 whitespace-nowrap text-sm text-gray-500")),
											h.Div(a.Attrs(), h.Text(fmt.Sprintf("%d", i+1))),
											h.Div(a.Attrs(a.Class("font-medium text-gray-900")), h.Text(score.Initials)),
											h.Div(a.Attrs(), h.Text(fmt.Sprintf("%d", score.Score))),
											h.Div(a.Attrs(), h.Text(score.Date.Format("Jan 02"))),
										))
									}
									return rows
								}()...,
							),
						),
					),
				),
			),
		),
	)
}

func renderDifficultyButton(d, current Difficulty) h.Element {
	classes := "px-4 py-2 rounded-md text-sm font-medium focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
	if d == current {
		classes += " bg-indigo-600 text-white"
	} else {
		classes += " bg-white text-gray-700 hover:bg-gray-50 border border-gray-300"
	}

	return h.Button(
		a.Attrs(a.Class(classes), a.OnClick(gt.SendBasicMessage("CHANGE_DIFFICULTY", int(d)))),
		h.Text(d.String()),
	)
}

func (game MemoryGame) renderStatus() h.Element {
	if game.HasWon() {
		return game.renderGameWon()
	}
	return renderGameOngoing()
}

func (game MemoryGame) renderGameWon() h.Element {
	if game.AskingForInitials {
		return h.Div(
			a.Attrs(a.Class("text-center space-y-4 bg-yellow-50 p-6 rounded-lg border border-yellow-200")),
			h.H2(a.Attrs(a.Class("text-3xl font-bold text-yellow-600")), h.Text("New High Score!")),
			h.P(a.Attrs(a.Class("text-gray-700")), h.Text("Enter your initials to join the leaderboard:")),
			h.Div(
				a.Attrs(a.Class("flex justify-center space-x-2")),
				h.Input(
					a.Attrs(
						a.Id("initials-input"),
						a.Type("text"),
						a.MaxLength(3),
						a.Class("w-20 text-center uppercase text-2xl font-bold border-2 border-yellow-400 rounded-md focus:ring-yellow-500 focus:border-yellow-500"),
						a.Placeholder("AAA"),
					),
				),
				h.Button(
					a.Attrs(a.Class("inline-flex items-center px-4 py-2 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-yellow-600 hover:bg-yellow-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-yellow-500"), a.OnClick(gt.SendBasicMessageWithValueFromInput("SUBMIT_INITIALS", "initials-input"))),
					h.Text("Submit"),
				),
			),
		)
	}

	return h.Div(
		a.Attrs(a.Class("text-center space-y-4")),
		h.H2(a.Attrs(a.Class("text-3xl font-bold text-green-600")), h.Text("Well Done! You Won!")),
		h.Button(
			a.Attrs(a.Class("inline-flex items-center px-4 py-2 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"), a.OnClick(gt.SendBasicMessageNoArgs("RESTART_GAME"))),
			h.Text("New Game"),
		),
	)
}

func renderGameOngoing() h.Element {
	return h.Div(a.Attrs(a.Class("text-center text-gray-500 italic")), h.Text("Keep going!"))
}

func (deck Deck) render(d Difficulty) h.Element {
	gridCols := "grid-cols-4"
	if d == Hard {
		gridCols = "grid-cols-6"
	}

	return h.Div(
		a.Attrs(a.Id("deck"), a.Class(fmt.Sprintf("grid %s gap-4", gridCols))),
		func() []h.Element {
			var elements []h.Element
			for index, card := range deck {
				elements = append(elements, card.renderContainer(index))
			}
			return elements
		}()...,
	)
}

func (card Card) renderContainer(index int) h.Element {
	return h.Div(
		a.Attrs(a.Class("relative")),
		card.render(),
		card.renderFlipButton(index),
	)
}

func (card Card) renderFlipButton(index int) h.Element {
	if !card.Matched {
		return h.Button(
			a.Attrs(a.Class("absolute inset-0 w-full h-full opacity-0 cursor-pointer"), a.OnClick(gt.SendBasicMessage("FLIP_CARD", index))),
			h.Text("Flip"),
		)
	}
	return h.Nothing(a.Attrs())
}

func (card Card) render() h.Element {
	return h.Div(
		a.Attrs(
			a.Class(card.getClass()),
		),
		h.Span(
			a.Attrs(a.Class("text-4xl font-bold")),
			card.getValue(),
		),
	)
}

func (card Card) getClass() string {
	classes := []string{"h-32 flex items-center justify-center rounded-lg shadow-md transition-all duration-300"}
	if card.Flipped {
		classes = append(classes, "bg-white text-indigo-600 transform rotate-y-180")
	} else {
		classes = append(classes, "bg-indigo-600 text-white")
	}
	if card.Matched {
		classes = append(classes, "opacity-50 cursor-default")
	} else {
		classes = append(classes, "cursor-pointer hover:shadow-lg")
	}

	// Join with a space
	return strings.Join(classes, " ")
}

func (card Card) getValue() h.Element {
	if card.Flipped {
		return h.Text(fmt.Sprintf("%d", card.Value))
	}
	return h.Nothing(a.Attrs())
}
