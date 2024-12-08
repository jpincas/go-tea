package main

import (
	"fmt"

	gotea "github.com/jpincas/go-tea"
	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
)

var counterMessages gt.MessageMap = gt.MessageMap{
	"INCREMENT_COUNTER": IncrementCounter,
}

func IncrementCounter(m gotea.Message, s gt.State) gt.Response {
	n := m.ArgsToInt()
	state := model(s)
	state.Counter = state.Counter + n
	return gt.Respond()
}

func renderCounter(counter int) h.Element {
	return h.Div(
		a.Attrs(a.Id("counter")),
		h.Button(
			a.Attrs(a.Class("counter-increment"), a.OnClick(gt.SendBasicMessage("INCREMENT_COUNTER", -1))),
			h.Text("Down"),
		),
		h.Div(
			a.Attrs(a.Class("counter-readout")),
			h.Text(fmt.Sprintf("%d", counter)),
		),
		h.Button(
			a.Attrs(a.Class("counter-increment"), a.OnClick(gt.SendBasicMessage("INCREMENT_COUNTER", 1))),
			h.Text("Up"),
		),
	)
}
