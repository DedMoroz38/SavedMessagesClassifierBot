package main

import (
	"log"
	"saved_messages_classifier/server"
)

func main() {
	bot, err := server.Serve()
	if err != nil {
		log.Fatalf("Error while connecting to bot: %v", err)
	}

	bot.Debug = true
	log.Printf("Authorized as @%s", bot.Self.UserName)

	server.HandleUpdates(bot)
}
