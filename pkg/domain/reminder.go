package domain

const MIN_START = 6
const MAX_START = 21
const MIN_PERIOD_REMINDER = 1

type Week struct {
	Monday    bool
	Tuesday   bool
	Wednesday bool
	Thursday  bool
	Friday    bool
	Saturday  bool
	Sunday    bool
}

type Reminder struct {
	When Week
	From int // 0-23 hours
	To   int // 0-23 hours
}

// AtLeastOneDay returns true if there's one day of the week that you want to be reminded of your
// workout
func (w Week) AtLeastOneDay() bool {
	return w.Monday || w.Tuesday || w.Wednesday || w.Thursday || w.Friday || w.Saturday || w.Sunday
}

// IsValid returns true if you are a normal folk that trains in reasonable hours
// such as 6am to 21pm. A reminder is valid if it's at least once a week.
// It is true that your workout could start at 1am (or be less than once a week)
// but this is a programming hassle I am not willing to deal with
func (r Reminder) IsValid() bool {
	if r.From < MIN_START || MAX_START < r.From {
		return false
	}
	if r.To < MIN_START+MIN_PERIOD_REMINDER || MAX_START+MIN_PERIOD_REMINDER < r.To {
		return false
	}
	return r.From < r.To && r.When.AtLeastOneDay()
}
