package activity

import (
	"log"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	d "github.com/nextuponstream/workoutReminderBot/pkg/domain"
)

func Handler(p d.Persistence, bot *tgbotapi.BotAPI, userMessage *tgbotapi.Message) {
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
	a := d.CreateActivity(activityName)

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
