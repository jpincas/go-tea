module gotea-starter

go 1.23.3

require (
	github.com/google/uuid v1.6.0
	github.com/jpincas/go-tea v0.0.0
)

require (
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/olahol/melody v1.2.1 // indirect
)

// For local development, point to the parent go-tea directory
replace github.com/jpincas/go-tea => ../
