package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository struct {
	Conn *mongo.Client
}

func NewUserRepo(conn *mongo.Client) IRepository {
	return &repository{
		Conn: conn,
	}
}

func (repo *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	collection := repo.Conn.Database("MessengerDB").Collection("users")
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1, "id": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return nil, err
	}
	_, err = collection.InsertOne(ctx, *user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (repo *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	collection := repo.Conn.Database("MessengerDB").Collection("users")
	var userRes = new(User)
	err := collection.FindOne(ctx, bson.D{{"email", email}}).Decode(userRes)
	if err != nil {
		return nil, err
	}
	return userRes, nil
}
