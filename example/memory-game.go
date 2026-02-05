package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/go-tea/attributes"
	"github.com/jpincas/go-tea/css"
	h "github.com/jpincas/go-tea/html"
)

var memoryGameMessages gt.MessageMap = gt.MessageMap{
	"FLIP_CARD":         FlipCard,
	"REMOVE_MATCHES":    RemoveMatches,
	"FLIP_ALL_BACK":     FlipAllBack,
	"RESTART_GAME":      RestartGame,
	"CHANGE_DIFFICULTY": ChangeDifficulty,
	"SUBMIT_INITIALS":   SubmitInitials,
}

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
)

func (d Difficulty) String() string {
	return [...]string{"Easy", "Medium", "Hard"}[d]
}

func (d Difficulty) Pairs() int {
	return [...]int{6, 8, 12}[d]
}

type MemoryGame struct {
	Deck              Deck
	LastAttemptedCard int
	TurnsTaken        int
	Score             int
	BestScore         int
	Difficulty        Difficulty
	Leaderboard       *Leaderboard
	AskingForInitials bool
}

func (m *MemoryGame) takeTurn() {
	m.TurnsTaken++
}

func (m *MemoryGame) incrementScore() {
	m.Score++
}

func (m *MemoryGame) resetScores() {
	m.TurnsTaken = 0
	m.Score = 0
}

func (m MemoryGame) HasWon() bool {
	for _, card := range m.Deck {
		if !card.Matched {
			return false
		}
	}
	return true
}

type Deck []Card

type Card struct {
	Value   int
	Flipped bool
	Matched bool
}

func FlipCard(m gt.Message, s gt.State) gt.Response {
	state := model(s)
	cardToFlip := m.ArgsToInt()

	state.MemoryGame.Deck.flipCard(cardToFlip)

	if state.MemoryGame.Deck.numberFlippedCards() == 2 {
		state.MemoryGame.takeTurn()

		if state.MemoryGame.Deck.hasFoundMatch() {
			state.MemoryGame.incrementScore()
			return gt.RespondWithDelayedNextMsg(gt.Message{Message: "REMOVE_MATCHES"}, 1*time.Second)
		}

		return gt.RespondWithDelayedNextMsg(gt.Message{Message: "FLIP_ALL_BACK"}, 1*time.Second)
	}

	return gt.Respond()
}

func ChangeDifficulty(m gt.Message, s gt.State) gt.Response {
	state := model(s)
	newDifficulty := Difficulty(m.ArgsToInt())
	state.MemoryGame.Difficulty = newDifficulty
	state.MemoryGame.Deck = NewDeck(newDifficulty.Pairs())
	state.MemoryGame.resetScores()
	return gt.Respond()
}

func RestartGame(_ gt.Message, s gt.State) gt.Response {
	state := model(s)
	state.MemoryGame.Deck = NewDeck(state.MemoryGame.Difficulty.Pairs())
	state.MemoryGame.resetScores()
	state.MemoryGame.AskingForInitials = false
	return gt.Respond()
}

func FlipAllBack(_ gt.Message, s gt.State) gt.Response {
	state := model(s)
	state.MemoryGame.Deck.flipAllBack()
	return gt.Respond()
}

func SubmitInitials(m gt.Message, s gt.State) gt.Response {
	state := model(s)
	initials := m.ArgsToString()

	// Basic validation
	if len(initials) > 3 {
		initials = initials[:3]
	}
	initials = strings.ToUpper(initials)

	score := HighScore{
		Initials: initials,
		Score:    state.MemoryGame.TurnsTaken,
		Date:     time.Now(),
	}

	state.MemoryGame.Leaderboard.AddScore(state.MemoryGame.Difficulty, score)
	state.MemoryGame.Leaderboard.Save()
	state.MemoryGame.AskingForInitials = false

	return gt.Respond()
}

func RemoveMatches(_ gt.Message, s gt.State) gt.Response {
	state := model(s)
	state.MemoryGame.Deck.removeMatches()

	if state.MemoryGame.HasWon() {
		if state.MemoryGame.BestScore == 0 || state.MemoryGame.TurnsTaken < state.MemoryGame.BestScore {
			state.MemoryGame.BestScore = state.MemoryGame.TurnsTaken
		}

		if state.MemoryGame.Leaderboard.IsHighScore(state.MemoryGame.Difficulty, state.MemoryGame.TurnsTaken) {
			state.MemoryGame.AskingForInitials = true
		}
	}

	return gt.Respond()
}

