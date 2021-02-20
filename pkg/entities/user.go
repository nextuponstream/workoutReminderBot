package entities

import (
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type User struct {
	ID               int    `json:"id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	UserName         string `json:"username"`
	RegistrationTime string `json:"registration_time"`
}

// CreateUser with registration time set to now
func CreateUser(u tgbotapi.User) User {
	return User{u.ID, u.FirstName, u.LastName, u.UserName, time.Now().String()}
}

// AddUserIfNotExists
func (p *Persistence) AddUserIfNotExists(u tgbotapi.User) error {
	return p.dp.AddUserIfNotExists(u)
}
