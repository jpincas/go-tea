package main

import (
	"fmt"

	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/go-tea/attributes"
	h "github.com/jpincas/go-tea/html"
)

// ============================================================================
// Home Page
// ============================================================================

func renderHome(m *Model) h.Element {
	return h.Div(a.Attrs(a.Class("space-y-8")),
		// Header
		h.Div(a.Attrs(a.Class("text-center space-y-2")),
			h.H1(a.Attrs(a.Class("text-4xl font-bold text-gray-900")),
				h.Text("Welcome to Gotea")),
			h.P(a.Attrs(a.Class("text-gray-600")),
				h.Text("The Elm Architecture for Go - server-side state with WebSocket updates"))),

		// Counter demo
		h.Div(a.Attrs(a.Class("bg-white p-6 rounded-lg shadow-sm border")),
			h.H2(a.Attrs(a.Class("text-lg font-semibold mb-4")),
				h.Text("Counter Example")),
			h.P(a.Attrs(a.Class("text-gray-600 text-sm mb-4")),
				h.Text("Each click sends a message to the server, which updates the state and re-renders.")),
			renderCounter(m.Counter)),

		// Features
		h.Div(a.Attrs(a.Class("grid grid-cols-1 md:grid-cols-2 gap-4")),
			featureCard("/form", "Form Handling",
				"Real-time form binding with automatic state sync"),
			featureCard("/component", "Components",
				"Reusable components with namespaced messages")))
}

func renderCounter(count int) h.Element {
	return h.Div(a.Attrs(a.Class("flex items-center gap-4")),
		h.Button(a.Attrs(
			a.Class("w-10 h-10 bg-red-500 hover:bg-red-600 text-white font-bold rounded"),
			a.OnClick(gt.SendBasicMessage("INCREMENT", -1))),
			h.Text("-")),
		h.Span(a.Attrs(a.Class("text-3xl font-mono font-bold w-16 text-center")),
			h.Text(fmt.Sprintf("%d", count))),
		h.Button(a.Attrs(
			a.Class("w-10 h-10 bg-green-500 hover:bg-green-600 text-white font-bold rounded"),
			a.OnClick(gt.SendBasicMessage("INCREMENT", 1))),
			h.Text("+")))
}

func featureCard(href, title, description string) h.Element {
	return h.A(a.Attrs(
		a.Href(href),
		a.Class("block p-4 bg-white rounded-lg shadow-sm border hover:shadow-md transition-shadow")),
		h.H3(a.Attrs(a.Class("font-semibold text-gray-900")),
			h.Text(title)),
		h.P(a.Attrs(a.Class("text-sm text-gray-600 mt-1")),
			h.Text(description)))
}

// ============================================================================
// Form Page
// ============================================================================

