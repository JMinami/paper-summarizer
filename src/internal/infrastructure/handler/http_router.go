package handler

import (
	"paper-summarizer/internal/config"
	"paper-summarizer/internal/container"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HttpRouter(
	conf *config.Config,
	container *container.Container,
) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/line/callback", lineHandler(conf, container))
	e.Logger.Fatal(e.Start(":8080"))
}
