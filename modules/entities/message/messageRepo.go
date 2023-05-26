package message

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	Conn *mongo.Client
}

func NewMessageRepo(conn *mongo.Client) *Repository {
	return &Repository{
		Conn: conn,
	}
}
