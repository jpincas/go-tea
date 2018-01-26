package types

import (
	"math/rand"
	"time"

	gotea "github.com/jpincas/go-tea"
)

type Deck []Card

func FlipAllBack(_ gotea.MessageArguments, s *gotea.Session) (gotea.State, *gotea.Message) {
	time.Sleep(1500 * time.Millisecond)
	state := s.State.(Model)
	state.Deck.flipAllBack()
	return state, nil
}

func RemoveMatches(_ gotea.MessageArguments, s *gotea.Session) (gotea.State, *gotea.Message) {
	time.Sleep(1500 * time.Millisecond)
	state := s.State.(Model)
	state.Deck.removeMatches()
	return state, nil
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
