package main

import (
	"encoding/json"
	"fmt"

	gt "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/msg"
	a "github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
)

type Form struct {
	Options           []string
	TextInput         string
	TextboxInput      string
	SelectInput       string
	IntInput          int
	FloatInput        float64
	CheckboxInput     bool
	RadioTextInput    string
	MultipleTextInput []string
}

func (f Form) String() string {
	// Encode to JSON with newlines and indents
	b, _ := json.MarshalIndent(f, "", "\n")
	return string(b)
}

func (f Form) renderValues() h.Element {
	return h.Div(
		a.Attrs(),
		h.H2(a.Attrs(), h.Text("Form Values")),
		h.Ul(
			a.Attrs(),
			h.Li(a.Attrs(), h.Text(fmt.Sprintf("TextInput: %s", f.TextInput))),
			h.Li(a.Attrs(), h.Text(fmt.Sprintf("TextboxInput: %s", f.TextboxInput))),
			h.Li(a.Attrs(), h.Text(fmt.Sprintf("SelectInput: %s", f.SelectInput))),
			h.Li(a.Attrs(), h.Text(fmt.Sprintf("IntInput: %d", f.IntInput))),
			h.Li(a.Attrs(), h.Text(fmt.Sprintf("FloatInput: %f", f.FloatInput))),
			h.Li(a.Attrs(), h.Text(fmt.Sprintf("CheckboxInput: %t", f.CheckboxInput))),
			h.Li(a.Attrs(), h.Text(fmt.Sprintf("RadioTextInput: %s", f.RadioTextInput))),
			h.Li(a.Attrs(), h.Text(fmt.Sprintf("MultipleTextInput: %v", f.MultipleTextInput))),
		),
	)
}

var formMessages gt.MessageMap = gt.MessageMap{
	"FORM_UPDATE": formUpdate,
}

func formUpdate(args json.RawMessage, s gt.State) gt.Response {
	state := model(s)

	if err := msg.DecodeStruct(args, &state.Form); err != nil {
		return gt.RespondWithError(err)
	}

	return gt.Respond()
}
