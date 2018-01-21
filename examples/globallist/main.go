package main

import (
	"math/rand"

	gotea "github.com/jpincas/go-tea"
)

// Model is the data to be maintained as state
// - REQUIRED by gotea runtime
type Model struct {
	// We don't want to accumulate the list individually on all sessions
	// so we use a POINTER to the global coordinate 'database'
	// Therefore all sessions will share a global coordinate list
	Coordinates *[]Coordinate
}

// InitialState should return an intial model
// - REQUIRED by gotea runtime
func (m Model) initialState() Model {
	return Model{
		Coordinates: &CoordinateDB,
	}
}

func init() {
	// Initialise the message map
	// - REQUIRED by gotea runtime
	// - but you could also add to this map in other files
	// - e.g. App.Messages["newMessage"] = newFunction
	gotea.App.Messages = map[string]func(map[string]interface{}, *gotea.Session){
		"add-coordinate": addCoordinate,
	}

	// create a seed for initial session state
	gotea.App.InitialSessionState = Model{
		Coordinates: &CoordinateDB,
	}

}

// APP SPECFIC

type Coordinate struct {
	X, Y int
}

// CoordinateDB is a simple database of coordinates
var CoordinateDB []Coordinate

// addCoordinate adds a random coordinate to the database
func addCoordinate(params map[string]interface{}, s *gotea.Session) {
	// create new coordinate, add to 'database'
	x := rand.Intn(100)
	y := rand.Intn(100)
	CoordinateDB = append(CoordinateDB, Coordinate{x, y})

	// broadcast to all active connections
	gotea.App.Broadcast()
}

// MAIN

// main starts the server
func main() {
	gotea.App.Start("../../dist")
}
