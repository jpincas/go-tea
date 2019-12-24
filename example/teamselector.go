package main

import (
	"encoding/json"

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

func teamSelectorSelectTag(args json.RawMessage, s gt.State) gt.Response {
	state := s.(*Model)

	var tag string
	if err := json.Unmarshal(args, &tag); err != nil {
		return gt.RespondWithError(state, err)
	}

	state.TeamSelector.SelectTag(tag)
	return gt.Respond(state)
}

func teamSelectorSearchInputUpdate(args json.RawMessage, s gt.State) gt.Response {
	state := s.(*Model)

	var input string
	if err := json.Unmarshal(args, &input); err != nil {
		return gt.RespondWithError(state, err)
	}

	state.TeamSelector.SuggestTags(input)
	return gt.Respond(state)
}

func teamSelectorRemoveTag(args json.RawMessage, s gt.State) gt.Response {
	state := s.(*Model)

	var tag string
	if err := json.Unmarshal(args, &tag); err != nil {
		return gt.RespondWithError(state, err)
	}

	state.TeamSelector.RemoveTag(tag)
	return gt.Respond(state)
}
