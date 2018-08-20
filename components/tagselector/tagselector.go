package tagselector

import gotea "github.com/jpincas/go-tea"

// Message Constants
const (
	MsgSelectTag         = "TAG.SELECT"
	MsgSearchInputUpdate = "SEARCHINPUT.UPDATE"
)

type Model struct {
	gotea.ComponentID
	SearchInput   string
	AvailableTags []string
	SuggestedTags []string
	SelectedTags  []string
}

func (m *Model) SelectTag(newTag string) {
	m.SelectedTags = append(m.SelectedTags, newTag)
}
