package main

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/xmayukx/straw/cmd"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TGBOTAPIKEY"))
	if err != nil {
		panic(err)
	}

	bot.Debug = false
	fmt.Print(bot.GetMe())
	cmd.BotInstance(bot)
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 20

	msgQueue := make(chan tgbotapi.MessageConfig)
	go cmd.ProcessMsg(&msgQueue)

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.ReplyToMessageID = update.Message.MessageID
		cmd.UpdateInstance(&update)
		msgQueue <- msg
	}

}
