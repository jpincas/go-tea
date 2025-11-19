package blocktrader

import (
	"fmt"
	"math/rand"
	"time"

	gt "github.com/jpincas/go-tea"
)

// Configuration
type Config struct {
	BoardSize         int
	AxisStartValue    float64
	TickDuration      time.Duration
	StartingCash      float64
	GameDurationTicks int
	Trend             []float64
}

var DefaultConfig = Config{
	BoardSize:         6,
	AxisStartValue:    10,
	TickDuration:      275 * time.Millisecond,
	StartingCash:      1000,
	GameDurationTicks: 350,
	Trend:             []float64{2, 1, 0.75, 0.5, 0.3, 0.2},
}

// State
type State struct {
	Timer                   int
	Cash                    float64
	Rows                    []float64
	Cols                    []float64
	OwnedBlocks             map[string]float64 // Key: "row:col", Value: boughtAt
	PreviousMutationAmounts []float64
	TrendTick               int
	RowColumn               int // 0 or 1
	Member                  int // Index of row/col being mutated
	GameActive              bool
	Trend                   []float64
}

type Model struct {
	State  State
	Config Config
}

func NewModel() Model {
	return Model{
		Config: DefaultConfig,
		State:  NewState(DefaultConfig),
	}
}

// Messages
var Messages = gt.MessageMap{
	"BLOCKTRADER_START_GAME":   StartGame,
	"BLOCKTRADER_STOP_GAME":    StopGame,
	"BLOCKTRADER_RESET_GAME":   ResetGame,
	"BLOCKTRADER_TICK":         Tick,
	"BLOCKTRADER_TOGGLE_BLOCK": ToggleBlock,
}

func StartGame(msg gt.Message, s gt.State) gt.Response {
	m := s.(interface{ GetBlocktrader() *Model }).GetBlocktrader()
	m.State = NewState(m.Config)
	m.State.GameActive = true
	return gt.RespondWithNextMsg(gt.Message{Message: "BLOCKTRADER_TICK"})
}

func StopGame(msg gt.Message, s gt.State) gt.Response {
	m := s.(interface{ GetBlocktrader() *Model }).GetBlocktrader()
	m.State.GameActive = false
	return gt.Respond()
}

func ResetGame(msg gt.Message, s gt.State) gt.Response {
	m := s.(interface{ GetBlocktrader() *Model }).GetBlocktrader()
	m.State = NewState(m.Config)
	return gt.Respond()
}

func Tick(msg gt.Message, s gt.State) gt.Response {
	m := s.(interface{ GetBlocktrader() *Model }).GetBlocktrader()

	if !m.State.GameActive {
		return gt.Respond()
	}

	m.tick()

	if m.State.Timer <= 0 || m.State.Cash <= 0 {
		m.endGame()
		return gt.Respond()
	}

	return gt.RespondWithDelayedNextMsg(gt.Message{Message: "BLOCKTRADER_TICK"}, m.Config.TickDuration)
}

func ToggleBlock(msg gt.Message, s gt.State) gt.Response {
	m := s.(interface{ GetBlocktrader() *Model }).GetBlocktrader()
	
	// Parse args: [row, col]
	args := msg.Arguments.([]interface{})
	row := int(args[0].(float64))
	col := int(args[1].(float64))

	m.toggleOwned(row, col)
	return gt.Respond()
}

// Logic

func NewState(c Config) State {
	rows := make([]float64, c.BoardSize)
	cols := make([]float64, c.BoardSize)
	for i := range rows {
		rows[i] = c.AxisStartValue
		cols[i] = c.AxisStartValue
	}

	return State{
		Timer:                   c.GameDurationTicks,
		Cash:                    c.StartingCash,
		Rows:                    rows,
		Cols:                    cols,
		OwnedBlocks:             make(map[string]float64),
		PreviousMutationAmounts: []float64{0},
		TrendTick:               0,
		RowColumn:               0,
		Member:                  0,
		Trend:                   c.Trend,
		GameActive:              false,
	}
}

func (m *Model) tick() {
	// Calculate mutation
	mutationSeed := rand.Float64()
	if rand.Intn(2) == 0 {
		mutationSeed = -mutationSeed
	}
	
	mutationAmount := m.calcMutationAmount(mutationSeed)

	// Update trend tick and selection
	if m.State.TrendTick >= len(m.State.Trend) {
		m.State.RowColumn = rand.Intn(2)
		m.State.Member = rand.Intn(m.Config.BoardSize)
		m.State.TrendTick = 0
	} else {
		m.State.TrendTick++
	}

	// Apply mutation
	if m.State.RowColumn > 0 {
		// Cols
		proposed := m.State.Cols[m.State.Member] + mutationAmount
		if proposed > 0 {
			m.State.Cols[m.State.Member] = proposed
		} else {
			m.State.Cols[m.State.Member] = 0
		}
	} else {
		// Rows
		proposed := m.State.Rows[m.State.Member] + mutationAmount
		if proposed > 0 {
			m.State.Rows[m.State.Member] = proposed
		} else {
			m.State.Rows[m.State.Member] = 0
		}
	}

	m.State.Timer--
}

func (m *Model) calcMutationAmount(thisMutation float64) float64 {
	// Add to history
	m.State.PreviousMutationAmounts = append([]float64{thisMutation}, m.State.PreviousMutationAmounts...)
	
	// Calculate weighted sum
	var sum float64
	for i, trendVal := range m.State.Trend {
		if i < len(m.State.PreviousMutationAmounts) {
			sum += trendVal * m.State.PreviousMutationAmounts[i]
		}
	}
	return sum
}

func (m *Model) endGame() {
	m.State.GameActive = false
	// Sell all blocks
	portfolioValue := m.portfolioValue()
	if portfolioValue > 0 {
		m.State.Cash += portfolioValue
	}
}

func (m *Model) portfolioValue() float64 {
	var total float64
	for key := range m.State.OwnedBlocks {
		row, col := parseBlockKey(key)
		val := m.blockValue(row, col)
		total += val
	}
	return total
}

func (m *Model) toggleOwned(row, col int) {
	key := blockKey(row, col)
	if _, owned := m.State.OwnedBlocks[key]; owned {
		m.sell(row, col)
	} else {
		m.buy(row, col)
	}
}

func (m *Model) buy(row, col int) {
	val := m.blockValue(row, col)
	if val > 0 && m.State.Cash >= val {
		m.State.OwnedBlocks[blockKey(row, col)] = val
		m.State.Cash -= val
	}
}

func (m *Model) sell(row, col int) {
	val := m.blockValue(row, col)
	key := blockKey(row, col)
	if _, owned := m.State.OwnedBlocks[key]; owned {
		delete(m.State.OwnedBlocks, key)
		m.State.Cash += val
	}
}

func (m *Model) blockValue(row, col int) float64 {
	return m.State.Rows[row] * m.State.Cols[col]
}

func blockKey(row, col int) string {
	return fmt.Sprintf("%d:%d", row, col)
}

func parseBlockKey(key string) (int, int) {
	var row, col int
	fmt.Sscanf(key, "%d:%d", &row, &col)
	return row, col
}
