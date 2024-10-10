package services

import (
	"log"
	"os/exec"
	"saved_messages_classifier/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func CallbackQueryHandler(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI, chatID int64) {
	clearTerminal()
	log.Println("CallbackQueryHandler")

	switch callbackQuery.Data {
	case config.CallbackQueries[config.AddFolder]:
		msg := tgbotapi.NewMessage(chatID, "Please send me the name of the folder.")
		_, err := bot.Send(msg)
		if err != nil {
			return
		}
	case config.CallbackQueries[config.AddFile]:
		msg := tgbotapi.NewMessage(chatID, "Please send me the file.")
		_, err := bot.Send(msg)
		if err != nil {
			return
		}
	}
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = exec.Command("clear").Stdout
	cmd.Run()
}
