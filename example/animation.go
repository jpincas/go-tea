package main

import (
	"fmt"
	"math/rand"

	gotea "github.com/jpincas/go-tea"
	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/htmlfunc/attributes"
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
		a.Attrs(a.Class("space-y-8")),
		// Header
		h.Div(
			a.Attrs(a.Class("text-center space-y-2")),
			h.H1(a.Attrs(a.Class("text-4xl font-bold text-stone-900"), a.Custom("style", "font-family: 'DM Serif Display', serif;")), h.Text("🏀 Server-Driven Animation")),
			h.P(a.Attrs(a.Class("text-stone-600 max-w-2xl mx-auto")), h.Text("A 30fps bouncing ball animation driven entirely by the server. Each frame is a WebSocket message that updates the state and re-renders.")),
		),

		// Controls
		h.Div(
			a.Attrs(a.Class("flex flex-wrap justify-center gap-3")),
			h.Button(
				a.Attrs(
					a.Class("inline-flex items-center gap-2 px-5 py-2.5 bg-emerald-500 hover:bg-emerald-600 text-white font-semibold rounded-xl border-2 border-stone-900 shadow-brutal-sm hover:shadow-brutal hover:-translate-x-0.5 hover:-translate-y-0.5 transition-all"),
					a.OnClick(gt.SendBasicMessageNoArgs("START_ANIMATION")),
				),
				h.Span(a.Attrs(), h.Text("▶")),
				h.Text("Start"),
			),
			h.Button(
				a.Attrs(
					a.Class("inline-flex items-center gap-2 px-5 py-2.5 bg-rose-500 hover:bg-rose-600 text-white font-semibold rounded-xl border-2 border-stone-900 shadow-brutal-sm hover:shadow-brutal hover:-translate-x-0.5 hover:-translate-y-0.5 transition-all"),
					a.OnClick(gt.SendBasicMessageNoArgs("STOP_ANIMATION")),
				),
				h.Span(a.Attrs(), h.Text("⏸")),
				h.Text("Stop"),
			),
			h.Button(
				a.Attrs(
					a.Class("inline-flex items-center gap-2 px-5 py-2.5 bg-white hover:bg-stone-50 text-stone-700 font-semibold rounded-xl border-2 border-stone-900 shadow-brutal-sm hover:shadow-brutal hover:-translate-x-0.5 hover:-translate-y-0.5 transition-all"),
					a.OnClick(gt.SendBasicMessageNoArgs("RESET_ANIMATION")),
				),
				h.Span(a.Attrs(), h.Text("↺")),
				h.Text("Reset"),
			),
		),

		// Coordinates display
		h.Div(
			a.Attrs(a.Class("text-center")),
			h.Span(
				a.Attrs(a.Class("inline-block px-4 py-2 bg-stone-100 rounded-lg border-2 border-stone-300 text-sm"), a.Custom("style", "font-family: 'JetBrains Mono', monospace;")),
				h.Text(fmt.Sprintf("x: %d, y: %d", animation.X, animation.Y)),
			),
		),

		// Animation canvas
		h.Div(
			a.Attrs(a.Class("flex justify-center")),
			h.Div(
				a.Attrs(
					a.Id("animation-background"),
					a.Class("relative rounded-2xl border-2 border-stone-900 shadow-brutal overflow-hidden"),
					a.Custom("style", fmt.Sprintf("width: %dpx; height: %dpx; background: linear-gradient(135deg, #fef3c7 0%%, #fde68a 50%%, #fcd34d 100%%);", animationBackgroundSize, animationBackgroundSize)),
				),
				// Ball
				h.Div(
					a.Attrs(
						a.Id("animation-ball"),
						a.Class("absolute rounded-full border-2 border-stone-900"),
						a.Custom("style", fmt.Sprintf("width: %dpx; height: %dpx; transform: translate(%dpx, %dpx); background: linear-gradient(135deg, #f97316 0%%, #ea580c 100%%); box-shadow: inset -4px -4px 0px rgba(0,0,0,0.2);", animationBallSize, animationBallSize, animation.TranslateX, animation.TranslateY)),
					),
				),
			),
		),

		// Explanatory note
		renderExplanatoryNote(
			"Server-Side Animation",
			`
			<p class="mb-3">This demonstrates a continuous animation loop driven by the server.</p>
			<ul class="list-disc pl-5 space-y-2">
				<li><strong class="text-stone-900">Game Loop:</strong> The server sends a <code class="bg-stone-200 px-1.5 py-0.5 rounded text-xs font-mono">NEXT_ANIMATION_FRAME</code> message to itself repeatedly using <code class="bg-stone-200 px-1.5 py-0.5 rounded text-xs font-mono">RespondWithDelayedNextMsg</code>.</li>
				<li><strong class="text-stone-900">State Update:</strong> On each tick, the ball's position is updated in the state.</li>
				<li><strong class="text-stone-900">Rendering:</strong> The new state is rendered and sent to the client. Morphdom ensures only the changed attributes (style) are updated.</li>
			</ul>
			`,
		),
	)
}
