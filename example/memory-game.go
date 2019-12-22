package main

import (
	"encoding/json"
	"math/rand"
	"time"

	gotea "github.com/jpincas/go-tea"
)

var memoryGameMessages gotea.MessageMap = gotea.MessageMap{
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

func FlipCard(args json.RawMessage, s gotea.State) (gotea.State, *gotea.Message, error) {
	state := s.(Model)

	// cast the argument to int - comes back from JS as float64
	var cardToFlip int
	if err := json.Unmarshal(args, &cardToFlip); err != nil {
		return state, nil, err
	}

	state.MemoryGame.Deck.flipCard(cardToFlip)

	if state.MemoryGame.Deck.numberFlippedCards() == 2 {
		state.MemoryGame.takeTurn()

		if state.MemoryGame.Deck.hasFoundMatch() {
			state.MemoryGame.incrementScore()
			return gotea.WithNextMsg(state, "REMOVE_MATCHES", nil)
		}

		return gotea.WithNextMsg(state, "FLIP_ALL_BACK", nil)
	}

	return gotea.WithNoMsg(state)
}

func RestartGame(_ json.RawMessage, s gotea.State) (gotea.State, *gotea.Message, error) {
	state := s.(Model)
	state.MemoryGame.Deck.reset()
	state.MemoryGame.resetScores()
	return gotea.WithNoMsg(state)
}

func FlipAllBack(_ json.RawMessage, s gotea.State) (gotea.State, *gotea.Message, error) {
	time.Sleep(500 * time.Millisecond)
	state := s.(Model)
	state.MemoryGame.Deck.flipAllBack()
	return gotea.WithNoMsg(state)
}

func RemoveMatches(_ json.RawMessage, s gotea.State) (gotea.State, *gotea.Message, error) {
	time.Sleep(500 * time.Millisecond)
	state := s.(Model)
	state.MemoryGame.Deck.removeMatches()

	if state.MemoryGame.HasWon() {
		if state.MemoryGame.BestScore == 0 || state.MemoryGame.TurnsTaken < state.MemoryGame.BestScore {
			state.MemoryGame.BestScore = state.MemoryGame.TurnsTaken
		}
	}

	return gotea.WithNoMsg(state)
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
