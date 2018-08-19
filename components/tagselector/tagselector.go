package tagselector

import (
	"encoding/json"

	"github.com/jpincas/go-tea"
)

type Model struct {
	SearchInput   string
	AvailableTags []string
	SuggestedTags []string
	SelectedTags  []string
}

var Messages = gotea.MessageMap{
	"TAGSELECTOR_SELECTTAG":          Handler,
	"TAGSELECTOR_SEARCHINPUT_UPDATE": Handler,
}

func Handler(args json.RawMessage, s gotea.State) (gotea.State, *gotea.Message, error) {
	return s, nil, nil
}

func (m *Model) SelectTag(newTag string) {
	m.SelectedTags = append(m.SelectedTags, newTag)
}
