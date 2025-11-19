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
		a.Attrs(a.Class("bg-gray-50 p-4 rounded-lg shadow-inner")),
		h.H2(a.Attrs(a.Class("text-lg font-medium text-gray-900 mb-4")), h.Text("Form Values")),
		h.Ul(
			a.Attrs(a.Class("space-y-2 text-sm text-gray-600")),
			h.Li(a.Attrs(), h.Text(fmt.Sprintf("TextInput: %s", f.TextInput))),
			h.Li(a.Attrs(), h.Text(fmt.Sprintf("TextboxInput: %s", f.TextboxInput))),
			h.Li(a.Attrs(), h.Text(fmt.Sprintf("SelectInput: %s", f.SelectInput))),
			h.Li(a.Attrs(), h.Text(fmt.Sprintf("IntInput: %d", f.IntInput))),
			h.Li(a.Attrs(), h.Text(fmt.Sprintf("FloatInput: %f", f.FloatInput))),
			h.Li(a.Attrs(), h.Text(fmt.Sprintf("CheckboxInput: %t", f.CheckboxInput))),
			h.Li(a.Attrs(), h.Text(fmt.Sprintf("RadioTextInput: %s", f.RadioTextInput))),
			h.Li(a.Attrs(), h.Text(fmt.Sprintf("MultipleTextInput: %v", f.MultipleTextInput))),
		),
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
		a.Attrs(a.Class("space-y-6")),
		renderExplanatoryNote(
			"Forms and Input Handling",
			`
			<p class="mb-2">This example shows how to handle form inputs and state.</p>
			<ul class="list-disc pl-5 space-y-1">
				<li><strong>Two-Way Binding:</strong> Input values are bound to the state. <code>INPUT</code> events trigger messages that update the state.</li>
				<li><strong>Form Submission:</strong> The <code>SUBMIT</code> event is captured to process the form data.</li>
				<li><strong>Validation:</strong> Input validation logic can be implemented in the update function before updating the state.</li>
			</ul>
			`,
		),
		h.H2(a.Attrs(a.Class("text-2xl font-bold text-gray-900")), h.Text("Form Example")),
		h.Div(
			a.Attrs(a.Class("grid grid-cols-1 md:grid-cols-2 gap-8")),
			h.Form(
				a.Attrs(a.Class("space-y-6"), a.Id("my-form")),
				
				h.Div(
					a.Attrs(),
					h.Label(a.Attrs(a.Class("block text-sm font-medium text-gray-700")), h.Text("Text Input")),
					h.Input(
						a.Attrs(
							a.Class("mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"),
							a.Type("text"),
							a.Placeholder("Some simple text input"),
							a.Value(f.TextInput),
							a.Name("textInput"),
							a.OnKeyUp(gt.BasicUpdateForm("FORM_UPDATE", "my-form")),
						),
					),
				),

				h.Div(
					a.Attrs(),
					h.Label(a.Attrs(a.Class("block text-sm font-medium text-gray-700")), h.Text("Simple Select")),
					h.Select(
						a.Attrs(
							a.Class("mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"),
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
					),
				),

				h.Div(
					a.Attrs(),
					h.Label(a.Attrs(a.Class("block text-sm font-medium text-gray-700")), h.Text("Multiple Select")),
					h.Select(
						a.Attrs(
							a.Class("mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"),
							a.Name("MultipleTextInput"), 
							a.OnChange(gt.BasicUpdateForm("FORM_UPDATE", "my-form")), 
							a.Size(4), 
							a.Multiple(true),
						),
						h.Option(a.Attrs(a.Value("first"), f.isSelected("first", f.MultipleTextInput)), h.Text("first")),
						h.Option(a.Attrs(a.Value("second"), f.isSelected("second", f.MultipleTextInput)), h.Text("second")),
						h.Option(a.Attrs(a.Value("third"), f.isSelected("third", f.MultipleTextInput)), h.Text("third")),
						h.Option(a.Attrs(a.Value("fourth"), f.isSelected("fourth", f.MultipleTextInput)), h.Text("fourth")),
					),
				),

				h.Div(
					a.Attrs(),
					h.Label(a.Attrs(a.Class("block text-sm font-medium text-gray-700")), h.Text("Text Area")),
					h.TextArea(
						a.Attrs(
							a.Class("mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"),
							a.OnKeyUp(gt.BasicUpdateForm("FORM_UPDATE", "my-form")), 
							a.Name("TextboxInput"), 
							a.Rows(5),
						),
						h.Text(f.TextboxInput),
					),
				),

				h.Div(
					a.Attrs(),
					h.Label(a.Attrs(a.Class("block text-sm font-medium text-gray-700")), h.Text("Radio")),
					h.Div(
						a.Attrs(a.Class("mt-2 space-y-2")),
						h.Div(
							a.Attrs(a.Class("flex items-center")),
							h.Input(
								a.Attrs(
									a.Class("h-4 w-4 border-gray-300 text-indigo-600 focus:ring-indigo-500"),
									a.Type("radio"), 
									a.Name("RadioTextInput"), 
									a.Value("male"), 
									a.OnChange(gt.BasicUpdateForm("FORM_UPDATE", "my-form")), 
									f.isChecked("male", f.RadioTextInput),
								),
							),
							h.Label(a.Attrs(a.Class("ml-3 block text-sm font-medium text-gray-700")), h.Text("Male")),
						),
						h.Div(
							a.Attrs(a.Class("flex items-center")),
							h.Input(
								a.Attrs(
									a.Class("h-4 w-4 border-gray-300 text-indigo-600 focus:ring-indigo-500"),
									a.Type("radio"), 
									a.Name("RadioTextInput"), 
									a.Value("female"), 
									a.OnChange(gt.BasicUpdateForm("FORM_UPDATE", "my-form")), 
									f.isChecked("female", f.RadioTextInput),
								),
							),
							h.Label(a.Attrs(a.Class("ml-3 block text-sm font-medium text-gray-700")), h.Text("Female")),
						),
					),
				),

				h.Div(
					a.Attrs(a.Class("flex items-start")),
					h.Div(
						a.Attrs(a.Class("flex h-5 items-center")),
						h.Input(
							a.Attrs(
								a.Class("h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"),
								a.Type("checkbox"), 
								a.Name("CheckboxInput"), 
								a.OnChange(gt.BasicUpdateForm("FORM_UPDATE", "my-form")), 
								f.isCheckedBool(f.CheckboxInput),
							),
						),
					),
					h.Div(
						a.Attrs(a.Class("ml-3 text-sm")),
						h.Label(a.Attrs(a.Class("font-medium text-gray-700")), h.Text("True?")),
					),
				),
			),
			
			f.renderValues(),
		),
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
