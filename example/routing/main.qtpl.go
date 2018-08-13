// This file is automatically generated by qtc from "main.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line main.qtpl:1
package main

//line main.qtpl:2
import "github.com/jpincas/go-tea"

//line main.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line main.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line main.qtpl:4
func StreamMain(qw422016 *qt422016.Writer, s gotea.Session) {
	//line main.qtpl:4
	qw422016.N().S(`

    <h1>Routing Example</h1>

    <h2>Current route: `)
	//line main.qtpl:8
	qw422016.E().S(s.State.(Model).Route)
	//line main.qtpl:8
	qw422016.N().S(`</h2>

    <div>
        <a class="gotea-link" href="/1">Page 1</a>
        <a class="gotea-link" href="/2">Page 2</a>   
    </div>

`)
//line main.qtpl:15
}

//line main.qtpl:15
func WriteMain(qq422016 qtio422016.Writer, s gotea.Session) {
	//line main.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line main.qtpl:15
	StreamMain(qw422016, s)
	//line main.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line main.qtpl:15
}

//line main.qtpl:15
func Main(s gotea.Session) string {
	//line main.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
	//line main.qtpl:15
	WriteMain(qb422016, s)
	//line main.qtpl:15
	qs422016 := string(qb422016.B)
	//line main.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
	//line main.qtpl:15
	return qs422016
//line main.qtpl:15
}