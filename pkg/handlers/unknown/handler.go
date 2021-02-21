package unknown

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

// Handler basic reply when unknown command is issued by user
func Handler(bot *tgbotapi.BotAPI, userMessage *tgbotapi.Message) {
	reply := "Unknown command"
	msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
	msg.ReplyToMessageID = userMessage.MessageID

	bot.Send(msg)
}
