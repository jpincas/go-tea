package main

import (
	"encoding/json"

	gt "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/example/tagselector"
	"github.com/jpincas/go-tea/msg"
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
	state := model(s)

	tag, err := msg.DecodeString(args)
	if err != nil {
		return gt.RespondWithError(err)
	}

	state.NameSelector.SelectTag(tag)
	return gt.Respond()
}

func nameSelectorRemoveTag(args json.RawMessage, s gt.State) gt.Response {
	state := model(s)

	tag, err := msg.DecodeString(args)
	if err != nil {
		return gt.RespondWithError(err)
	}

	state.NameSelector.RemoveTag(tag)
	return gt.Respond()
}

func nameSelectorSearchInputUpdate(args json.RawMessage, s gt.State) gt.Response {
	state := model(s)

	input, err := msg.DecodeString(args)
	if err != nil {
		return gt.RespondWithError(err)
	}

	state.NameSelector.SuggestTags(input)
	return gt.Respond()
}
