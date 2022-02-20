package mongo

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	d "github.com/nextuponstream/workoutReminderBot/pkg/domain"
)

// GetActivity retrieves one activity by name from the mongo database and any errors encountered
func (m *Mongo) GetActivity(activityName string) (d.Activity, error) {
	filter := bson.D{{"name", activityName}}

	var activity d.Activity
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

// ActivityExists returns true if activity exists in mongo database and any error encountered
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

// AddActivityIfNotExists adds activity in mongo database and returns any error encountered
func (m *Mongo) AddActivityIfNotExists(activity d.Activity) error {
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

// getActivities
func (m *Mongo) getActivities() *mongo.Collection {
	return m.database.Collection("activities")
}
