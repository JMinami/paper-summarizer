package client

import (
	"paper-summarizer/internal/domain"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type (
	LineClient struct {
		bot   *linebot.Client
		event *linebot.Event
	}
)

var _ (domain.MessageSender) = (*LineClient)(nil)

func NewLineClient(bot *linebot.Client, event *linebot.Event) *LineClient {
	return &LineClient{
		bot:   bot,
		event: event,
	}
}

func (c *LineClient) Send(message string) error {
	call := c.bot.PushMessage(
		c.event.Source.UserID,
		linebot.NewTextMessage(message),
	)
	if _, err := call.Do(); err != nil {
		return err
	}
	return nil
}
