package main

import (
	gotea "github.com/jpincas/go-tea"
)

// Model is the data to be maintained as state
// - REQUIRED by gotea runtime
type Model struct {
	Deck              Deck
	LastAttemptedCard int
	TurnsTaken        int
	Score             int
}

func init() {
	// Initialise the message map
	// - REQUIRED by gotea runtime
	// - but you could also add to this map in other files
	// - e.g. App.Messages["newMessage"] = newFunction
	gotea.App.Messages = map[string]func(map[string]interface{}, *gotea.Session){
		"flipcard":      flipCard,
		"flipAllBack":   flipAllBack,
		"removeMatches": removeMatches,
	}

	// create a seed for initial session state
	gotea.App.InitialSessionState = Model{
		Deck:              newDeck(4),
		TurnsTaken:        0,
		LastAttemptedCard: 5, //hack
		Score:             0,
	}

}

// main starts the server
func main() {
	gotea.App.Start("../../dist")
}
