package mongo

import (
	"context"
	"errors"
	"log"
	"strconv"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/nextuponstream/workoutReminderBot/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetActivity from the mongo db activities collection
func (m *Mongo) GetUser(id string) (domain.User, error) {
	filter := bson.D{{"id", id}}

	var user domain.User
	err := m.getActivities().FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return user, err
		}
		log.Fatal(err)
	}

	return user, err
}

// AddUserIfNotExists with some relevant details
func (m *Mongo) AddUserIfNotExists(user tgbotapi.User) error {
	exists, err := m.userExists(user)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user already exists")
	}

	usr := domain.CreateUser(user)
	_, err = m.getUsers().InsertOne(context.TODO(), usr)

	return err
}

// getUsers collection from which you can insert a user via InsertOne
func (m *Mongo) getUsers() *mongo.Collection {
	return m.database.Collection("users")
}

// userExists
func (m *Mongo) userExists(user tgbotapi.User) (bool, error) {
	_, err := m.GetUser(strconv.Itoa(user.ID))
	isMissing := err == mongo.ErrNoDocuments
	if isMissing {
		return false, nil
	} else if err != nil {
		return false, err
	} else { // found, err == nil
		return true, err
	}
}
