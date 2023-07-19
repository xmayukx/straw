package handlers

import (
	"log"

	"strings"

	"github.com/kkdai/youtube/v2"
	"github.com/xmayukx/straw/utils"
)

func VideoHandler(url string, userID string) string {

	if strings.Contains(url, "youtu.be") {
		videoID, err := youtube.ExtractVideoID(url)
		if err != nil {
			log.Panic(err)
		}
		return utils.YoutubeDownload(videoID, userID)

	} else if strings.Contains(url, "instagram") {
		return "Coming Soon!"
	} else {
		return "I don't know that link."
	}

}
