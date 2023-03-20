package line_handlers

import (
	"paper-summarizer/internal/application"
	"paper-summarizer/internal/container"
	"paper-summarizer/internal/infrastructure/client"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func LinePaperSummarizer(
	diContainer *container.Container,
	keyword string,
	maxNum int,
) LineHandlerFunc {
	return func(bot *linebot.Client, event *linebot.Event) error {
		lineClient, err := client.NewLineClientBuilder().SetBot(bot).SetEvent(event).Build()
		if err != nil {
			return err
		}

		paperSummarySendService := application.NewPaperSummarySendService(
			diContainer.PaperService, diContainer.PaperFormatterService, lineClient,
		)
		if err := paperSummarySendService.Send(keyword, maxNum); err != nil {
			return err
		}
		return nil
	}
}
