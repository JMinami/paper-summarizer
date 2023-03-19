package line_handlers

import (
	"paper-summarizer/internal/application"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func LinePaperSummarizer(
	bot *linebot.Client,
	event *linebot.Event,
	paperService *application.PaperService,
	paperFormatService *application.PaperFormatService,
	keyword string,
	maxNum int,
) error {
	translationPapers, err := paperService.TranslatePapers(keyword, maxNum)
	if err != nil {
		return err
	}

	message := paperFormatService.FormatPaper(translationPapers)

	call := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message))
	if _, err := call.Do(); err != nil {
		return err
	}

	return nil
}
