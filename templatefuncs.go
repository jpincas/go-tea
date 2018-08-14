package gotea

import (
	"fmt"
	"html/template"
)

var TemplateHelpers = template.FuncMap{
	"sendMessage": SendMessage,
}

func SendMessage(msg string, args interface{}) template.JS {
	s := fmt.Sprintf("gotea.sendMessage('%s', %v)", msg, args)
	return template.JS(s)
}
