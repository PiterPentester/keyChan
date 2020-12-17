package main

import (
	"log"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	// Load env variables
	var (
		port      = os.Getenv("PORT")
		publicURL = os.Getenv("PUBLIC_URL")     // you must add it to your config vars
		token     = os.Getenv("TELEGRAM_TOKEN") // you must add it to your config vars
	)
	// Set webhook
	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}
	// Init settings
	pref := tb.Settings{
		Token:  token,
		Poller: webhook,
	}
	// Init new bot
	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
	// Reply to "/hello" command
	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hi!")
	})

	b.Start()

}
