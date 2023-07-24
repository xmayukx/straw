package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xmayukx/straw/handlers"
)

var bot *tgbotapi.BotAPI
var update *tgbotapi.Update

func UpdateInstance(u *tgbotapi.Update) {
	update = u
}

func BotInstance(b *tgbotapi.BotAPI) {
	bot = b
}

func ProcessMsg(msgQueue *chan tgbotapi.MessageConfig) {

	for msg := range *msgQueue {

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg.Text = "Hello, I'm Straw. I can download videos from YouTube and Instagram. Send me a link to a video and I'll download it for you."
			case "help":
				msg.Text = "Send me a link to a video and I'll download it for you."
			default:
				msg.Text = "I don't know that command."
			}
			if _, err := bot.Send(msg); err != nil {
				msg.Text = "Something went wrong."
				bot.Send(msg)
			}
		} else if strings.Contains(update.Message.Text, "https") || strings.Contains(update.Message.Text, "http") || strings.Contains(update.Message.Text, "www") {

			msg.Text = "⬇️Your video is downloading..."
			if _, err := bot.Send(msg); err != nil {
				msg.Text = "Something went wrong."
				bot.Send(msg)
			}
			filePath := handlers.VideoHandler(update.Message.Text, bot.Self.FirstName)

			file, err := os.Open(filePath)
			if err != nil {
				fmt.Println("Error opening file ", err.Error())
				return
			}

			buffer := make([]byte, 1024)

			var fileBytes []byte

			for {
				n, err := file.Read(buffer)
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Println("Error reading file ", err.Error())
					return
				}

				fileBytes = append(fileBytes, buffer[:n]...)
			}

			videoConfig := tgbotapi.NewVideo(update.Message.Chat.ID, tgbotapi.FileBytes{
				Name:  file.Name() + `.mp4`,
				Bytes: fileBytes,
			})
			file.Close()
			videoConfig.Caption = "🎥 Here's your video!"
			videoConfig.ReplyToMessageID = update.Message.MessageID
			if _, err := bot.Send(videoConfig); err != nil {
				msg.Text = "Something went wrong."
				bot.Send(msg)
			}
		} else {
			msg.Text = "Please provide a valid URL or cammand (/start or /help)."
			if _, err := bot.Send(msg); err != nil {
				msg.Text = "Something went wrong."
				bot.Send(msg)
			}
		}

	}
}
