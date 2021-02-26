package handler

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/nextuponstream/workoutReminderBot/pkg/domain"
)

// TODO cancel map of user reminders

// RemindMe handles all /remindme commands of a telegram user
func RemindMe(p domain.Persistence, bot *tgbotapi.BotAPI, userMessage *tgbotapi.Message) (chan struct{}, domain.Reminder, error) {
	reminder := domain.Reminder{}
	msg := userMessage.Text
	tokens := strings.Split(msg, " ")
	if len(tokens) < 5 {
		reply := "Not enough arguments in reminder"
		msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
		msg.ReplyToMessageID = userMessage.MessageID
		bot.Send(msg)
		return nil, reminder, errors.New("Not enough arguments")
	}
	reminder.RoutineName = tokens[1]
	from, err := strconv.Atoi(tokens[2])
	if err != nil {
		reply := "Cannot read `from`"
		msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
		msg.ReplyToMessageID = userMessage.MessageID
		bot.Send(msg)
		return nil, reminder, errors.New(reply)
	}
	reminder.From = from
	to, err := strconv.Atoi(tokens[3])
	if err != nil {
		reply := "Cannot read `to`"
		msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
		msg.ReplyToMessageID = userMessage.MessageID
		bot.Send(msg)
		return nil, reminder, errors.New(reply)
	}
	reminder.To = to

	reminder.When = ParseWeek(tokens[4:])
	if !reminder.IsValid() {
		reply := "Could not parse reminder, please set your reminder between 6am-21pm and use at least one of: mo/tu/we/th/fr/sa/su"
		msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
		msg.ReplyToMessageID = userMessage.MessageID
		bot.Send(msg)
		return nil, reminder, errors.New("Invalid weekdays for reminder")
	}

	cancel := make(chan struct{})
	go remindUser(reminder, cancel, p, bot, userMessage)

	return cancel, reminder, nil
}

// ParseWeek set week with arguments
func ParseWeek(args []string) domain.Week {
	week := domain.Week{}
	weekdays := make(map[string]int)
	weekdays["mo"] = 0
	weekdays["tu"] = 1
	weekdays["we"] = 2
	weekdays["th"] = 3
	weekdays["fr"] = 4
	weekdays["sa"] = 5
	weekdays["su"] = 6
	for _, day := range args {
		if index, ok := weekdays[day]; ok {
			week.Week[index] = true
		}
	}
	return week
}

// remindUser waits until timeout of reminder to remind user of his routine. Can be cancelled with cancel channel.
func remindUser(reminder domain.Reminder, cancel chan struct{}, p domain.Persistence, bot *tgbotapi.BotAPI, userMessage *tgbotapi.Message) {
	for {
		user, err := p.GetUser(strconv.Itoa(userMessage.From.ID))
		if err != nil {
			log.Print(err)
			reply := "Couldn't get your timezone from the database"
			msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
			msg.ReplyToMessageID = userMessage.MessageID
			bot.Send(msg)
			break
		}
		location, err := time.LoadLocation(user.Timezone)
		if err != nil {
			reply := "Couldn't parse your timezone"
			msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
			msg.ReplyToMessageID = userMessage.MessageID
			bot.Send(msg)
			break
		}

		now := time.Now().In(location) // time in user location
		duration, err := GetRemainingTime(now, reminder)
		if err != nil {
			log.Println(err)
			reply := "Bad reminder"
			msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
			msg.ReplyToMessageID = userMessage.MessageID
			bot.Send(msg)
			break
		}

		remainingMinutes := int(duration.Minutes()) % 60
		reply := fmt.Sprintf("Your reminder was set and the next one will come in %dh%dm",
			int(duration.Hours()), remainingMinutes)
		msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
		msg.ReplyToMessageID = userMessage.MessageID
		bot.Send(msg)

		select {
		case <-time.After(duration):
			reply = RemindMessage(reminder)
		case <-cancel:
			reply = "Reminder was cancelled"
			// TODO fetch all exercise once timeout is reached
			break
		}

		msg = tgbotapi.NewMessage(userMessage.Chat.ID, reply)
		msg.ReplyToMessageID = userMessage.MessageID
		bot.Send(msg)
	}
}

// GetRemainingTime returns a duration until next programmed reminder and any encountered error
func GetRemainingTime(now time.Time, reminder domain.Reminder) (time.Duration, error) {
	todayDay := int((now.Weekday() - 1) % 7) // from monday = 0
	for todayDay < 0 {
		todayDay += 7
	}
	duration := TimeUntil(reminder.From, now)
	day := todayDay
	if duration > time.Hour*24 {
		day = (todayDay + 1) % 7
	}

	for {
		if reminder.When.Week[day] { // found
			return duration + AddDays(todayDay, day), nil
		}
		day = (day + 1) % 7
	}
}

// AddDays returns a time duration for how many remaining days there is to wait.
// Note: returns 0 if from == to
func AddDays(from int, to int) time.Duration {
	total := time.Duration(0)
	for day := from; day != to; day = (day + 1) % 7 {
		total = total + time.Hour*24
	}

	return total
}

// TimeUntil returns the duration until some hour on the same day
func TimeUntil(hour int, now time.Time) time.Duration {
	reminder := time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, now.Location())
	if !now.Before(reminder) && now != reminder {
		reminder = reminder.Add(time.Hour * 24 * 7)
		return reminder.Sub(now)
	}
	return reminder.Sub(now)
}

// NextDayToRemind returns how many days remains until the next reminder
func RemainingDayBeforeReminder(today int, reminder domain.Reminder) int {
	remindDay := today
	isNextDay := false
	for i := 0; i < 14; i++ {
		nextDay := today + i + 1
		isNextDay = reminder.When.Week[nextDay%7]
		if isNextDay {
			remindDay = nextDay
			return remindDay - today
		}
	}
	return -1
}

// RemindMessage displays useful information to the user
func RemindMessage(reminder domain.Reminder) string {
	days := ""
	// Note: for me a week starts on monday and not sunday
	w := reminder.When
	for day := time.Monday; day <= time.Saturday; day++ {
		if w.Week[day-1] {
			days = days + day.String() + ", "
		}
	}
	if w.Week[6] {
		days = days + time.Sunday.String() + ", "
	}
	days = strings.TrimSuffix(days, ", ")
	return "You will be reminded on: " + days
}
