package domain

type (
	Papers []Paper
	Paper  struct {
		Title   string
		Summary string
		URL     string
	}
	PaperFetcher interface {
		Fetch(RequestFetch) (Papers, error)
	}
	RequestFetch struct {
		Keyword string
		MaxNum  int
	}
)
