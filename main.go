package main

import (
	"time"
  "log"
	"os"

	"github.com/joho/godotenv"
	tele "gopkg.in/tucnak/telebot.v3"
)

func main() {
	// Load env variables
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}
	// Get token from env
	token := os.Getenv("TELEGRAM_TOKEN")
	// Init bot settings
	b, err := tele.NewBot(tele.Settings{
		Token: token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Println(err)
		return
	}
	// Hanle "/start" command
	b.Handle("/start", func (c tele.Context) error {
	  return c.Reply("Hello!")
  })
  // Start bot
	b.Start()
}
