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

	gotea.RegisterMessages(
		InputPasswordOne,
		InputPasswordTwo,
	)

	// function that returns a new session
	gotea.App.NewSession = func() gotea.Session {
		return gotea.Session{
			State: Model{
				PasswordOne: "",
				PasswordTwo: "",
			},
		}
	}
	gotea.App.Config.AppPort = 8081
}

func InputPasswordOne(args gotea.MessageArguments) gotea.Message {
	return gotea.NewMsg(inputPasswordOne, args)
}

func InputPasswordTwo(args gotea.MessageArguments) gotea.Message {
	return gotea.NewMsg(inputPasswordTwo, args)
}

func inputPasswordOne(args gotea.MessageArguments, s *gotea.Session) {
	// !!!!!!!!!!!!!!!!!!!!!!!!!
	// We need to find a much neater solution for this
	castModel := s.State.(Model)
	pointerToCastModel := &castModel
	arguments := args.([]interface {})
	pointerToCastModel.PasswordOne = arguments[0].(string)
	s.State = *pointerToCastModel
}

func inputPasswordTwo(args gotea.MessageArguments, s *gotea.Session) {
	// !!!!!!!!!!!!!!!!!!!!!!!!!
	// We need to find a much neater solution for this
	castModel := s.State.(Model)
	pointerToCastModel := &castModel
	arguments := args.([]interface {})
	pointerToCastModel.PasswordTwo = arguments[0].(string)
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
