package main

import (
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("MyCalendarBotToken"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			currentUser, err := FindUser(update.Message.Chat.ID)
			log.Printf("%v, %v", currentUser.UserID, err)
			//currentEvents := pullCurrentEvents(currentUser)

			if update.Message.IsCommand() {
				if update.Message.Command() == startCommand.Command {

					msg.ReplyMarkup = GenerateCalendar(time.Now().Year(), time.Now().Month())

				} /* else if update.Message.Command() == createEventCommand.Command {
					//newEvent := createEvent(currentUser)
				} else if update.Message.Command() == deleteEventCommand.Command {

				} else if update.Message.Command() == editEventCommand.Command {

				} else if update.Message.Command() == weekLayoutCommand.Command {

				} else if update.Message.Command() == dayLayoutCommand.Command {

				}*/
			}

			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}

		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "teste")
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}
