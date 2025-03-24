package telegram

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
	
	"tg_bot/telegram/events"
)

func Update(bot *Bot) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.API.GetUpdatesChan(u)

	// Обработка входящих сообщений
	for update := range updates {
		if update.Message == nil {
			continue
		}
		
		text := strings.TrimSpace(update.Message.Text)
		chatID := update.Message.Chat.ID

		switch {
		case text == "/help" || text == "/start" :
			events.SendHelp(bot.API, chatID)
		case strings.HasPrefix(text, "/letter "):
			letter := strings.TrimSpace(strings.TrimPrefix(text, "/letter "))
			if len([]rune(letter)) == 1 {
				events.SendLetter(bot.API, bot.DB, chatID, letter)
			} else {
				events.SendMessage(bot.API, chatID, "Укажите одну букву, например: /letter А")
			}
		default:
			events.SendSearch(bot.API, bot.DB, chatID, text)
		}
	}
}