func NewDeck(n int) (deck Deck) {
	cardValues := append([]int{}, append(rand.Perm(n), rand.Perm(n)...)...)
	for _, v := range cardValues {
		deck = append(deck, Card{v, false, false})
	}
	return
}

func (deck Deck) flipAllBack() {
	for i := range deck {
		deck[i].Flipped = false
	}
}

func (deck Deck) flipCard(i int) {
	deck[i].Flipped = !deck[i].Flipped
}

func (deck Deck) onlyFlipped() (flippedCards Deck) {
	// filter to only flipped cards
	for _, card := range deck {
		if card.Flipped {
			flippedCards = append(flippedCards, card)
		}
	}
	return
}

func (deck Deck) numberFlippedCards() int {
	return len(deck.onlyFlipped())
}

func (deck Deck) hasFoundMatch() bool {
	// if there are exactly 2 cards flipped
	// and their valus match
	flippedCards := deck.onlyFlipped()
	if len(flippedCards) == 2 {
		if flippedCards[0].Value == flippedCards[1].Value {
			return true
		}
	}
	return false
}

func (deck Deck) removeMatches() {
	for i, card := range deck {
		if card.Flipped {
			deck[i].Matched = true
			deck[i].Flipped = false
		}
	}
}

func (deck Deck) reset() {
	for i := range deck {
		deck[i].Matched = false
		deck[i].Flipped = false
	}
}

// Render
func (game MemoryGame) render() h.Element {
	return h.Div(a.Attrs(
		a.Class("space-y-8")),
		// Header
		h.Div(a.Attrs(
			a.Class("text-center space-y-2")),
			h.H1(a.Attrs(
				a.Class("text-4xl font-bold text-stone-900"),
				a.Style(css.FontFamily("'DM Serif Display', serif"))),
				h.Text("🎴 Memory Game")),
			h.P(a.Attrs(
				a.Class("text-stone-600")),
				h.Text("Find all matching pairs in as few turns as possible!"))),
		// Difficulty selector
		h.Div(a.Attrs(
			a.Class("flex justify-center gap-2")),
			renderDifficultyButton(Easy, game.Difficulty),
			renderDifficultyButton(Medium, game.Difficulty),
			renderDifficultyButton(Hard, game.Difficulty)),
		// Stats bar
		h.Div(a.Attrs(
			a.Class("grid grid-cols-3 gap-4")),
			statCard("🔄", "Turns", fmt.Sprintf("%d", game.TurnsTaken), "amber"),
			statCard("✅", "Pairs Found", fmt.Sprintf("%d", game.Score), "emerald"),
			statCard("🏆", "Best Score", func() string {
				if game.BestScore == 0 {
					return "—"
				}
				return fmt.Sprintf("%d", game.BestScore)
			}(), "sky")),
		// Game board
		game.Deck.render(game.Difficulty),
		// Status
		game.renderStatus(),
		// Explanatory note
		renderExplanatoryNote(
			"Memory Game Architecture",
			`
			<p class="mb-3">The Memory Game demonstrates a more complex component with its own state and logic.</p>
			<ul class="list-disc pl-5 space-y-2">
				<li><strong class="text-stone-900">Component Structure:</strong> The <code class="bg-stone-200 px-1.5 py-0.5 rounded text-xs font-mono">MemoryGame</code> struct encapsulates the game state (Deck, Score, Turns).</li>
				<li><strong class="text-stone-900">Message Handling:</strong> Game-specific messages (like <code class="bg-stone-200 px-1.5 py-0.5 rounded text-xs font-mono">FLIP_CARD</code>) are handled by the main update function.</li>
				<li><strong class="text-stone-900">Delayed Messages:</strong> When cards don't match, a delayed message flips them back after 1 second using <code class="bg-stone-200 px-1.5 py-0.5 rounded text-xs font-mono">RespondWithDelayedNextMsg</code>.</li>
			</ul>
			`),
		// Leaderboard
		game.renderLeaderboard())
}

func statCard(icon, label, value, color string) h.Element {
	bgColors := map[string]string{
		"amber":   "bg-amber-50 border-amber-300",
		"emerald": "bg-emerald-50 border-emerald-300",
		"sky":     "bg-sky-50 border-sky-300",
	}
	return h.Div(a.Attrs(
		a.Class(fmt.Sprintf("text-center p-4 rounded-xl border-2 %s", bgColors[color]))),
		h.Div(a.Attrs(
			a.Class("text-2xl mb-1")),
			h.Text(icon)),
		h.Div(a.Attrs(
			a.Class("text-xs font-medium text-stone-500 uppercase tracking-wide")),
			h.Text(label)),
		h.Div(a.Attrs(
			a.Class("text-3xl font-bold text-stone-900 mt-1"),
			a.Style(css.FontFamily("'JetBrains Mono', monospace"))),
			h.Text(value)))
}

