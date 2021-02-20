package mongo

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	e "github.com/nextuponstream/workoutReminderBot/pkg/entities"
)

// GetActivities
func (m *Mongo) getActivities() *mongo.Collection {
	return m.database.Collection("activities")
}

// GetActivity from the mongo db activities collection
func (m *Mongo) GetActivity(activityName string) (e.Activity, error) {
	filter := bson.D{{"name", activityName}}

	var activity e.Activity
	err := m.getActivities().FindOne(context.TODO(), filter).Decode(&activity)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return activity, err
		}
		log.Fatal(err)
	}

	return activity, err
}

// ActivityExists returns true if activity exists
func (m *Mongo) activityExists(activityName string) (bool, error) {
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

// AddActivityIfNotExists in mongo db in activities collection
func (m *Mongo) AddActivityIfNotExists(activity e.Activity) error {
	exists, err := m.activityExists(activity.Name)
	if err != nil {
		log.Fatal(err)
	}

	if exists {
		return errors.New("activity already exists")
	}

	_, err = m.getActivities().InsertOne(context.TODO(), activity)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
