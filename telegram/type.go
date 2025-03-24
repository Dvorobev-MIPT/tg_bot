package telegram

import (
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
)

type Bot struct {
	API *tgbotapi.BotAPI
	DB  *sql.DB
}
