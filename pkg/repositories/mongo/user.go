package mongo

import (
	"context"
	"log"

	"github.com/nextuponstream/workoutReminderBot/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetActivity from the mongo db activities collection
func (m *Mongo) GetUser(telegramId string) (domain.User, error) {
	//filter := bson.D{{"timezone", "CET"}} //why did it work
	filter := bson.D{{"telegram_id", telegramId}}

	var user domain.User
	err := m.getUsers().FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return user, err
		}
		log.Fatal(err)
	}

	return user, nil
}

// UpsertUser with some relevant details
func (m *Mongo) UpsertUser(user domain.User) error {
	opts := options.Replace().SetUpsert(true)
	filter := bson.D{{"telegram_id", user.TelegramId}}
	replacement := bson.D{
		{"telegram_id", user.TelegramId},
		{"first_name", user.FirstName},
		{"last_name", user.LastName},
		{"username", user.UserName},
		{"registration_time", user.RegistrationTime},
		{"timezone", user.Timezone}}
	_, err := m.getUsers().ReplaceOne(context.TODO(), filter, replacement, opts)

	return err
}

// getUsers collection from which you can insert a user via InsertOne
func (m *Mongo) getUsers() *mongo.Collection {
	return m.database.Collection("users")
}
