# How to use the bot
## Bot commands
... you can issue as a telegram user:
| Command | description |
|---|----|
|`/help` | instructions on bot usage |
| `/exercise <activity name> [<r/s/l/d/n value>]` | create exercise for activity and optionnaly indicate its reps, sets, length, duration and notes |
|`/viewexercises`| view all exercises created by you |
|`/activity <name> <a description>` | provide a description for an activity |
|`/viewactivity <name>`| view an activity description |
|`/workout <name> <activities>` | create a workout made of activities |
|`/routine <name> <workouts>` | create a routine made of workouts |
|`/remindme <routine> <from> <to> <when: mo/tu/we/th/fr/sa/su>` | sends a reminder to maintain the routine |

# Setup
## BotFather
### Register bot
Go on telegram, start a chat with BotFather and create a new bot via `/newBot`. Set up a bot name and the bot user name.
### Command hints
With BotFather, set commands with `/setCommands` and paste:
```
help - instructions on bot usage
exercise - create exercise for activity and optionnaly indicate its reps, sets, length, duration and notes (e.g. /exercise <activity name> <r, s, l, d or n value>)
viewexercises - view all exercises created by you
activity - create an activity for your next workout with an optionnal description (e.g. /activity push-ups let's f*cking goooooo!)
viewactivity - view an activity description
workout - create a workout made of activities (e.g. /workout gitStrong pushups run)
routine - create a routine made of workouts (e.g. /routine gitStrong)
remindme - sends a reminder to maintain the routine (e.g. /remindme gitStrong 16 21 tu we th fr sa)
```
## Environnement file
Create your own environnement file `.env` by following the example file `example.env`.
## Start
```bash
git clone git@github.com:nextuponstream/workoutReminderBot.git
cd workoutReminderBot/scripts
sh runBot.sh
```