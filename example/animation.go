package main

import (
	"fmt"
	"math/rand"

	gotea "github.com/jpincas/go-tea"
	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/htmlfunc/attributes"
	"github.com/jpincas/htmlfunc/css"
	h "github.com/jpincas/htmlfunc/html"
)

const (
	animationBackgroundSize = 500
	animationBallSize       = 20
	animationFrameDelay     = 33
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
	"START_ANIMATION":      StartAnimation,
	"NEXT_ANIMATION_FRAME": NextAnimationFrame,
	"STOP_ANIMATION":       StopAnimation,
	"RESET_ANIMATION":      ResetAnimation,
}

func StartAnimation(_ gotea.Message, s gt.State) gt.Response {
	state := model(s)
	state.Animation.Stop = false
	return gt.RespondWithDelayedNextMsg(gt.Message{Message: "NEXT_ANIMATION_FRAME"}, animationFrameDelay)
}

func NextAnimationFrame(_ gotea.Message, s gt.State) gt.Response {
	state := model(s)

	if state.Animation.Stop {
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

	return gt.RespondWithDelayedNextMsg(gt.Message{Message: "NEXT_ANIMATION_FRAME"}, animationFrameDelay)
}

func StopAnimation(_ gotea.Message, s gt.State) gt.Response {
	state := model(s)
	state.Animation.Stop = true
	return gt.Respond()
}

func ResetAnimation(_ gotea.Message, s gt.State) gt.Response {
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
		a.Attrs(a.Class("space-y-6")),
		renderExplanatoryNote(
			"Server-Side Animation",
			`
			<p class="mb-2">This demonstrates a continuous animation loop driven by the server.</p>
			<ul class="list-disc pl-5 space-y-1">
				<li><strong>Game Loop:</strong> The server sends a <code>TICK</code> message to itself repeatedly using <code>RespondWithDelayedNextMsg</code>.</li>
				<li><strong>State Update:</strong> On each tick, the ball's position is updated in the state.</li>
				<li><strong>Rendering:</strong> The new state is rendered and sent to the client. GoTea's virtual DOM diffing ensures only the changed attributes (style) are updated in the DOM, resulting in smooth animation.</li>
			</ul>
			`,
		),
		h.H2(a.Attrs(a.Class("text-2xl font-bold text-gray-900")), h.Text("A Server-Driven Animation")),
		h.P(a.Attrs(a.Class("text-gray-600")), h.Text("A 30fps bouncing ball animation being driven entirely by the server. You'd probably never want to do this, but is fun to know that you can! Clicking 'start' fires off a never-ending sequence of messages with a 33ms delay between each. Each message sets the x and y coordinates of the ball and the scene is rerendered by the Gotea runtime.")),
		
		h.Div(
			a.Attrs(a.Class("flex space-x-4")),
			h.Button(
				a.Attrs(a.Class("inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"), a.OnClick(gt.SendBasicMessageNoArgs("START_ANIMATION"))),
				h.Text("Start Animation"),
			),
			h.Button(
				a.Attrs(a.Class("inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"), a.OnClick(gt.SendBasicMessageNoArgs("STOP_ANIMATION"))),
				h.Text("Stop Animation"),
			),
			h.Button(
				a.Attrs(a.Class("inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md shadow-sm text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"), a.OnClick(gt.SendBasicMessageNoArgs("RESET_ANIMATION"))),
				h.Text("Reset Animation"),
			),
		),

		h.P(a.Attrs(a.Class("font-mono text-sm text-gray-500")), h.Text(fmt.Sprintf("Coordinates: (%d, %d)", animation.X, animation.Y))),
		
		h.Div(
			a.Attrs(
				a.Id("animation-background"), 
				a.Class("relative bg-gray-200 rounded-lg border border-gray-300 overflow-hidden"),
				a.Style(css.Width(fmt.Sprintf("%dpx", animationBackgroundSize)), css.Height(fmt.Sprintf("%dpx", animationBackgroundSize))),
			),
			h.Div(
				a.Attrs(
					a.Id("animation-ball"), 
					a.Class("absolute bg-indigo-600 rounded-full shadow-lg"),
					a.Style(
						css.Width(fmt.Sprintf("%dpx", animationBallSize)), 
						css.Height(fmt.Sprintf("%dpx", animationBallSize)),
						css.Transform(fmt.Sprintf("translate(%dpx, %dpx)", animation.TranslateX, animation.TranslateY)),
					),
				),
			),
		),
	)
}
