package main

import (
	"time"
)

// Types

type Card struct {
	Value   int
	Flipped bool
	Matched bool
}

// Messages

func flipCard(params map[string]interface{}, s *Session) {
	cardToFlip := int(params["cardIndex"].(float64))
	s.State.Deck.flipCard(cardToFlip)

	// if this is the second card of the pair being flipped
	if s.State.Deck.numberFlippedCards() == 2 {
		s.State.TurnsTaken++
		if s.State.Deck.hasFoundMatch() {
			s.State.Score++
			go func() {
				time.Sleep(1500 * time.Millisecond)
				Msg{"removeMatches", params}.Process(s)
			}()
		} else {
			go func() {
				time.Sleep(1500 * time.Millisecond)
				Msg{"flipAllBack", params}.Process(s)
			}()
		}
	}

}
