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
	s := fmt.Sprintf("gotea.sendMessage('%s', %v)", msg, args)
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
