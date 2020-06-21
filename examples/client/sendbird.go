package main

import "github.com/maxence-charriere/go-app/v6/pkg/app"

type frontend struct {
	app.Compo
	messages []message
}

type message struct {
	userID       int
	userNickname string
	message      string
}

func (f *frontend) Render() app.UI {
	var nodes []app.Node

	nodes = append(
		nodes,
		app.H1().Body(app.Text("Sendbird Go Client")),
	)

	for _, message := range f.messages {
		nodes = append(
			nodes,
			app.Div().Body(
				app.Span().Text(message.userNickname),
				app.Span().Text(": "),
				app.Span().Text(message.message),
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
	//f.name = src.Get("value").String()
	f.Update()
}
