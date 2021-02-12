module github.com/nextuponstream/workoutReminderBot

go 1.15

replace (
	github.com/nextuponstream/workoutReminderBot/handlers/help v1.0.0 => ./handlers/help
	github.com/nextuponstream/workoutReminderBot/handlers/unknown v1.0.0 => ./handlers/unknown
	github.com/nextuponstream/workoutReminderBot/handlers/create/activity v1.0.0 => ./handlers/create/activity
	github.com/nextuponstream/workoutReminderBot/entities v1.0.0 => ./entities
	github.com/nextuponstream/workoutReminderBot/repositories/mongo v1.0.0 => ./repositories/mongo
)

require (
	github.com/Syfaro/telegram-bot-api v4.6.4+incompatible
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible // indirect
	github.com/joho/godotenv v1.3.0
	github.com/nextuponstream/workoutReminderBot/handlers/help v1.0.0
	github.com/nextuponstream/workoutReminderBot/handlers/unknown v1.0.0
	github.com/nextuponstream/workoutReminderBot/handlers/create/activity v1.0.0
	github.com/nextuponstream/workoutReminderBot/entities v1.0.0
	github.com/nextuponstream/workoutReminderBot/repositories/mongo v1.0.0
)