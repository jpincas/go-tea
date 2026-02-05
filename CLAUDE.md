# Gotea

Go library implementing **The Elm Architecture (TEA)** for server-side rendered web applications. State lives on the server, communication happens via WebSockets, and DOM updates use morphdom for efficient patching.

**Key Insight:** The browser is a "dumb terminal" - all application logic and state management happens server-side. The client only sends messages and receives HTML to patch into the DOM.

## Quick Start

```bash
cd example && go run .
# Visit http://localhost:8080
```

## Architecture

```
Browser                          Server (runtime.go)
────────                         ───────────────────
HTML + gotea.js                  Per-Session State
    │                                │
    │ WebSocket {message, args}      │
    └───────────────────────────────►│ handleMessage() → Update() lookup
                                     │ Handler mutates state
    ◄───────────────────────────────┤│ state.Render() → HTML
    │ HTML Response                  │
    │                                │
morphdom patches DOM                 │ If Persistable: STATE_SNAPSHOT
```

## Core Interfaces

```go
// State - every application must implement
type State interface {
    Routable                           // Embed gt.Router for this
    Init(uuid.UUID) State              // Initialize state for new session
    Update() MessageMap                // Return all message handlers
    Render() []byte                    // Render current state to HTML
    RenderError(error) []byte          // Render error state
}

// Persistable - optional, enables state restoration across server restarts
type Persistable interface {
    Serialize() ([]byte, error)        // Convert state to JSON for localStorage
    Deserialize([]byte) error          // Restore state from JSON
}
```

## Message Handling

```go
// Handler signature
type MessageHandler func(m Message, s State) Response

// Type assertion helper (define at top of each file)
func model(s gt.State) *Model {
    return s.(*Model)
}

// Handler example
func MyHandler(m gt.Message, s gt.State) gt.Response {
    state := model(s)
    // Mutate state...
    return gt.Respond()
}

// Response types
gt.Respond()                                           // Basic re-render
gt.RespondWithError(err)                               // Calls RenderError()
gt.RespondWithNextMsg(gt.Message{...})                 // Chain message immediately
gt.RespondWithDelayedNextMsg(msg, 33*time.Millisecond) // Chain with delay (for game loops)

// Accessing arguments
m.ArgsToInt()           // JSON number → int
m.ArgsToString()        // JSON string → string
m.MustDecodeArgs(&obj)  // Decode to struct (panics on error)
```

## HTML Rendering

```go
import (
    a "github.com/jpincas/go-tea/attributes"
    h "github.com/jpincas/go-tea/html"
    "github.com/jpincas/go-tea/css"
)

// Element construction
h.Div(a.Attrs(a.Class("container"), a.Id("main")),
    h.H1(a.Attrs(), h.Text("Title")),
)

// Conditional rendering
h.Div(...).RenderIf(condition)

// Raw HTML (be careful with XSS)
h.UnsafeRaw("<strong>bold</strong>")

// Inline styles via css package
a.Style(css.FontFamily("'JetBrains Mono', monospace"), css.Flex_("1"))
```

## Triggering Messages from HTML

```go
// Basic message with args
a.OnClick(gt.SendBasicMessage("ACTION", 42))

// No args
a.OnClick(gt.SendBasicMessageNoArgs("ACTION"))

// Get value from input field
a.OnKeyUp(gt.SendBasicMessageWithValueFromInput("SEARCH", "input-id"))

// Serialize entire form
a.OnChange(gt.BasicUpdateForm("FORM_UPDATE", "form-id"))

// Conditional
a.OnClick(gt.IfElse("condition", msg1, msg2))
```

## Application Setup

