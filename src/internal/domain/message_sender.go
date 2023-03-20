package domain

type (
	MessageSender interface {
		Send(message string) error
	}
)
