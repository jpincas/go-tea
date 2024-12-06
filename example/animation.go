package main

import (
	"encoding/json"
	"fmt"
	"math/rand"

	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/htmlfunc/attributes"
	"github.com/jpincas/htmlfunc/css"
	h "github.com/jpincas/htmlfunc/html"
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

func (animation Animation) render() h.Element {
	return h.Div(
		a.Attrs(),
		h.H2(a.Attrs(), h.Text("A Server-Driven Animation")),
		h.P(a.Attrs(), h.Text("A 30fps bouncing ball animation being driven entirely by the server. You'd probably never want to do this, but is fun to know that you can! Clicking 'start' fires off a never-ending sequence of messages with a 33ms delay between each. Each message sets the x and y coordinates of the ball and the scene is rerendered by the Gotea runtime.")),
		h.Button(
			a.Attrs(a.OnClick(gt.SendMessageNoArgs("START_ANIMATION"))),
			h.Text("Start Animation"),
		),
		h.Button(
			a.Attrs(a.OnClick(gt.SendMessageNoArgs("STOP_ANIMATION"))),
			h.Text("Stop Animation"),
		),
		h.Button(
			a.Attrs(a.OnClick(gt.SendMessageNoArgs("RESET_ANIMATION"))),
			h.Text("Reset Animation"),
		),
		h.P(a.Attrs(), h.Text(fmt.Sprintf("Coordinates: (%d, %d)", animation.X, animation.Y))),
		h.Div(
			a.Attrs(a.Id("animation-background")),
			h.Div(
				a.Attrs(a.Id("animation-ball"), a.Style(
					css.Transform(fmt.Sprintf("translate(%dpx, %dpx)", animation.TranslateX, animation.TranslateY)),
				)),
			),
		),
	)
}
