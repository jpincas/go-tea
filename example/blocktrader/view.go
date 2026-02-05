package blocktrader

import (
	"fmt"
	"math"

	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/go-tea/attributes"
	"github.com/jpincas/go-tea/css"
	h "github.com/jpincas/go-tea/html"
)

func (m *Model) Render() h.Element {
	customCSS := `
.block {
	min-width: 0;
	min-height: 0;
}
`
	return h.Div(a.Attrs(
		a.Class("space-y-8")),
		h.Style(a.Attrs(), h.UnsafeRaw(customCSS)),

		// Header
		h.Div(a.Attrs(
			a.Class("text-center space-y-2")),
			h.H1(a.Attrs(
				a.Class("text-4xl font-bold text-stone-900"),
				a.Style(css.FontFamily("'DM Serif Display', serif"))),
				h.Text("📈 Blocktrader")),
			h.P(a.Attrs(
				a.Class("text-stone-600")),
				h.Text("Buy low, sell high! Watch the market shift as row and column values fluctuate."))),

		// Explanatory note
		renderExplanatoryNote(
			"How to Play",
			`
			<p class="mb-3">A port of a Javascript trading game to GoTea demonstrating server-side game loops.</p>
			<ul class="list-disc pl-5 space-y-2">
				<li><strong class="text-stone-900">Trading:</strong> Click blocks to buy (white) or sell (yellow). Block values = row × column multipliers.</li>
				<li><strong class="text-stone-900">Game Loop:</strong> The server simulates market changes using delayed messages at ~30fps.</li>
				<li><strong class="text-stone-900">State Persistence:</strong> Your portfolio persists to <code class="bg-stone-200 px-1.5 py-0.5 rounded text-xs font-mono">localStorage</code> across refreshes.</li>
			</ul>
			`),

		// Stats bar
		renderStatsBar(m),

		// Game board
		renderBoard(m))
}

func renderExplanatoryNote(title, content string) h.Element {
	return h.Details(a.Attrs(
		a.Class("bg-amber-50 border-2 border-stone-900 rounded-xl shadow-brutal-sm overflow-hidden")),
		h.Summary(a.Attrs(
			a.Class("cursor-pointer px-6 py-4 font-semibold text-stone-900 hover:bg-amber-100 transition-colors flex items-center gap-2")),
			h.Span(a.Attrs(
				a.Class("text-lg")),
				h.Text("📖")),
			h.Text(title)),
		h.Div(a.Attrs(
			a.Class("px-6 pb-5 text-stone-700 text-sm leading-relaxed border-t-2 border-stone-200 pt-4")),
			h.UnsafeRaw(content)))
}

func renderStatsBar(m *Model) h.Element {
	portfolioVal := m.portfolioValue()
	totalVal := m.State.Cash + portfolioVal
	profitLoss := totalVal - m.Config.StartingCash

	return h.Div(a.Attrs(
		a.Class("bg-gradient-to-r from-amber-50 to-orange-50 p-5 rounded-xl border-2 border-stone-900 shadow-brutal-sm")),
		h.Div(a.Attrs(
			a.Class("flex flex-wrap items-center justify-between gap-4")),

			// KPIs
			h.Div(a.Attrs(
				a.Class("flex flex-wrap gap-6")),
				kpi("Cash", m.State.Cash, "text-stone-900"),
				kpi("Portfolio", portfolioVal, "text-violet-700"),
				kpi("Total Value", totalVal, "text-stone-900"),
				kpiWithColor("P/L", profitLoss)),

			// Controls
			renderControls(m)))
}

func kpi(label string, value float64, colorClass string) h.Element {
	return h.Div(a.Attrs(
		a.Class("flex flex-col")),
		h.Span(a.Attrs(
			a.Class("text-xs text-stone-500 uppercase font-semibold tracking-wide")),
			h.Text(label)),
		h.Span(a.Attrs(
			a.Class(fmt.Sprintf("text-xl font-bold %s", colorClass)),
			a.Style(css.FontFamily("'JetBrains Mono', monospace"))),
			h.Text(fmt.Sprintf("$%.0f", value))))
}

func kpiWithColor(label string, value float64) h.Element {
	colorClass := "text-stone-900"
	prefix := ""
	if value > 0 {
		colorClass = "text-emerald-600"
		prefix = "+"
	} else if value < 0 {
		colorClass = "text-rose-600"
	}

	return h.Div(a.Attrs(
		a.Class("flex flex-col")),
		h.Span(a.Attrs(
			a.Class("text-xs text-stone-500 uppercase font-semibold tracking-wide")),
			h.Text(label)),
		h.Span(a.Attrs(
			a.Class(fmt.Sprintf("text-xl font-bold %s", colorClass)),
			a.Style(css.FontFamily("'JetBrains Mono', monospace"))),
			h.Text(fmt.Sprintf("%s$%.0f", prefix, math.Abs(value)))))
}

