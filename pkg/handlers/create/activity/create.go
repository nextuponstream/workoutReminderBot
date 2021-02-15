package activity

import (
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	e "github.com/nextuponstream/workoutReminderBot/pkg/entities"
)

func Handler(bot *tgbotapi.BotAPI, userMessage *tgbotapi.Message) {
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
	a := e.Create(activityName)

	if len(tokens) > 2 {
		description := strings.Replace(usrMsg, "/activity"+sep+activityName+sep, "", 1)
		a.Description = description
	}

	err := e.InsertActivity(a)
	if err != nil {
		reply = err.Error()
	} else {
		reply = "Activity " + activityName + " was created"
	}

	msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
	msg.ReplyToMessageID = userMessage.MessageID
	bot.Send(msg)
}
