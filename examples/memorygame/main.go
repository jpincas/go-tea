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

	// set messages
	gotea.App.Messages["flipcard"] = flipCard
	gotea.App.Messages["flipAllBack"] = flipAllBack
	gotea.App.Messages["removeMatches"] = removeMatches

	// function that returns a new session
	gotea.App.NewSession = func() gotea.Session {
		return gotea.Session{
			State: Model{
				Deck:              newDeck(4),
				TurnsTaken:        0,
				LastAttemptedCard: 5, //hack
				Score:             0,
			},
		}
	}

}

// main starts the server
func main() {
	gotea.App.Start("../../dist")
}
