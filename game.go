package main

import (
	"math/rand"
	"time"
)

type Card struct {
	Value   int
	Flipped bool
	Matched bool
}

type Deck []Card

func newDeck(n int) (deck Deck) {
	cardValues := append([]int{}, append(rand.Perm(n), rand.Perm(n)...)...)
	for _, v := range cardValues {
		deck = append(deck, Card{v, false, false})
	}
	return
}

func (deck Deck) flipAllBack() {
	for i, _ := range deck {
		deck[i].Flipped = false
	}
}

func (deck Deck) flipCard(i int) {
	deck[i].Flipped = !deck[i].Flipped
}

func flipCard(params map[string]interface{}) {
	cardToFlip := int(params["cardIndex"].(float64))
	State.Deck.flipCard(cardToFlip)

	// if this is the second card of the pair being flipped
	if State.Deck.numberFlippedCards() == 2 {
		State.TurnsTaken++
		if State.Deck.hasFoundMatch() {
			State.Score++
			go func() {
				time.Sleep(1500 * time.Millisecond)
				processMessage(MsgPayload{"removeMatches", params})
			}()
		} else {
			go func() {
				time.Sleep(1500 * time.Millisecond)
				processMessage(MsgPayload{"flipAllBack", params})
			}()
		}
	}

}

func flipAllBack(params map[string]interface{}) {
	State.Deck.flipAllBack()
}

func (deck Deck) onlyFlipped() (flippedCards Deck) {
	// filter to only flipped cards
	for _, card := range State.Deck {
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

func removeMatches(params map[string]interface{}) {
	State.Deck.removeMatches()
}

func (deck Deck) removeMatches() {
	for i, card := range deck {
		if card.Flipped {
			deck[i].Matched = true
			deck[i].Flipped = false
		}
	}
}

func (model Model) HasWon() bool {
	for _, card := range model.Deck {
		if !card.Matched {
			return false
		}
	}
	return true
}
