package application

import (
	"fmt"
	"paper-summarizer/internal/domain"
)

type PaperSummarySendService struct {
	paperService      *PaperService
	paperFormatServie *PaperFormatService
	messageSender     domain.MessageSender
}

func NewPaperSummarySendService(
	paperService *PaperService,
	paperFormatService *PaperFormatService,
	messageSender domain.MessageSender,
) *PaperSummarySendService {
	return &PaperSummarySendService{
		paperService:      paperService,
		paperFormatServie: paperFormatService,
		messageSender:     messageSender,
	}
}

func (s *PaperSummarySendService) Send(
	keyword string,
	maxNum int,
) error {
	translationPapers, err := s.paperService.TranslatePapers(keyword, maxNum)
	if err != nil {
		return fmt.Errorf("translate papers error -> %s", err)
	}

	if err := s.messageSender.Send(fmt.Sprintf("キーワード:[%s]の論文の要約を%d件表示します", keyword, maxNum)); err != nil {
		return fmt.Errorf("send error -> %s", err)
	}

	for _, translatedPaper := range translationPapers {
		message := s.paperFormatServie.FormatPaper(domain.Papers{translatedPaper})
		strs := s.splitStringByRunes(message, 4999)

		for _, str := range strs {
			if err := s.messageSender.Send(str); err != nil {
				return fmt.Errorf("send error2 -> %s", err)
			}
		}
	}

	return nil

}

func (s *PaperSummarySendService) splitStringByRunes(str string, chunkSize int) []string {
	runes := []rune(str)
	chunks := make([]string, 0)

	for i := 0; i < len(runes); i += chunkSize {
		end := i + chunkSize
		if end > len(runes) {
			end = len(runes)
		}
		chunks = append(chunks, string(runes[i:end]))
	}

	return chunks
}
