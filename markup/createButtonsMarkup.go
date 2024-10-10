package markup

import (
	"saved_messages_classifier/constants"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var AddFolderButton = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(constants.AddFolderButton),
	),
)
