package handler

import (
	"log"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/nextuponstream/workoutReminderBot/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

// Activity persist an activity created by the telegram user
func Activity(p domain.Persistence, bot *tgbotapi.BotAPI, userMessage *tgbotapi.Message) {
	err := p.AddUserIfNotExists(*userMessage.From)
	if err != nil {
		log.Fatal(err)
	}

	var reply string

	usrMsg := userMessage.Text
	sep := " "

	tokens := strings.Split(usrMsg, sep)
	if len(tokens) < 2 {
		reply = "Please provide a name to the activity"
		msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
		msg.ReplyToMessageID = userMessage.MessageID
		bot.Send(msg)
		return
	}

	activityName := tokens[1]
	a := domain.CreateActivity(activityName)

	if len(tokens) > 2 {
		description := strings.Replace(usrMsg, "/activity"+sep+activityName+sep, "", 1)
		a.Description = description
	}

	err = p.AddActivityIfNotExists(a)
	if err != nil {
		reply = err.Error()
	} else {
		reply = "Activity " + activityName + " was created"
	}

	msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
	msg.ReplyToMessageID = userMessage.MessageID
	bot.Send(msg)
}

// ActivityView responds to user requesting to view an activity description
func ActivityView(p domain.Persistence, bot *tgbotapi.BotAPI, userMessage *tgbotapi.Message) {
	var reply string

	sep := " "

	tokens := strings.Split(userMessage.Text, sep)
	if len(tokens) < 2 {
		reply = "Please provide a name to the activity you want to view"
		msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
		msg.ReplyToMessageID = userMessage.MessageID
		bot.Send(msg)
		return
	}

	activityName := tokens[1]
	activity, err := p.ViewActivity(activityName)
	if err == mongo.ErrNoDocuments {
		reply = "This activity doesn't exist"
	} else if err != nil {
		reply = "Something went wrong"
	} else {
		reply = "Description: " + activity.Description
	}

	msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
	msg.ReplyToMessageID = userMessage.MessageID
	bot.Send(msg)
}
