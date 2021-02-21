package domain

import tgbotapi "github.com/Syfaro/telegram-bot-api"

type Exercise struct {
	Reps     int     `json:"reps"`
	Set      int     `json:"set"`
	Length   float32 `json:"length"`
	Duration string  `json:"duration"`
	Notes    string  `json:"notes"`
}

// AddExerciseIfNotExists
func (p *Persistence) AddExerciseIfNotExists(e Exercise, bot tgbotapi.User, a Activity) error {
	return p.gp.AddExerciseIfNotExists(e, bot, a)
}
