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
		a.Attrs(),
		h.H2(a.Attrs(), h.Text("Form Values")),
		h.Ul(
			a.Attrs(),
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
func (form Form) render() h.Element {
	return h.Div(
		a.Attrs(),
		h.H1(a.Attrs(), h.Text("Form Example")),
		h.Div(
			a.Attrs(a.Class("row")),
			h.Form(
				a.Attrs(a.Class("equalchild"), a.Id("my-form")),
				h.H2(a.Attrs(), h.Text("Text Input")),
				h.Input(
					a.Attrs(
						a.Class("input"),
						a.Type("text"),
						a.Placeholder("Some simple text input"),
						a.Value(form.TextInput),
						a.Name("textInput"),
						a.OnKeyUp(gt.BasicUpdateForm("FORM_UPDATE", "my-form")),
					),
				),
				h.H2(a.Attrs(), h.Text("Simple Select")),
				h.Select(
					a.Attrs(a.Name("selectInput"), a.OnChange(gt.BasicUpdateForm("FORM_UPDATE", "my-form"))),
					func() []h.Element {
						var options []h.Element
						for _, option := range form.Options {
							options = append(options, h.Option(
								a.Attrs(
									a.Value(option),
									func() a.Attribute {
										if form.SelectInput == option {
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
				h.H2(a.Attrs(), h.Text("Multiple Select")),
				h.Select(
					a.Attrs(a.Name("MultipleTextInput"), a.OnChange(gt.BasicUpdateForm("FORM_UPDATE", "my-form")), a.Size(4), a.Multiple(true)),
					h.Option(a.Attrs(a.Value("first"), form.isSelected("first", form.MultipleTextInput)), h.Text("first")),
					h.Option(a.Attrs(a.Value("second"), form.isSelected("second", form.MultipleTextInput)), h.Text("second")),
					h.Option(a.Attrs(a.Value("third"), form.isSelected("third", form.MultipleTextInput)), h.Text("third")),
					h.Option(a.Attrs(a.Value("fourth"), form.isSelected("fourth", form.MultipleTextInput)), h.Text("fourth")),
				),
				h.H2(a.Attrs(), h.Text("Text Area")),
				h.TextArea(
					a.Attrs(a.OnKeyUp(gt.BasicUpdateForm("FORM_UPDATE", "my-form")), a.Name("TextboxInput"), a.Rows(10), a.Cols(30)),
					h.Text(form.TextboxInput),
				),
				h.H2(a.Attrs(), h.Text("Radio")),
				h.Input(
					a.Attrs(a.Type("radio"), a.Name("RadioTextInput"), a.Value("male"), a.OnChange(gt.BasicUpdateForm("FORM_UPDATE", "my-form")), form.isChecked("male", form.RadioTextInput)),
				),
				h.Text("Male"),
				h.Input(
					a.Attrs(a.Type("radio"), a.Name("RadioTextInput"), a.Value("female"), a.OnChange(gt.BasicUpdateForm("FORM_UPDATE", "my-form")), form.isChecked("female", form.RadioTextInput)),
				),
				h.Text("Female"),
				h.H2(a.Attrs(), h.Text("Checkbox")),
				h.Input(
					a.Attrs(a.Type("checkbox"), a.Name("CheckboxInput"), a.OnChange(gt.BasicUpdateForm("FORM_UPDATE", "my-form")), form.isCheckedBool(form.CheckboxInput)),
				),
				h.Text("True?"),
			),
			h.Div(
				a.Attrs(a.Class("equalchild")),
				h.H2(a.Attrs(), h.Text("State")),
				form.renderValues(),
			),
		),
	)
}

func (form Form) isSelected(value string, selectedValues []string) a.Attribute {
	for _, v := range selectedValues {
		if v == value {
			return a.Selected(true)
		}
	}
	return a.Selected(false)
}

func (form Form) isChecked(value string, selectedValue string) a.Attribute {
	if value == selectedValue {
		return a.Checked(true)
	}
	return a.Checked(false)
}

func (form Form) isCheckedBool(checked bool) a.Attribute {
	if checked {
		return a.Checked(true)
	}
	return a.Checked(false)
}
