package blocktrader

import (
	"fmt"
	"math"

	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
)

func (m *Model) Render() h.Element {
	css := `
.blocktrader-container {
	font-family: 'Courier New', Courier, monospace;
}
.block {
	min-width: 0;
	min-height: 0;
}
`
	return h.Div(
		a.Attrs(a.Class("blocktrader-container p-4")),
		h.Style(a.Attrs(), h.UnsafeRaw(css)),
		renderExplanatoryNote(
			"Blocktrader Game",
			`
			<p class="mb-2">A port of a Javascript game to GoTea.</p>
			<ul class="list-disc pl-5 space-y-1">
				<li><strong>Game Loop:</strong> The game loop is simulated on the server using delayed messages.</li>
				<li><strong>State Persistence:</strong> The game state is persisted to <code>localStorage</code> and restored on reload.</li>
				<li><strong>Complex UI:</strong> The UI is built using <code>htmlfunc</code> with dynamic styles for the grid.</li>
			</ul>
			`,
		),
		renderHeader(m),
		renderBoard(m),
	)
}

func renderExplanatoryNote(title, content string) h.Element {
	return h.Details(
		a.Attrs(a.Class("bg-blue-50 border border-blue-200 rounded-md mb-6")),
		h.Summary(
			a.Attrs(a.Class("cursor-pointer p-4 font-bold text-blue-900 hover:bg-blue-100 rounded-t-md outline-none")),
			h.Text(fmt.Sprintf("ℹ️ %s", title)),
		),
		h.Div(
			a.Attrs(a.Class("p-4 pt-0 text-blue-800 text-sm leading-relaxed")),
			h.UnsafeRaw(content),
		),
	)
}

func renderHeader(m *Model) h.Element {
	portfolioVal := m.portfolioValue()
	totalVal := m.State.Cash + portfolioVal
	profitLoss := totalVal - m.Config.StartingCash

	return h.Div(
		a.Attrs(a.Class("flex flex-col space-y-4 mb-4")),
		h.Div(
			a.Attrs(a.Class("flex justify-between items-center")),
			h.Div(
				a.Attrs(a.Class("flex space-x-6")),
				kpi("Cash", m.State.Cash),
				kpi("Portfolio", portfolioVal),
				kpi("Total Value", totalVal),
				kpi("Profit/Loss", profitLoss),
			),
			renderControls(m),
		),
	)
}

func kpi(label string, value float64) h.Element {
	return h.Div(
		a.Attrs(a.Class("flex flex-col")),
		h.Span(a.Attrs(a.Class("text-xs text-gray-500 uppercase")), h.Text(label)),
		h.Span(a.Attrs(a.Class("text-xl font-bold font-mono")), h.Text(fmt.Sprintf("$%.0f", value))),
	)
}

func renderControls(m *Model) h.Element {
	return h.Div(
		a.Attrs(a.Class("flex items-center space-x-4")),
		h.Div(
			a.Attrs(a.Class("text-2xl font-mono font-bold")),
			h.Text(fmt.Sprintf("%d", m.State.Timer)),
		),
		h.Button(
			a.Attrs(
				a.Class("bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"),
				a.OnClick(gt.SendBasicMessage("BLOCKTRADER_START_GAME", nil)),
			),
			h.Text("Start"),
		).RenderIf(!m.State.GameActive),
		h.Button(
			a.Attrs(
				a.Class("bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"),
				a.OnClick(gt.SendBasicMessage("BLOCKTRADER_STOP_GAME", nil)),
			),
			h.Text("Stop"),
		).RenderIf(m.State.GameActive),
		h.Button(
			a.Attrs(
				a.Class("bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded"),
				a.OnClick(gt.SendBasicMessage("BLOCKTRADER_RESET_GAME", nil)),
			),
			h.Text("Reset"),
		),
	)
}

func renderBoard(m *Model) h.Element {
	var rows h.Elements
	for r := 0; r < m.Config.BoardSize; r++ {
		var cols h.Elements
		for c := 0; c < m.Config.BoardSize; c++ {
			cols = append(cols, renderBlock(m, r, c))
		}
		rows = append(rows, h.Div(
			a.Attrs(a.Class("flex flex-row w-full"), a.Custom("style", fmt.Sprintf("flex: %.2f", m.State.Rows[r]))),
			cols...,
		))
	}

	return h.Div(
		a.Attrs(a.Class("flex flex-col w-full h-[600px] border-2 border-gray-300")),
		rows...,
	)
}

func renderBlock(m *Model, row, col int) h.Element {
	val := m.blockValue(row, col)
	if val == 0 {
		return h.Div(a.Attrs(a.Class("flex-grow"))) // Invisible placeholder
	}

	key := blockKey(row, col)
	boughtAt, owned := m.State.OwnedBlocks[key]
	profit := val - boughtAt
	
	canAfford := m.State.Cash >= val
	
	classes := "block relative border border-gray-400 m-0.5 cursor-pointer transition-all duration-200 select-none flex items-center justify-center overflow-hidden group"
	if owned {
		classes += " bg-yellow-200"
		if profit > 0 {
			classes += " text-green-700"
		} else if profit < 0 {
			classes += " text-red-700"
		}
	} else {
		if canAfford {
			classes += " bg-gray-100 hover:bg-gray-200"
		} else {
			classes += " bg-gray-300 opacity-50 cursor-not-allowed"
		}
	}

	// Highlight recently mutated
	justMutated := m.State.GameActive && 
		((m.State.RowColumn == 1 && m.State.Member == col) || 
		 (m.State.RowColumn == 0 && m.State.Member == row))
	
	if justMutated {
		classes += " ring-2 ring-blue-400 z-10"
	}

	fontSize := math.Min(1.0, (val/100.0)+0.2)

	return h.Div(
		a.Attrs(
			a.Class(classes),
			a.Custom("style", fmt.Sprintf("flex: %.2f; font-size: %.2fem", m.State.Cols[col], fontSize)),
			a.OnClick(gt.SendMessage(gt.Message{
				Message: "BLOCKTRADER_TOGGLE_BLOCK",
				Arguments: []interface{}{float64(row), float64(col)},
			})),
		),
		h.Div(
			a.Attrs(a.Class("text-center")),
			h.Span(a.Attrs(a.Class("font-bold")), h.Text(fmt.Sprintf("%.0f", math.Floor(val)))),
			h.Div(
				a.Attrs(a.Class("text-xs")),
				h.Text(fmt.Sprintf("%+.0f", math.Floor(profit))),
			).RenderIf(owned),
		),
	)
}

const css = `
.block {
	min-width: 0;
	min-height: 0;
}
`
