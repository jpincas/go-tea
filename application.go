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
	Templates     *template.Template
	Messages      map[string]func(MessageArguments, *Session)
	TemplateFuncs template.FuncMap
	Sessions      SessionStore
	NewSession    func() Session
	Config        Config
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

func (app *Application) parseTemplates() {
	// parse app specific templates
	// therefore, your app templates should go in the folder 'templates'
	// in the root directory of your app
	app.Templates = template.Must(template.New("main").Funcs(app.TemplateFuncs).ParseGlob("templates/*.html"))
}

// Start runs the application server
func (app Application) Start(distDirectory string) {

	// must leave this until here
	// relies on stuff that happens in init
	App.parseTemplates()

	fs := http.FileServer(http.Dir(distDirectory))
	http.HandleFunc("/server", handler)
	http.Handle("/", fs)
	log.Println("Staring gotea app server...")
	http.ListenAndServe(fmt.Sprintf(":%v", app.Config.AppPort), nil)

}

// FUTURE ROUTING CODE

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
