package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Model struct {
	Deck              Deck
	DeckFlippedStatus []bool
	LastAttemptedCard int
	TurnsTaken        int
	Score             int
}

var State Model
var Templates *template.Template
var Messages map[string]func(map[string]interface{}, *websocket.Conn)

func init() {

	// parse templates
	Templates = template.Must(template.New("main").ParseGlob("components/**/*.html"))

	// init the message map
	Messages = map[string]func(map[string]interface{}, *websocket.Conn){
		"flipcard":      flipCard,
		"flipAllBack":   flipAllBack,
		"removeMatches": removeMatches,
	}

	// declare the inital state of the model
	// on startup
	State = Model{
		Deck:              newDeck(10),
		TurnsTaken:        0,
		LastAttemptedCard: 11, //hack
		Score:             0,
	}
}

func main() {

	fs := http.FileServer(http.Dir("dist"))

	http.HandleFunc("/server", handler)
	http.Handle("/", fs)

	log.Println("Staring server...")
	http.ListenAndServe(":8080", nil)
}
