# Gotea - Development Guide

## Project Overview

Gotea is a Go library implementing **The Elm Architecture (TEA)** for server-side rendered web applications. It enables building dynamic SPA-like applications using Go for logic and state management, with minimal JavaScript. State lives on the server, communication happens via WebSockets, and DOM updates use morphdom for efficient patching.

**Key Insight:** This is NOT a traditional web framework. The browser is essentially a "dumb terminal" - all application logic and state management happens server-side. The client only sends messages and receives HTML to patch into the DOM.

## Quick Start

```bash
# Run the example application
cd example && go run .

# Visit http://localhost:8080
```

## Architecture

### Complete Message Flow

```
┌─────────────────────────────────────────────────────┐
│           BROWSER (Client)                           │
├─────────────────────────────────────────────────────┤
│  HTML Document                                       │
│  ├─ Event Listeners (onclick, onkeyup, etc.)        │
│  └─ gotea.js                                         │
│     ├─ WebSocket connection to /server              │
│     ├─ sendMessage() → JSON to server               │
│     ├─ morphdom → patches DOM with new HTML         │
│     └─ localStorage → state persistence             │
└────────────┬──────────────────────────────────────┬─┘
             │ WebSocket                            │
             │ {message, args}                      │ HTML Response
             ▼                                       │
┌────────────────────────────────────────────────────┴─┐
│          SERVER (runtime.go)                         │
├──────────────────────────────────────────────────────┤
│  Per-Session State                                   │
│  ├─ handleMessage() deserializes JSON               │
│  ├─ message.process() looks up handler              │
│  ├─ Handler executes, mutates state                 │
│  ├─ state.Render() → full HTML                      │
│  └─ If Persistable: sends STATE_SNAPSHOT            │
└──────────────────────────────────────────────────────┘
```

### Core Flow (Detailed)

1. **Initial HTTP Request:** Browser requests `/`, server creates session cookie with UUID, calls `model.Init(sessionID)`, renders initial HTML
2. **WebSocket Connection:** Browser connects to `/server?whence=/current/path`, server associates connection with session state
3. **User Interaction:** Click/keyup/change triggers inline JS like `gotea.sendMessage({message: "ACTION", args: 42})`
4. **Message Processing:** Server deserializes JSON, looks up handler in `Update()` MessageMap, executes handler
5. **State Mutation:** Handler mutates state, returns `gt.Response`
6. **Re-render:** Unless `BlockRerender: true`, server calls `state.Render()` and sends HTML
7. **DOM Patch:** Client receives HTML, morphdom patches only changed nodes (preserves focus, input state)
8. **Optional Persistence:** If state implements `Persistable`, snapshot sent to client's localStorage

### Key Files

| File | Purpose | Lines |
|------|---------|-------|
| `runtime.go` | Core engine: State interface, Message loop, WebSocket, Router, Session management | ~450 |
| `component.go` | ComponentID abstraction for message namespacing | ~22 |
| `templatehelpers.go` | Helper functions generating JS message calls | ~93 |
| `js/gotea.js` | Client-side WebSocket, morphdom, routing, persistence | ~200 |
| `msg/helpers.go` | Message argument decoding utilities | ~20 |
| `tester/tester.go` | Test utilities for component testing | ~65 |

## Core Interfaces

### State Interface

Every application must implement this interface:

```go
type State interface {
    Routable                           // Embedded: provides routing methods
    Init(uuid.UUID) State              // Initialize state for new session
    Update() MessageMap                // Return all message handlers
    Render() []byte                    // Render current state to HTML
    RenderError(error) []byte          // Render error state
}
```

### Routable Interface (embedded in State)

```go
type Routable interface {
    OnRouteChange(path string)         // Called when route changes
    RenderRoute(s State) []byte        // Render the current route
    SetNewRoute(route string)          // Update current route
    GetRoute() string                  // Get current route
}
```

### Persistable Interface (optional)

Implement to enable state restoration across server restarts:

```go
type Persistable interface {
    Serialize() ([]byte, error)        // Convert state to JSON for localStorage
    Deserialize([]byte) error          // Restore state from JSON
}
```

## Message System

### Message Structure

```go
type Message struct {
    Message       string `json:"message"`       // Handler name (e.g., "INCREMENT")
    Arguments     any    `json:"args"`          // Arbitrary arguments
    Identifier    string `json:"identifier"`    // Optional identifier
    BlockRerender bool   `json:"blockRerender"` // If true, don't re-render after handling
}
```

