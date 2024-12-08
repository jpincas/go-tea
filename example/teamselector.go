package main

import (
	gt "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/example/tagselector"
)

var teamSelector = tagselector.Model{
	ComponentID:    "TEAMSELECTOR",
	AvailableTags:  []string{"Arsenal", "Man City", "Real Madrid"},
	NoMatchMessage: "No teams start with those letters. Start again!",
}

var teamSelectorMessages = map[string]gt.MessageHandler{
	tagselector.MsgSelectTag:         teamSelectorSelectTag,
	tagselector.MsgSearchInputUpdate: teamSelectorSearchInputUpdate,
	tagselector.MsgRemoveTag:         teamSelectorRemoveTag,
}

func teamSelectorSelectTag(m gt.Message, s gt.State) gt.Response {
	state := model(s)
	tag := m.ArgsToString()
	state.TeamSelector.SelectTag(tag)
	return gt.Respond()
}

func teamSelectorSearchInputUpdate(m gt.Message, s gt.State) gt.Response {
	state := model(s)
	input := m.ArgsToString()
	state.TeamSelector.SuggestTags(input)
	return gt.Respond()
}

func teamSelectorRemoveTag(m gt.Message, s gt.State) gt.Response {
	state := model(s)
	tag := m.ArgsToString()
	state.TeamSelector.RemoveTag(tag)
	return gt.Respond()
}
