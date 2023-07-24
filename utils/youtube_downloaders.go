package utils

import (
	"io"
	"log"

	"os"

	"github.com/kkdai/youtube/v2"
)

func YoutubeDownload(videoID string) string {

	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err.Error())
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err.Error())
	}

	filePath := `C:\Users\hazar\Downloads\strawvids\` + `youtube_` + video.ID + `_` + `.mp4`
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		log.Panic(err.Error())
	}
	return filePath
}
