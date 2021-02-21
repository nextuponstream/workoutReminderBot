package domain

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

// InitDatabase will persist all entities and relationships through one unique interface
func InitDatabase(dpersistence DocPersistence, gpersistence GraphPersistence) Persistence {
	return Persistence{dpersistence, gpersistence}
}

type Persistence struct {
	dp DocPersistence
	gp GraphPersistence
}

type PersistenceI interface {
	AddUserIfNotExists(tgbotapi.User) error
	ViewActivity(string) error
	GetExercises(tgbotapi.User) ([]Exercise, error)
}

type GraphPersistence interface {
	AddExerciseIfNotExists(Exercise, tgbotapi.User, Activity) error
	GetExercises(tgbotapi.User) ([]Exercise, error)
}

type DocPersistence interface {
	AddUserIfNotExists(tgbotapi.User) error
	AddActivityIfNotExists(Activity) error
	GetActivity(activityName string) (Activity, error)
}
