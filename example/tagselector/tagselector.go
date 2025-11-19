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
		a.Attrs(a.Class("bg-white p-6 rounded-lg shadow-md space-y-6")),
		h.Div(
			a.Attrs(a.Class("space-y-4")),
			h.Input(
				a.Attrs(
					a.Id(searchInputID),
					a.Class("w-full px-4 py-2 border border-gray-300 rounded-md focus:ring-indigo-500 focus:border-indigo-500 shadow-sm"),
					a.Type("text"),
					a.Placeholder("Start typing to see tags"),
					a.Value(selector.SearchInput),
					a.OnKeyUp(gt.SendBasicMessageWithValueFromInput(msgUpdateSearchInput, searchInputID)),
				),
			),
			h.Ul(
				a.Attrs(a.Class("space-y-2 max-h-40 overflow-y-auto")),
				func() []h.Element {
					var elements []h.Element
					for _, tag := range selector.SuggestedTags {
						elements = append(elements, h.Li(
							a.Attrs(a.Class("px-4 py-2 bg-gray-50 hover:bg-indigo-50 rounded-md cursor-pointer transition-colors duration-150"), a.OnClick(gt.SendBasicMessage(msgSelectTag, tag))),
							h.Text(tag),
						))
					}
					return elements
				}()...,
			),
			func() h.Element {
				if selector.ShowNoMatchMessage() {
					return h.P(a.Attrs(a.Class("text-sm text-red-500")), h.Text(selector.NoMatchMessage))
				}
				return h.Nothing(a.Attrs())
			}(),
		),
		h.Div(
			a.Attrs(a.Class("border-t border-gray-200 pt-4")),
			h.H4(a.Attrs(a.Class("text-sm font-medium text-gray-500 mb-3")), h.Text("Selected Tags:")),
			h.Ul(
				a.Attrs(a.Class("flex flex-wrap gap-2")),
				func() []h.Element {
					var elements []h.Element
					for _, tag := range selector.SelectedTags {
						elements = append(elements, h.Li(
							a.Attrs(a.Class("inline-flex items-center px-3 py-1 rounded-full text-sm font-medium bg-indigo-100 text-indigo-800 cursor-pointer hover:bg-indigo-200 transition-colors duration-150"), a.OnClick(gt.SendBasicMessage(msgRemoveTag, tag))),
							h.Text(tag),
							h.Span(a.Attrs(a.Class("ml-2 text-indigo-500 hover:text-indigo-700")), h.Text("×")),
						))
					}
					return elements
				}()...,
			),
		),
	)
}
