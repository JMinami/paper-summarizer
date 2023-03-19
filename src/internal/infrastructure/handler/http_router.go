package handler

import (
	"paper-summarizer/internal/config"
	"paper-summarizer/internal/container"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func HttpRouter(
	conf *config.Config,
	bot *linebot.Client,
	container *container.Container,
) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/line/callback", lineHandler(conf, bot, container))
	e.Logger.Fatal(e.Start(":8080"))
}
