package main

import (
	"encoding/json"
	"fmt"

	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
)

type Form struct {
	Options           []string
	TextInput         string
	TextboxInput      string
	SelectInput       string
	IntInput          int
	FloatInput        float64
	CheckboxInput     bool
	RadioTextInput    string
	MultipleTextInput []string
}

func (f Form) String() string {
	// Encode to JSON with newlines and indents
	b, _ := json.MarshalIndent(f, "", "\n")
	return string(b)
}

func (f Form) renderValues() h.Element {
	return h.Div(
		a.Attrs(a.Class("bg-gradient-to-br from-emerald-50 to-teal-50 p-6 rounded-xl border-2 border-stone-900 shadow-brutal-sm")),
		h.H2(a.Attrs(a.Class("text-lg font-bold text-stone-900 mb-4 flex items-center gap-2")),
			h.Span(a.Attrs(), h.Text("📊")),
			h.Text("Live Form State"),
		),
		h.Div(
			a.Attrs(a.Class("space-y-3 text-sm"), a.Custom("style", "font-family: 'JetBrains Mono', monospace;")),
			formValueRow("textInput", f.TextInput),
			formValueRow("textboxInput", f.TextboxInput),
			formValueRow("selectInput", f.SelectInput),
			formValueRow("intInput", fmt.Sprintf("%d", f.IntInput)),
			formValueRow("floatInput", fmt.Sprintf("%.2f", f.FloatInput)),
			formValueRow("checkboxInput", fmt.Sprintf("%t", f.CheckboxInput)),
			formValueRow("radioTextInput", f.RadioTextInput),
			formValueRow("multipleTextInput", fmt.Sprintf("%v", f.MultipleTextInput)),
		),
	)
}

func formValueRow(key, value string) h.Element {
	displayValue := value
	if displayValue == "" {
		displayValue = "—"
	}
	return h.Div(
		a.Attrs(a.Class("flex items-start gap-2")),
		h.Span(a.Attrs(a.Class("text-emerald-600 font-semibold shrink-0")), h.Text(key+":")),
		h.Span(a.Attrs(a.Class("text-stone-700 break-all")), h.Text(displayValue)),
	)
}

var formMessages gt.MessageMap = gt.MessageMap{
	"FORM_UPDATE": formUpdate,
}

func formUpdate(m gt.Message, s gt.State) gt.Response {
	state := model(s)
	m.MustDecodeArgs(&state.Form)
	return gt.Respond()
}

