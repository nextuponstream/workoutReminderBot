package entities

type Exercise struct {
	Repeat   int     `json:"repeat"`
	Length   float32 `json:"length"`
	Duration int     `json:"duration"`
	Notes    string  `json:"notes"`
}

// InsertExercise save the exercise in the graph database
func InsertExercise(e Exercise) error {
	return gp.InsertExercise(e)
}

func (e *Exercise) SetDuration(d int) {
	e.Duration = d
}

func (e *Exercise) SetRepeat(r int) {
	e.Repeat = r
}

func (e *Exercise) SetNotes(n string) {
	e.Notes = n
}
