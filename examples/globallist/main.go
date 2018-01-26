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

	// message function map
	gotea.App.Messages["AddCoordinate"] = addCoordinate

	// session state seeder
	gotea.App.NewSession = func() gotea.Session {
		return gotea.Session{
			State: Model{
				Coordinates: &CoordinateDB,
			},
		}
	}

	// main view renderer
	gotea.App.RenderView = WriteMain

}

// APP SPECFIC

type Coordinate struct {
	X, Y int
}

// CoordinateDB is a simple database of coordinates
var CoordinateDB []Coordinate

// addCoordinate adds a random coordinate to the database
func addCoordinate(_ gotea.MessageArguments, s *gotea.Session) (gotea.State, *gotea.Message) {
	// create new coordinate, add to 'database'
	x := rand.Intn(100)
	y := rand.Intn(100)
	CoordinateDB = append(CoordinateDB, Coordinate{x, y})

	// broadcast to all active connections
	gotea.App.Broadcast()

	// since we don't mutate the session state or need to add any more messages
	return s.State, nil
}

// Message generator

func AddCoordinate() gotea.Message {
	return gotea.Message{
		FuncCode:  "AddCoordinate",
		Arguments: nil,
	}
}

// MAIN

// main starts the server
func main() {
	gotea.App.Start("../../dist")
}
