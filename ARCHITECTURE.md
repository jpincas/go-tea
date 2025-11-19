# Gotea Architecture

## Overview
Gotea is an experimental Go library implementing **The Elm Architecture (TEA)** for server-side rendered web applications. It enables building dynamic SPAs using Go for logic and state management, with minimal JavaScript.

## Core Concepts
- **Server-Side State**: Application state (`Model`) resides on the server.
- **Websockets**: Real-time communication via `melody`.
- **Morphdom**: Efficient DOM patching on the client.
- **Routing**: Server-side routing logic with client-side history management.

## Key Components
- **`runtime.go`**: Core engine (State, Message loop, Websockets).
- **`js/gotea.js`**: Client-side websocket and DOM patching logic.
- **`component.go`**: Component abstraction.

## Data Flow
1.  **Init**: Server initializes state and renders HTML.
2.  **Connect**: Client connects via Websocket.
3.  **Interaction**: User triggers event -> JS sends JSON message to Server.
4.  **Update**: Server handles message -> Updates State -> Renders new HTML.
5.  **Patch**: Server sends HTML -> Client patches DOM with `morphdom`.

## Project Structure
- `/`: Core library files.
- `/js`: Client-side JavaScript.
- `/example`: Demo application.
