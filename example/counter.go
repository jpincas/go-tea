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
		a.Attrs(a.Id("counter"), a.Class("flex items-center space-x-4 p-4 bg-gray-50 rounded-lg shadow-sm w-fit")),
		h.Button(
			a.Attrs(a.Class("bg-red-500 hover:bg-red-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline transition duration-150 ease-in-out"), a.OnClick(gt.SendBasicMessage("INCREMENT_COUNTER", -1))),
			h.Text("Down"),
		),
		h.Div(
			a.Attrs(a.Class("text-2xl font-mono font-bold text-gray-800 w-12 text-center")),
			h.Text(fmt.Sprintf("%d", counter)),
		),
		h.Button(
			a.Attrs(a.Class("bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline transition duration-150 ease-in-out"), a.OnClick(gt.SendBasicMessage("INCREMENT_COUNTER", 1))),
			h.Text("Up"),
		),
	)
}
