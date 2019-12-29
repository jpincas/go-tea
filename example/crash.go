package main

import (
	"encoding/json"

	gt "github.com/jpincas/go-tea"
)

var crashMessages gt.MessageMap = gt.MessageMap{
	"CRASH_ME": CrashMe,
}

func CrashMe(args json.RawMessage, s gt.State) gt.Response {
	panic("I told you not to hit the big red button")
	return gt.Respond()
}
