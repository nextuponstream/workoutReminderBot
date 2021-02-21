package main

import (
	"log"
	"os"
	"strings"

	domain "github.com/nextuponstream/workoutReminderBot/pkg/domain"
	activity "github.com/nextuponstream/workoutReminderBot/pkg/handlers/activity"
	exercise "github.com/nextuponstream/workoutReminderBot/pkg/handlers/exercise"
	help "github.com/nextuponstream/workoutReminderBot/pkg/handlers/help"
	unknown "github.com/nextuponstream/workoutReminderBot/pkg/handlers/unknown"
	mongo "github.com/nextuponstream/workoutReminderBot/pkg/repositories/mongo"
	neo4j "github.com/nextuponstream/workoutReminderBot/pkg/repositories/neo4j"

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
			go help.Handler(bot, userMessage)
		case "activity":
			go activity.Handler(p, bot, userMessage)
		case "viewactivity":
			go activity.HandlerView(p, bot, userMessage)
		case "exercise":
			go exercise.Handler(p, bot, userMessage)
		default:
			go unknown.Handler(bot, userMessage)
		}
	}
}