```go
func main() {
    app := gt.NewApp(&Model{})
    app.Start(8080, "static")  // port, static files directory
}

type Model struct {
    gt.Router              // MUST embed for routing
    sessionID uuid.UUID
    // Your state fields...
}

func (m *Model) Init(sid uuid.UUID) gt.State {
    model := &Model{sessionID: sid}

    // Register routes
    model.Register("/", homeHandler)
    model.Register("/game", gameHandler)

    return model
}

func (m *Model) Update() gt.MessageMap {
    return gt.MergeMaps(
        featureAMessages,
        featureBMessages,
        myComponent.UniqueMsgMap(componentMessages),  // Namespaced
    )
}
```

## Routing

```go
// Routes registered in Init()
model.Register("/path", func(s gt.State) []byte {
    return renderView().Bytes()
})

// OnRouteChange hook - called on every navigation
func (m *Model) OnRouteChange(path string) {
    p, _ := url.Parse(path)
    id := m.RouteParam("id")  // ?id=123

    if p.Path == "/game" {
        m.Game.Reset()
    }
}

// Links are auto-intercepted by gotea.js (SPA routing)
h.A(a.Attrs(a.Href("/game")), h.Text("Play"))
```

## Components with Namespacing

Use `ComponentID` to prevent message collisions when reusing components:

```go
type TagSelector struct {
    gt.ComponentID
    Tags []string
}

// Create instances with unique IDs
nameSelector := TagSelector{ComponentID: "name-selector"}
teamSelector := TagSelector{ComponentID: "team-selector"}

// Register with namespace prefix
func (m *Model) Update() gt.MessageMap {
    return gt.MergeMaps(
        nameSelector.UniqueMsgMap(tagMessages),  // "name-selector_SELECT"
        teamSelector.UniqueMsgMap(tagMessages),  // "team-selector_SELECT"
    )
}

// In component render
ts.UniqueMsg("SELECT")     // Message name with prefix
ts.UniqueID("search-box")  // HTML ID with prefix
```

## Broadcasting

For chat, multiplayer, real-time updates:

```go
var app *gt.Application

func main() {
    app = gt.NewApp(&Model{})
    app.Start(8080, "static")
}

func sendMessage(m gt.Message, s gt.State) gt.Response {
    // Update shared state...
    app.Broadcast()  // Re-renders ALL connected clients
    return gt.Respond()
}
```

## Testing

```go
func TestCounter(t *testing.T) {
    session := tester.NewSession(t, &Model{})

    // IMPORTANT: JSON numbers are float64
    session.Dispatch("INCREMENT_COUNTER", float64(1))

    state := session.GetState().(*Model)
    if state.Counter != 1 {
        t.Errorf("Expected 1, got %d", state.Counter)
    }
}
```

## Project Structure

```
/                       # Core library
  runtime.go           # Core engine (~480 lines)
  component.go         # ComponentID namespacing
  templatehelpers.go   # JS message construction
/html                  # h.Div(), h.Span(), etc.
/attributes            # a.Class(), a.OnClick(), etc.
/css                   # css.FontFamily(), css.Flex_(), etc.
/js/gotea.js           # Client-side WebSocket, morphdom, routing
/tester                # Test utilities
/example               # Demo app with multiple features
```

## Dependencies

- `github.com/olahol/melody` - WebSocket handling
- `github.com/google/uuid` - Session IDs
- morphdom (JS, bundled in gotea.js) - DOM patching

## Gotchas

1. **JSON numbers are float64** - In tests/handlers, pass `float64(1)` not `1` for numeric args
2. **Embed gt.Router** - Your Model struct MUST embed `gt.Router`
3. **Register routes in Init()** - Routes must be registered in `Init()`, not at package level
4. **Pointer receivers** - All State interface methods should use pointer receivers
5. **DelayedNextMsg runs async** - Chained messages run in goroutines, state is shared
6. **morphdom preserves focus** - Input focus and selection survive re-renders
7. **Messages are SCREAMING_SNAKE_CASE** - Convention for message names
8. **BlockRerender: true** - Use on messages that shouldn't trigger UI refresh

## Running Tests

```bash
go test ./...
cd example && go test ./...
```
