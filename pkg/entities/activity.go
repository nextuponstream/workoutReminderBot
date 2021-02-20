package entities

type Activity struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func CreateActivity(name string) Activity {
	a := Activity{}
	a.Name = name
	return a
}

func (p *Persistence) AddActivityIfNotExists(a Activity) error {
	return p.dp.AddActivityIfNotExists(a)
}

func (p *Persistence) ViewActivity(activityName string) (Activity, error) {
	return p.dp.GetActivity(activityName)
}

func (a *Activity) SetDescription(desc string) {
	a.Description = desc
}
