package handler

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/nextuponstream/workoutReminderBot/pkg/domain"
)

// Routine handles all /routine commands from the telegram user
func Routine(p domain.Persistence, bot *tgbotapi.BotAPI, userMessage *tgbotapi.Message) {
	// TODO implement
	reply := "unimplemented"
	msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
	msg.ReplyToMessageID = userMessage.MessageID
	bot.Send(msg)
}