func (game MemoryGame) renderLeaderboard() h.Element {
	scores := game.Leaderboard.GetScores(game.Difficulty)
	if len(scores) == 0 {
		return h.Nothing(a.Attrs())
	}

	return h.Div(a.Attrs(
		a.Class("bg-gradient-to-br from-amber-50 to-orange-50 border-2 border-stone-900 rounded-2xl p-6 shadow-brutal-sm")),
		h.H3(a.Attrs(
			a.Class("text-xl font-bold text-stone-900 mb-4 flex items-center gap-2")),
			h.Span(a.Attrs(),
				h.Text("🏆")),
			h.Text(fmt.Sprintf("%s Leaderboard", game.Difficulty))),
		h.Div(a.Attrs(
			a.Class("overflow-hidden rounded-xl border-2 border-stone-900")),
			// Header
			h.Div(a.Attrs(
				a.Class("bg-stone-900 text-white grid grid-cols-4 gap-4 px-4 py-3 text-xs font-semibold uppercase tracking-wider")),
				h.Div(a.Attrs(),
					h.Text("#")),
				h.Div(a.Attrs(),
					h.Text("Player")),
				h.Div(a.Attrs(),
					h.Text("Turns")),
				h.Div(a.Attrs(),
					h.Text("Date"))),
			// Rows
			h.Div(a.Attrs(
				a.Class("bg-white divide-y-2 divide-stone-200")),
				func() []h.Element {
					var rows []h.Element
					for i, score := range scores {
						medalEmoji := ""
						if i == 0 {
							medalEmoji = "🥇 "
						} else if i == 1 {
							medalEmoji = "🥈 "
						} else if i == 2 {
							medalEmoji = "🥉 "
						}
						rows = append(rows, h.Div(a.Attrs(
							a.Class("grid grid-cols-4 gap-4 px-4 py-3 text-sm")),
							h.Div(a.Attrs(
								a.Class("font-mono text-stone-500")),
								h.Text(fmt.Sprintf("%s%d", medalEmoji, i+1))),
							h.Div(a.Attrs(
								a.Class("font-bold text-stone-900")),
								h.Text(score.Initials)),
							h.Div(a.Attrs(
								a.Class("font-mono")),
								h.Text(fmt.Sprintf("%d", score.Score))),
							h.Div(a.Attrs(
								a.Class("text-stone-500")),
								h.Text(score.Date.Format("Jan 02")))))
					}
					return rows
				}()...)))
}

func renderDifficultyButton(d, current Difficulty) h.Element {
	if d == current {
		return h.Button(a.Attrs(
			a.Class("px-5 py-2 rounded-full text-sm font-semibold bg-stone-900 text-white border-2 border-stone-900"),
			a.OnClick(gt.SendBasicMessage("CHANGE_DIFFICULTY", int(d)))),
			h.Text(d.String()))
	}
	return h.Button(a.Attrs(
		a.Class("px-5 py-2 rounded-full text-sm font-medium text-stone-600 bg-white border-2 border-stone-300 hover:border-stone-900 hover:text-stone-900 transition-colors"),
		a.OnClick(gt.SendBasicMessage("CHANGE_DIFFICULTY", int(d)))),
		h.Text(d.String()))
}

func (game MemoryGame) renderStatus() h.Element {
	if game.HasWon() {
		return game.renderGameWon()
	}
	return renderGameOngoing()
}

