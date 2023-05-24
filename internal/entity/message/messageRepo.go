package message

import "go.mongodb.org/mongo-driver/mongo"

type MessageRepo struct {
	Conn *mongo.Client
}

func NewMessageRepo(conn *mongo.Client) *MessageRepo {
	return &MessageRepo{
		Conn: conn,
	}
}