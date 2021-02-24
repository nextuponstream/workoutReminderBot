package handler

import (
	"errors"
	"strings"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/nextuponstream/workoutReminderBot/pkg/domain"
)

// TODO cancel map of user reminders

// RemindMe handles all /remindme commands of a telegram user
func RemindMe(p domain.Persistence, bot *tgbotapi.BotAPI, userMessage *tgbotapi.Message) {
	// TODO parse user message to build reminder
	reminder := domain.Reminder{}
	// TODO if ok reminder then create goroutine
	for {
		duration, err := GetRemainingTime(time.Now(), reminder)
		if err != nil {
			reply := "Bad reminder"
			msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
			msg.ReplyToMessageID = userMessage.MessageID
			bot.Send(msg)
			return
		}

		reply := "Reminder was cancelled"
		select {
		case <-time.After(duration):
			reply = RemindMessage(reminder)
			// TODO fetch all exercise once timeout is reached and add to reply
			// TODO case cancel
			// case <-cancel
			// break;
		}

		msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
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
			daysDuration, err := AddDays(todayDay, day)
			if err != nil {
				return 0, err
			}

			return duration + daysDuration, nil
		}
		day = (day + 1) % 7
	}
}

// AddDays returns a time duration for how many remaining days there is to wait.
// Note: returns 0 if from == to
func AddDays(from int, to int) (time.Duration, error) {
	if from > to {
		return 0, errors.New("Invalid arguments: from > to")
	}
	total := time.Duration(0)
	for day := from; day != to; day++ {
		total = total + time.Hour*24
	}

	return total, nil
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