### Accessing Arguments

```go
func MyHandler(m gt.Message, s gt.State) gt.Response {
    // For simple types
    intVal := m.ArgsToInt()              // JSON numbers → int
    strVal := m.ArgsToString()           // JSON string → string

    // For complex types (forms, objects)
    var myStruct MyType
    m.MustDecodeArgs(&myStruct)          // Panics on error

    // Or with error handling
    err := m.DecodeArgs(&myStruct)
}
```

### Response Types

```go
// Basic response - re-renders state
return gt.Respond()

// Error response - calls RenderError()
return gt.RespondWithError(err)

// Chain another message immediately
return gt.RespondWithNextMsg(gt.Message{Message: "NEXT_ACTION"})

// Chain with delay (for animations, game loops)
return gt.RespondWithDelayedNextMsg(
    gt.Message{Message: "NEXT_FRAME"},
    33*time.Millisecond,  // ~30fps
)
```

### Message Handler Signature

```go
type MessageHandler func(m Message, s State) Response

// Convention: type assertion helper at top of file
func model(s gt.State) *Model {
    return s.(*Model)
}

func MyHandler(m gt.Message, s gt.State) gt.Response {
    state := model(s)
    // ... mutate state ...
    return gt.Respond()
}
```

## Coding Conventions

### HTML Rendering with htmlfunc

```go
import (
    a "github.com/jpincas/htmlfunc/attributes"
    h "github.com/jpincas/htmlfunc/html"
)

func render() h.Element {
    return h.Div(
        a.Attrs(a.Class("container"), a.Id("main")),
        h.H1(a.Attrs(), h.Text("Title")),
        h.P(a.Attrs(), h.Text("Content")),
    )
}

// Convert to bytes for Render()
func (m *Model) Render() []byte {
    return renderPage(m).Bytes()
}
```

### Styling with Tailwind CSS

```go
h.Button(
    a.Attrs(
        a.Class("bg-indigo-600 hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded"),
        a.OnClick(gt.SendBasicMessage("MY_MESSAGE", nil)),
    ),
    h.Text("Click me"),
)
```

### Triggering Messages from HTML

```go
// Basic message with args
a.OnClick(gt.SendBasicMessage("ACTION", 42))

// Message with no args
a.OnClick(gt.SendBasicMessageNoArgs("ACTION"))

// Get value from specific input field
a.OnKeyUp(gt.SendBasicMessageWithValueFromInput("SEARCH", "search-input"))

// Get value from the triggering input itself
a.OnKeyUp(gt.SendBasicMessageWithValueFromThisInput("TYPING"))

// Serialize entire form
a.OnChange(gt.BasicUpdateForm("FORM_UPDATE", "form-id"))

// Conditional message
a.OnClick(gt.IfElse("condition", msg1, msg2))
```

## Application Setup

### Creating and Starting the App

```go
func main() {
    // Create application with your model
    app := gt.NewApp(&Model{})

    // Start server on port with static file directory
    app.Start(8080, "static")
}
```

### Model Definition

```go
type Model struct {
    gt.Router              // Embed for routing support
    sessionID uuid.UUID    // Track session

    // Feature-specific state (flat structure)
    Counter    int
    Form       FormState
    MemoryGame MemoryGameState
    // etc.
}

func (m *Model) Init(sid uuid.UUID) gt.State {
    model := &Model{sessionID: sid}

    // Register routes
    model.Register("/", homeHandler)
    model.Register("/game", gameHandler)

    // Initialize feature state
    model.MemoryGame = NewMemoryGame()

    return model
}
```

### Merging Message Maps

```go
func (m *Model) Update() gt.MessageMap {
    return gt.MergeMaps(
        counterMessages,
        formMessages,
        memoryGameMessages,
        // Namespaced component messages
        myComponent.UniqueMsgMap(componentMessages),
    )
}
```

## Routing

### Route Registration

```go
func (m *Model) Init(sid uuid.UUID) gt.State {
    model := &Model{sessionID: sid}

    model.Register("/", func(s gt.State) []byte {
        return renderHome().Bytes()
    })

    model.Register("/game", func(s gt.State) []byte {
        state := s.(*Model)
        return state.MemoryGame.render().Bytes()
    })

    return model
}
```

### Route Change Hook

Called whenever route changes (link clicks, browser back/forward):

```go
func (m *Model) OnRouteChange(path string) {
    // Parse path and query params
    p, _ := url.Parse(path)

    // Access query parameters
    id := m.RouteParam("id")  // from ?id=123

    // Route-specific initialization
    if p.Path == "/game" {
        m.MemoryGame.Reset()
    }
}
```

