package main

import (
	gotea "github.com/jpincas/go-tea"
)

// Model is the data to be maintained as state
// - REQUIRED by gotea runtime
type Model struct {
	PasswordOne string
	PasswordTwo string
}

func init() {

	// set messages
	gotea.App.Messages["input-password-one"] = inputPasswordOne
	gotea.App.Messages["input-password-two"] = inputPasswordTwo

	// create a seed for initial session state
	gotea.App.InitialSessionState = Model{
		PasswordOne: "",
		PasswordTwo: "",
	}

}
func inputPasswordOne(params gotea.MsgTag, s *gotea.Session) {
	// s.State.(Model).PasswordOne = params["input"].(string)
	// !!!!!!!!!!!!!!!!!!!!!!!!!
	// We need to find a much neater solution for this
	castModel := s.State.(Model)
	pointerToCastModel := &castModel
	pointerToCastModel.PasswordOne = params["input"].(string)
	s.State = *pointerToCastModel
}

func inputPasswordTwo(params gotea.MsgTag, s *gotea.Session) {
	castModel := s.State.(Model)
	pointerToCastModel := &castModel
	pointerToCastModel.PasswordTwo = params["input"].(string)
	s.State = *pointerToCastModel
}



// main starts the server
func main() {
	gotea.App.Start("../../dist")
}
