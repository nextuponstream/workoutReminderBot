package main

import (
	"log"
	"os"
	"strings"

	entities "github.com/nextuponstream/workoutReminderBot/pkg/entities"
	activity "github.com/nextuponstream/workoutReminderBot/pkg/handlers/activity"
	exercise "github.com/nextuponstream/workoutReminderBot/pkg/handlers/exercise"
	help "github.com/nextuponstream/workoutReminderBot/pkg/handlers/help"
	unknown "github.com/nextuponstream/workoutReminderBot/pkg/handlers/unknown"
	mongo "github.com/nextuponstream/workoutReminderBot/pkg/repositories/mongo"
	neo4j "github.com/nextuponstream/workoutReminderBot/pkg/repositories/neo4j"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const BOT_TOKEN_STR = "BOT_TOKEN"
const MDB_U_STR = "MONGO_INITDB_ROOT_USERNAME"
const MDB_PW_STR = "MONGO_INITDB_ROOT_PASSWORD"
const MDB_DB_NAME_STR = "MONGO_INITDB_DATABASE"
const MDB_URI_STR = "MDB_URI"
const NEO4J_AUTH_STR = "NEO4J_AUTH"

func main() {
	mdbUser := os.Getenv(MDB_U_STR)
	mdbPw := os.Getenv(MDB_PW_STR)
	mdbName := os.Getenv(MDB_DB_NAME_STR)
	mdbUri := os.Getenv(MDB_URI_STR)

	mongoDb := mongo.CreateMongoDb(mdbUser, mdbPw, mdbName, mdbUri)
	defer mongoDb.Disconnect()

	// TODO env
	user := "neo4j"
	pw := os.Getenv(NEO4J_AUTH_STR)
	pw = strings.TrimPrefix(pw, "neo4j/")
	neo4jGdp := neo4j.Create(user, pw)
	defer neo4jGdp.Close()

	// context
	p := entities.InitDatabase(&mongoDb, &neo4jGdp)

	// telegram connection
	botToken := os.Getenv(BOT_TOKEN_STR)
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
