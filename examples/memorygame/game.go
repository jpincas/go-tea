package main

func (model Model) takeTurn() {
	model.TurnsTaken++
}

func (model Model) incrementScore() {
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
