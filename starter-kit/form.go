package main

// FormState holds the form data. Field names must match the HTML name attributes.
// The form is automatically serialized by gt.BasicUpdateForm and decoded via MustDecodeArgs.
type FormState struct {
	Name     string   // text input
	Email    string   // email input
	Selected string   // select dropdown
	Message  string   // textarea
	Agreed   bool     // checkbox
	Options  []string // available options for select (not from form)
}
