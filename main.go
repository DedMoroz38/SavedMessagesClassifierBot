package main

import (
	"log"
	"saved_messages_classifier/controllers"
	"saved_messages_classifier/server"
)

func main() {
	bot, err := server.InitBot()
	if err != nil {
		log.Fatalf("Error while connecting to bot: %v", err)
	}

	log.Printf("Authorized as @%s", bot.Self.UserName)

	controllers.HandleUpdates(bot)
}
