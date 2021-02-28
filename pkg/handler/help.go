package handler

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

// Handler reply with advice on bot usage
func Help(bot *tgbotapi.BotAPI, userMessage *tgbotapi.Message) {
	reply := "help - instructions on bot usage\n" +
		"exercise - create exercise for activity and optionnaly indicate its reps, sets, length, " +
		"duration and notes (e.g. /exercise <activity name> [<r/s/l/d/n value>])\n" +
		"activity - create an activity for your next workout with an optionnal description " +
		"(e.g. /activity push-ups let's f*cking goooooo!)\n" +
		"viewactivity - view an activity description" +
		"workout - create a workout made of activities (e.g. /workout gitStrong pushups run)" +
		"routine - create a routine made of workouts (e.g. /routine gitStrong)" +
		"remindme - sends a reminder to maintain the routine (e.g. /remindme gitStrong 16 21 tu we th fr sa)" +
		"cancel - cancel all reminders or for all mentionned routines (e.g. /cancel OR /cancel gitStrong)"

	msg := tgbotapi.NewMessage(userMessage.Chat.ID, reply)
	msg.ReplyToMessageID = userMessage.MessageID

	bot.Send(msg)
}
