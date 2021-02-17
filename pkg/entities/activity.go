package entities

type Activity struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func Create(name string) Activity {
	a := Activity{}
	a.Name = name
	return a
}

func InsertActivity(activity Activity) error {
	return p.InsertActivity(activity)
}

func ViewActivity(activityName string) (Activity, error) {
	return p.GetActivity(activityName)
}

func (a *Activity) SetDescription(desc string) {
	a.Description = desc
}
