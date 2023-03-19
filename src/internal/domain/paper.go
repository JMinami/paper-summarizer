package domain

type (
	Papers []Paper
	Paper  struct {
		Title   string
		Summary string
		URL     string
	}
)
