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

// Message generator

func FlipCard(args gotea.MessageArguments) gotea.Message {
	return gotea.Message{
		Func:      flipCard,
		FuncCode:  "FlipCard",
		Arguments: args,
	}
}

// Messages

func flipCard(args gotea.MessageArguments, s *gotea.Session) {

	cardToFlip := int(args.(float64))

	s.State.(Model).Deck.flipCard(cardToFlip)

	// if this is the second card of the pair being flipped
	if s.State.(Model).Deck.numberFlippedCards() == 2 {

		// !!!!!!!!!!!!!!!!!!!!!!!!!
		// We need to find a much neater solution for this
		castModel := s.State.(Model)
		pointerToCastModel := &castModel
		// the function can now modify state
		pointerToCastModel.takeTurn()
		// and assign it back (after dereferencing)
		s.State = *pointerToCastModel

		if s.State.(Model).Deck.hasFoundMatch() {

			// !!!!!!!!!!!!!!!!!!!!!!!!!
			// We need to find a much neater solution for this
			castModel := s.State.(Model)
			pointerToCastModel := &castModel
			pointerToCastModel.incrementScore()
			s.State = *pointerToCastModel

			go func() {
				time.Sleep(1500 * time.Millisecond)
				gotea.Message{
					Func:      removeMatches,
					Arguments: nil,
				}.Process(s)
			}()
		} else {
			go func() {
				time.Sleep(1500 * time.Millisecond)
				gotea.Message{
					Func:      flipAllBack,
					Arguments: nil,
				}.Process(s)
			}()
		}
	}

}