### Client-Side Routing

Links are automatically intercepted by gotea.js:

```go
// Internal link - handled by Gotea (no page reload)
h.A(a.Attrs(a.Href("/game")), h.Text("Play Game"))

// External link - normal browser navigation
h.A(a.Attrs(a.Href("https://example.com"), a.Class("external")), h.Text("External"))
```

## Components

### ComponentID for Namespacing

Prevents message collisions when using multiple instances of same component:

```go
type TagSelector struct {
    gt.ComponentID
    AvailableTags []string
    SelectedTags  []string
    SearchInput   string
}

// In parent Init:
nameSelector := TagSelector{ComponentID: "name-selector"}
teamSelector := TagSelector{ComponentID: "team-selector"}

// Component messages (local names)
var tagSelectorMessages = gt.MessageMap{
    "SELECT_TAG": selectTagHandler,
    "REMOVE_TAG": removeTagHandler,
}

// Register with namespace prefix
func (m *Model) Update() gt.MessageMap {
    return gt.MergeMaps(
        m.NameSelector.UniqueMsgMap(tagSelectorMessages),  // "name-selector_SELECT_TAG"
        m.TeamSelector.UniqueMsgMap(tagSelectorMessages),  // "team-selector_SELECT_TAG"
    )
}

// In component render, use unique message names
func (ts *TagSelector) Render() h.Element {
    return h.Button(
        a.Attrs(a.OnClick(gt.SendBasicMessage(ts.UniqueMsg("SELECT_TAG"), tag))),
        h.Text("Select"),
    )
}

// Also use UniqueID for HTML element IDs
h.Input(a.Attrs(a.Id(ts.UniqueID("search-input"))))
```

### Component Handler Pattern

```go
func selectTagHandler(m gt.Message, s gt.State) gt.Response {
    state := model(s)
    tag := m.ArgsToString()

    // Determine which component based on message identifier or prefix
    // Then call component method
    state.NameSelector.SelectTag(tag)

    return gt.Respond()
}
```

## Advanced Patterns

### Animation Loop (Server-Driven)

```go
var animationMessages = gt.MessageMap{
    "START_ANIMATION":     startAnimation,
    "NEXT_FRAME":          nextFrame,
    "STOP_ANIMATION":      stopAnimation,
}

func startAnimation(m gt.Message, s gt.State) gt.Response {
    state := model(s)
    state.Animation.Running = true

    // Start the loop
    return gt.RespondWithDelayedNextMsg(
        gt.Message{Message: "NEXT_FRAME"},
        33*time.Millisecond,  // ~30fps
    )
}

func nextFrame(m gt.Message, s gt.State) gt.Response {
    state := model(s)

    if !state.Animation.Running {
        return gt.Respond()  // Stop the loop
    }

    // Update animation state
    state.Animation.X += state.Animation.VelocityX
    state.Animation.Y += state.Animation.VelocityY

    // Bounce off walls
    if state.Animation.X <= 0 || state.Animation.X >= 100 {
        state.Animation.VelocityX *= -1
    }

    // Continue the loop
    return gt.RespondWithDelayedNextMsg(
        gt.Message{Message: "NEXT_FRAME"},
        33*time.Millisecond,
    )
}
```

### Broadcasting to All Clients

For chat, multiplayer games, or real-time updates:

```go
// Store app reference globally
var app *gt.Application

func main() {
    app = gt.NewApp(&Model{})
    app.Start(8080, "static")
}

// Global shared state (careful with concurrency!)
var chatMessages []ChatMessage
var mutex sync.Mutex

func sendChatMessage(m gt.Message, s gt.State) gt.Response {
    state := model(s)

    mutex.Lock()
    chatMessages = append(chatMessages, ChatMessage{
        User:    state.Username,
        Content: m.ArgsToString(),
    })
    mutex.Unlock()

    // Re-render ALL connected clients
    app.Broadcast()

    return gt.Respond()
}
```

### Form Handling

