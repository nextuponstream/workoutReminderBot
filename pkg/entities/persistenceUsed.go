package entities

var p Persistence
var gp GraphPersistence

func InitDatabase(persistence Persistence, gpersistence GraphPersistence) {
	p = persistence
	gp = gpersistence
}

type Persistence interface {
	InsertActivity(Activity) error
	GetActivity(activityName string) (Activity, error)
}

type GraphPersistence interface {
	InsertExercise(Exercise) error
}
