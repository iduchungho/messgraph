package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var root *mongo.Client
var lock = &sync.Mutex{}

type Database struct {
	conn *mongo.Client
}

func getURI() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	cluster := os.Getenv("DB_CLUSTER")

	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", user, password, cluster)
	return uri
}

func newConnectMongo() *mongo.Client {
	// Set Client Options
	uri := getURI()
	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func newConn() *mongo.Client {
	if root == nil {
		lock.Lock()
		defer lock.Unlock()
		if root == nil {
			root = newConnectMongo()
			fmt.Println("mongodb connection successfully")
		} else {
			return root
		}
	}
	return root
}

func CreateConnectDatabase() *Database {
	return &Database{
		conn: root,
	}
}

func (db *Database) GetConnect() *mongo.Client {
	return db.conn
}

func (db *Database) CloseConnect(context context.Context) {
	db.conn.Disconnect(context)
}