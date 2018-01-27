package main

import (
	"log"

	gotea "github.com/jpincas/go-tea"
)

type Model struct {
	Test string
}

func init() {

	// message function map
	gotea.App.Messages["SubmitForm"] = submitForm

	// session state seeder
	gotea.App.NewSession = func() gotea.Session {
		return gotea.Session{
			State: Model{
				Test: "test",
			},
		}
	}

	// main view renderer
	gotea.App.RenderView = WriteMain

}

// APP SPECFIC

type Person struct {
	FirstName, LastName string
}

func SubmitForm() gotea.Message {
	return gotea.Message{
		FuncCode: "SubmitForm",
	}
}

func submitForm(formData gotea.MessageArguments, s *gotea.Session) (gotea.State, *gotea.Message) {
	person := formData.(Person)
	log.Println(person)
	return s.State, nil
}

// MAIN

// main starts the server
func main() {
	gotea.App.Start("../../dist")
}
