package exercise

import (
	"errors"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/nextuponstream/workoutReminderBot/pkg/entities"
)

// Handler creates an exercise
func Handler(p entities.Persistence, bot *tgbotapi.BotAPI, userMessage *tgbotapi.Message) {
	err := p.AddUserIfNotExists(*userMessage.From)
	if err != nil {
		log.Fatal(err)
	}

	var reply string

	usrMsg := userMessage.Text
	sep := " "

	activityName := strings.TrimPrefix(usrMsg, "/activity ")
	activityName = strings.ReplaceAll(activityName, "\n", sep)
	tokens := strings.Split(activityName, sep)
	if len(tokens) < 2 {
		reply = "Please provide the activity name to create an exercise for"
		msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
		msg.ReplyToMessageID = userMessage.MessageID
		bot.Send(msg)
		return
	}

	ex, err := GetExercice(usrMsg)
	if err != nil {
		reply = "Activity details could not be parsed"
		msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
		msg.ReplyToMessageID = userMessage.MessageID
		bot.Send(msg)
		return
	}

	activityName = tokens[1]
	a := entities.CreateActivity(activityName)
	err = p.AddExerciseIfNotExists(ex, *userMessage.From, a)
	if err != nil {
		reply = "an error occured while creating exercise"
		msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
		msg.ReplyToMessageID = userMessage.MessageID
		bot.Send(msg)
		return
	}

	reply = "successfully inserted exercise for " + activityName
	msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
	msg.ReplyToMessageID = userMessage.MessageID
	bot.Send(msg)
}

// getExercice from user message
func GetExercice(usrMsg string) (entities.Exercise, error) {
	ex := entities.Exercise{}
	sep := " "

	lines := strings.Split(usrMsg, "\n")
	for _, line := range lines[1:] {
		tokens := strings.TrimPrefix(line, sep)
		c := tokens[0]
		arg := strings.SplitAfterN(line, sep, 2)[1]
		switch c {
		case 'r':
			num, err := strconv.ParseFloat(arg, 32)
			if err != nil {
				return ex, err
			}
			ex.Reps = int(num)
			break
		case 's':
			num, err := strconv.ParseFloat(arg, 32)
			if err != nil {
				return ex, err
			}
			ex.Set = int(num)
			break
		case 'l':
			num, err := strconv.ParseFloat(arg, 32)
			if err != nil {
				return ex, err
			}
			ex.Length = float32(num)
		case 'd':
			ex.Duration = arg
		case 'n':
			ex.Notes = arg
		default:
			return ex, errors.New("bad argument")
		}
	}

	return ex, nil
}
