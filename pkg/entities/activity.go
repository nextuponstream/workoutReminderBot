package entities

type Activity struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Duration    int    `json:"duration,omitempty"`
	Repeat      int    `json:"repeat,omitempty"`
	Notes       string `json:"notes,omitempty"`
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

func (a *Activity) SetDuration(d int) {
	a.Duration = d
}

func (a *Activity) SetRepeat(r int) {
	a.Repeat = r
}

func (a *Activity) SetNotes(n string) {
	a.Notes = n
}
