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

	// Reply to "/start" command
	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, "Usage:")
		b.Send(m.Sender, "/memp - generate memorable password")
		b.Send(m.Sender, "/abrp - generate abracadabra password")
		b.Send(m.Sender, "/hello - return 'Hi!'")
	})
	// Reply to "/hello" command
	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hi!")
	})

	// Reply to "/memp" command
	b.Handle("/memp", func(m *tb.Message) {
		// get words
		wrds, _ := GetRandWords(4)
		// generating memorable password
		b.Send(m.Sender, "Memorable password:")
		b.Send(m.Sender, GenMemorablePass(wrds))
	})

	// Reply to "/abrp" command
	b.Handle("/abrp", func(m *tb.Message) {
		b.Send(m.Sender, "Abracadabra pass:")
		b.Send(m.Sender, String(20))
	})

	b.Start()

}
