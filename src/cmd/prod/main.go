package main

import (
	"log"
	"paper-summarizer/internal/config"
	"paper-summarizer/internal/infrastructure/handler"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	c := config.New()
	bot, err := linebot.New(c.LineMessagingAPIChannelSecret, c.LineMessagingAPIChannelAccessToken)
	if err != nil {
		log.Fatal(err)
	}

	handler.HttpRouter(bot)
}
