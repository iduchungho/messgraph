package message

import "time"

type MessageService struct {
	Repository *MessageRepo
	timeout     time.Duration
}

func NewMessageService(messageRepo *MessageRepo) *MessageService {
	return &MessageService{
		Repository: messageRepo,
		timeout: time.Duration(2) * time.Second,
	}
}