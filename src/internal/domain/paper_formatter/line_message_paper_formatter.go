package paper_formatter

import (
	"paper-summarizer/internal/domain"
)

type (
	LineMessagePaperFormatter struct{}
)

var _ (domain.PaperFormatter) = (*LineMessagePaperFormatter)(nil)

func NewLineMessagePaperFormatter() *LineMessagePaperFormatter {
	return &LineMessagePaperFormatter{}
}

func (f *LineMessagePaperFormatter) Format(papers domain.Papers) string {
	var formattedPapers string
	for _, paper := range papers {
		formattedPapers += "Title: " + paper.Title + "\nSummary: " + paper.Summary + "\nURL: " + paper.URL + "\n\n"
	}
	return formattedPapers
}
