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

	reminders := make(map[int]map[domain.Reminder](chan struct{}))
	// TODO restore reminders on shutdown by parsing the graph database

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
		case "timezone":
			go handler.Timezone(p, bot, userMessage)
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
			cancel, reminder, err := handler.RemindMe(p, bot, userMessage)
			if err == nil {
				if _, ok := reminders[userMessage.From.ID]; !ok {
					reminders[userMessage.From.ID] = make(map[domain.Reminder]chan struct{})
				}
				if _, ok := reminders[userMessage.From.ID][reminder]; !ok {
					reminders[userMessage.From.ID][reminder] = cancel
				}
			}
		case "cancel":
			go handler.Cancel(reminders, p, bot, userMessage)
		default:
			go handler.Unknown(bot, userMessage)
		}
	}
}
