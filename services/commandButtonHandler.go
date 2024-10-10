package services

import (
	"saved_messages_classifier/constants"
	"saved_messages_classifier/markup"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func CommandHandler(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	switch message.Command() {
	case "start":
		msg := tgbotapi.NewMessage(message.Chat.ID, constants.GreatingMessage)
		msg.ReplyMarkup = markup.AddFolderButton

		_, err := bot.Send(msg)
		if err != nil {
			return
		}
	}
}
