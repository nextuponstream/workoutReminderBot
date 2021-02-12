# How to use the bot
## Basic commands
| Command | description |
|---|----|
|`/help` | instructions on bot usage |
|`/activity <name of the activity> <a description>` | create an activity for your next workout |
|`/viewActivity <name of the activity>`| view an activity description |
|`/createWorkout` | not implemented |
|`/startWorkout` | not implemented |

# Setup
## BotFather
### Register bot
Go on telegram, contact BotFather and create a new bot via `/newBot`. Set up a bot name and the bot user name.
### Command hints
With BotFather, set commands with `/setCommands` and paste:
```
help - instructions on bot usage
activity - create an activity for your next workout with an optionnal description (e.g. /activity push-ups let's f*cking goooooo!)
viewactivity - view an activity description
```
## Environnement file
Create your own environnement file `bot.env` by following the file example `example.bot.env`.
## Start
```bash
git clone git@github.com:nextuponstream/workoutReminderBot.git
cd workoutReminderBot/scripts
sh runBot.sh
```