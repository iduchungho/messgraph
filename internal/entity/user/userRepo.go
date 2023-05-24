package user

import "go.mongodb.org/mongo-driver/mongo"

type UserRepo struct {
	Conn *mongo.Client
}

func NewUserRepo(conn *mongo.Client) *UserRepo{
	return &UserRepo{
		Conn: conn,
	}
}