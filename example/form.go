package main

import (
	"encoding/json"

	gotea "github.com/jpincas/go-tea"
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

var formMessages gotea.MessageMap = gotea.MessageMap{
	"FORM_UPDATE": formUpdate,
}

func formUpdate(args json.RawMessage, s gotea.State) (gotea.State, *gotea.Message, error) {
	state := s.(Model)

	if err := json.Unmarshal(args, &state.Form); err != nil {
		return state, nil, err
	}

	return state, nil, nil
}
