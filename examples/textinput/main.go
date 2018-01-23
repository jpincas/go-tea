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

	gotea.App.Config.AppPort = 8081
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

type PasswordConstraint struct {
	Satisfied bool
	Description string
}

func (model Model) PasswordErrors() []PasswordConstraint {
	errors := make([]PasswordConstraint, 0)
	errors = append(errors, 
		PasswordConstraint{
			Satisfied: len(model.PasswordOne) >= 8,
			Description: "Password must be more than 8 characters long.",
		})

	errors = append(errors, 
		PasswordConstraint{
			Satisfied: len(model.PasswordOne) <= 10,
			Description: "Password must be 10 characters or fewer",
		})
		
	errors = append(errors, 
		PasswordConstraint{
			Satisfied: s.ToLower(model.PasswordOne) != model.PasswordOne,
			Description: "Password must contain one upper-case letter",
		})

	errors = append(errors, 
		PasswordConstraint{
			Satisfied: s.ToUpper(model.PasswordOne) != model.PasswordOne,
			Description: "Password must contain one lower-case letter",
		})

	errors = append(errors, 
		PasswordConstraint{
			Satisfied: s.ContainsAny(model.PasswordOne, "0123456789"),
			Description: "Password must contain a number",
		})

	symbols := "!@%^&*;,."
	errors = append(errors, 
		PasswordConstraint{
			Satisfied: s.ContainsAny(model.PasswordOne, symbols),
			Description: "Password must contain one of the following symbols: " + symbols,
		})

	return errors
}



// main starts the server
func main() {
	gotea.App.Start("../../dist")
}
