package main

import (
	"strings"

	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/go-tea/attributes"
	h "github.com/jpincas/go-tea/html"
)

// TagSelector is a reusable component that demonstrates ComponentID namespacing.
// ComponentID ensures that multiple instances of this component can coexist
// without message collisions.
type TagSelector struct {
	gt.ComponentID        // Embed for namespacing
	Available      []string
	Suggestions    []string
	Selected       []string
	SearchQuery    string
}

// NewTagSelector creates a new instance with a unique ID
func NewTagSelector(id string, available []string) TagSelector {
	return TagSelector{
		ComponentID: gt.ComponentID(id),
		Available:   available,
	}
}

// SelectTag adds a tag to selected and clears search
func (ts *TagSelector) SelectTag(tag string) {
	// Avoid duplicates
	for _, t := range ts.Selected {
		if t == tag {
			return
		}
	}
	ts.Selected = append(ts.Selected, tag)
	ts.SearchQuery = ""
	ts.Suggestions = nil
}

// RemoveTag removes a tag from selected
func (ts *TagSelector) RemoveTag(tag string) {
	var filtered []string
	for _, t := range ts.Selected {
		if t != tag {
			filtered = append(filtered, t)
		}
	}
	ts.Selected = filtered
}

// Search filters available tags by query
func (ts *TagSelector) Search(query string) {
	ts.SearchQuery = query
	if query == "" {
		ts.Suggestions = nil
		return
	}

	var matches []string
	for _, tag := range ts.Available {
		if strings.Contains(strings.ToLower(tag), strings.ToLower(query)) {
			// Don't suggest already-selected tags
			alreadySelected := false
			for _, s := range ts.Selected {
				if s == tag {
					alreadySelected = true
					break
				}
			}
			if !alreadySelected {
				matches = append(matches, tag)
			}
		}
	}
	ts.Suggestions = matches
}

// Render returns the component's HTML.
// Uses UniqueMsg and UniqueID for namespaced identifiers.
func (ts TagSelector) Render() h.Element {
	// Namespaced message names: "my-selector_SELECT", "my-selector_REMOVE", etc.
	msgSelect := ts.UniqueMsg("SELECT")
	msgRemove := ts.UniqueMsg("REMOVE")
	msgSearch := ts.UniqueMsg("SEARCH")

	// Namespaced element ID
	inputID := ts.UniqueID("search-input")

	return h.Div(a.Attrs(a.Class("space-y-4")),
		// Search input
		h.Input(a.Attrs(
			a.Id(inputID),
			a.Class("w-full px-4 py-2 border rounded focus:ring-2 focus:ring-blue-500"),
			a.Type("text"),
			a.Placeholder("Search tags..."),
			a.Value(ts.SearchQuery),
			// SendBasicMessageWithValueFromInput gets the input value and sends it
			a.OnKeyUp(gt.SendBasicMessageWithValueFromInput(msgSearch, inputID)))),

		// Suggestions dropdown
		h.Div(a.Attrs(a.Class("space-y-1"))).RenderIf(len(ts.Suggestions) == 0),
		h.Div(a.Attrs(a.Class("border rounded divide-y")),
			renderSuggestions(ts.Suggestions, msgSelect)...).
			RenderIf(len(ts.Suggestions) > 0),

		// Selected tags
		h.Div(a.Attrs(a.Class("flex flex-wrap gap-2")),
			renderSelectedTags(ts.Selected, msgRemove)...))
}

func renderSuggestions(suggestions []string, msgSelect string) []h.Element {
	var els []h.Element
	for _, tag := range suggestions {
		els = append(els,
			h.Div(a.Attrs(
				a.Class("px-3 py-2 hover:bg-blue-50 cursor-pointer"),
				a.OnClick(gt.SendBasicMessage(msgSelect, tag))),
				h.Text(tag)))
	}
	return els
}

func renderSelectedTags(selected []string, msgRemove string) []h.Element {
	var els []h.Element
	for _, tag := range selected {
		els = append(els,
			h.Span(a.Attrs(
				a.Class("inline-flex items-center gap-1 px-3 py-1 bg-blue-100 text-blue-800 rounded-full text-sm")),
				h.Text(tag),
				h.Button(a.Attrs(
					a.Class("ml-1 text-blue-600 hover:text-blue-800"),
					a.OnClick(gt.SendBasicMessage(msgRemove, tag))),
					h.Text("x"))))
	}
	return els
}
