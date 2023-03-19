package domain

type (
	PaperFormatter interface {
		Format(papers Papers) string
	}
)
