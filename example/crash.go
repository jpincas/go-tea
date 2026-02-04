package main

import (
	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
)

var crashMessages gt.MessageMap = gt.MessageMap{
	"CRASH_ME": CrashMe,
}

func CrashMe(_ gt.Message, s gt.State) gt.Response {
	panic("I told you not to hit the big red button")
}

func renderCrash() h.Element {
	return h.Button(
		a.Attrs(
			a.Id("crash"),
			a.Class("inline-flex items-center gap-2 px-4 py-2 bg-rose-600 hover:bg-rose-700 text-white font-semibold rounded-lg border-2 border-stone-900 shadow-brutal-sm hover:shadow-brutal hover:-translate-x-0.5 hover:-translate-y-0.5 transition-all"),
			a.OnClick(gt.SendBasicMessageNoArgs("CRASH_ME")),
		),
		h.Span(a.Attrs(), h.Text("💥")),
		h.Text("Trigger Panic"),
	)
}
