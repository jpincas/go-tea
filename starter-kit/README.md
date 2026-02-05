# Gotea Starter Kit

A minimal but complete Gotea application demonstrating all core patterns.

## Quick Start

```bash
# Install JS dependencies and build
npm install
npm run build

# Run the Go server
go run .

# Visit http://localhost:8080
```

## Project Structure

```
├── main.go         # App setup, Model struct, State interface, routing
├── messages.go     # All message handlers
├── views.go        # Render functions for each page
├── component.go    # Reusable component with ComponentID
├── form.go         # Form state struct
├── js/
│   └── main.js     # Imports gotea client
├── static/
│   ├── main.js     # Built JS (after npm run build)
│   └── styles.css  # Custom CSS
├── package.json    # JS build config
└── go.mod          # Go module
```

## Patterns Demonstrated

### 1. Counter (messages.go, views.go)
Basic message handling with `SendBasicMessage` and `ArgsToInt()`.

### 2. Form Handling (form.go, views.go)
Two-way form binding using `BasicUpdateForm` and `MustDecodeArgs`.

### 3. Routing (main.go)
Multiple routes registered in `Init()`, rendered via `RenderRoute()`.

### 4. Reusable Component (component.go)
`ComponentID` for namespaced messages preventing collisions.

### 5. Conditional Rendering (views.go)
Using `RenderIf()` for conditional elements.

## Key Files to Modify

- **Add state fields**: `main.go` → `Model` struct
- **Add message handlers**: `messages.go`
- **Add views**: `views.go`
- **Add routes**: `main.go` → `Init()`
- **Register messages**: `main.go` → `Update()`

## Common Operations

### Adding a New Page

1. Add route in `Init()`:
```go
model.Register("/newpage", func(s gt.State) []byte {
    return renderNewPage(model).Bytes()
})
```

2. Add render function in `views.go`
3. Add nav link in `renderNav()`

### Adding a New Message

1. Add handler in `messages.go`:
```go
var myMessages = gt.MessageMap{
    "MY_ACTION": myHandler,
}

func myHandler(m gt.Message, s gt.State) gt.Response {
    state := model(s)
    // Mutate state...
    return gt.Respond()
}
```

2. Register in `Update()`:
```go
return gt.MergeMaps(
    counterMessages,
    myMessages, // Add here
)
```

3. Trigger from HTML:
```go
a.OnClick(gt.SendBasicMessage("MY_ACTION", args))
```

## Gotchas

- JSON numbers are `float64` — use `m.ArgsToInt()` not type assertion
- Model MUST embed `gt.Router`
- Routes must be registered in `Init()`, not at package level
- Use pointer receivers for all State interface methods
