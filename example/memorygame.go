package main

import (
	"encoding/json"
	"math/rand"
	"time"

	gotea "github.com/jpincas/go-tea"
)

type MemoryGame struct {
	Deck              Deck
	LastAttemptedCard int
	TurnsTaken        int
	Score             int
}

func (m *MemoryGame) takeTurn() {
	m.TurnsTaken++
}

func (m *MemoryGame) incrementScore() {
	m.Score++
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

	state.MemoryGame.Deck.flipCard(int(cardToFlip))

	if state.MemoryGame.Deck.numberFlippedCards() == 2 {
		state.MemoryGame.takeTurn()

		if state.MemoryGame.Deck.hasFoundMatch() {
			state.MemoryGame.incrementScore()

			return state, &gotea.Message{
				Message:   "REMOVE_MATCHES",
				Arguments: nil,
			}, nil
		}

		return state, &gotea.Message{
			Message:   "FLIP_ALL_BACK",
			Arguments: nil,
		}, nil
	}

	return state, nil, nil
}

func FlipAllBack(_ json.RawMessage, s gotea.State) (gotea.State, *gotea.Message, error) {
	time.Sleep(500 * time.Millisecond)
	state := s.(Model)
	state.MemoryGame.Deck.flipAllBack()
	return state, nil, nil
}

func RemoveMatches(_ json.RawMessage, s gotea.State) (gotea.State, *gotea.Message, error) {
	time.Sleep(500 * time.Millisecond)
	state := s.(Model)
	state.MemoryGame.Deck.removeMatches()
	return state, nil, nil
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
