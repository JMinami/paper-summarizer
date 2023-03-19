package domain

type (
	PaperFetcher interface {
		Fetch(RequestFetch) (Papers, error)
	}
	RequestFetch struct {
		Keyword string
		MaxNum  int
	}
)
