package domain

import tgbotapi "github.com/Syfaro/telegram-bot-api"

type Exercise struct {
	Reps     int     `json:"reps"`
	Set      int     `json:"set"`
	Length   float32 `json:"length"`   // in km
	Duration string  `json:"duration"` // dd:hh:mm:ss
	Notes    string  `json:"notes"`
	Activity string  `json:"activity,omitempty"`
}

// AddExerciseIfNotExists adds an exercise to the graph database and returns any error encountered
func (p *Persistence) AddExerciseIfNotExists(e Exercise, bot tgbotapi.User, a Activity) error {
	return p.gp.AddExerciseIfNotExists(e, bot, a)
}

// GetExercises retrieves exercises from the graph database and returns any error encountered
func (p *Persistence) GetExercises(u tgbotapi.User) ([]Exercise, error) {
	return p.gp.GetExercises(u)
}
