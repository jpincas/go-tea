package main

import (
	"encoding/json"
	"fmt"

	gt "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/msg"
	a "github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
)

var counterMessages gt.MessageMap = gt.MessageMap{
	"INCREMENT_COUNTER": IncrementCounter,
}

func IncrementCounter(args json.RawMessage, s gt.State) gt.Response {
	state := model(s)

	n, err := msg.DecodeInt(args)
	if err != nil {
		return gt.RespondWithError(err)
	}

	state.Counter = state.Counter + n
	return gt.Respond()
}

func renderCounter(counter int) h.Element {
	return h.Div(
		a.Attrs(a.Id("counter")),
		h.Button(
			a.Attrs(a.Class("counter-increment"), a.OnClick(gt.SendMessage("INCREMENT_COUNTER", -1))),
			h.Text("Down"),
		),
		h.Div(
			a.Attrs(a.Class("counter-readout")),
			h.Text(fmt.Sprintf("%d", counter)),
		),
		h.Button(
			a.Attrs(a.Class("counter-increment"), a.OnClick(gt.SendMessage("INCREMENT_COUNTER", 1))),
			h.Text("Up"),
		),
	)
}
