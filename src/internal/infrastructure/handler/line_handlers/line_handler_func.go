package line_handlers

import "github.com/line/line-bot-sdk-go/v7/linebot"

type LineHandlerFunc func(*linebot.Client, *linebot.Event) error
