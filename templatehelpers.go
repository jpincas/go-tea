package gotea

import (
	"encoding/json"
	"fmt"
)

// These functions are intended to be used by templates
// when constructing the JS messages that are sent from the browser

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
