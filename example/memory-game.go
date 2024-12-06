package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"

	gt "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/msg"
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

func FlipCard(args json.RawMessage, s gt.State) gt.Response {
	state := model(s)

	cardToFlip, err := msg.DecodeInt(args)
	if err != nil {
		return gt.RespondWithError(err)
	}

	state.MemoryGame.Deck.flipCard(cardToFlip)

	if state.MemoryGame.Deck.numberFlippedCards() == 2 {
		state.MemoryGame.takeTurn()

		if state.MemoryGame.Deck.hasFoundMatch() {
			state.MemoryGame.incrementScore()
			return gt.RespondWithDelayedNextMsg("REMOVE_MATCHES", nil, 1000)
		}

		return gt.RespondWithDelayedNextMsg("FLIP_ALL_BACK", nil, 1000)
	}

	return gt.Respond()
}

func RestartGame(_ json.RawMessage, s gt.State) gt.Response {
	state := model(s)
	state.MemoryGame.Deck.reset()
	state.MemoryGame.resetScores()
	return gt.Respond()
}

func FlipAllBack(_ json.RawMessage, s gt.State) gt.Response {
	state := model(s)
	state.MemoryGame.Deck.flipAllBack()
	return gt.Respond()
}

func RemoveMatches(_ json.RawMessage, s gt.State) gt.Response {
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
		a.Attrs(),
		h.H2(a.Attrs(), h.Text(fmt.Sprintf("Turns Taken: %d", game.TurnsTaken))),
		h.H2(a.Attrs(), h.Text(fmt.Sprintf("Pairs Found: %d", game.Score))),
		h.H2(a.Attrs(), h.Text(fmt.Sprintf("Best Score: %d", game.BestScore))),
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
		a.Attrs(),
		h.H2(a.Attrs(), h.Text("Well Done! You Won!")),
		h.Button(
			a.Attrs(a.OnClick(gt.SendMessageNoArgs("RESTART_GAME"))),
			h.Text("New Game"),
		),
	)
}

func renderGameOngoing() h.Element {
	return h.Div(a.Attrs(), h.Text("Keep going!"))
}

func (deck Deck) render() h.Element {
	return h.Div(
		a.Attrs(a.Id("deck")),
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
		a.Attrs(a.Class("card-container")),
		card.render(),
		card.renderFlipButton(index),
	)
}

func (card Card) renderFlipButton(index int) h.Element {
	if !card.Matched {
		return h.Button(
			a.Attrs(a.Class("flipcard"), a.OnClick(gt.SendMessage("FLIP_CARD", index))),
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
			a.Attrs(a.Class("value")),
			card.getValue(),
		),
	)
}

func (card Card) getClass() string {
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

func (card Card) getValue() h.Element {
	if card.Flipped {
		return h.Text(fmt.Sprintf("%d", card.Value))
	}
	return h.Nothing(a.Attrs())
}