// Render
func (f Form) render() h.Element {
	return h.Div(
		a.Attrs(a.Class("space-y-8")),
		// Header
		h.Div(
			a.Attrs(a.Class("text-center space-y-2")),
			h.H1(a.Attrs(a.Class("text-4xl font-bold text-stone-900"), a.Custom("style", "font-family: 'DM Serif Display', serif;")), h.Text("📝 Form Handling")),
			h.P(a.Attrs(a.Class("text-stone-600")), h.Text("Real-time form state binding — watch the state update as you type!")),
		),
		renderExplanatoryNote(
			"Forms and Input Handling",
			`
			<p class="mb-3">This example shows how to handle form inputs and state.</p>
			<ul class="list-disc pl-5 space-y-2">
				<li><strong class="text-stone-900">Two-Way Binding:</strong> Input values are bound to the state. Events trigger messages that update the state.</li>
				<li><strong class="text-stone-900">Form Serialization:</strong> The entire form is serialized and sent with each change using <code class="bg-stone-200 px-1.5 py-0.5 rounded text-xs font-mono">BasicUpdateForm</code>.</li>
				<li><strong class="text-stone-900">Validation:</strong> Input validation logic can be implemented in the update function before updating the state.</li>
			</ul>
			`,
		),
		h.Div(
			a.Attrs(a.Class("grid grid-cols-1 lg:grid-cols-2 gap-8")),
			// Form column
			h.Form(
				a.Attrs(a.Class("space-y-5 bg-white p-6 rounded-xl border-2 border-stone-900 shadow-brutal-sm"), a.Id("my-form")),
				h.H3(a.Attrs(a.Class("text-lg font-bold text-stone-900 mb-4")), h.Text("Input Fields")),

				formField("Text Input", h.Input(
					a.Attrs(
						a.Class("mt-1 block w-full rounded-lg px-4 py-2.5 text-stone-900 placeholder-stone-400"),
						a.Type("text"),
						a.Placeholder("Type something here..."),
						a.Value(f.TextInput),
						a.Name("textInput"),
						a.OnKeyUp(gt.BasicUpdateForm("FORM_UPDATE", "my-form")),
					),
				)),

				formField("Select", h.Select(
					a.Attrs(
						a.Class("mt-1 block w-full rounded-lg px-4 py-2.5 text-stone-900 bg-white"),
						a.Name("selectInput"),
						a.OnChange(gt.BasicUpdateForm("FORM_UPDATE", "my-form")),
					),
					func() []h.Element {
						var options []h.Element
						for _, option := range f.Options {
							options = append(options, h.Option(
								a.Attrs(
									a.Value(option),
									func() a.Attribute {
										if f.SelectInput == option {
											return a.Selected(true)
										}
										return a.Selected(false)
									}(),
								),
								h.Text(option),
							))
						}
						return options
					}()...,
				)),

				formField("Multiple Select", h.Select(
					a.Attrs(
						a.Class("mt-1 block w-full rounded-lg px-4 py-2 text-stone-900 bg-white"),
						a.Name("MultipleTextInput"),
						a.OnChange(gt.BasicUpdateForm("FORM_UPDATE", "my-form")),
						a.Size(4),
						a.Multiple(true),
					),
					h.Option(a.Attrs(a.Value("first"), f.isSelected("first", f.MultipleTextInput)), h.Text("first")),
					h.Option(a.Attrs(a.Value("second"), f.isSelected("second", f.MultipleTextInput)), h.Text("second")),
					h.Option(a.Attrs(a.Value("third"), f.isSelected("third", f.MultipleTextInput)), h.Text("third")),
					h.Option(a.Attrs(a.Value("fourth"), f.isSelected("fourth", f.MultipleTextInput)), h.Text("fourth")),
				)),

				formField("Text Area", h.TextArea(
					a.Attrs(
						a.Class("mt-1 block w-full rounded-lg px-4 py-2.5 text-stone-900 placeholder-stone-400"),
						a.OnKeyUp(gt.BasicUpdateForm("FORM_UPDATE", "my-form")),
						a.Name("TextboxInput"),
						a.Rows(4),
						a.Placeholder("Enter longer text here..."),
					),
					h.Text(f.TextboxInput),
				)),

				// Radio buttons
				h.Div(
					a.Attrs(a.Class("space-y-2")),
					h.Label(a.Attrs(a.Class("block text-sm font-semibold text-stone-700")), h.Text("Radio Selection")),
					h.Div(
						a.Attrs(a.Class("flex gap-4 mt-2")),
						radioOption("Male", "male", "RadioTextInput", f.RadioTextInput == "male"),
						radioOption("Female", "female", "RadioTextInput", f.RadioTextInput == "female"),
					),
				),

				// Checkbox
				h.Div(
					a.Attrs(a.Class("flex items-center gap-3 p-3 bg-stone-50 rounded-lg")),
					h.Input(
						a.Attrs(
							a.Class("h-5 w-5 rounded border-2 border-stone-900 text-emerald-600 focus:ring-emerald-500"),
							a.Type("checkbox"),
							a.Name("CheckboxInput"),
							a.OnChange(gt.BasicUpdateForm("FORM_UPDATE", "my-form")),
							f.isCheckedBool(f.CheckboxInput),
						),
					),
					h.Label(a.Attrs(a.Class("font-medium text-stone-700")), h.Text("Toggle this checkbox")),
				),
			),

			// Values column
			f.renderValues(),
		),
	)
}

func formField(label string, input h.Element) h.Element {
	return h.Div(
		a.Attrs(a.Class("space-y-1")),
		h.Label(a.Attrs(a.Class("block text-sm font-semibold text-stone-700")), h.Text(label)),
		input,
	)
}

func radioOption(label, value, name string, checked bool) h.Element {
	return h.Label(
		a.Attrs(a.Class("flex items-center gap-2 cursor-pointer")),
		h.Input(
			a.Attrs(
				a.Class("h-4 w-4 border-2 border-stone-900 text-emerald-600 focus:ring-emerald-500"),
				a.Type("radio"),
				a.Name(name),
				a.Value(value),
				a.OnChange(gt.BasicUpdateForm("FORM_UPDATE", "my-form")),
				func() a.Attribute {
					if checked {
						return a.Checked(true)
					}
					return a.Checked(false)
				}(),
			),
		),
		h.Span(a.Attrs(a.Class("text-sm font-medium text-stone-700")), h.Text(label)),
	)
}

func (f Form) isSelected(value string, selectedValues []string) a.Attribute {
	for _, v := range selectedValues {
		if v == value {
			return a.Selected(true)
		}
	}
	return a.Selected(false)
}

func (f Form) isChecked(value string, selectedValue string) a.Attribute {
	if value == selectedValue {
		return a.Checked(true)
	}
	return a.Checked(false)
}

func (f Form) isCheckedBool(checked bool) a.Attribute {
	if checked {
		return a.Checked(true)
	}
	return a.Checked(false)
}
