package exercise

import (
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

// Handler creates an exercise
func Handler(bot *tgbotapi.BotAPI, userMessage *tgbotapi.Message) {
	var reply string

	// TODO new user handling

	usrMsg := userMessage.Text
	sep := " "

	// HAHA:

	tokens := strings.Split(usrMsg, sep)
	if len(tokens) < 2 {
		reply = "Please provide the activity name to create an exercise for"
		msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
		msg.ReplyToMessageID = userMessage.MessageID
		bot.Send(msg)
		return
	}

	activityName := tokens[1]

	for _, token := range tokens[2:] {
		c := token[0]
		num, err := strconv.Atoi(token[1:])
		if err != nil {
			log.Fatal(err)
		}
		switch c {
		case 'n':
			log.Println(activityName, "has to be done", num, "in a row")
			break
		case 'r':
			log.Println(activityName, "with", num, "repeats")
			break
		case 'l':
			log.Println(activityName, "runs for", num, "kms")
		default:
			reply = "unknown argument"
			msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
			msg.ReplyToMessageID = userMessage.MessageID
			bot.Send(msg)
			return
		}
	}

	msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
	msg.ReplyToMessageID = userMessage.MessageID
	bot.Send(msg)
}
