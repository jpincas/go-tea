package gotea

import (
	"fmt"
	"html/template"
)

// Gotea message construction helpers
var TemplateFuncs = map[string]interface{}{
	"goteaMessage":  SendMessage,
	"goteaMessage_": SendMessageNoArgs,
	"goteaForm":     SubmitForm,
	"goteaValue":    SendMessageWithInputValue,
	"memberString":  MemberString,
}

func SendMessage(msg string, args interface{}) template.JS {
	s := fmt.Sprintf("gotea.sendMessage('%s', %s)", msg, argsToJSON(args))
	return template.JS(s)
}

func SendMessageNoArgs(msg string) template.JS {
	s := fmt.Sprintf("gotea.sendMessage('%s', null)", msg)
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
