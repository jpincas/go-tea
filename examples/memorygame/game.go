package main

func (model Model) HasWon() bool {
	for _, card := range model.Deck {
		if !card.Matched {
			return false
		}
	}
	return true
}
