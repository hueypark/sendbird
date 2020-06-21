// +build wasm

package main

import (
	"log"
	"time"

	"github.com/hueypark/sendbird"
	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

var (
	applicationID string
	apiToken      string
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	sb, err := sendbird.NewSendbird(applicationID, apiToken)
	if err != nil {
		log.Fatalln(err)
	}

	res, err := sb.Messages(sendbird.OpenChannels, "mars2", time.Now().Unix())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)

	f := &frontend{}
	f.messages = []message{
		{
			userID:       1,
			userNickname: "Huey",
			message:      "Hi",
		},
		{
			userID:       2,
			userNickname: "Park",
			message:      "How are you?",
		},
	}

	app.Route("/", f)
	app.Run()
}
