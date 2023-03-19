package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	linebot "github.com/line/line-bot-sdk-go/v7/linebot"
)

func lineHandler(bot *linebot.Client) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		events, err := bot.ParseRequest(ctx.Request())
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				return ctx.String(http.StatusBadRequest, "Invalid signature")
			}
			return ctx.String(http.StatusBadRequest, "Error parsing request")
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					log.Println(message)
				}
			}
		}
		return ctx.String(http.StatusOK, "OK")
	}
}
