package paper_formatter

import (
	"testing"

	"paper-summarizer/internal/domain"
)

func TestLineMessagePaperFormatter_Format(t *testing.T) {
	papers := domain.Papers{
		{
			Title:   "Paper 1",
			Summary: "This is the summary of Paper 1",
			URL:     "https://example.com/paper1",
		},
		{
			Title:   "Paper 2",
			Summary: "This is the summary of Paper 2",
			URL:     "https://example.com/paper2",
		},
	}
	formatter := NewLineMessagePaperFormatter()
	result := formatter.Format(papers)
	expected := "Title: Paper 1\nSummary: This is the summary of Paper 1\nURL: https://example.com/paper1\n\n" +
		"Title: Paper 2\nSummary: This is the summary of Paper 2\nURL: https://example.com/paper2\n\n"
	if result != expected {
		t.Errorf("Result should match expected output. Expected: %q, Got: %q", expected, result)
	}
}