func renderFormPage(m *Model) h.Element {
	return h.Div(a.Attrs(a.Class("space-y-6")),
		h.H1(a.Attrs(a.Class("text-3xl font-bold text-gray-900")),
			h.Text("Form Handling")),
		h.P(a.Attrs(a.Class("text-gray-600")),
			h.Text("Form state is synced to the server on every change.")),

		h.Div(a.Attrs(a.Class("grid grid-cols-1 lg:grid-cols-2 gap-6")),
			// Form inputs
			h.Form(a.Attrs(
				a.Class("bg-white p-6 rounded-lg shadow-sm border space-y-4"),
				a.Id("my-form")),

				formField("Text Input",
					h.Input(a.Attrs(
						a.Class("w-full px-3 py-2 border rounded focus:ring-2 focus:ring-blue-500"),
						a.Type("text"),
						a.Name("Name"),
						a.Placeholder("Enter your name..."),
						a.Value(m.Form.Name),
						a.OnKeyUp(gt.BasicUpdateForm("FORM_UPDATE", "my-form"))))),

				formField("Email",
					h.Input(a.Attrs(
						a.Class("w-full px-3 py-2 border rounded focus:ring-2 focus:ring-blue-500"),
						a.Type("email"),
						a.Name("Email"),
						a.Placeholder("you@example.com"),
						a.Value(m.Form.Email),
						a.OnKeyUp(gt.BasicUpdateForm("FORM_UPDATE", "my-form"))))),

				formField("Select",
					h.Select(a.Attrs(
						a.Class("w-full px-3 py-2 border rounded"),
						a.Name("Selected"),
						a.OnChange(gt.BasicUpdateForm("FORM_UPDATE", "my-form"))),
						renderOptions(m.Form.Options, m.Form.Selected)...)),

				formField("Message",
					h.TextArea(a.Attrs(
						a.Class("w-full px-3 py-2 border rounded focus:ring-2 focus:ring-blue-500"),
						a.Name("Message"),
						a.Rows(3),
						a.Placeholder("Your message..."),
						a.OnKeyUp(gt.BasicUpdateForm("FORM_UPDATE", "my-form"))),
						h.Text(m.Form.Message))),

				// Checkbox with conditional rendering
				h.Div(a.Attrs(a.Class("flex items-center gap-2")),
					h.Input(a.Attrs(
						a.Class("w-4 h-4"),
						a.Type("checkbox"),
						a.Name("Agreed"),
						a.OnChange(gt.BasicUpdateForm("FORM_UPDATE", "my-form")),
						checkedIf(m.Form.Agreed))),
					h.Label(a.Attrs(a.Class("text-sm text-gray-700")),
						h.Text("I agree to the terms")))),

			// Live state display
			h.Div(a.Attrs(a.Class("bg-gray-50 p-6 rounded-lg border")),
				h.H3(a.Attrs(a.Class("font-semibold mb-4")),
					h.Text("Live Form State")),
				h.Div(a.Attrs(a.Class("space-y-2 font-mono text-sm")),
					stateRow("Name", m.Form.Name),
					stateRow("Email", m.Form.Email),
					stateRow("Selected", m.Form.Selected),
					stateRow("Message", m.Form.Message),
					stateRow("Agreed", fmt.Sprintf("%t", m.Form.Agreed))),
				// Conditional rendering example
				h.Div(a.Attrs(a.Class("mt-4 p-3 bg-green-100 text-green-800 rounded"))).
					RenderIf(m.Form.Agreed),
				h.Span(a.Attrs(),
					h.Text("Terms accepted!")).RenderIf(m.Form.Agreed))))
}

func formField(label string, input h.Element) h.Element {
	return h.Div(a.Attrs(a.Class("space-y-1")),
		h.Label(a.Attrs(a.Class("block text-sm font-medium text-gray-700")),
			h.Text(label)),
		input)
}

func stateRow(key, value string) h.Element {
	if value == "" {
		value = "(empty)"
	}
	return h.Div(a.Attrs(a.Class("flex gap-2")),
		h.Span(a.Attrs(a.Class("text-gray-500")), h.Text(key+":")),
		h.Span(a.Attrs(a.Class("text-gray-900")), h.Text(value)))
}

func renderOptions(options []string, selected string) []h.Element {
	var els []h.Element
	els = append(els, h.Option(a.Attrs(a.Value("")), h.Text("Select...")))
	for _, opt := range options {
		attrs := a.Attrs(a.Value(opt))
		if opt == selected {
			attrs = a.Attrs(a.Value(opt), a.Selected(true))
		}
		els = append(els, h.Option(attrs, h.Text(opt)))
	}
	return els
}

func checkedIf(checked bool) a.Attribute {
	if checked {
		return a.Checked(true)
	}
	return a.Checked(false)
}

// ============================================================================
// Component Page
// ============================================================================

func renderComponentPage(m *Model) h.Element {
	return h.Div(a.Attrs(a.Class("space-y-6")),
		h.H1(a.Attrs(a.Class("text-3xl font-bold text-gray-900")),
			h.Text("Reusable Components")),
		h.P(a.Attrs(a.Class("text-gray-600")),
			h.Text("Components use ComponentID to namespace their messages, preventing collisions.")),

		h.Div(a.Attrs(a.Class("bg-white p-6 rounded-lg shadow-sm border")),
			h.H2(a.Attrs(a.Class("text-lg font-semibold mb-4")),
				h.Text("Tag Selector Component")),
			m.Selector.Render()),

		h.Div(a.Attrs(a.Class("bg-gray-50 p-4 rounded-lg border mt-4")),
			h.H3(a.Attrs(a.Class("font-semibold mb-2")),
				h.Text("Selected Tags:")),
			h.Div(a.Attrs(a.Class("font-mono text-sm")),
				h.Text(fmt.Sprintf("%v", m.Selector.Selected))).
					RenderIf(len(m.Selector.Selected) > 0),
			h.Span(a.Attrs(a.Class("text-gray-500 italic")),
				h.Text("None selected")).
				RenderIf(len(m.Selector.Selected) == 0)))
}
