package controllers

import (
	"fmt"
	"log"
	"saved_messages_classifier/constants"
	"saved_messages_classifier/services"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type MessageController struct {
	Env *Env
	Bot *tgbotapi.BotAPI
}

func (env *Env) MessageHandler(message *tgbotapi.Message, bot *tgbotapi.BotAPI) { // TODO: pass message text only
	log.Println("MessageHandler")
	MessageController := &MessageController{
		Env: env,
		Bot: bot,
	}

	switch message.Text {
	case constants.AddFolderButton:
		MessageController.AddFolderAction(message.Chat.ID)
	default:
		MessageController.DefaultAction(message.Chat.ID, message.Text)
	}
}

func (mc *MessageController) DefaultAction(chatId int64, messageText string) {
	bot := mc.Bot
	stateService := services.StateService{RedisClient: mc.Env.RedisClinet}
	folderService := services.FolderService{RedisClient: mc.Env.RedisClinet}
	chatIdStr := strconv.FormatInt(chatId, 10)

	state, err := stateService.GetState(chatIdStr)
	if err != nil {
		fmt.Println("Error getting state:", err)
		return
	}

	switch state {
	case constants.AddFolder:
		err = folderService.CreateFolder(messageText)
		msg := tgbotapi.NewMessage(chatId, "Added new folder: "+messageText)
		_, err = bot.Send(msg)
		if err != nil {
			return
		}
		err = stateService.DeleteState(chatIdStr)
		if err != nil {
			fmt.Println("Error deleting state:", err)
			return
		}
	default:
		msg := tgbotapi.NewMessage(chatId, "Which folder do you want to add this message to?")
		_, err = bot.Send(msg)
		if err != nil {
			return
		}
	}

}

func (mc *MessageController) AddFolderAction(chatId int64) {
	bot := mc.Bot
	stateService := services.StateService{RedisClient: mc.Env.RedisClinet}
	chatIdStr := strconv.FormatInt(chatId, 10)
	fmt.Println("chat id: ", chatIdStr)

	err := stateService.SetState(chatIdStr, constants.AddFolder)
	if err != nil {
		fmt.Println("Error setting state:", err)
		return
	}

	msg := tgbotapi.NewMessage(chatId, "Enter folder name")
	_, err = bot.Send(msg)
	if err != nil {
		return
	}
}
