package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	gt "github.com/jpincas/go-tea"
	a "github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
)

var usernames = make(map[uuid.UUID]string)
var messages []Message

var chatMessages gt.MessageMap = gt.MessageMap{
	"SEND_MESSAGE": SendMessage,
	"SET_USERNAME": SetUsername,
}

type Chat struct {
	Username string
	Messages *[]Message
}

type Message struct {
	TimeStamp time.Time
	User      string
	Text      string
}

func (chat *Chat) AddMessage(message Message) {
	messages = append(messages, message)
}

func SetUsername(m gt.Message, s gt.State) gt.Response {
	state := model(s)
	input := m.ArgsToString()

	// Cache the username against the sid so we can use it in init
	usernames[state.sessionID] = input
	state.Chat.Username = input

	return gt.Respond()
}

func SendMessage(m gt.Message, s gt.State) gt.Response {
	state := model(s)
	input := m.ArgsToString()

	messages = append(messages, Message{
		TimeStamp: time.Now(),
		User:      usernames[state.sessionID],
		Text:      input,
	})

	app.Broadcast()

	return gt.Respond()
}

func (chat Chat) render() h.Element {
	return h.Div(
		a.Attrs(a.Class("space-y-6")),
		// Header
		h.Div(
			a.Attrs(a.Class("text-center space-y-2")),
			h.H1(a.Attrs(a.Class("text-4xl font-bold text-stone-900"), a.Custom("style", "font-family: 'DM Serif Display', serif;")), h.Text("💬 Chat Room")),
			h.P(a.Attrs(a.Class("text-stone-600")), h.Text("Real-time chat with broadcasting — messages sync instantly across all connected clients!")),
		),

		// Username setup
		h.Div(
			a.Attrs(a.Class("bg-gradient-to-r from-violet-50 to-purple-50 p-5 rounded-xl border-2 border-stone-900 shadow-brutal-sm")),
			h.Div(
				a.Attrs(a.Class("flex flex-col sm:flex-row items-stretch sm:items-end gap-3")),
				h.Div(
					a.Attrs(a.Class("flex-grow")),
					h.Label(a.Attrs(a.For("usernameInput"), a.Class("block text-sm font-semibold text-stone-700 mb-1")), h.Text("Your Display Name")),
					h.Input(a.Attrs(
						a.Type("text"),
						a.Id("usernameInput"),
						a.Value(chat.Username),
						a.Placeholder("Enter a username.."),
						a.Class("block w-full rounded-lg px-4 py-2.5 text-stone-900"),
					)),
				),
				h.Button(
					a.Attrs(
						a.OnClick(gt.SendBasicMessageWithValueFromInput("SET_USERNAME", "usernameInput")),
						a.Class("px-5 py-2.5 bg-violet-600 hover:bg-violet-700 text-white font-semibold rounded-lg border-2 border-stone-900 shadow-brutal-sm hover:shadow-brutal hover:-translate-x-0.5 hover:-translate-y-0.5 transition-all whitespace-nowrap"),
					),
					h.Text("Set Name"),
				),
			),
		),

		// Messages area
		h.Div(
			a.Attrs(a.Id("messages"), a.Class("bg-stone-100 p-4 rounded-xl border-2 border-stone-300 h-80 overflow-y-auto space-y-3")),
			func() []h.Element {
				if len(*chat.Messages) == 0 {
					return []h.Element{
						h.Div(
							a.Attrs(a.Class("flex items-center justify-center h-full text-stone-400")),
							h.Text("No messages yet. Start the conversation! 👋"),
						),
					}
				}
				var elements []h.Element
				for _, msg := range *chat.Messages {
					isOwn := msg.User == chat.Username
					containerClass := "flex justify-start"
					bubbleClass := "bg-white border-stone-300"
					textClass := "text-stone-800"

					if isOwn {
						containerClass = "flex justify-end"
						bubbleClass = "bg-emerald-100 border-emerald-400"
						textClass = "text-emerald-900"
					}

					elements = append(elements, h.Div(
						a.Attrs(a.Class(containerClass)),
						h.Div(
							a.Attrs(a.Class(fmt.Sprintf("max-w-xs lg:max-w-md px-4 py-2.5 rounded-xl border-2 %s", bubbleClass))),
							h.Div(
								a.Attrs(a.Class("flex justify-between items-baseline mb-1 gap-3")),
								h.Span(a.Attrs(a.Class(fmt.Sprintf("font-bold text-xs %s", textClass))), h.Text(msg.User)),
								h.Span(a.Attrs(a.Class("text-xs text-stone-400"), a.Custom("style", "font-family: 'JetBrains Mono', monospace;")), h.Text(msg.TimeStamp.Format("15:04"))),
							),
							h.P(a.Attrs(a.Class(fmt.Sprintf("text-sm %s", textClass))), h.Text(msg.Text)),
						),
					))
				}
				return elements
			}()...,
		),

		// Message input
		h.Div(
			a.Attrs(a.Class("bg-white p-4 rounded-xl border-2 border-stone-900 shadow-brutal-sm")),
			h.Div(
				a.Attrs(a.Class("flex items-center gap-2 mb-3")),
				h.Span(a.Attrs(a.Class("text-sm text-stone-500")), h.Text("Posting as")),
				h.Span(a.Attrs(a.Class("text-sm font-bold text-stone-900 bg-stone-100 px-2 py-0.5 rounded")), h.Text(func() string {
					if chat.Username == "" {
						return "Anonymous"
					}
					return chat.Username
				}())),
			),
			h.Div(
				a.Attrs(a.Class("flex gap-2")),
				h.Input(a.Attrs(
					a.Type("text"),
					a.Id("messageInput"),
					a.Placeholder("Type your message.."),
					a.Class("flex-grow rounded-lg px-4 py-2.5 text-stone-900"),
				)),
				h.Button(
					a.Attrs(
						a.OnClick(gt.SendBasicMessageWithValueFromInput("SEND_MESSAGE", "messageInput")),
						a.Class("px-5 py-2.5 bg-emerald-500 hover:bg-emerald-600 text-white font-semibold rounded-lg border-2 border-stone-900 shadow-brutal-sm hover:shadow-brutal hover:-translate-x-0.5 hover:-translate-y-0.5 transition-all"),
					),
					h.Text("Send →"),
				),
			),
		),

		// Explanatory note
		renderExplanatoryNote(
			"Real-time Chat",
			`
			<p class="mb-3">A simple chat application demonstrating shared state and broadcasting.</p>
			<ul class="list-disc pl-5 space-y-2">
				<li><strong class="text-stone-900">Shared State:</strong> The list of messages is a global variable shared across all sessions.</li>
				<li><strong class="text-stone-900">Broadcasting:</strong> When a new message is received, the server calls <code class="bg-stone-200 px-1.5 py-0.5 rounded text-xs font-mono">app.Broadcast()</code> to re-render all connected clients.</li>
				<li><strong class="text-stone-900">Session State:</strong> The username is stored in the session-specific state.</li>
			</ul>
			`,
		),
	)
}
