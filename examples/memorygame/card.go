package main

import (
	"time"

	gotea "github.com/jpincas/go-tea"
)

// Types

type Card struct {
	Value   int
	Flipped bool
	Matched bool
}

// Messages

func flipCard(params map[string]interface{}, s *gotea.Session) {
	cardToFlip := int(params["cardIndex"].(float64))
	s.State.(Model).Deck.flipCard(cardToFlip)

	// if this is the second card of the pair being flipped
	if s.State.(Model).Deck.numberFlippedCards() == 2 {
		s.State.(Model).takeTurn()
		if s.State.(Model).Deck.hasFoundMatch() {
			s.State.(Model).incrementScore()
			go func() {
				time.Sleep(1500 * time.Millisecond)
				gotea.Msg{"removeMatches", params}.Process(s)
			}()
		} else {
			go func() {
				time.Sleep(1500 * time.Millisecond)
				gotea.Msg{"flipAllBack", params}.Process(s)
			}()
		}
	}

}
