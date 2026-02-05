package main

import (
	"fmt"

	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/go-tea/attributes"
	h "github.com/jpincas/go-tea/html"
)

var counterMessages gt.MessageMap = gt.MessageMap{
	"INCREMENT_COUNTER": IncrementCounter,
}

func IncrementCounter(m gt.Message, s gt.State) gt.Response {
	n := m.ArgsToInt()
	state := model(s)
	state.Counter = state.Counter + n
	return gt.Respond()
}

func renderCounter(counter int) h.Element {
	return h.Div(
		a.Attrs(a.Id("counter"), a.Class("flex items-center gap-3")),
		h.Button(
			a.Attrs(
				a.Class("w-12 h-12 bg-rose-500 hover:bg-rose-600 text-white font-bold text-xl rounded-xl border-2 border-stone-900 shadow-brutal-sm hover:shadow-brutal hover:-translate-x-0.5 hover:-translate-y-0.5 transition-all"),
				a.OnClick(gt.SendBasicMessage("INCREMENT_COUNTER", -1)),
			),
			h.Text("−"),
		),
		h.Div(
			a.Attrs(a.Class("w-20 h-12 flex items-center justify-center bg-white border-2 border-stone-900 rounded-xl shadow-brutal-sm"), a.Custom("style", "font-family: 'JetBrains Mono', monospace;")),
			h.Span(a.Attrs(a.Class("text-2xl font-bold text-stone-900")), h.Text(fmt.Sprintf("%d", counter))),
		),
		h.Button(
			a.Attrs(
				a.Class("w-12 h-12 bg-emerald-500 hover:bg-emerald-600 text-white font-bold text-xl rounded-xl border-2 border-stone-900 shadow-brutal-sm hover:shadow-brutal hover:-translate-x-0.5 hover:-translate-y-0.5 transition-all"),
				a.OnClick(gt.SendBasicMessage("INCREMENT_COUNTER", 1)),
			),
			h.Text("+"),
		),
	)
}
