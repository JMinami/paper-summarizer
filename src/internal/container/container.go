package container

import (
	"paper-summarizer/internal/application"
	"paper-summarizer/internal/domain/paper_formatter"
	"paper-summarizer/internal/infrastructure/client"
)

type Container struct {
	PaperService          *application.PaperService
	PaperFormatterService *application.PaperFormatService
}

func NewContainler() (*Container, error) {
	translator := client.NewGCPTranslator()
	fetcher := client.NewArxivClient()
	formatter := paper_formatter.NewLineMessagePaperFormatter()

	paperService := application.NewPaperService(translator, fetcher)
	paperFormatService := application.NewPaperFormatter(formatter)

	return &Container{
		PaperService:          paperService,
		PaperFormatterService: paperFormatService,
	}, nil
}
