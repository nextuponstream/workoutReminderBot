package neo4j

import (
	"log"

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
func (n *Neo4j) GetExercises(user tgbotapi.User) ([]domain.Exercise, error) {
	session := n.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	//  https://stackoverflow.com/a/24016201
	query :=
		"MATCH (u:User { tid: $userId })\n" +
			"MATCH (a:Activity)\n" +
			"MATCH (u)-[e:EXERCISE]->(a)\n" +
			"RETURN a.name, e.reps, e.set, e.length, e.duration, e.notes"
	params := map[string]interface{}{
		"userId": user.ID,
	}

	records, err := session.Run(query, params)
	if err != nil {
		return []domain.Exercise{}, err
	}

	exercices := []domain.Exercise{}

	for records.Next() {
		re := records.Record()
		log.Print(re)
		ex := domain.Exercise{}
		if activityName, ok := re.Get("a.name"); ok {
			ex.Activity = activityName.(string)
		}
		if reps, ok := re.Get("e.reps"); ok {
			ex.Reps = int(reps.(int64))
		}
		if set, ok := re.Get("e.set"); ok {
			ex.Set = int(set.(int64))
		}
		if length, ok := re.Get("e.length"); ok {
			ex.Length = float32(length.(float64))
		}
		if duration, ok := re.Get("e.duration"); ok {
			ex.Duration = duration.(string)
		}
		if notes, ok := re.Get("e.notes"); ok {
			ex.Notes = notes.(string)
		}

		exercices = append(exercices, ex)
	}

	return exercices, nil
}
