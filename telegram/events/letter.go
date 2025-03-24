package events

import (
	"fmt"
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	
	"tg_bot/date_base"	
)

// \letter - Вывод списка ФИО преподавателей по первой букве фамилии

func SendLetter(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64, letter string) {
	teachers, err := date_base.GetTeachersByLetter(db, letter)
	if err != nil {
		SendMessage(bot, chatID, "Ошибка при поиске преподавателей")
		return
	}
	
	if len(teachers) == 0 {
		message := fmt.Sprintf("Преподаватели на букву '%s' не найдены",
							   letter)
		SendMessage(bot, chatID, message)
		return
	}

	message:= fmt.Sprintf("Преподаватели на букву '%s':\n\n", letter)
	for i, teacher_ := range teachers {	
		message += fmt.Sprintf("%d. %s\n", i+1, teacher_.Name)
	}
	SendMessage(bot, chatID, message)
}