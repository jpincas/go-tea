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

    <form>
        First name:<br>
        <input type="text" name="firstname" value="Mickey">
        <br>
        Last name:<br>
        <input type="text" name="lastname" value="Mouse">
        <br><br>
        <button
            type='button' 
            class='gotea-form-submit'
            data-msg='`)
	//line main.qtpl:16
	qw422016.E().S(gotea.Msg(SubmitForm()))
	//line main.qtpl:16
	qw422016.N().S(`'>
                Add Person
        </button>
    </form>

    <div>
        <h1>Here's who we know about...</h1>
        <ul>
            `)
	//line main.qtpl:24
	for _, person := range *s.State.(Model).People {
		//line main.qtpl:24
		qw422016.N().S(`
            <li>`)
		//line main.qtpl:25
		qw422016.E().S(person.FirstName)
		//line main.qtpl:25
		qw422016.N().S(` `)
		//line main.qtpl:25
		qw422016.E().S(person.LastName)
		//line main.qtpl:25
		qw422016.N().S(`</li>
            `)
		//line main.qtpl:26
	}
	//line main.qtpl:26
	qw422016.N().S(`
        </ul>

    </div>

`)
//line main.qtpl:31
}

//line main.qtpl:31
func WriteMain(qq422016 qtio422016.Writer, s gotea.Session) {
	//line main.qtpl:31
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line main.qtpl:31
	StreamMain(qw422016, s)
	//line main.qtpl:31
	qt422016.ReleaseWriter(qw422016)
//line main.qtpl:31
}

//line main.qtpl:31
func Main(s gotea.Session) string {
	//line main.qtpl:31
	qb422016 := qt422016.AcquireByteBuffer()
	//line main.qtpl:31
	WriteMain(qb422016, s)
	//line main.qtpl:31
	qs422016 := string(qb422016.B)
	//line main.qtpl:31
	qt422016.ReleaseByteBuffer(qb422016)
	//line main.qtpl:31
	return qs422016
//line main.qtpl:31
}