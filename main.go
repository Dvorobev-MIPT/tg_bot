package main


import (
	"tg_bot/date_base"
	"tg_bot/telegram"
)

func main() {
	const token = "{YOUR_TOKEN}"	// замените на {YOUR_TOKEN}" ваш токен

	db, err := date_base.DbConnect()
	if err != nil {
		return
	}
	defer db.Close()

	bot, err := telegram.NewBot(token, db)
	if err != nil {
		return
	}
	telegram.Update(bot)
}

