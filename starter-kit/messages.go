package main

import (
	gt "github.com/jpincas/go-tea"
)

// ============================================================================
// Counter Messages
// ============================================================================

var counterMessages = gt.MessageMap{
	"INCREMENT": incrementCounter,
}

func incrementCounter(m gt.Message, s gt.State) gt.Response {
	// ArgsToInt() safely converts JSON number (float64) to int
	n := m.ArgsToInt()
	state := model(s)
	state.Counter += n
	return gt.Respond()
}

// ============================================================================
// Form Messages
// ============================================================================

var formMessages = gt.MessageMap{
	"FORM_UPDATE": formUpdate,
}

func formUpdate(m gt.Message, s gt.State) gt.Response {
	state := model(s)
	// MustDecodeArgs decodes the JSON args directly into the struct
	m.MustDecodeArgs(&state.Form)
	return gt.Respond()
}

// ============================================================================
// Tag Selector Messages (Component)
// ============================================================================

// These are the "base" message handlers. They get namespaced via UniqueMsgMap.
var tagSelectorMessages = gt.MessageMap{
	"SELECT": tagSelect,
	"REMOVE": tagRemove,
	"SEARCH": tagSearch,
}

func tagSelect(m gt.Message, s gt.State) gt.Response {
	state := model(s)
	tag := m.ArgsToString()
	state.Selector.SelectTag(tag)
	return gt.Respond()
}

func tagRemove(m gt.Message, s gt.State) gt.Response {
	state := model(s)
	tag := m.ArgsToString()
	state.Selector.RemoveTag(tag)
	return gt.Respond()
}

func tagSearch(m gt.Message, s gt.State) gt.Response {
	state := model(s)
	query := m.ArgsToString()
	state.Selector.Search(query)
	return gt.Respond()
}
