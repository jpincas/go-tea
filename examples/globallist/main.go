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

func init() {

	gotea.App.Messages["add-coordinate"] = addCoordinate

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
func addCoordinate(_ gotea.MsgTag, s *gotea.Session) {
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
