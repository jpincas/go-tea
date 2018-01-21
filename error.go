package gotea

import (
	"bytes"
	"html/template"
	"log"

	"github.com/gorilla/websocket"
)

func init() {
	// parse the framework error template
	App.ErrorTemplate = template.Must(template.New("error").Parse(`
		<h1>Whoops!</h1>
		<h2>There was a gotea runtime error</h2>
		<hr />
		<p>{{ .ErrorMessage }}</p>
		`))
}

// renderError renders a friendly error message in the browser
func renderError(conn *websocket.Conn, err error) {
	if conn == nil {
		// Oops! There's no socket connection
		log.Println("Could not render, no socket to render to")
		return
	}
	tpl := bytes.Buffer{}

	templateData := struct {
		ErrorMessage string
	}{
		err.Error(),
	}

	App.ErrorTemplate.Execute(&tpl, templateData)
	conn.WriteMessage(1, tpl.Bytes())
}
