package tagselector

import (
	"strings"

	gotea "github.com/jpincas/go-tea"
)

// Message Constants
const (
	MsgSelectTag         = "TAG.SELECT"
	MsgRemoveTag         = "TAG.REMOVE"
	MsgSearchInputUpdate = "SEARCHINPUT.UPDATE"
)

type Model struct {
	gotea.ComponentID
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
