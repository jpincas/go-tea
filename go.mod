module github.com/jpincas/go-tea

go 1.23.3

require (
	github.com/gorilla/websocket v1.5.3
	github.com/jpincas/htmlfunc v0.0.0-00010101000000-000000000000 // Add this line
	github.com/satori/go.uuid v1.2.0
)

replace github.com/jpincas/htmlfunc => ../htmlfunc // Move this line outside of the require block

require gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
