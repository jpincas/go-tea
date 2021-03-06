package main

import (
	"encoding/json"
	"math/rand"

	gt "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/msg"
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
