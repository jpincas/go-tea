package main

import (
	"encoding/json"

	gotea "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/components/tagselector"
)

var teamSelector = tagselector.Model{
	ComponentID:    "TEAMSELECTOR",
	AvailableTags:  []string{"Arsenal", "Man City", "Real Madrid"},
	NoMatchMessage: "No teams start with those letters. Start again!",
}

var teamSelectorMessages = map[string]gotea.MessageHandler{
	tagselector.MsgSelectTag:         teamSelectorSelectTag,
	tagselector.MsgSearchInputUpdate: teamSelectorSearchInputUpdate,
	tagselector.MsgRemoveTag:         teamSelectorRemoveTag,
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

func teamSelectorSearchInputUpdate(args json.RawMessage, s gotea.State) (gotea.State, *gotea.Message, error) {
	state := s.(Model)

	var input string
	if err := json.Unmarshal(args, &input); err != nil {
		return state, nil, err
	}

	state.TeamSelector.SuggestTags(input)
	return state, nil, nil
}

func teamSelectorRemoveTag(args json.RawMessage, s gotea.State) (gotea.State, *gotea.Message, error) {
	state := s.(Model)

	var tag string
	if err := json.Unmarshal(args, &tag); err != nil {
		return state, nil, err
	}

	state.TeamSelector.RemoveTag(tag)
	return state, nil, nil
}
