package domain

import (
	"strconv"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID               primitive.ObjectID `bson:"_id"`
	TelegramId       string             `json:"telegram_id"`
	FirstName        string             `json:"first_name"`
	LastName         string             `json:"last_name"`
	UserName         string             `json:"username"`
	RegistrationTime string             `json:"registration_time"`
	Timezone         string             `json:"timezone"`
}

// CreateUser with registration time set to now
func CreateUser(u tgbotapi.User, timezone string) (User, error) {
	_, err := time.LoadLocation(timezone)
	if err != nil {
		return User{}, err
	}
	user := User{
		TelegramId:       strconv.Itoa(u.ID),
		FirstName:        u.FirstName,
		LastName:         u.LastName,
		UserName:         u.UserName,
		RegistrationTime: time.Now().String(),
		Timezone:         timezone,
	}
	return user, nil
}

// UpsertUser
func (p *Persistence) UpsertUser(u User) error {
	return p.dp.UpsertUser(u)
}

// GetUser
func (p *Persistence) GetUser(id string) (User, error) {
	return p.dp.GetUser(id)
}
