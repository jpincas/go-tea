package gotea

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func init() {
	App.IndexTemplate = template.Must(template.New("index").Parse(index))
}

// Config stores configuration options for the application
type Config struct {
	AppPort int
}

// Application is the top-level representation of your application
type Application struct {
	// template for rendering index html and js
	IndexTemplate *template.Template
	// template for rendering framework errors
	ErrorTemplate       *template.Template
	Templates           *template.Template
	Messages            map[string]func(MsgTag, *Session)
	Sessions            SessionStore
	InitialSessionState State
	Config              Config
}

// App instantiates a new application and makes it globally available
var App Application

// Broadcast re-renders every active session
func (app Application) Broadcast() {
	for _, session := range app.Sessions {
		if session.Conn != nil {
			session.render()
		}
	}
}

// Start runs the application server
func (app Application) Start(distDirectory string) {

	fs := http.FileServer(http.Dir(distDirectory))
	http.HandleFunc("/server", handler)
	http.Handle("/", fs)
	log.Println("Staring gotea app server...")
	http.ListenAndServe(fmt.Sprintf(":%v", app.Config.AppPort), nil)

	// TODO:
	// USe this code for routing
	// http.HandleFunc("/server", handler)
	// http.HandleFunc("/", renderIndex)
	// log.Println("Staring gotea app server...")
	// http.ListenAndServe(fmt.Sprintf(":%v", app.Config.AppPort), nil)
}

// func renderIndex(w http.ResponseWriter, r *http.Request) {

// 	templateData := struct {
// 		AppTitle string
// 	}{
// 		"gotea App",
// 	}

// 	App.IndexTemplate.Execute(w, templateData)

// }
