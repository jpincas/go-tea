# Go-Tea

Go-tea is an implementation of a TEA (The Elm Architecture) or Redux-like system, in Go, using server rendered HTML, websockets and messages to create a SPA-like experience without having to write Javascript.

## How it works

First, read the code in `runtime.go` which is short, heavily commented, and walks you through the system step by step.  Here's a brief, high-level summary:

1 - You define what your application state looks like with a stanard Go struct.
2 - You specify what this state should be initialised to at the start of a user session.
3 - You specify a render function that renders your app state to HTML.
4 - You specify a number of 'messages' along with their handler functions.  These will be responsibly for modifying application state.
5 - A browser connects to your go-tea application.  A small amount of JS does the work of establishing a websocket connection with the Go application.
6 - The go-tea runtime initiates a 'session' for the client and will use the open websocket for all further communication.
7 - The application sends an initial HTML render down the websocket to the client.
8 - The browser display this HTML
9 - The user interacts with your site/app in some way, a 'message' is triggered (again, a small amount of JS is necessary for this, but you don't have to touch it).
10 - This message is sent up the wire to the go-tea runtime.
11 - The go-tea runtime executes the message handler attached to the message, modifying application state.
12 - The application state is rerendered and sent back down the wire to the client.
13 - The incoming DOM is patched onto the existing DOM using Morphdom, giving a seamless update to the UI.

## JS

Go-tea requires a small amount of JS to work - you'll find it in the `js` directory.  Application authors don't have to touch it.  If you do any work on it, you need to rebuild the `dist` directory with `npm run build`.

## Example

The `/example` repo demonstrates many of go-tea's capabilities, including routing, components and forms.  First, build the static assets (again, `npm run build`), then run the example with `go run *.go`.  Be warned: if you are working on gotea core AND the example, and you make changes to the go-tea core JS, you need to first rebuild that, THEN rebuild the example JS, since it includes the core JS. Clear?

