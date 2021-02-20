package activity

import (
	"strings"

	"go.mongodb.org/mongo-driver/mongo"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	e "github.com/nextuponstream/workoutReminderBot/pkg/entities"
)

func HandlerView(p e.Persistence, bot *tgbotapi.BotAPI, userMessage *tgbotapi.Message) {
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
		reply = err.Error()
	} else {
		reply = "Description: " + activity.Description
	}

	msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
	msg.ReplyToMessageID = userMessage.MessageID
	bot.Send(msg)
}