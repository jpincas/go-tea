package main

import (
	"encoding/json"

	gt "github.com/jpincas/go-tea"
)

type Animation struct {
	X, Y                   int
	XDirection, YDirection bool
}

var animationMessages gt.MessageMap = gt.MessageMap{
	"START_ANIMATION": StartAnimation,
}

func StartAnimation(_ json.RawMessage, s gt.State) gt.Response {
	state := model(s)

	if state.Animation.X == 100 {
		state.Animation.XDirection = false
	} else if state.Animation.X == 0 {
		state.Animation.XDirection = true
	}

	if state.Animation.Y == 100 {
		state.Animation.YDirection = false
	} else if state.Animation.Y == 0 {
		state.Animation.YDirection = true
	}

	if state.Animation.XDirection {
		state.Animation.X++
	} else {
		state.Animation.X--
	}

	if state.Animation.YDirection {
		state.Animation.Y++
	} else {
		state.Animation.Y--
	}

	return gt.RespondWithDelayedNextMsg("START_ANIMATION", nil, 500)
}
