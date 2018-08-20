package main

import (
	"encoding/json"

	gotea "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/components/tagselector"
)

var nameSelector = tagselector.Model{
	Component: gotea.Component{
		UniqueID: "NAMESELECTOR",
	},
	AvailableTags: []string{"Jon", "Allan", "Piotr"},
}

var teamSelector = tagselector.Model{
	Component: gotea.Component{
		UniqueID: "TEAMSELECTOR",
	},
	AvailableTags: []string{"Arsenal", "Man City", "Real Madrid"},
}

var nameSelectorMsgMap = gotea.MessageMap{
	nameSelector.UniqueMsg(tagselector.MsgSelectTag): nameSelectorSelectTag,
}

var teamSelectorMsgMap = gotea.MessageMap{
	teamSelector.UniqueMsg(tagselector.MsgSelectTag): teamSelectorSelectTag,
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

func teamSelectorSelectTag(args json.RawMessage, s gotea.State) (gotea.State, *gotea.Message, error) {
	state := s.(Model)

	var tag string
	if err := json.Unmarshal(args, &tag); err != nil {
		return state, nil, err
	}

	state.TeamSelector.SelectTag(tag)
	return state, nil, nil
}
