package handler

import (
	"log"
	"net/http"
	"paper-summarizer/internal/config"
	"paper-summarizer/internal/container"
	"paper-summarizer/internal/infrastructure/handler/line_handlers"

	"github.com/labstack/echo/v4"
	linebot "github.com/line/line-bot-sdk-go/v7/linebot"
)

func lineHandler(conf *config.Config, container *container.Container) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		bot, err := linebot.New(conf.LineMessagingAPIChannelSecret, conf.LineMessagingAPIChannelAccessToken)
		if err != nil {
			return err
		}
		events, err := bot.ParseRequest(ctx.Request())
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				return ctx.String(http.StatusBadRequest, "Invalid signature")
			}
			return ctx.String(http.StatusBadRequest, "Error parsing request")
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				var err error
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					err = line_handlers.LinePaperSummarizer(container, message.Text, conf.PaperMaxNum)(bot, event)
				}
				if err != nil {
					log.Println(err)
					ctx.String(http.StatusInternalServerError, err.Error())
				}
			}
		}
		return ctx.String(http.StatusOK, "OK")
	}
}