```go
type FormState struct {
    Name       string   `json:"name"`
    Email      string   `json:"email"`
    Subscribe  bool     `json:"subscribe"`
    Country    string   `json:"country"`
    Interests  []string `json:"interests"`  // Multi-select
}

var formMessages = gt.MessageMap{
    "FORM_UPDATE": formUpdate,
    "FORM_SUBMIT": formSubmit,
}

func formUpdate(m gt.Message, s gt.State) gt.Response {
    state := model(s)
    m.MustDecodeArgs(&state.Form)  // Deserialize entire form
    return gt.Respond()
}

// In template:
h.Form(
    a.Attrs(a.Id("my-form")),
    h.Input(a.Attrs(
        a.Type("text"),
        a.Name("name"),
        a.Value(form.Name),
        a.OnKeyUp(gt.BasicUpdateForm("FORM_UPDATE", "my-form")),
    )),
    h.Select(
        a.Attrs(
            a.Name("country"),
            a.OnChange(gt.BasicUpdateForm("FORM_UPDATE", "my-form")),
        ),
        // options...
    ),
)
```

### Silent Updates (BlockRerender)

For updates that shouldn't trigger UI refresh:

```go
// From server-initiated message
return gt.RespondWithNextMsg(gt.Message{
    Message:       "BACKGROUND_SAVE",
    BlockRerender: true,
})

// From client
msg := gt.Message{
    Message:       "LOG_ANALYTICS",
    Args:          eventData,
    BlockRerender: true,
}
```

### State Persistence

```go
func (m *Model) Serialize() ([]byte, error) {
    // Only persist important state, not UI ephemera
    snapshot := struct {
        Counter    int
        HighScores []Score
        // NOT: SearchInput, Animation state, etc.
    }{
        Counter:    m.Counter,
        HighScores: m.HighScores,
    }
    return json.Marshal(snapshot)
}

func (m *Model) Deserialize(data []byte) error {
    var snapshot struct {
        Counter    int
        HighScores []Score
    }
    if err := json.Unmarshal(data, &snapshot); err != nil {
        return err
    }
    m.Counter = snapshot.Counter
    m.HighScores = snapshot.HighScores
    return nil
}
```

## Testing

### Using the Tester Package

```go
func TestCounter(t *testing.T) {
    session := tester.NewSession(t, &Model{})

    // IMPORTANT: Use float64 for numbers (JSON convention)
    session.Dispatch("INCREMENT_COUNTER", float64(1))

    // Assert state
    state := session.GetState().(*Model)
    if state.Counter != 1 {
        t.Errorf("Expected 1, got %d", state.Counter)
    }

    // Check rendered output
    html := session.Render()
    if !strings.Contains(html, "1") {
        t.Error("Missing expected content")
    }
}

func TestFormUpdate(t *testing.T) {
    session := tester.NewSession(t, &Model{})

    // Complex args as struct
    session.Dispatch("FORM_UPDATE", map[string]interface{}{
        "name":  "John",
        "email": "john@example.com",
    })

    state := session.GetState().(*Model)
    if state.Form.Name != "John" {
        t.Errorf("Expected John, got %s", state.Form.Name)
    }
}
```

## Project Structure

```
/                       # Core library
  runtime.go           # Main runtime engine (State, Message, Router, WebSocket)
  component.go         # ComponentID abstraction
  templatehelpers.go   # JS message construction helpers
/js                    # Client-side JavaScript
  gotea.js            # WebSocket, morphdom, routing, persistence
/msg                   # Message decoding helpers
/tester               # Test utilities
/example              # Demo application
  main.go            # Model definition, routes, app startup
  counter.go         # Simple counter feature
  memory-game.go     # Complex game with delayed messages
  form.go            # Form handling example
  animation.go       # Server-driven animation
  chat.go            # Broadcasting example
  /static            # Static assets (JS, CSS)
  /blocktrader       # Package-level component example
  /tagselector       # Reusable component example
```

## Dependencies

- `github.com/olahol/melody` - WebSocket handling
- `github.com/google/uuid` - Session IDs
- `github.com/jpincas/htmlfunc` - HTML rendering (html, attributes, css packages)
- `morphdom` (JS, bundled) - Efficient DOM patching

## Running Tests

```bash
go test ./...
cd example && go test ./...
```

## Common Gotchas

1. **JSON numbers are float64:** In tests, pass `float64(1)` not `1` for numeric args
2. **State must be pointer receiver:** All State interface methods should use pointer receivers
3. **Embed gt.Router:** Don't forget to embed `gt.Router` in your Model struct
4. **Register routes in Init:** Routes must be registered in `Init()`, not at package level
5. **morphdom preserves focus:** Input focus and selection are preserved across re-renders
6. **Messages are uppercase:** Convention is `SCREAMING_SNAKE_CASE` for message names
7. **DelayedNextMsg runs in goroutine:** Chained messages are async, state is shared
