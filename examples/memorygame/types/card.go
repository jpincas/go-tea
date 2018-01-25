package types

import (
	gotea "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/examples/memorygame/messages"
)

type Card struct {
	Value   int
	Flipped bool
	Matched bool
}

func FlipCard(args gotea.MessageArguments, s *gotea.Session) (gotea.State, *gotea.Message) {
	// cast the argument to int - comes back from JS as float64
	cardToFlip := int(args.(float64))
	state := s.State.(Model)

	state.Deck.flipCard(cardToFlip)

	if state.Deck.numberFlippedCards() == 2 {
		state.takeTurn()

		if state.Deck.hasFoundMatch() {
			state.incrementScore()
			nextMsg := messages.RemoveMatches()
			return state, &nextMsg
		}

		nextMsg := messages.FlipAllBack()
		return state, &nextMsg
	}

	return state, nil
}
