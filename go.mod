module github.com/jpincas/go-tea

go 1.23.3

require github.com/jpincas/htmlfunc v0.0.0-00010101000000-000000000000 // Add this line

replace github.com/jpincas/htmlfunc => ../htmlfunc // Move this line outside of the require block

require github.com/olahol/melody v1.2.1

require github.com/gorilla/websocket v1.5.0 // indirect
