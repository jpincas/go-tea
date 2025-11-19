package tester

import (
	"testing"

	"github.com/google/uuid"
	gt "github.com/jpincas/go-tea"
)

// TestSession holds the state of a GoTea application for testing
type TestSession struct {
	State gt.State
	t     *testing.T
}

// NewSession creates a new test session with the given model
func NewSession(t *testing.T, model gt.State) *TestSession {
	// Initialize with a dummy session ID
	state := model.Init(uuid.New())
	return &TestSession{
		State: state,
		t:     t,
	}
}

// Dispatch sends a message to the application state
func (s *TestSession) Dispatch(msgName string, args any) *TestSession {
	// Construct the message
	msg := gt.Message{
		Message: msgName,
		Arguments: args,
	}

	// Find the handler
	// We check the state's Update map
	handler, found := s.State.Update()[msgName]
	if !found {
		s.t.Fatalf("Message handler not found for message: %s", msgName)
	}

	// Execute the handler
	response := handler(msg, s.State)
	if response.Error != nil {
		s.t.Fatalf("Error processing message %s: %v", msgName, response.Error)
	}

	// If there's a next message, process it too (recursively? or just queue it?)
	// For simple unit testing, we might want to stop here or allow manual processing of next msg.
	// For now, let's just process the immediate update.
	// If we wanted to simulate the runtime fully, we'd handle NextMsg here.
	// But often in tests we want to assert state after the first message.
	
	return s
}

// Render returns the rendered HTML as a string
func (s *TestSession) Render() string {
	return string(s.State.Render())
}

// GetState returns the current state
func (s *TestSession) GetState() gt.State {
	return s.State
}
