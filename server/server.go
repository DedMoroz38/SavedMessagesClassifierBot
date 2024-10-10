package server

import (
	"errors"
	"log"
	"os"

	"saved_messages_classifier/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func Serve() (*tgbotapi.BotAPI, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	botToken := os.Getenv("TOKEN")
	if botToken == "" {
		return nil, errors.New("Failed to get bot token")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return bot, nil
}

func HandleUpdates(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("Error getting updates: %v", err)
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			services.CallbackQueryHandler(update.CallbackQuery, bot, update.CallbackQuery.Message.Chat.ID)
			continue
		}

		if update.Message == nil {
			services.MessageHandler(update.Message, bot)
		}

		if update.Message.IsCommand() {
			services.CommandHandler(update.Message, bot)
			continue
		}
	}

}
