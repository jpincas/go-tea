package main

import (
	"strings"
	"testing"

	"github.com/jpincas/go-tea/tester"
)

func TestCounter(t *testing.T) {
	// Initialize the test session with the main Model
	// Note: In a real app you might want to test components in isolation,
	// but here we test via the main model as it embeds the Counter.
	session := tester.NewSession(t, &Model{})

	// Initial state check
	state := session.GetState().(*Model)
	if state.Counter != 0 {
		t.Errorf("Expected initial counter value to be 0, got %d", state.Counter)
	}

	// Test INCREMENT
	// Note: We pass float64 because the runtime expects JSON numbers (which are floats)
	session.Dispatch("INCREMENT_COUNTER", float64(1))
	state = session.GetState().(*Model)
	if state.Counter != 1 {
		t.Errorf("Expected counter value to be 1 after increment, got %d", state.Counter)
	}

	// Check Rendered HTML
	html := session.Render()
	if !strings.Contains(html, "1") {
		t.Errorf("Expected rendered HTML to contain '1', got: %s", html)
	}

	// Test DECREMENT (using INCREMENT_COUNTER with -1)
	session.Dispatch("INCREMENT_COUNTER", float64(-1))
	state = session.GetState().(*Model)
	if state.Counter != 0 {
		t.Errorf("Expected counter value to be 0 after decrement, got %d", state.Counter)
	}
}
