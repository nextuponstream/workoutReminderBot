package handler

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/nextuponstream/workoutReminderBot/pkg/domain"
)

// Cancel handles /cancel [<routine>] commands to cancel routine by name
func Cancel(reminders map[int]map[domain.Reminder](chan struct{}), p domain.Persistence, bot *tgbotapi.BotAPI, userMessage *tgbotapi.Message) {
	// TODO parse routine name if any

	// cancel all routine from user
	if _, ok := reminders[userMessage.From.ID]; ok {
		for _, cancel := range reminders[userMessage.From.ID] {
			cancel <- struct{}{}
		}
	}

	reply := "All routines were cancelled"
	msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
	msg.ReplyToMessageID = userMessage.MessageID
	bot.Send(msg)
}
