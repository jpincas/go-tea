package main

import (
	"encoding/json"

	gotea "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/components/tagselector"
)

// Name Selector (instantiation of Tag Selector)

var nameSelector = tagselector.Model{
	ComponentID:    "NAMESELECTOR",
	AvailableTags:  []string{"Jon", "Allan", "Piotr"},
	NoMatchMessage: "No matching names found. Please try again",
}

var nameSelectorMessages = map[string]gotea.MessageHandler{
	tagselector.MsgSelectTag:         nameSelectorSelectTag,
	tagselector.MsgSearchInputUpdate: nameSelectorSearchInputUpdate,
	tagselector.MsgRemoveTag:         nameSelectorRemoveTag,
}

func nameSelectorSelectTag(args json.RawMessage, s gotea.State) (gotea.State, *gotea.Message, error) {
	state := s.(Model)

	var tag string
	if err := json.Unmarshal(args, &tag); err != nil {
		return state, nil, err
	}

	state.NameSelector.SelectTag(tag)
	return state, nil, nil
}

func nameSelectorRemoveTag(args json.RawMessage, s gotea.State) (gotea.State, *gotea.Message, error) {
	state := s.(Model)

	var tag string
	if err := json.Unmarshal(args, &tag); err != nil {
		return state, nil, err
	}

	state.NameSelector.RemoveTag(tag)
	return state, nil, nil
}

func nameSelectorSearchInputUpdate(args json.RawMessage, s gotea.State) (gotea.State, *gotea.Message, error) {
	state := s.(Model)

	var input string
	if err := json.Unmarshal(args, &input); err != nil {
		return state, nil, err
	}

	state.NameSelector.SuggestTags(input)
	return state, nil, nil
}
