package unknown

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func Handler(bot *tgbotapi.BotAPI, userMessage *tgbotapi.Message) {
	reply := "Unknown command"
	msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
	msg.ReplyToMessageID = userMessage.MessageID

	bot.Send(msg)
}
