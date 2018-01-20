package main

type Model struct {
	Deck              Deck
	DeckFlippedStatus []bool
	LastAttemptedCard int
	TurnsTaken        int
	Score             int
}

func initialState() Model {
	return Model{
		Deck:              newDeck(10),
		TurnsTaken:        0,
		LastAttemptedCard: 11, //hack
		Score:             0,
	}
}

func init() {

	// init the message map
	App.Messages = map[string]func(map[string]interface{}, *Session){
		"flipcard":      flipCard,
		"flipAllBack":   flipAllBack,
		"removeMatches": removeMatches,
	}
}
