package main

import (
	"encoding/json"

	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
)

var crashMessages gt.MessageMap = gt.MessageMap{
	"CRASH_ME": CrashMe,
}

func CrashMe(args json.RawMessage, s gt.State) gt.Response {
	panic("I told you not to hit the big red button")
	return gt.Respond()
}

func renderCrash() h.Element {
	return h.Div(
		a.Attrs(a.Id("crash")),
		h.P(a.Attrs(), h.Text("In order to stop errors in client-app code killing the Gotea runtime, crash protection is included!")),
		h.Button(
			a.Attrs(a.Class("crash-button"), a.OnClick(gt.SendMessageNoArgs("CRASH_ME"))),
			h.Text("Crash!"),
		),
	)
}
