package neo4j

import (
	"errors"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/nextuponstream/workoutReminderBot/pkg/domain"
)

func (n *Neo4j) AddExerciseIfNotExists(e domain.Exercise, user tgbotapi.User, a domain.Activity) error {
	session := n.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	//  https://stackoverflow.com/a/24016201
	queryExercise :=
		"MERGE (u:User { tid: $userId })\n" +
			"MERGE (a:Activity { name: $activityName })\n" +
			"MERGE (u)-[:EXERCISE { " +
			"reps: $reps, " +
			"set: $set, " +
			"length: $length, " +
			"duration: $duration, " +
			"notes: $notes " +
			"}]->(a)"
	queries := []string{queryExercise}
	params := map[string]interface{}{
		"userId":       user.ID,
		"activityName": a.Name,
		"reps":         e.Reps,
		"set":          e.Set,
		"length":       e.Length,
		"duration":     e.Duration,
		"notes":        e.Notes,
	}

	for _, query := range queries {
		_, err := session.Run(query, params)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetExercises retrieves all exercises created by user
func GetExercises(user tgbotapi.User) ([]domain.Exercise, error) {
	// TODO get all exercises from user
	/*
	   re := record.(*db.Record)
	   	if id, ok := re.Get("n.id"); ok {
	   		item.Id = id.(int64)
	   	} else {
	   		return item, errors.New("could not find id field")
	   	}
	   	if name, ok := re.Get("n.name"); ok {
	   		item.Name = name.(string)
	   	} else {
	   		return item, errors.New("could not find name field")
	   	}

	*/
	return []domain.Exercise{}, errors.New("not implemented")
}
