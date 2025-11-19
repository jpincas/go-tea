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
		renderExplanatoryNote(
			"Reusable Components",
			`
			<p class="mb-2">This example demonstrates how to create reusable components in GoTea.</p>
			<ul class="list-disc pl-5 space-y-1">
				<li><strong>Component Model:</strong> The <code>TagSelector</code> is a struct that holds its own state.</li>
				<li><strong>Instantiation:</strong> We create two instances of the <code>TagSelector</code> in the main <code>Model</code>, each with different configuration.</li>
				<li><strong>Message Routing:</strong> Messages are routed to the correct component instance (though in this simple example, we just have separate message handlers).</li>
			</ul>
			`,
		),
		h.Div(
			a.Attrs(a.Class("text-center")),
			h.H2(a.Attrs(a.Class("text-2xl font-bold text-gray-900")), h.Text("Components Demo")),
			h.P(a.Attrs(a.Class("mt-2 text-gray-600")), h.Text("Two instantiations of the 'tag selector' component, running side-by-side")),
		),
		h.Div(
			a.Attrs(a.Class("grid grid-cols-1 md:grid-cols-2 gap-8")),
			h.Div(
				a.Attrs(a.Class("space-y-4")),
				h.H3(a.Attrs(a.Class("text-lg font-medium text-gray-900")), h.Text("Select a Name")),
				nameSelector.Render(),
			),
			h.Div(
				a.Attrs(a.Class("space-y-4")),
				h.H3(a.Attrs(a.Class("text-lg font-medium text-gray-900")), h.Text("Select a Team")),
				tagSelector.Render(),
			),
		),
	)
}
