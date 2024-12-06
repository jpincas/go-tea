package main

import (
	"encoding/json"
	"fmt"
	"time"

	gt "github.com/jpincas/go-tea"
	"github.com/jpincas/go-tea/msg"
	a "github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
)

var chatMessages gt.MessageMap = gt.MessageMap{
	"SEND_MESSAGE": SendMessage,
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

var messages []Message

func randomChatUsername() string {
	return fmt.Sprintf("User%d", time.Now().Unix())
}

func (chat *Chat) AddMessage(message Message) {
	messages = append(messages, message)
}

func SendMessage(args json.RawMessage, s gt.State) gt.Response {
	state := model(s)

	input, err := msg.DecodeString(args)
	if err != nil {
		return gt.RespondWithError(err)
	}

	messages = append(messages, Message{
		TimeStamp: time.Now(),
		User:      state.Chat.Username,
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
			a.Attrs(a.OnClick(gt.SendMessageWithInputValue("SEND_MESSAGE", "messageInput")), a.Class("send-button")),
			h.Text("Send"),
		),
	)
}
