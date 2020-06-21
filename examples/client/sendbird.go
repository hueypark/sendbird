package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/hueypark/sendbird"
	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

var (
	applicationID string
	apiToken      string

	// channelURL for tempolary test.
	channelURL = "sendbird"

	// users for tempolary test.
	users = []sendbird.User{
		{
			UserID:   "charles",
			Nickname: "charles",
		},
		{
			UserID:   "david",
			Nickname: "david",
		},
		{
			UserID:   "hueypark",
			Nickname: "hueypark",
		},
		{
			UserID:   "jay",
			Nickname: "jay",
		},
		{
			UserID:   "john",
			Nickname: "john",
		},
	}
)

type frontend struct {
	app.Compo
	sb            *sendbird.Sendbird
	user          *sendbird.User
	lastMessageID int64
	messages      []sendbird.Message
}

func newFrontend() (*frontend, error) {
	f := &frontend{}

	var err error
	f.user, err = getRandomUser()
	if err != nil {
		return nil, err
	}

	f.sb, err = sendbird.NewSendbird(applicationID, apiToken)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (f *frontend) Render() app.UI {
	var nodes []app.Node

	nodes = append(
		nodes,
		app.H1().Body(app.Text("Sendbird Go Client")),
	)

	log.Println(f.messages)

	var messages []sendbird.Message
	var err error
	if f.lastMessageID == 0 {
		f.messages, err = f.sb.Messages(
			sendbird.OpenChannels,
			channelURL,
			sendbird.MessageReadTimestamp,
			time.Now().Unix()*1000,
		)
	} else {
		messages, err = f.sb.Messages(
			sendbird.OpenChannels,
			channelURL,
			sendbird.MessageReadID,
			f.lastMessageID,
		)
	}
	if err != nil {
		log.Fatalln(err)
	}

	f.updateLastMessageID(messages)

	f.messages = append(f.messages, messages...)

	for _, message := range f.messages {
		nodes = append(
			nodes,
			app.Div().Body(
				app.Span().Text(message.User.Nickname),
				app.Span().Text(": "),
				app.Span().Text(message.Message),
			),
		)
	}

	nodes = append(
		nodes,
		app.Div().Body(
			app.Input().
				AutoFocus(true).
				OnChange(f.OnChat),
		),
	)

	return app.Div().Body(
		app.Main().Body(nodes...),
	)
}

func (f *frontend) OnChat(src app.Value, e app.Event) {
	message := src.Get("value").String()

	_, err := f.sb.SendMessage(sendbird.OpenChannels, channelURL, f.user.UserID, message)
	if err != nil {
		log.Println(err)
	}

	f.Update()
}

func (f *frontend) updateLastMessageID(messages []sendbird.Message) {
	for _, message := range messages {
		if f.lastMessageID <= message.MessageID {
			f.lastMessageID = message.MessageID
		}
	}
}

func getRandomUser() (*sendbird.User, error) {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(users))
	for i, user := range users {
		if i != idx {
			continue
		}

		return &user, nil
	}

	return nil, fmt.Errorf("can not get random user")
}
