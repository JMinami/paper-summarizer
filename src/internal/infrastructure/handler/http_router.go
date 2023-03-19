package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func HttpRouter(bot *linebot.Client) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/line/callback", lineHandler(bot))
	e.Logger.Fatal(e.Start(":8080"))
}
