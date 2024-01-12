package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var startCommand = tgbotapi.BotCommand{
	Command:     "start",
	Description: "Send default message",
}

/*var createEventCommand = tgbotapi.BotCommand{
	Command:     "createEvent",
	Description: "Create an event",
}

var deleteEventCommand = tgbotapi.BotCommand{
	Command:     "deleteEvent",
	Description: "Delete an event",
}

var editEventCommand = tgbotapi.BotCommand{
	Command:     "editEvent",
	Description: "Edit an event",
}

var weekLayoutCommand = tgbotapi.BotCommand{
	Command:     "weekLayout",
	Description: "Change layout to show events of the week",
}

var dayLayoutCommand = tgbotapi.BotCommand{
	Command:     "dayLayout",
	Description: "Change layout to show events of the day",
}*/

///func startCommandFunc()
