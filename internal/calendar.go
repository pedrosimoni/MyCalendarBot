package internal

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

/*type event struct {
	User        user
	Day         int
	Month       int
	Year        int
	Hour        int
	Minutes     int
	Title       string
	Description string
	Notify      bool
}*/

func GenerateCalendar(year int, month time.Month) (calendar tgbotapi.InlineKeyboardMarkup) {
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	lastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, time.Local)

	row := GenerateRow(firstDay, lastDay)

	calendar = tgbotapi.NewInlineKeyboardMarkup(

		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("S", "S"),
			tgbotapi.NewInlineKeyboardButtonData("M", "M"),
			tgbotapi.NewInlineKeyboardButtonData("T", "T"),
			tgbotapi.NewInlineKeyboardButtonData("W", "W"),
			tgbotapi.NewInlineKeyboardButtonData("T", "T"),
			tgbotapi.NewInlineKeyboardButtonData("F", "F"),
			tgbotapi.NewInlineKeyboardButtonData("S", "S"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(row[0], "s1"),
			tgbotapi.NewInlineKeyboardButtonData(row[1], "m1"),
			tgbotapi.NewInlineKeyboardButtonData(row[2], "tu1"),
			tgbotapi.NewInlineKeyboardButtonData(row[3], "w1"),
			tgbotapi.NewInlineKeyboardButtonData(row[4], "th1"),
			tgbotapi.NewInlineKeyboardButtonData(row[5], "f1"),
			tgbotapi.NewInlineKeyboardButtonData(row[6], "sa1"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(row[7], "su2"),
			tgbotapi.NewInlineKeyboardButtonData(row[8], "m2"),
			tgbotapi.NewInlineKeyboardButtonData(row[9], "tu2"),
			tgbotapi.NewInlineKeyboardButtonData(row[10], "w2"),
			tgbotapi.NewInlineKeyboardButtonData(row[11], "th2"),
			tgbotapi.NewInlineKeyboardButtonData(row[12], "f2"),
			tgbotapi.NewInlineKeyboardButtonData(row[13], "sa2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(row[14], "su3"),
			tgbotapi.NewInlineKeyboardButtonData(row[15], "m3"),
			tgbotapi.NewInlineKeyboardButtonData(row[16], "tu3"),
			tgbotapi.NewInlineKeyboardButtonData(row[17], "w3"),
			tgbotapi.NewInlineKeyboardButtonData(row[18], "th3"),
			tgbotapi.NewInlineKeyboardButtonData(row[19], "f3"),
			tgbotapi.NewInlineKeyboardButtonData(row[20], "sa3"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(row[21], "su4"),
			tgbotapi.NewInlineKeyboardButtonData(row[22], "m4"),
			tgbotapi.NewInlineKeyboardButtonData(row[23], "tu4"),
			tgbotapi.NewInlineKeyboardButtonData(row[24], "w4"),
			tgbotapi.NewInlineKeyboardButtonData(row[25], "th4"),
			tgbotapi.NewInlineKeyboardButtonData(row[26], "f4"),
			tgbotapi.NewInlineKeyboardButtonData(row[27], "sa4"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(row[28], "su5"),
			tgbotapi.NewInlineKeyboardButtonData(row[29], "m5"),
			tgbotapi.NewInlineKeyboardButtonData(row[30], "tu5"),
			tgbotapi.NewInlineKeyboardButtonData(row[31], "w5"),
			tgbotapi.NewInlineKeyboardButtonData(row[32], "th5"),
			tgbotapi.NewInlineKeyboardButtonData(row[33], "f5"),
			tgbotapi.NewInlineKeyboardButtonData(row[34], "sa5"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(row[35], "su6"),
			tgbotapi.NewInlineKeyboardButtonData(row[36], "m6"),
			tgbotapi.NewInlineKeyboardButtonData(row[37], "tu6"),
			tgbotapi.NewInlineKeyboardButtonData(row[38], "w6"),
			tgbotapi.NewInlineKeyboardButtonData(row[39], "th6"),
			tgbotapi.NewInlineKeyboardButtonData(row[40], "f6"),
			tgbotapi.NewInlineKeyboardButtonData(row[41], "sa6"),
		),
	)

	return calendar
}

func GenerateRow(firstDay, lastDay time.Time) (row [43]string) {
	offset := 0
	switch firstDay.Weekday().String() {

	case "Monday":
		offset = 1
	case "Tuesday":
		offset = 2
	case "Wednesday":
		offset = 3
	case "Thursday":
		offset = 4
	case "Friday":
		offset = 5
	case "Saturday":
		offset = 6
	default:
		offset = 0
	}

	days := lastDay.Day()

	for i := 0; i < 43; i++ {
		if i >= offset && i+offset-1 <= days {
			row[i+offset-1] = fmt.Sprintf("%v", i)
		} else {
			row[i] = "_"
		}
	}

	return row

}
