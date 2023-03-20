package client

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type (
	LineClientBuilder struct {
		bot   *linebot.Client
		event *linebot.Event
	}
)

func NewLineClientBuilder() *LineClientBuilder {
	return &LineClientBuilder{
		bot:   nil,
		event: nil,
	}
}

func (b *LineClientBuilder) SetBot(bot *linebot.Client) *LineClientBuilder {
	b.bot = bot
	return b
}

func (b *LineClientBuilder) SetEvent(event *linebot.Event) *LineClientBuilder {
	b.event = event
	return b
}

func (b *LineClientBuilder) Build() (*LineClient, error) {
	if b.bot == nil {
		return nil, fmt.Errorf("bot is not set")
	}
	if b.event == nil {
		return nil, fmt.Errorf("event is not set")
	}

	return NewLineClient(b.bot, b.event), nil
}
