package main

// Model is the data to be maintained as state
// - REQUIRED by gotea runtime
type Model struct {
	Deck              Deck
	DeckFlippedStatus []bool
	LastAttemptedCard int
	TurnsTaken        int
	Score             int
}

// InitialState should return an intial model
// - REQUIRED by gotea runtime
func initialState() Model {
	return Model{
		Deck:              newDeck(10),
		TurnsTaken:        0,
		LastAttemptedCard: 11, //hack
		Score:             0,
	}
}

func init() {
	// Initialise the message map
	// - REQUIRED by gotea runtime
	// - but you could also add to this map in other files
	// - e.g. App.Messages["newMessage"] = newFunction
	App.Messages = map[string]func(map[string]interface{}, *Session){
		"flipcard":      flipCard,
		"flipAllBack":   flipAllBack,
		"removeMatches": removeMatches,
	}
}
