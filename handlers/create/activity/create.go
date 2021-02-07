package activity

import (
	"log"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func Handler(bot *tgbotapi.BotAPI, userMessage *tgbotapi.Message) {
	reply := "What do you need?"

	msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
	msg.ReplyToMessageID = userMessage.MessageID

	bot.Send(msg)

	log.Println("Command:", "help")
	log.Println("Reply:", reply)
}
