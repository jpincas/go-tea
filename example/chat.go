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
		h.H2(a.Attrs(a.Class("text-2xl font-bold text-gray-900")), h.Text("Chat Room")),
		
		h.Div(
			a.Attrs(a.Class("bg-white p-4 rounded-lg shadow-sm border border-gray-200 space-y-4")),
			h.Div(
				a.Attrs(a.Class("flex items-end space-x-2")),
				h.Div(
					a.Attrs(a.Class("flex-grow")),
					h.Label(a.Attrs(a.For("usernameInput"), a.Class("block text-sm font-medium text-gray-700 mb-1")), h.Text("Enter your username:")),
					h.Input(a.Attrs(a.Type("text"), a.Id("usernameInput"), a.Value(chat.Username), a.Class("block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"))),
				),
				h.Button(
					a.Attrs(a.OnClick(gt.SendBasicMessageWithValueFromInput("SET_USERNAME", "usernameInput")), a.Class("inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500")),
					h.Text("Set Username"),
				),
			),
		),

		h.Div(
			a.Attrs(a.Id("messages"), a.Class("bg-gray-50 p-4 rounded-lg border border-gray-200 h-96 overflow-y-auto space-y-4")),
			func() []h.Element {
				var elements []h.Element
				for _, msg := range *chat.Messages {
					containerClass := "flex justify-start"
					bubbleClass := "bg-white text-gray-800 border border-gray-200"
					
					if msg.User == chat.Username {
						containerClass = "flex justify-end"
						bubbleClass = "bg-indigo-100 text-indigo-900 border border-indigo-200"
					}
					
					elements = append(elements, h.Div(
						a.Attrs(a.Class(containerClass)),
						h.Div(
							a.Attrs(a.Class(fmt.Sprintf("max-w-xs lg:max-w-md px-4 py-2 rounded-lg shadow-sm %s", bubbleClass))),
							h.Div(
								a.Attrs(a.Class("flex justify-between items-baseline mb-1 space-x-2")),
								h.Span(a.Attrs(a.Class("font-bold text-xs")), h.Text(msg.User)),
								h.Span(a.Attrs(a.Class("text-xs opacity-75")), h.Text(msg.TimeStamp.Format("15:04"))),
							),
							h.P(a.Attrs(a.Class("text-sm")), h.Text(msg.Text)),
						),
					))
				}
				return elements
			}()...,
		),

		h.Div(
			a.Attrs(a.Class("bg-white p-4 rounded-lg shadow-sm border border-gray-200 space-y-4")),
			h.P(a.Attrs(a.Class("text-sm text-gray-500 italic")), h.Text(fmt.Sprintf("Posting as %s", chat.Username))),
			h.Div(
				a.Attrs(a.Class("flex space-x-2")),
				h.Input(a.Attrs(a.Type("text"), a.Id("messageInput"), a.Placeholder("Type your message..."), a.Class("block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"))),
				h.Button(
					a.Attrs(a.OnClick(gt.SendBasicMessageWithValueFromInput("SEND_MESSAGE", "messageInput")), a.Class("inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500")),
					h.Text("Send"),
				),
			),
		),
	)
}
