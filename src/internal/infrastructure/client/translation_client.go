package client

import (
	"context"
	"paper-summarizer/src/internal/domain"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

type (
	GCPTranslator struct {
	}
)

var _ domain.Translator = (*GCPTranslator)(nil)

func NewGCPTranslator() *GCPTranslator {
	return &GCPTranslator{}
}

func (t *GCPTranslator) TranslateText(texts []string) ([]string, error) {
	client, err := translate.NewClient(context.Background())
	if err != nil {
		return nil, err
	}
	defer client.Close()

	targetLang := language.Japanese

	resp, err := client.Translate(context.Background(), texts, targetLang, nil)
	if err != nil {
		return nil, err
	}

	translatedTexts := make([]string, len(resp))
	for i, translated := range resp {
		translatedTexts[i] = translated.Text
	}
	return translatedTexts, nil
}
