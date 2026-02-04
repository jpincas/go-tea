package main

import (
	"fmt"
	"sync"

	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/htmlfunc/attributes"
	"github.com/jpincas/htmlfunc/css"
	h "github.com/jpincas/htmlfunc/html"
)

// Canvas dimensions
const canvasSize = 32

// Color palette
var colorPalette = []string{
	"#000000", // black
	"#FFFFFF", // white
	"#EF4444", // red
	"#F97316", // orange
	"#EAB308", // yellow
	"#22C55E", // green
	"#3B82F6", // blue
	"#A855F7", // purple
}

// Global shared state for the canvas
var canvasPixels = make(map[string]string) // "x:y" -> "#hexcolor"
var canvasMutex sync.Mutex

// PixelCanvas holds the per-session state for the pixel canvas feature
type PixelCanvas struct {
	SelectedColor string
}

func NewPixelCanvas() PixelCanvas {
	return PixelCanvas{
		SelectedColor: "#000000", // Default to black
	}
}

func pixelKey(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

// Message handlers
var pixelCanvasMessages = gt.MessageMap{
	"PAINT_PIXEL":  paintPixel,
	"SELECT_COLOR": selectColor,
	"CLEAR_CANVAS": clearCanvas,
}

func paintPixel(m gt.Message, s gt.State) gt.Response {
	state := model(s)

	var args struct {
		X int `json:"x"`
		Y int `json:"y"`
	}
	m.MustDecodeArgs(&args)

	canvasMutex.Lock()
	canvasPixels[pixelKey(args.X, args.Y)] = state.PixelCanvas.SelectedColor
	canvasMutex.Unlock()

	app.Broadcast()
	return gt.Respond()
}

func selectColor(m gt.Message, s gt.State) gt.Response {
	state := model(s)
	color := m.ArgsToString()

	// Validate color is in palette
	for _, c := range colorPalette {
		if c == color {
			state.PixelCanvas.SelectedColor = color
			break
		}
	}

	return gt.Respond()
}

func clearCanvas(m gt.Message, s gt.State) gt.Response {
	canvasMutex.Lock()
	canvasPixels = make(map[string]string)
	canvasMutex.Unlock()

	app.Broadcast()
	return gt.Respond()
}

// Rendering
func (pc *PixelCanvas) render() h.Element {
	return h.Div(
		a.Attrs(a.Class("space-y-8")),

		// Header
		h.Div(
			a.Attrs(a.Class("text-center space-y-2")),
			h.H1(a.Attrs(a.Class("text-4xl font-bold text-stone-900"), a.Custom("style", "font-family: 'DM Serif Display', serif;")), h.Text("🎨 Collaborative Pixel Canvas")),
			h.P(a.Attrs(a.Class("text-stone-600")),
				h.Text("Paint pixels together in real-time! Open in multiple tabs to see changes sync instantly."),
			),
		),

		// Color palette
		h.Div(
			a.Attrs(a.Class("bg-gradient-to-r from-pink-50 to-purple-50 p-5 rounded-xl border-2 border-stone-900 shadow-brutal-sm")),
			h.Div(
				a.Attrs(a.Class("flex flex-wrap items-center justify-between gap-4")),
				h.Div(
					a.Attrs(a.Class("flex items-center gap-3")),
					h.Span(a.Attrs(a.Class("text-sm font-semibold text-stone-700")), h.Text("Color:")),
					h.Div(
						a.Attrs(a.Class("flex gap-2")),
						renderColorPalette(pc.SelectedColor)...,
					),
				),
				h.Div(
					a.Attrs(a.Class("flex items-center gap-2")),
					h.Span(a.Attrs(a.Class("text-sm text-stone-500")), h.Text("Selected:")),
					h.Div(
						a.Attrs(
							a.Class("w-8 h-8 rounded-lg border-2 border-stone-900 shadow-brutal-sm"),
							a.Custom("style", fmt.Sprintf("background-color: %s", pc.SelectedColor)),
						),
					),
					h.Span(a.Attrs(a.Class("text-xs text-stone-400"), a.Custom("style", "font-family: 'JetBrains Mono', monospace;")), h.Text(pc.SelectedColor)),
				),
			),
		),

		// Canvas
		h.Div(
			a.Attrs(a.Class("flex justify-center")),
			h.Div(
				a.Attrs(
					a.Class("inline-block border-2 border-stone-900 rounded-xl overflow-hidden shadow-brutal"),
					a.Custom("style", "display: grid; grid-template-columns: repeat(32, 1fr); gap: 0; background: white;"),
				),
				renderCanvasGrid()...,
			),
		),

		// Clear button
		h.Div(
			a.Attrs(a.Class("flex justify-center")),
			h.Button(
				a.Attrs(
					a.OnClick(gt.SendBasicMessageNoArgs("CLEAR_CANVAS")),
					a.Class("inline-flex items-center gap-2 px-5 py-2.5 bg-rose-500 hover:bg-rose-600 text-white font-semibold rounded-xl border-2 border-stone-900 shadow-brutal-sm hover:shadow-brutal hover:-translate-x-0.5 hover:-translate-y-0.5 transition-all"),
				),
				h.Span(a.Attrs(), h.Text("🗑")),
				h.Text("Clear Canvas"),
			),
		),

		// Instructions
		renderExplanatoryNote(
			"How it works",
			`
			<ul class="list-disc pl-5 space-y-2">
				<li><strong class="text-stone-900">Shared State:</strong> The canvas is stored on the server and shared by all connected users.</li>
				<li><strong class="text-stone-900">Broadcasting:</strong> When any user paints a pixel, <code class="bg-stone-200 px-1.5 py-0.5 rounded text-xs font-mono">app.Broadcast()</code> re-renders for all clients.</li>
				<li><strong class="text-stone-900">Concurrency:</strong> A mutex protects the canvas map from race conditions.</li>
				<li><strong class="text-stone-900">Per-session:</strong> Each user has their own selected color stored in session state.</li>
			</ul>
			`,
		),
	)
}

func renderColorPalette(selectedColor string) []h.Element {
	var elements []h.Element
	for _, color := range colorPalette {
		selectedClass := ""
		if color == selectedColor {
			selectedClass = " ring-2 ring-offset-2 ring-stone-900 scale-110"
		}

		borderClass := "border-stone-400"
		if color == "#FFFFFF" {
			borderClass = "border-stone-500"
		}

		elements = append(elements, h.Button(
			a.Attrs(
				a.OnClick(gt.SendBasicMessage("SELECT_COLOR", color)),
				a.Class(fmt.Sprintf("w-8 h-8 rounded-lg cursor-pointer border-2 %s hover:scale-110 transition-all%s", borderClass, selectedClass)),
				a.Custom("style", fmt.Sprintf("background-color: %s", color)),
				a.Title(color),
			),
		))
	}
	return elements
}

func renderCanvasGrid() []h.Element {
	canvasMutex.Lock()
	defer canvasMutex.Unlock()

	var elements []h.Element
	for y := 0; y < canvasSize; y++ {
		for x := 0; x < canvasSize; x++ {
			color := canvasPixels[pixelKey(x, y)]
			if color == "" {
				color = "#FFFFFF" // Default to white
			}

			elements = append(elements, h.Div(
				a.Attrs(
					a.Class("cursor-crosshair hover:opacity-75 transition-opacity"),
					a.Style(
						css.Width("12px"),
						css.Height("12px"),
						css.BackgroundColor(color),
					),
					a.OnClick(gt.SendBasicMessage("PAINT_PIXEL", map[string]int{"x": x, "y": y})),
				),
			))
		}
	}
	return elements
}
