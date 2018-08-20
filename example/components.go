package main

import (
	"encoding/json"

	gotea "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/components/tagselector"
)

// Name Selector (instantiation of Tag Selector)

var nameSelector = tagselector.Model{
	ComponentID:   "NAMESELECTOR",
	AvailableTags: []string{"Jon", "Allan", "Piotr"},
}

var nameSelectorMessages = map[string]gotea.MessageHandler{
	tagselector.MsgSelectTag: nameSelectorSelectTag,
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

// Team Selector (instantiation of Tag Selector)

var teamSelector = tagselector.Model{
	ComponentID:   "TEAMSELECTOR",
	AvailableTags: []string{"Arsenal", "Man City", "Real Madrid"},
}

var teamSelectorMessages = map[string]gotea.MessageHandler{
	tagselector.MsgSelectTag: teamSelectorSelectTag,
}

func teamSelectorSelectTag(args json.RawMessage, s gotea.State) (gotea.State, *gotea.Message, error) {
	state := s.(Model)

	var tag string
	if err := json.Unmarshal(args, &tag); err != nil {
		return state, nil, err
	}

	state.TeamSelector.SelectTag(tag)
	return state, nil, nil
}
