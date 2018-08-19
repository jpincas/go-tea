package main

import (
	"encoding/json"

	gotea "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/components/tagselector"
)

func init() {
	// Have to specify tag selector messages
	tagselector.Messages["TAGSELECTOR_SELECTTAG"] = func(args json.RawMessage, s gotea.State) (gotea.State, *gotea.Message, error) {
		state := s.(Model)

		var tag string
		if err := json.Unmarshal(args, &tag); err != nil {
			return state, nil, err
		}

		state.TagSelector.SelectTag(tag)
		return state, nil, nil
	}
}
