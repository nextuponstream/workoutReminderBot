package main

import (
	"log"
	"os"

	help "github.com/nextuponstream/workoutReminderBot/handlers/help"
	unknown "github.com/nextuponstream/workoutReminderBot/handlers/unknown"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/joho/godotenv"
)

const BOT_TOKEN_STRING = "BOT_TOKEN"

func main() {
	godotenv.Load(".env")
	botToken := os.Getenv(BOT_TOKEN_STRING)
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		userMessage := update.Message
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			go help.Handler(bot, userMessage)
		default:
			go unknown.Handler(bot, userMessage)
		}
	}
}
