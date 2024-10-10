package services

import (
	"saved_messages_classifier/constants"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func MessageHandler(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	switch message.Text {
	case constants.AddFolderButton:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Hello! How can I help you?")
		_, err := bot.Send(msg)
		if err != nil {
			return
		}
	}
}
