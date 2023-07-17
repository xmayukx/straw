package setup

import (
	"io/ioutil"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xmayukx/straw/handlers"
)

var bot *tgbotapi.BotAPI
var update *tgbotapi.Update

func Update(u *tgbotapi.Update) {
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
		} else {

			filePath := handlers.VideoHandler(update.Message.Text)

			videoBytes, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Panic(err)
			}

			videoConfig := tgbotapi.NewVideo(update.Message.Chat.ID, tgbotapi.FileBytes{
				Name:  "video.mp4",
				Bytes: videoBytes,
			})
			videoConfig.Caption = "Here's your video!"
			videoConfig.ReplyToMessageID = update.Message.MessageID

			if _, err := bot.Send(videoConfig); err != nil {
				msg.Text = "Something went wrong."
				bot.Send(msg)
			}
		}

	}
}
