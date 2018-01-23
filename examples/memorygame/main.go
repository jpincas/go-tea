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
	// gotea.App.Messages["flipcard"] = flipCard
	// gotea.App.Messages["flipAllBack"] = flipAllBack
	// gotea.App.Messages["removeMatches"] = removeMatches

	//gotea.App.Messages[FlipCard(nil).FuncCode] = FlipCard(nil).Func
	//gotea.App.TemplateFuncs[FlipCard(nil).FuncCode] = FlipCard

	gotea.RegisterMessages(
		FlipCard,
	)

	// create a seed for initial session state
	gotea.App.InitialSessionState = Model{
		Deck:              newDeck(4),
		TurnsTaken:        0,
		LastAttemptedCard: 5, //hack
		Score:             0,
	}

	gotea.ParseTemplates()

}

// main starts the server
func main() {
	gotea.App.Start("../../dist")
}
