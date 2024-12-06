package main

import (
	"encoding/json"

	gt "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/example/tagselector"
	"github.com/jpincas/go-tea/msg"
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

func renderComponents(nameSelector, tagselector tagselector.Model) h.Element {
	return h.Div(
		a.Attrs(),
		h.H2(a.Attrs(), h.Text("Components Demo")),
		h.P(a.Attrs(), h.Text("Two instantiations of the 'tag selector' component, running side-by-side")),
		h.Div(
			a.Attrs(),
			h.Div(
				a.Attrs(),
				h.H3(a.Attrs(), h.Text("Select a Name")),
				renderTagSelector(nameSelector),
			),
			h.Div(
				a.Attrs(),
				h.H3(a.Attrs(), h.Text("Select a Team")),
				renderTagSelector(tagselector),
			),
		),
	)
}

func renderTagSelector(selector tagselector.Model) h.Element {
	msgSelectTag := selector.UniqueMsg("TAG.SELECT")
	msgUpdateSearchInput := selector.UniqueMsg("SEARCHINPUT.UPDATE")
	msgRemoveTag := selector.UniqueMsg("TAG.REMOVE")
	searchInputID := selector.UniqueID("search-input")

	return h.Div(
		a.Attrs(a.Class("tagselector tagselector-container")),
		h.Div(
			a.Attrs(a.Class("tagselector-suggestedtags")),
			h.Input(
				a.Attrs(
					a.Id(searchInputID),
					a.Class("input"),
					a.Type("text"),
					a.Placeholder("Start typing to see tags"),
					a.Value(selector.SearchInput),
					a.OnKeyUp(gt.SendMessageWithInputValue(msgUpdateSearchInput, searchInputID)),
				),
			),
			h.Ul(
				a.Attrs(a.Class("tagselector-tagslist tagselector-suggestedtagslist")),
				func() []h.Element {
					var elements []h.Element
					for _, tag := range selector.SuggestedTags {
						elements = append(elements, h.Li(
							a.Attrs(a.Class("tagselector-tag tagselector-suggestedtag"), a.OnClick(gt.SendMessage(msgSelectTag, tag))),
							h.Text(tag),
						))
					}
					return elements
				}()...,
			),
			func() h.Element {
				if selector.ShowNoMatchMessage() {
					return h.P(a.Attrs(), h.Text(selector.NoMatchMessage))
				}
				return h.Nothing(a.Attrs())
			}(),
		),
		h.Div(
			a.Attrs(a.Class("tagselector-selectedtags")),
			h.H4(a.Attrs(a.Class("tagselector-selectedtagstitle")), h.Text("Selected Tags:")),
			h.Ul(
				a.Attrs(a.Class("tagselector-tagslist tagselector-selectedtagslist")),
				func() []h.Element {
					var elements []h.Element
					for _, tag := range selector.SelectedTags {
						elements = append(elements, h.Li(
							a.Attrs(a.Class("tagselector-tag tagselector-selectedtag"), a.OnClick(gt.SendMessage(msgRemoveTag, tag))),
							h.Text(tag),
						))
					}
					return elements
				}()...,
			),
		),
	)
}
