package gotea

import (
	"fmt"
)

// These functions are intended to be used by templates
// when constructing the JS messages that are sent from the browser

const (
	// These function names are used in the JS code that is sent to the browser
	sendMessageFuncName              = "sendMessage"
	sendMessageWithValueFuncName     = "sendMessageWithValueFromInput"
	sendMessageWithThisValueFuncName = "sendMessageWithValueFromThisInput"
	updateFormFuncName               = "updateFormState"
)

func constructFuncName(funcName string) string {
	return fmt.Sprintf("gotea.%s", funcName)
}

// Some higher level constructors

func SendBasicMessage(msg string, args any) string {
	return SendMessage(
		Message{
			Message:   msg,
			Arguments: args,
		},
	)
}

func SendBasicMessageNoArgs(msg string) string {
	return SendMessage(
		Message{
			Message: msg,
		},
	)
}

func SendBasicMessageWithValueFromInput(msg string, inputID string) string {
	return SendMessageWithValueFromInput(
		Message{
			Message: msg,
		},
		inputID,
	)
}

func BasicUpdateForm(msg string, formID string) string {
	return UpdateFormState(
		Message{
			Message: msg,
		},
		formID,
	)
}

// Some lower level constructors

func SendMessage(m Message) string {
	return fmt.Sprintf(`%s(%s)`, constructFuncName(sendMessageFuncName), m.toJson())
}

func SendMessageWithValueFromInput(m Message, inputID string) string {
	return fmt.Sprintf(`%s(%s, "%s")`, constructFuncName(sendMessageWithValueFuncName), m.toJson(), inputID)
}

func SendMessageWithValueFromThisInput(m Message) string {
	return fmt.Sprintf(`%s(%s)`, constructFuncName(sendMessageWithThisValueFuncName), m.toJson())
}

func UpdateFormState(m Message, formID string) string {
	return fmt.Sprintf(`%s(%s, "%s")`, constructFuncName(updateFormFuncName), m.toJson(), formID)
}

// func argsToJSON(args interface{}) string {
// 	if _, isString := args.(string); isString {
// 		return fmt.Sprintf(`JSON.stringify("%s")`, args)
// 	}

// 	jsonData, err := json.Marshal(args)
// 	if err != nil {
// 		return "null"
// 	}
// 	jsonStr := string(jsonData)
// 	return jsonStr
// }
