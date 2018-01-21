package main

import (
	s "strings"
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

func (model Model) PasswordErrors() []string {
	errors := make([]string, 0)

	if len(model.PasswordOne) < 8{
		errors = append(errors, "Password must be more than 8 characters long.")
	}
	if len(model.PasswordOne) > 10{
		errors = append(errors, "Password must be 10 characters or fewer")
	}
	if s.ToLower(model.PasswordOne) == model.PasswordOne {
		errors = append(errors, "Password must contain one upper-case letter")
	}
	if s.ToUpper(model.PasswordOne) == model.PasswordOne {
		errors = append(errors, "Password must contain one lower-case letter")
	}
	if !(s.ContainsAny(model.PasswordOne, "0123456789")) {
		errors = append(errors, "Password must contain a number")
	}
	symbols := "!@%^&*;,."
	if !(s.ContainsAny(model.PasswordOne, symbols)) {
		errors = append(errors, "Password must contain one of the following symbols: " + symbols)
	}
	return errors
}



// main starts the server
func main() {
	gotea.App.Start("../../dist")
}
