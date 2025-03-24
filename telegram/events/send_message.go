package events

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Функция для отправки сообщений
func SendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	message := tgbotapi.NewMessage(chatID, text)
	bot.Send(message)
}