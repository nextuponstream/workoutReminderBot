package mongo

import (
	"context"
	"log"
	"time"

	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	client   *mongo.Client
	ctx      context.Context
	database *mongo.Database
}

func CreateMongoDb(user string, pw string, name string, uri string) Mongo {
	// mongo db connection
	// ref: https://www.mongodb.com/golang
	authenticationURI := "mongodb://" + user + ":" + pw + "@" + uri
	client, err := mongo.NewClient(options.Client().ApplyURI(authenticationURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(name)

	return Mongo{client, ctx, database}
}

func (m *Mongo) Disconnect() {
	m.client.Disconnect(m.ctx)
}
