package main

// Model is the data to be maintained as state
// - REQUIRED by gotea runtime
type Model struct {
	Deck              Deck
	LastAttemptedCard int
	TurnsTaken        int
	Score             int
}

func (model *Model) takeTurn() {
	model.TurnsTaken++
}

func (model *Model) incrementScore() {
	model.Score++
}

func (model Model) HasWon() bool {
	for _, card := range model.Deck {
		if !card.Matched {
			return false
		}
	}
	return true
}
