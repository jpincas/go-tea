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
	"goteaValue":   SendMessageWithInputValue,
	"memberString": MemberString,
}

// Gotea message construction helpers

func SendMessage(msg string, args interface{}) template.JS {
	s := fmt.Sprintf("gotea.sendMessage('%s', %s)", msg, argsToJSON(args))
	return template.JS(s)
}

func SubmitForm(msg string, formID string) template.JS {
	s := fmt.Sprintf("gotea.submitForm('%s', '%s')", msg, formID)
	return template.JS(s)
}

func SendMessageWithInputValue(msg string, inputID string) template.JS {
	s := fmt.Sprintf("gotea.sendMessageWithValue('%s', '%s')", msg, inputID)
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

// SquashFuncMaps is a helper to combine multiple template FuncMaps
func SquashFuncMaps(funcMaps ...template.FuncMap) template.FuncMap {
	masterMap := template.FuncMap{}

	for _, thisMap := range funcMaps {
		for k, v := range thisMap {
			masterMap[k] = v
		}
	}

	return masterMap
}

// General Helpers

// MemberString returns whether a target string is a member of a slice of strings
func MemberString(target string, list []string) bool {
	for _, member := range list {
		if member == target {
			return true
		}
	}

	return false
}
