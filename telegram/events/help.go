package events

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// \help - вывод информации о командах боту

// Новые функции для обработки команд
func SendHelp(bot *tgbotapi.BotAPI, chatID int64) {
	helpText := 
	`Доступные команды:
	/help - показать это сообщение;
	/letter [Буква] - показать всех преподавателей на указанную букву фамилии;
	[ФИО] - поиск преподавателя по имени (например: Иванов Иван Иванович).

	Будьте внимательны:
	Любой ввод, отличный от \help и \letter, воспринимается как ФИО.

	Так же примите во внимание, что даты выводятся в формате yyyy-mm-dd.
	`
	
	SendMessage(bot, chatID, helpText)
}
