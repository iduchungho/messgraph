package message

import "time"

type Service struct {
	Repository *Repository
	timeout    time.Duration
}

func NewMessageService(messageRepo *Repository) *Service {
	return &Service{
		Repository: messageRepo,
		timeout:    time.Duration(2) * time.Second,
	}
}
