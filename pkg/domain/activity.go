package domain

type Activity struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// CreateActivity is a constructor for the activity entity
func CreateActivity(name string) Activity {
	a := Activity{}
	a.Name = name
	return a
}

// AddActivityIfNotExists adds an activity to the database and returns any encountered errors
func (p *Persistence) AddActivityIfNotExists(a Activity) error {
	return p.dp.AddActivityIfNotExists(a)
}

// ViewActivity search an activity in the database by name and returns any encountered errors
func (p *Persistence) ViewActivity(name string) (Activity, error) {
	return p.dp.GetActivity(name)
}
