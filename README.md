# 🍵 Gotea

### Status: works well, but still experimental, API in constant flux and needs a lot of polishing.  If you're interested in using/helping, probably best to contact me.

## What?

Gotea is a small, experimental library that loosely implements a form of TEA (The Elm Architecture) or a Redux-like system, but in Go, using server-side state, server-rendered HTML, websockets and messages - it lets you create SPA-like applications without having to write a single line of Javascript.  

Since coming up with the idea for Gotea, a few similar frameworks in other languages/platforms have appeared, chiefly inspired by Pheonix's LiveView.  It might be worth checking out some of their documentation/talks to get a high-level picture of this type of architecture.

- [LiveView for Pheonix](https://hexdocs.pm/phoenix_live_view/Phoenix.LiveView.html)
- [StimulusReflex for Ruby](https://github.com/hopsoft/stimulus_reflex)
- [Blazor for .NET](https://dotnet.microsoft.com/apps/aspnet/web-apps/blazor)
- [phpx for PHP](https://freek.dev/1254-introducing-phpx-implementing-phoenix-liveview-in-php)
- [LiveView for Crystal](https://github.com/jgaskins/live_view)

It's encouraging to see genuine interest in this type of architecture and Pheonix's success in creating something quite mature and powerful makes me confident that we can do the same thing in Go - after all Go has been proven to handle millions of concurrent websocket connections smoothly.

## Why?

My team and I have found TEA (and particularly Elm) to be a joy for writing client-side apps.  But we've also come to think that not all use cases are best served by a full-on SPA.  This is particurly so when the 'app' cannot do much without talking to a server.  For 'server heavy' applications where each interaction is normally accompanied by a round trip to the server, I'm always tempted to just use a traditional server-rendered HTML architecture, but the static nature of the UI and eventual need to start mixing in dynamic bits with JS is a pain point. Gotea's sweetspot is applications that need a complex, dynamic UI but are limited in what they can do in isolation from the server.  You definitely wouldn't, for example, write a game with Gotea, but a server-logic-heavy admin panel would be a great fit.

## How it works

If you prefer to read code, just skim through `runtime.go` which will get you up to speed in a couple of hundred lines.

1.  You define what your application state looks like with a standard Go struct.
2.  You specify what this state should be initialised to at the start of a user session.
3.  You provide a 'view' function for rendering your state (you can use Go templates, 3rd-party templates, string concatenation etc)
4.  You specify a number of 'messages' along with their handler functions.  These will be responsible for modifying application state.
5.  A browser connects to your go-tea application and is immediately served with static HTML as determined by the initial state and template combo.  
6.  A small amount of JS does the work of establishing a websocket connection with the Go application.
7.  The Gotea runtime initiates a 'session' for the client and will use the open websocket for all further communication.
8.  The user interacts with your site/app in some way, a 'message' is triggered (again, a small amount of JS is necessary for this, but you don't have to touch it).
9.  This message is sent through the websocket to the Gotea runtime.
10.  The Gotea runtime executes the message handler attached to the message, modifying application state.
11.  The application state is rerendered and sent back down the websocket to the client.
12.  The incoming DOM is patched onto the existing DOM using Morphdom, giving a fast and seamless update to the UI.

## JS

Go-tea requires a small amount of JS to work - you'll find it in the `js` directory.  Application authors don't have to touch it.  If you do any work on it, you need to rebuild the `dist` directory with `npm run build` and run `go generate` to hardcode the resulting JS into the Go program.


Warning: Parcel seems to be a bit messed up when building the example, it can nuke the `package.json` and/or delete Parcel itself from `node_modules`.  Be careful not to check in a nuked `package.json`

Be warned: if you are working on gotea core AND the example, and you make changes to the go-tea core JS, you need to first rebuild that, THEN rebuild the example JS, since it includes the core JS. Clear?


## Routing

Gotea has (experimental) support for routes built in.  You can specify a special, 'route change' function which modifies application state and/or the rendering template according to any custom logic.  This function is fired at startup and on any route changes.  On the client side, the Gotea JS 'upgrades' any internal links it finds to use the Gotea 'route change' method instead of a browser request and refresh and takes care of browser history.

## Example

The `/example` repo demonstrates many of go-tea's capabilities, including routing, components and forms.  Run it and go to `localhost:8080`  



