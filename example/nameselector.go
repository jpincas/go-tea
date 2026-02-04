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
		a.Attrs(a.Class("space-y-8")),
		// Header
		h.Div(
			a.Attrs(a.Class("text-center space-y-2")),
			h.H1(a.Attrs(a.Class("text-4xl font-bold text-stone-900"), a.Custom("style", "font-family: 'DM Serif Display', serif;")), h.Text("🧩 Reusable Components")),
			h.P(a.Attrs(a.Class("text-stone-600")), h.Text("Two instantiations of the same 'TagSelector' component, each with their own state.")),
		),

		// Components grid
		h.Div(
			a.Attrs(a.Class("grid grid-cols-1 md:grid-cols-2 gap-6")),
			// Name selector
			h.Div(
				a.Attrs(a.Class("space-y-3")),
				h.Div(
					a.Attrs(a.Class("flex items-center gap-2")),
					h.Span(a.Attrs(a.Class("text-2xl")), h.Text("👤")),
					h.H3(a.Attrs(a.Class("text-lg font-bold text-stone-900")), h.Text("Select a Name")),
				),
				nameSelector.Render(),
			),
			// Team selector
			h.Div(
				a.Attrs(a.Class("space-y-3")),
				h.Div(
					a.Attrs(a.Class("flex items-center gap-2")),
					h.Span(a.Attrs(a.Class("text-2xl")), h.Text("⚽")),
					h.H3(a.Attrs(a.Class("text-lg font-bold text-stone-900")), h.Text("Select a Team")),
				),
				tagSelector.Render(),
			),
		),

		// Explanatory note
		renderExplanatoryNote(
			"Reusable Components",
			`
			<p class="mb-3">This example demonstrates how to create reusable components in GoTea.</p>
			<ul class="list-disc pl-5 space-y-2">
				<li><strong class="text-stone-900">Component Model:</strong> The <code class="bg-stone-200 px-1.5 py-0.5 rounded text-xs font-mono">TagSelector</code> is a struct that holds its own state.</li>
				<li><strong class="text-stone-900">Instantiation:</strong> We create two instances with different configuration in the main Model.</li>
				<li><strong class="text-stone-900">Message Namespacing:</strong> Each component uses <code class="bg-stone-200 px-1.5 py-0.5 rounded text-xs font-mono">UniqueMsg()</code> and <code class="bg-stone-200 px-1.5 py-0.5 rounded text-xs font-mono">UniqueID()</code> to prevent collisions.</li>
			</ul>
			`,
		),
	)
}
