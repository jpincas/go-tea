package main

import (
	"encoding/json"

	gt "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/msg"
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
