package client

import (
	"paper-summarizer/internal/domain"
	"testing"
)

func TestFetchPapers(t *testing.T) {
	sut := NewArxivClient()
	maxNum := 10
	papers, err := sut.Fetch(domain.RequestFetch{
		Keyword: "quantum computing",
		MaxNum:  maxNum,
	})
	if err != nil {
		t.Fatalf("Error fetching papers: %v", err)
	}

	// 1つ以上,maxnum以下の論文が返されることを確認
	if len(papers) < 1 && len(papers) > maxNum {
		t.Fatalf("Expected at least one paper, got %d", len(papers))
	}

	// 各論文のTitle、Summary、URLが空でないことを確認
	for _, paper := range papers {
		if paper.Title == "" {
			t.Errorf("Expected non-empty title, got an empty string")
		}
		if paper.Summary == "" {
			t.Errorf("Expected non-empty summary, got an empty string")
		}
		if paper.URL == "" {
			t.Errorf("Expected non-empty URL, got an empty string")
		}
	}
}
