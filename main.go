package main

import (
	"flag"
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var token string

	// CLI arguments
	flag.StringVar(&token, "token", "", "Token for Telegram bot")
	flag.Parse()

	// Checking whether the token is empty
	if token == "" {
		log.Fatal("Token is empty")
	}

	// Connect to bot with token
	bot, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
	}

	// Start handler
	bot.Handle("/start", func(m *tb.Message) {
		bot.Send(m.Chat, `Hello! I'm Telegram bot for PlantUML
Read the documentation from https://plantuml.com/
And send me Plantuml code`)
	})

	// Processing Plantuml code
	bot.Handle(tb.OnText, func(m *tb.Message) {
		photoURL := genPlantumlLink(m.Text)
		bot.Send(m.Chat, photoURL)
	})

	// Start bot
	bot.Start()
}
