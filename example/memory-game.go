package main

import (
	"fmt"
	"math/rand"
	"strings"

	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
)

var memoryGameMessages gt.MessageMap = gt.MessageMap{
	"FLIP_CARD":      FlipCard,
	"REMOVE_MATCHES": RemoveMatches,
	"FLIP_ALL_BACK":  FlipAllBack,
	"RESTART_GAME":   RestartGame,
}

type MemoryGame struct {
	Deck              Deck
	LastAttemptedCard int
	TurnsTaken        int
	Score             int
	BestScore         int
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

func RestartGame(_ gt.Message, s gt.State) gt.Response {
	state := model(s)
	state.MemoryGame.Deck.reset()
	state.MemoryGame.resetScores()
	return gt.Respond()
}

func FlipAllBack(_ gt.Message, s gt.State) gt.Response {
	state := model(s)
	state.MemoryGame.Deck.flipAllBack()
	return gt.Respond()
}

func RemoveMatches(_ gt.Message, s gt.State) gt.Response {
	state := model(s)
	state.MemoryGame.Deck.removeMatches()

	if state.MemoryGame.HasWon() {
		if state.MemoryGame.BestScore == 0 || state.MemoryGame.TurnsTaken < state.MemoryGame.BestScore {
			state.MemoryGame.BestScore = state.MemoryGame.TurnsTaken
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
		game.Deck.render(),
		game.renderStatus(),
	)
}

func (game MemoryGame) renderStatus() h.Element {
	if game.HasWon() {
		return renderGameWon()
	}
	return renderGameOngoing()
}

func renderGameWon() h.Element {
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

func (deck Deck) render() h.Element {
	return h.Div(
		a.Attrs(a.Id("deck"), a.Class("grid grid-cols-4 gap-4")),
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
