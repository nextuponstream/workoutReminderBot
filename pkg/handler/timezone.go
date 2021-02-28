package handler

import (
	"strings"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/nextuponstream/workoutReminderBot/pkg/domain"
)

// Timezone handles /timezone command to set the user timezone
func Timezone(p domain.Persistence, bot *tgbotapi.BotAPI, userMessage *tgbotapi.Message) {
	tokens := strings.Split(userMessage.Text, " ")
	if len(tokens) < 2 {
		// TODO retrieve user timezone if stored
		reply := "not implemented"
		msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
		msg.ReplyToMessageID = userMessage.MessageID
		bot.Send(msg)
		return
	}

	tempUser, err := domain.CreateUser(*userMessage.From, "")
	if err != nil {
		reply := "Something went wrong"
		msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
		msg.ReplyToMessageID = userMessage.MessageID
		bot.Send(msg)
		return
	}

	userInput := tokens[1]
	timezone, err := time.LoadLocation(userInput)
	if err != nil {
		reply := "Error while parsing timezone"
		msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
		msg.ReplyToMessageID = userMessage.MessageID
		bot.Send(msg)
		return
	}

	tempUser.Timezone = userInput
	err = p.UpsertUser(tempUser)
	if err != nil {
		reply := "add Something went wrong"
		msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
		msg.ReplyToMessageID = userMessage.MessageID
		bot.Send(msg)
		return
	}

	reply := "Your timezone is set to " + timezone.String()
	msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
	msg.ReplyToMessageID = userMessage.MessageID
	bot.Send(msg)
}
