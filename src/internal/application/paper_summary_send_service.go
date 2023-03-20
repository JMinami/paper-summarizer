package application

import "paper-summarizer/internal/domain"

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
		return err
	}

	message := s.paperFormatServie.FormatPaper(translationPapers)

	if err := s.messageSender.Send(message); err != nil {
		return err
	}
	return nil
}
