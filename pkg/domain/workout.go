package domain

type Workout struct {
	Name       string
	Activities []Activity
}

// IsValid returns true if your workout is searchable and has at least one activity
func (w Workout) IsValid() bool {
	return w.Name != "" && len(w.Activities) > 0
}
