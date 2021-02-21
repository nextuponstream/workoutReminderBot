package handler

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

// Unknown handler when unknown command is issued by user
func Unknown(bot *tgbotapi.BotAPI, userMessage *tgbotapi.Message) {
	reply := "Unknown command"
	msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
	msg.ReplyToMessageID = userMessage.MessageID

	bot.Send(msg)
}
