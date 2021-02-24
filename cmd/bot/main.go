package main

import (
	"log"
	"os"
	"strings"

	"github.com/nextuponstream/workoutReminderBot/pkg/domain"
	"github.com/nextuponstream/workoutReminderBot/pkg/handler"
	"github.com/nextuponstream/workoutReminderBot/pkg/repositories/mongo"
	"github.com/nextuponstream/workoutReminderBot/pkg/repositories/neo4j"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func main() {
	mdbUser := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	mdbPw := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	mdbName := os.Getenv("MONGO_INITDB_DATABASE")
	mdbUri := os.Getenv("MDB_URI")

	mongoDb := mongo.Create(mdbUser, mdbPw, mdbName, mdbUri)
	defer mongoDb.Disconnect()

	user := os.Getenv("NEO4J_USER")
	pw := os.Getenv("NEO4J_AUTH")
	ndbUri := os.Getenv("GDB_URI")
	pw = strings.TrimPrefix(pw, "neo4j/")
	neo4jGdp := neo4j.Create(user, pw, ndbUri)
	defer neo4jGdp.Close()

	// context
	p := domain.InitDatabase(&mongoDb, &neo4jGdp)

	// telegram connection
	botToken := os.Getenv("BOT_TOKEN")
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
			go handler.Help(bot, userMessage)
		case "activity":
			go handler.Activity(p, bot, userMessage)
		case "viewactivity":
			go handler.ActivityView(p, bot, userMessage)
		case "exercise":
			go handler.Exercise(p, bot, userMessage)
		case "viewexercises":
			go handler.ExercisesView(p, bot, userMessage)
		case "workout":
			go handler.Workout(p, bot, userMessage)
		case "routine":
			go handler.Routine(p, bot, userMessage)
		case "remindme":
			go handler.RemindMe(p, bot, userMessage)
		default:
			go handler.Unknown(bot, userMessage)
		}
	}
}
