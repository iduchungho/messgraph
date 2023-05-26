package message

import "context"

type IRepository interface {
	Save(ctx context.Context, message *Message) (*Message, error)
	GetByUsername(ctx context.Context, username string) ([]Message, error)
}

type IService interface {
	GetByUsername(ctx context.Context, username string) ([]Message, error)
}
