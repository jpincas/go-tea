package gotea

import (
	"fmt"
	"html/template"
)

// TemplateHelpers are functions that can be used in Go templates
// to conveniently generate HTML/JS for emitting go-tea events.
// Either parse this map directly into your templates,
// or combine with your existing funcMap
var TemplateHelpers = template.FuncMap{
	"goteaMessage": SendMessage,
	"goteaLink":    Link,
	"goteaForm":    SubmitForm,
}

func SendMessage(msg string, args interface{}) template.JS {
	s := fmt.Sprintf("gotea.sendMessage('%s', %s)", msg, argsToJSON(args))
	return template.JS(s)
}

func SubmitForm(msg string, formID string) template.JS {
	s := fmt.Sprintf("gotea.submitForm('%s', '%s')", msg, formID)
	return template.JS(s)
}

func Link(href, text, extraClasses string) template.HTML {
	s := fmt.Sprintf("<a class='gotea-link %s' href='%s'>%s</a>", extraClasses, href, text)
	return template.HTML(s)
}

func argsToJSON(args interface{}) string {
	switch args.(type) {
	case int:
		return fmt.Sprintf("%v", args)
	case float32:
		return fmt.Sprintf("%v", args)
	case float64:
		return fmt.Sprintf("%v", args)
	case bool:
		trueOrFalse := args.(bool)
		if trueOrFalse {
			return "true"
		} else {
			return "false"
		}
	default:
		// Everything else as a string
		return fmt.Sprintf(`'"%v"'`, args)
	}
}
