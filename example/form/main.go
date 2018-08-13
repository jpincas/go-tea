package main

import (
	gotea "github.com/jpincas/go-tea"
)

type Model struct {
	People *[]Person
}

func init() {

	// session state seeder
	gotea.App.NewSession = func() gotea.Session {
		return gotea.Session{
			State: Model{
				People: &PeopleDB,
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

var PeopleDB []Person

func SubmitForm() gotea.Message {
	return gotea.Message{
		FuncCode: "SubmitForm",
		Func:     submitForm,
	}
}

func submitForm(formData gotea.MessageArguments, s *gotea.Session) (gotea.State, *gotea.Message) {
	// JSON objects are unmarshalled as map[string]interface{}
	// cast the messageArguments to that format
	personData := formData.(map[string]interface{})

	// now assign the Person struct by casting 1 by 1
	newPerson := Person{
		FirstName: personData["firstname"].(string),
		LastName:  personData["lastname"].(string),
	}

	PeopleDB = append(PeopleDB, newPerson)

	return s.State, nil
}

// MAIN

// main starts the server
func main() {
	gotea.App.Start("../../dist")
}
