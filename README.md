# Straw
A telegram bot written in Golang for downloading Instagram and Youtube video.

### Requirements
- Golang installed on your machine (https://golang.org/)
- A Telegram bot token (you can get it from the BotFather)
- Internet connection to download videos from YouTube and Instagram
---
### Get Your Telegram Bot Token
To get your Telegram bot token, you'll need to create a new bot on Telegram. Follow these steps:

1. Open Telegram and search for the "BotFather" bot.
2. Start a chat with the BotFather and use the /newbot command to create a new bot.
3. Follow the instructions and provide a name and username for your bot.
3. Once the bot is created, the BotFather will provide you with a token. Keep this token safe; you'll need it to run the bot.
---
### Installation and Setup
1. Clone the repo

   ```
   git clone https://github.com/xmayukx/straw.git
   cd straw
   ```
2. Install dependencies:
   ```
   go mod tidy
   ```
3. Now create a file and name it `.env` to store your bot token. In terminal write the command:
   ```
   touch .env
   ```
4. In `.env` set the bot token like this:
   ```
   TGBOTAPIKEY=<YOUR_TOKEN>
   ```
5. Now open the terminal and type the following command to spin up the bot locally:
     ````
     go run main.go
     ````
---

# Using the Bot
- Open Telegram and search for your bot using the username you provided during bot creation.
- Send a message mentioning your bot and then type a YouTube or Instagram video link as a query in the chat. The bot will automatically process the inline query and provide you with the video to download.
- For example, you can mention the bot using `@YourBotUsername` and then type a YouTube or Instagram video link like `https://www.youtube.com/watch?v=example_video_id` or `https://www.instagram.com/p/example_post_id/` in the chat.

The bot should respond with the video file, and you can download it directly from the chat.
