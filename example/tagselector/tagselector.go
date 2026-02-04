package tagselector

import (
	"strings"

	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
)

// Message Constants
const (
	MsgSelectTag         = "TAG.SELECT"
	MsgRemoveTag         = "TAG.REMOVE"
	MsgSearchInputUpdate = "SEARCHINPUT.UPDATE"
)

type Model struct {
	gt.ComponentID
	SearchInput    string
	AvailableTags  []string
	SuggestedTags  []string
	SelectedTags   []string
	NoMatchMessage string
}

func (m *Model) SelectTag(newTag string) {
	m.SelectedTags = append(m.SelectedTags, newTag)
	m.SuggestTags("")
}

func (m *Model) RemoveTag(tagToRemove string) {
	var newSelectedTags []string
	for _, tag := range m.SelectedTags {
		if tag != tagToRemove {
			newSelectedTags = append(newSelectedTags, tag)
		}
	}

	m.SelectedTags = newSelectedTags
}

func (m *Model) SuggestTags(input string) {
	m.SearchInput = input

	if input == "" {
		m.SuggestedTags = []string{}
	} else {
		m.SuggestedTags = matchTags(m.AvailableTags, input)
	}
}

func (m Model) ShowNoMatchMessage() bool {
	return m.SearchInput != "" && len(m.SuggestedTags) == 0
}

func matchTags(list []string, toMatch string) (matched []string) {
	for _, tag := range list {
		if strings.Contains(strings.ToLower(tag), strings.ToLower(toMatch)) {
			matched = append(matched, tag)
		}
	}

	return
}

func (selector Model) Render() h.Element {
	msgSelectTag := selector.UniqueMsg("TAG.SELECT")
	msgUpdateSearchInput := selector.UniqueMsg("SEARCHINPUT.UPDATE")
	msgRemoveTag := selector.UniqueMsg("TAG.REMOVE")
	searchInputID := selector.UniqueID("search-input")

	return h.Div(
		a.Attrs(a.Class("bg-white p-5 rounded-xl border-2 border-stone-900 shadow-brutal-sm space-y-4")),
		// Search input
		h.Div(
			a.Attrs(a.Class("space-y-3")),
			h.Input(
				a.Attrs(
					a.Id(searchInputID),
					a.Class("w-full px-4 py-2.5 rounded-lg text-stone-900 placeholder-stone-400"),
					a.Type("text"),
					a.Placeholder("Start typing to search..."),
					a.Value(selector.SearchInput),
					a.OnKeyUp(gt.SendBasicMessageWithValueFromInput(msgUpdateSearchInput, searchInputID)),
				),
			),
			// Suggestions list
			h.Ul(
				a.Attrs(a.Class("space-y-1 max-h-32 overflow-y-auto")),
				func() []h.Element {
					var elements []h.Element
					for _, tag := range selector.SuggestedTags {
						elements = append(elements, h.Li(
							a.Attrs(
								a.Class("px-4 py-2 bg-stone-50 hover:bg-emerald-100 rounded-lg cursor-pointer transition-colors font-medium text-stone-700 hover:text-emerald-800 border border-transparent hover:border-emerald-300"),
								a.OnClick(gt.SendBasicMessage(msgSelectTag, tag)),
							),
							h.Text(tag),
						))
					}
					return elements
				}()...,
			),
			// No match message
			func() h.Element {
				if selector.ShowNoMatchMessage() {
					return h.P(a.Attrs(a.Class("text-sm text-rose-600 bg-rose-50 px-3 py-2 rounded-lg")), h.Text(selector.NoMatchMessage))
				}
				return h.Nothing(a.Attrs())
			}(),
		),
		// Selected tags
		h.Div(
			a.Attrs(a.Class("border-t-2 border-dashed border-stone-200 pt-4")),
			h.H4(a.Attrs(a.Class("text-xs font-semibold text-stone-500 uppercase tracking-wide mb-3")), h.Text("Selected:")),
			h.Ul(
				a.Attrs(a.Class("flex flex-wrap gap-2")),
				func() []h.Element {
					if len(selector.SelectedTags) == 0 {
						return []h.Element{
							h.Li(a.Attrs(a.Class("text-sm text-stone-400 italic")), h.Text("None selected")),
						}
					}
					var elements []h.Element
					for _, tag := range selector.SelectedTags {
						elements = append(elements, h.Li(
							a.Attrs(
								a.Class("inline-flex items-center gap-1 px-3 py-1.5 rounded-full text-sm font-semibold bg-emerald-100 text-emerald-800 border-2 border-emerald-300 cursor-pointer hover:bg-rose-100 hover:text-rose-800 hover:border-rose-300 transition-colors"),
								a.OnClick(gt.SendBasicMessage(msgRemoveTag, tag)),
							),
							h.Text(tag),
							h.Span(a.Attrs(a.Class("text-xs opacity-60")), h.Text("✕")),
						))
					}
					return elements
				}()...,
			),
		),
	)
}
