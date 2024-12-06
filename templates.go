package gotea

import (
	"encoding/json"
	"fmt"
)

// Gotea message construction helpers
var TemplateFuncs = map[string]interface{}{
	"goteaMessage":  SendMessage,
	"goteaMessage_": SendMessageNoArgs,
	"goteaForm":     SubmitForm,
	"goteaValue":    SendMessageWithInputValue,
	"memberString":  MemberString,
}

func SendMessage(msg string, args interface{}) string {
	s := fmt.Sprintf(`gotea.sendMessage("%s", %s)`, msg, argsToJSON(args))
	return s
}

func SendMessageNoArgs(msg string) string {
	s := fmt.Sprintf(`gotea.sendMessage("%s", null)`, msg)
	return s
}

func SubmitForm(msg string, formID string) string {
	s := fmt.Sprintf(`gotea.submitForm("%s", "%s")`, msg, formID)
	return s
}

func SendMessageWithInputValue(msg string, inputID string) string {
	s := fmt.Sprintf(`gotea.sendMessageWithValue("%s", "%s")`, msg, inputID)
	return s
}

func argsToJSON(args interface{}) string {
	if _, isString := args.(string); isString {
		return fmt.Sprintf(`JSON.stringify("%s")`, args)
	}

	jsonData, err := json.Marshal(args)
	if err != nil {
		return "null"
	}
	jsonStr := string(jsonData)
	return jsonStr
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
