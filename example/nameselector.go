package main

import (
	"encoding/json"

	gt "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/example/tagselector"
)

// Name Selector (instantiation of Tag Selector)

var nameSelector = tagselector.Model{
	ComponentID:    "NAMESELECTOR",
	AvailableTags:  []string{"Jon", "Allan", "Piotr"},
	NoMatchMessage: "No matching names found. Please try again",
}

var nameSelectorMessages = map[string]gt.MessageHandler{
	tagselector.MsgSelectTag:         nameSelectorSelectTag,
	tagselector.MsgSearchInputUpdate: nameSelectorSearchInputUpdate,
	tagselector.MsgRemoveTag:         nameSelectorRemoveTag,
}

func nameSelectorSelectTag(args json.RawMessage, s gt.State) gt.Response {
	state := s.(*Model)

	var tag string
	if err := json.Unmarshal(args, &tag); err != nil {
		return gt.RespondWithError(state, err)
	}

	state.NameSelector.SelectTag(tag)
	return gt.Respond(state)
}

func nameSelectorRemoveTag(args json.RawMessage, s gt.State) gt.Response {
	state := s.(*Model)

	var tag string
	if err := json.Unmarshal(args, &tag); err != nil {
		return gt.RespondWithError(state, err)
	}

	state.NameSelector.RemoveTag(tag)
	return gt.Respond(state)
}

func nameSelectorSearchInputUpdate(args json.RawMessage, s gt.State) gt.Response {
	state := s.(*Model)

	var input string
	if err := json.Unmarshal(args, &input); err != nil {
		return gt.RespondWithError(state, err)
	}

	state.NameSelector.SuggestTags(input)
	return gt.Respond(state)
}
