package application

import "paper-summarizer/internal/domain"

type PaperService struct {
	translator domain.Translator
	fetcher    domain.PaperFetcher
}

func NewPaperService(translator domain.Translator, fetcher domain.PaperFetcher) *PaperService {
	return &PaperService{
		translator: translator,
		fetcher:    fetcher,
	}
}

func (s *PaperService) TranslatePapers(keyword string, maxNum int) (domain.Papers, error) {
	papers, err := s.fetcher.Fetch(domain.RequestFetch{
		Keyword: keyword,
		MaxNum:  maxNum,
	})
	if err != nil {
		return nil, err
	}
	var res domain.Papers
	for _, paper := range papers {
		texts := []string{
			paper.Title, paper.Summary,
		}
		translatedTexts, err := s.translator.TranslateText(texts)
		if err != nil {
			continue
		}
		res = append(res, domain.Paper{
			Title:   translatedTexts[0],
			Summary: translatedTexts[1],
			URL:     paper.URL,
		})
	}
	return res, nil
}