func renderControls(m *Model) h.Element {
	return h.Div(a.Attrs(
		a.Class("flex items-center gap-3")),

		// Timer display
		h.Div(a.Attrs(
			a.Class("flex items-center gap-2 px-4 py-2 bg-stone-900 rounded-lg")),
			h.Span(a.Attrs(
				a.Class("text-xs text-stone-400 uppercase")),
				h.Text("Time")),
			h.Span(a.Attrs(
				a.Class("text-2xl font-bold text-white tabular-nums"),
				a.Style(css.FontFamily("'JetBrains Mono', monospace"))),
				h.Text(fmt.Sprintf("%d", m.State.Timer)))),

		// Start button
		h.Button(a.Attrs(
			a.Class("px-5 py-2.5 bg-emerald-500 hover:bg-emerald-600 text-white font-semibold rounded-xl border-2 border-stone-900 shadow-brutal-sm hover:shadow-brutal hover:-translate-x-0.5 hover:-translate-y-0.5 transition-all"),
			a.OnClick(gt.SendBasicMessage("BLOCKTRADER_START_GAME", nil))),
			h.Text("▶ Start")).RenderIf(!m.State.GameActive),

		// Stop button
		h.Button(a.Attrs(
			a.Class("px-5 py-2.5 bg-rose-500 hover:bg-rose-600 text-white font-semibold rounded-xl border-2 border-stone-900 shadow-brutal-sm hover:shadow-brutal hover:-translate-x-0.5 hover:-translate-y-0.5 transition-all"),
			a.OnClick(gt.SendBasicMessage("BLOCKTRADER_STOP_GAME", nil))),
			h.Text("⏸ Pause")).RenderIf(m.State.GameActive),

		// Reset button
		h.Button(a.Attrs(
			a.Class("px-5 py-2.5 bg-stone-500 hover:bg-stone-600 text-white font-semibold rounded-xl border-2 border-stone-900 shadow-brutal-sm hover:shadow-brutal hover:-translate-x-0.5 hover:-translate-y-0.5 transition-all"),
			a.OnClick(gt.SendBasicMessage("BLOCKTRADER_RESET_GAME", nil))),
			h.Text("↺ Reset")))
}

func renderBoard(m *Model) h.Element {
	var rows h.Elements
	for r := 0; r < m.Config.BoardSize; r++ {
		var cols h.Elements
		for c := 0; c < m.Config.BoardSize; c++ {
			cols = append(cols, renderBlock(m, r, c))
		}
		rows = append(rows, h.Div(a.Attrs(
			a.Class("flex flex-row w-full"),
			a.Style(css.Flex_(fmt.Sprintf("%.2f", m.State.Rows[r])))),
			cols...))
	}

	return h.Div(a.Attrs(
		a.Class("bg-white rounded-xl border-2 border-stone-900 shadow-brutal overflow-hidden")),
		h.Div(a.Attrs(
			a.Class("flex flex-col w-full h-[600px] p-2")),
			rows...))
}

func renderBlock(m *Model, row, col int) h.Element {
	val := m.blockValue(row, col)
	if val == 0 {
		return h.Div(a.Attrs(a.Class("flex-grow")))
	}

	key := blockKey(row, col)
	boughtAt, owned := m.State.OwnedBlocks[key]
	profit := val - boughtAt
	canAfford := m.State.Cash >= val

	// Base classes
	classes := "block relative m-0.5 cursor-pointer transition-all duration-150 select-none flex items-center justify-center overflow-hidden rounded-lg border-2"

	if owned {
		// Owned blocks - amber/yellow theme
		classes += " bg-gradient-to-br from-amber-100 to-yellow-200 border-amber-400"
		if profit > 0 {
			classes += " ring-2 ring-emerald-400"
		} else if profit < 0 {
			classes += " ring-2 ring-rose-400"
		}
	} else {
		if canAfford {
			classes += " bg-stone-50 border-stone-300 hover:bg-stone-100 hover:border-stone-400 hover:scale-[1.02]"
		} else {
			classes += " bg-stone-200 border-stone-300 opacity-40 cursor-not-allowed"
		}
	}

	// Highlight recently mutated row/col
	justMutated := m.State.GameActive &&
		((m.State.RowColumn == 1 && m.State.Member == col) ||
			(m.State.RowColumn == 0 && m.State.Member == row))

	if justMutated {
		classes += " ring-2 ring-violet-500 ring-offset-1 z-10"
	}

	fontSize := math.Min(1.0, (val/100.0)+0.25)

	// Determine text colors based on ownership and profit
	valueColor := "text-stone-700"
	profitColor := "text-stone-500"
	if owned {
		valueColor = "text-stone-900"
		if profit > 0 {
			profitColor = "text-emerald-600 font-semibold"
		} else if profit < 0 {
			profitColor = "text-rose-600 font-semibold"
		}
	}

	return h.Div(a.Attrs(
		a.Class(classes),
		a.Style(
			css.Flex_(fmt.Sprintf("%.2f", m.State.Cols[col])),
			css.FontSize(fmt.Sprintf("%.2fem", fontSize))),
		a.OnClick(gt.SendMessage(gt.Message{
			Message:   "BLOCKTRADER_TOGGLE_BLOCK",
			Arguments: []interface{}{float64(row), float64(col)},
		}))),
		h.Div(a.Attrs(
			a.Class("text-center"),
			a.Style(css.FontFamily("'JetBrains Mono', monospace"))),
			h.Span(a.Attrs(
				a.Class(fmt.Sprintf("font-bold %s", valueColor))),
				h.Text(fmt.Sprintf("%.0f", math.Floor(val)))),
			h.Div(a.Attrs(
				a.Class(fmt.Sprintf("text-xs %s", profitColor))),
				h.Text(fmt.Sprintf("%+.0f", math.Floor(profit)))).RenderIf(owned)))
}
