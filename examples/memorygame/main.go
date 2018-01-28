package main

import (
	gotea "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/examples/memorygame/types"
	"github.com/jpincas/go-tea/examples/memorygame/views"
)

func init() {

	// 2) register the function that returns a new session
	gotea.App.NewSession = func() gotea.Session {
		return gotea.Session{
			State: types.Model{
				Deck:              types.NewDeck(4),
				TurnsTaken:        0,
				LastAttemptedCard: 5, //hack
				Score:             0,
			},
		}
	}

	// 3) tell the app which render function to use to render the base view
	gotea.App.RenderView = views.WriteMain

}

// main starts the server
func main() {
	gotea.App.Start("../../dist")
}
