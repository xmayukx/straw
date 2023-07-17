package handlers

import (
	"log"

	"strings"

	"github.com/kkdai/youtube/v2"
	"github.com/xmayukx/straw/utils"
)

func VideoHandler(url string) string {

	if strings.Contains(url, "youtu.be") {
		videoID, err := youtube.ExtractVideoID(url)
		if err != nil {
			log.Panic(err)
		}
		return utils.YoutubeDownload(videoID)

	} else if strings.Contains(url, "instagram") {
		return "Comming Soon!"
	} else {
		return "I don't know that link."
	}

}
