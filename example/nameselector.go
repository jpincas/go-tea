package main

import (
	gt "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/example/tagselector"
	a "github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
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

func nameSelectorSelectTag(m gt.Message, s gt.State) gt.Response {
	state := model(s)
	tag := m.ArgsToString()
	state.NameSelector.SelectTag(tag)
	return gt.Respond()
}

func nameSelectorRemoveTag(m gt.Message, s gt.State) gt.Response {
	state := model(s)
	tag := m.ArgsToString()
	state.NameSelector.RemoveTag(tag)
	return gt.Respond()
}

func nameSelectorSearchInputUpdate(m gt.Message, s gt.State) gt.Response {
	state := model(s)
	input := m.ArgsToString()
	state.NameSelector.SuggestTags(input)
	return gt.Respond()
}

func renderComponents(nameSelector, tagSelector tagselector.Model) h.Element {
	return h.Div(
		a.Attrs(),
		h.H2(a.Attrs(), h.Text("Components Demo")),
		h.P(a.Attrs(), h.Text("Two instantiations of the 'tag selector' component, running side-by-side")),
		h.Div(
			a.Attrs(),
			h.Div(
				a.Attrs(),
				h.H3(a.Attrs(), h.Text("Select a Name")),
				nameSelector.Render(),
			),
			h.Div(
				a.Attrs(),
				h.H3(a.Attrs(), h.Text("Select a Team")),
				tagSelector.Render(),
			),
		),
	)
}
