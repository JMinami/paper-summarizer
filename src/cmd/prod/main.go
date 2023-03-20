package main

import (
	"log"
	"paper-summarizer/internal/config"
	"paper-summarizer/internal/container"
	"paper-summarizer/internal/infrastructure/handler"
)

func main() {
	c := config.New()

	diContainer, err := container.NewContainler()
	if err != nil {
		log.Fatal(err)
	}
	handler.HttpRouter(c, diContainer)
}
