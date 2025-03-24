package telegram

import (
	"database/sql"
	"log"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
)

func BotInitialize(token string, db *sql.DB) (*Bot, error) {
	// Инициализация бота
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	return &Bot{
		API: bot,
		DB:  db,
	}, nil
}

func NewBot(token string, db *sql.DB) (*Bot, error) {
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}

	botAPI.Debug = true

	return &Bot{
		API: botAPI,
		DB:  db,
	}, nil
}