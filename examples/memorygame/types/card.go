package types

import (
	gotea "github.com/jpincas/go-tea"
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

			return state, gotea.NextMsg(RemoveMatchesMsg())
		}

		return state, gotea.NextMsg(FlipAllBackMsg())
	}

	return state, nil
}

func FlipCardMsg(index int) gotea.Message {
	return gotea.Message{
		FuncCode:  "FlipCard",
		Arguments: index,
		Func:      FlipCard,
	}
}
