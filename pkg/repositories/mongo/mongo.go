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

// Create mongo database with user, password, database name and address
func Create(user string, password string, name string, uri string) Mongo {
	// ref: https://www.mongodb.com/golang
	authenticationURI := "mongodb://" + user + ":" + password + "@" + uri
	client, err := mongo.NewClient(options.Client().ApplyURI(authenticationURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(name)

	return Mongo{client, ctx, database}
}

// Disconnect defer ressources liberation of the mongo database
func (m *Mongo) Disconnect() {
	m.client.Disconnect(m.ctx)
}
