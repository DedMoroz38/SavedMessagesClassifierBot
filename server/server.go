package server

import (
	"errors"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func InitBot() (*tgbotapi.BotAPI, error) {
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

	return bot, nil
}
