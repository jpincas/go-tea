package main

import (
	"encoding/json"

	gt "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/msg"
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
