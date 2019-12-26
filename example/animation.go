package main

import (
	"encoding/json"
	"math/rand"

	gt "github.com/jpincas/go-tea"
)

type Animation struct {
	Stop                     bool
	X, Y                     int
	XDirection, YDirection   bool
	BackgroundSize, BallSize int
	TranslateX, TranslateY   int
	IncrementX, IncrementY   int
}

var animationMessages gt.MessageMap = gt.MessageMap{
	"START_ANIMATION": StartAnimation,
	"STOP_ANIMATION":  StopAnimation,
	"RESET_ANIMATION": ResetAnimation,
}

func StartAnimation(_ json.RawMessage, s gt.State) gt.Response {
	state := model(s)

	if state.Animation.Stop {
		state.Animation.Stop = false
		return gt.Respond()
	}

	if state.Animation.X >= 100 {
		state.Animation.XDirection = false
		state.Animation.IncrementX = rand.Intn(5)
	} else if state.Animation.X <= 0 {
		state.Animation.XDirection = true
		state.Animation.IncrementX = rand.Intn(5)
	}

	if state.Animation.Y >= 100 {
		state.Animation.YDirection = false
		state.Animation.IncrementY = rand.Intn(5)
	} else if state.Animation.Y <= 0 {
		state.Animation.YDirection = true
		state.Animation.IncrementY = rand.Intn(5)
	}

	if state.Animation.XDirection {
		state.Animation.X = state.Animation.X + state.Animation.IncrementX
	} else {
		state.Animation.X = state.Animation.X - state.Animation.IncrementX
	}

	if state.Animation.YDirection {
		state.Animation.Y = state.Animation.Y + state.Animation.IncrementY
	} else {
		state.Animation.Y = state.Animation.Y - state.Animation.IncrementY
	}

	state.Animation.TranslateX = translate(state.Animation.X, state.Animation.BackgroundSize, state.Animation.BallSize)
	state.Animation.TranslateY = translate(state.Animation.Y, state.Animation.BackgroundSize, state.Animation.BallSize)

	return gt.RespondWithDelayedNextMsg("START_ANIMATION", nil, 33)
}

func StopAnimation(_ json.RawMessage, s gt.State) gt.Response {
	state := model(s)
	state.Animation.Stop = true
	return gt.Respond()
}

func ResetAnimation(_ json.RawMessage, s gt.State) gt.Response {
	state := model(s)
	state.Animation.Stop = true
	state.Animation.X = 50
	state.Animation.Y = 50
	state.Animation.XDirection = true
	state.Animation.YDirection = true
	state.Animation.TranslateX = translate(state.Animation.X, state.Animation.BackgroundSize, state.Animation.BallSize)
	state.Animation.TranslateY = translate(state.Animation.Y, state.Animation.BackgroundSize, state.Animation.BallSize)
	return gt.Respond()
}

func translate(co, backgroundSize, ballSize int) int {
	return int((float64(co) / float64(100)) * float64(backgroundSize-ballSize))
}
