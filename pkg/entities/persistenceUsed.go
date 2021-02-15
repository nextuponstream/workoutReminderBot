package entities

var p Persistence

func InitDatabase(persistence Persistence) {
	p = persistence
}

type Persistence interface {
	InsertActivity(Activity) error
	GetActivity(activityName string) (Activity, error)
}
