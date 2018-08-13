package main

import (
	gotea "github.com/jpincas/go-tea"
)

type Model struct {
	Route string
}

func init() {

	// session state seeder
	gotea.App.NewSession = func() gotea.Session {
		return gotea.Session{
			State: Model{
				Route: "/",
			},
		}
	}

	// main view renderer
	gotea.App.RenderView = WriteMain

}

// main starts the server
func main() {
	gotea.App.Start("../../dist")
}
