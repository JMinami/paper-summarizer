package domain

type (
	Translator interface {
		TranslateText(texts []string) ([]string, error)
	}
)
