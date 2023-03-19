package application

import "paper-summarizer/internal/domain"

type PaperFormatService struct {
	formatter domain.PaperFormatter
}

func NewPaperFormatter(formatter domain.PaperFormatter) *PaperFormatService {
	return &PaperFormatService{
		formatter: formatter,
	}
}

func (s *PaperFormatService) FormatPaper(papers domain.Papers) string {
	return s.formatter.Format(papers)
}
