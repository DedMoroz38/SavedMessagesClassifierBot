package controllers

import (
	"log"
	"saved_messages_classifier/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
)

type Env struct {
	RedisClinet   *redis.Client
	PosgresClinet *db.PostgresClient
}

func HandleUpdates(bot *tgbotapi.BotAPI) {
	db := db.NewRedisClient()
	postgres

	env := &Env{RedisClinet: db}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("Error getting updates: %v", err)
	}

	for update := range updates {
		// if update.CallbackQuery != nil {
		// 	CallbackQueryHandler(update.CallbackQuery, bot, update.CallbackQuery.Message.Chat.ID)
		// 	continue
		// }
		if update.Message.IsCommand() {
			CommandHandler(update.Message, bot)
			continue
		}

		if update.Message != nil {
			env.MessageHandler(update.Message, bot)
		}

	}
}
