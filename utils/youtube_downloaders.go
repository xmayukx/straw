package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

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

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	downloadsDir := filepath.Join(currentDir, "downloads")

	if err := os.MkdirAll(downloadsDir, 0755); err != nil {
		fmt.Println("Error: ", err.Error())
	}

	filePath := filepath.Join(downloadsDir, "youtube_"+video.ID+".mp4")

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