func (game MemoryGame) renderGameWon() h.Element {
	if game.AskingForInitials {
		return h.Div(a.Attrs(
			a.Class("text-center space-y-4 bg-gradient-to-br from-amber-100 to-yellow-100 p-8 rounded-2xl border-2 border-stone-900 shadow-brutal")),
			h.Div(a.Attrs(
				a.Class("text-5xl mb-2")),
				h.Text("🎉")),
			h.H2(a.Attrs(
				a.Class("text-3xl font-bold text-stone-900"),
				a.Style(css.FontFamily("'DM Serif Display', serif"))),
				h.Text("New High Score!")),
			h.P(a.Attrs(
				a.Class("text-stone-600")),
				h.Text("Enter your initials to join the leaderboard:")),
			h.Div(a.Attrs(
				a.Class("flex justify-center gap-3 pt-2")),
				h.Input(a.Attrs(
					a.Id("initials-input"),
					a.Type("text"),
					a.MaxLength(3),
					a.Class("w-24 text-center uppercase text-2xl font-bold border-2 border-stone-900 rounded-xl bg-white shadow-brutal-sm"),
					a.Style(css.FontFamily("'JetBrains Mono', monospace")),
					a.Placeholder("AAA"))),
				h.Button(a.Attrs(
					a.Class("px-6 py-3 bg-emerald-500 hover:bg-emerald-600 text-white font-semibold rounded-xl border-2 border-stone-900 shadow-brutal-sm hover:shadow-brutal hover:-translate-x-0.5 hover:-translate-y-0.5 transition-all"),
					a.OnClick(gt.SendBasicMessageWithValueFromInput("SUBMIT_INITIALS", "initials-input"))),
					h.Text("Submit →"))))
	}

	return h.Div(a.Attrs(
		a.Class("text-center space-y-4 bg-gradient-to-br from-emerald-100 to-green-100 p-8 rounded-2xl border-2 border-stone-900 shadow-brutal")),
		h.Div(a.Attrs(
			a.Class("text-5xl mb-2")),
			h.Text("🎊")),
		h.H2(a.Attrs(
			a.Class("text-3xl font-bold text-stone-900"),
			a.Style(css.FontFamily("'DM Serif Display', serif"))),
			h.Text("You Won!")),
		h.P(a.Attrs(
			a.Class("text-stone-600")),
			h.Text(fmt.Sprintf("Completed in %d turns", game.TurnsTaken))),
		h.Button(a.Attrs(
			a.Class("px-6 py-3 bg-stone-900 hover:bg-stone-800 text-white font-semibold rounded-xl border-2 border-stone-900 shadow-brutal-sm hover:shadow-brutal hover:-translate-x-0.5 hover:-translate-y-0.5 transition-all"),
			a.OnClick(gt.SendBasicMessageNoArgs("RESTART_GAME"))),
			h.Text("Play Again →")))
}

func renderGameOngoing() h.Element {
	return h.Div(a.Attrs(
		a.Class("text-center py-4")),
		h.P(a.Attrs(
			a.Class("text-stone-500 font-medium")),
			h.Text("Find all the matching pairs! 🎯")))
}

func (deck Deck) render(d Difficulty) h.Element {
	gridCols := "grid-cols-4"
	if d == Hard {
		gridCols = "grid-cols-6"
	}

	return h.Div(a.Attrs(
		a.Id("deck"),
		a.Class(fmt.Sprintf("grid %s gap-3", gridCols))),
		func() []h.Element {
			var elements []h.Element
			for index, card := range deck {
				elements = append(elements, card.renderContainer(index))
			}
			return elements
		}()...)
}

func (card Card) renderContainer(index int) h.Element {
	return h.Div(a.Attrs(
		a.Class("relative aspect-square")),
		card.render(),
		card.renderFlipButton(index))
}

func (card Card) renderFlipButton(index int) h.Element {
	if !card.Matched {
		return h.Button(a.Attrs(
			a.Class("absolute inset-0 w-full h-full opacity-0 cursor-pointer"),
			a.OnClick(gt.SendBasicMessage("FLIP_CARD", index))),
			h.Text("Flip"))
	}
	return h.Nothing(a.Attrs())
}

// Card emoji mapping for a more playful look
var cardEmojis = []string{"🍎", "🍋", "🍇", "🍊", "🫐", "🍓", "🥝", "🍑", "🍒", "🥭", "🍍", "🥥"}

func (card Card) render() h.Element {
	return h.Div(a.Attrs(
		a.Class(card.getClass())),
		h.Span(a.Attrs(
			a.Class("text-4xl")),
			card.getValue()))
}

func (card Card) getClass() string {
	base := "h-full w-full flex items-center justify-center rounded-xl border-2 border-stone-900 transition-all duration-200"

	if card.Matched {
		return base + " bg-stone-100 opacity-40 cursor-default"
	}

	if card.Flipped {
		return base + " bg-white shadow-brutal"
	}

	return base + " bg-gradient-to-br from-violet-500 to-purple-600 cursor-pointer hover:shadow-brutal hover:-translate-x-0.5 hover:-translate-y-0.5"
}

func (card Card) getValue() h.Element {
	if card.Flipped {
		emoji := cardEmojis[card.Value%len(cardEmojis)]
		return h.Text(emoji)
	}
	return h.Span(a.Attrs(
		a.Class("text-2xl text-white/30")),
		h.Text("?"))
}
