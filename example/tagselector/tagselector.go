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
