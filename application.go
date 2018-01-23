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
	ErrorTemplate *template.Template
	// holder for app-specific templates
	Templates *template.Template
	// map of function names -> function
	// used by incoming messages to call functions by name
	Messages map[string]func(MessageArguments, *Session)
	// map of message generators to be used in templates
	TemplateFuncs template.FuncMap
	// list of active sessions
	Sessions SessionStore
	// state seeder for new sessions
	// - injected by the app
	NewSession func() Session
	// top level configuration
	Config Config
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

// parseTemplates will parse app specific templates
// this is called AFTER all the init functions
// because the app.TemplateFuncs must be set up first
func (app *Application) parseTemplates() {
	// your app templates should go in the folder 'templates'
	// in the root directory of your app
	app.Templates = template.Must(template.New("main").Funcs(app.TemplateFuncs).ParseGlob("templates/*.html"))
}

// Start runs the application server
func (app Application) Start(distDirectory string) {
	// parse templates here
	// because template function map must be set up already
	App.parseTemplates()

	fs := http.FileServer(http.Dir(distDirectory))
	http.HandleFunc("/server", handler)
	http.Handle("/", fs)
	log.Println("Staring gotea app server...")
	http.ListenAndServe(fmt.Sprintf(":%v", app.Config.AppPort), nil)
}

// FUTURE ROUTING CODE - LEAVE FOR NOW

// TODO:
// USe this code for routing
// http.HandleFunc("/server", handler)
// http.HandleFunc("/", renderIndex)
// log.Println("Staring gotea app server...")
// http.ListenAndServe(fmt.Sprintf(":%v", app.Config.AppPort), nil)

// func renderIndex(w http.ResponseWriter, r *http.Request) {

// 	templateData := struct {
// 		AppTitle string
// 	}{
// 		"gotea App",
// 	}

// 	App.IndexTemplate.Execute(w, templateData)

// }
