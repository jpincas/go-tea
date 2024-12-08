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
		a.Attrs(),
		h.H2(a.Attrs(), h.Text("Chat Room")),
		h.Div(
			a.Attrs(),
			h.Label(a.Attrs(a.For("usernameInput")), h.Text("Enter your username:")),
			h.Input(a.Attrs(a.Type("text"), a.Id("usernameInput"), a.Value(chat.Username), a.Class("username-input"))),
			h.Button(
				a.Attrs(a.OnClick(gt.SendBasicMessageWithValueFromInput("SET_USERNAME", "usernameInput")), a.Class("set-username-button")),
				h.Text("Set Username"),
			),
		),
		h.Div(
			a.Attrs(a.Id("messages")),
			func() []h.Element {
				var elements []h.Element
				for _, msg := range *chat.Messages {
					class := "message-right"
					if msg.User == chat.Username {
						class = "message-left"
					}
					elements = append(elements, h.Div(
						a.Attrs(a.Class(class)),
						h.Span(a.Attrs(a.Class("message-username")), h.Text(msg.User)),
						h.Span(a.Attrs(a.Class("message-timestamp")), h.Text(msg.TimeStamp.Format("15:04:05"))),
						h.P(a.Attrs(a.Class("message-text")), h.Text(msg.Text)),
					))
				}
				return elements
			}()...,
		),
		h.P(a.Attrs(a.Class("username-display")), h.Text(fmt.Sprintf("Posting as %s", chat.Username))),
		h.Input(a.Attrs(a.Type("text"), a.Id("messageInput"), a.Placeholder("Your message"), a.Class("message-input"))),
		h.Button(
			a.Attrs(a.OnClick(gt.SendBasicMessageWithValueFromInput("SEND_MESSAGE", "messageInput")), a.Class("send-button")),
			h.Text("Send"),
		),
	)
}
