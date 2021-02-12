package mongo

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	e "github.com/nextuponstream/workoutReminderBot/entities"
)

// from https://www.mongodb.com/golang

func (m *Mongo) GetActivity(activityName string) (e.Activity, error) {
	collection := m.database.Collection("activities")
	filter := bson.D{{"name", activityName}}

	var activity e.Activity
	err := collection.FindOne(context.TODO(), filter).Decode(&activity)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return activity, err
		}
		log.Fatal(err)
	}

	return activity, err
}

func (m *Mongo) ActivityExists(activityName string) (bool, error) {
	_, err := m.GetActivity(activityName)
	isMissing := err == mongo.ErrNoDocuments
	if isMissing {
		return false, nil
	} else if err != nil {
		return false, err
	} else { // found, err == nil
		return true, err
	}
}

func (m *Mongo) InsertActivity(activity e.Activity) error {
	exists, err := m.ActivityExists(activity.Name)
	if err != nil {
		log.Fatal(err)
	}

	if exists {
		return errors.New("activity already exists")
	}

	collection := m.database.Collection("activities")
	_, err = collection.InsertOne(context.TODO(), activity)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
