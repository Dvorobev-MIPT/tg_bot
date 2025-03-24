package events

import (
	"fmt"
	"strings"
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	
	"tg_bot/date_base"
	"tg_bot/teacher"
)

/* 
*  	\search - поиск преподавателя по ФИО (полному)
*	 Если преподаватель не найден - выводит до 3 наиболее близких ФИО
*/

func SendSearch(bot *tgbotapi.BotAPI, db *sql.DB, chatID int64, query string) {
	// Получаем ФИО преподавателя
	fullName := strings.TrimSpace(query)

	// Ищем преподавателя в базе данных
	teacher_, err := date_base.GetTeacherInfo(db, fullName)
	if err != nil {
		// Если преподаватель не найден, ищем похожие имена
		suggestions, err := teacher.FindSimilarTeachers(db, fullName)
		if err != nil {
			SendMessage(bot, chatID, "Произошла ошибка при поиске преподавателя.")
		}

		if len(suggestions) > 0 {
			message := "Преподаватель не найден. Возможно, вы имели в виду:\n"
			for i, name := range suggestions {
				message += fmt.Sprintf("%d. %s\n", i+1, name)
			}
			SendMessage(bot, chatID, message)
		} else {
			SendMessage(bot, chatID, "Преподаватель не найден.")
		}
		return
	}

	// Формируем ответ
	message := teacher.FormatTeacherInfo(teacher_)

	// Отправляем ответ пользователю
	SendMessage(bot, chatID, message)
}

