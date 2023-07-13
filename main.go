package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TGBOTAPIKEY"))
	if err != nil {
		panic(err)
	}
	bot.Debug = false
	fmt.Println(bot.GetMe())

	log.Printf("Authorized on account %s", bot.Self.UserName)
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {

		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		// if !update.Message.IsCommand() { // ignore any non-command Messages
		// 	continue
		// }
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		switch update.Message.Command() {
		case "help":
			msg.Text = "I understand /sayhi and /status."
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."
		default:
			msg.Text = "I don't know that command"
		}

		// fmt.Printf("Message from user: %v\n", update.Message.Text)

		// msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {

			panic(err)
		}
	}

}
