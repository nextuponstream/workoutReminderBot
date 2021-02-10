package activity

type Activity struct {
	Name        string
	Description string
	Duration    int
	Repeat      int
	Notes       string
}

var activities = []Activity{} // TODO persistence with database

func Create(name string) Activity {
	a := Activity{}
	a.Name = name
	return a
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
