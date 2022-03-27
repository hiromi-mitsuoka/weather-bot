package main

import (
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

// NOTE: LINE BOT docs(https://github.com/line/line-bot-sdk-go)
func main() {
	bot, err := linebot.New(
		// How to set to Heroku (
		// https://devcenter.heroku.com/ja/articles/config-vars#set-a-config-var
		// https://poster.ooo/howto/check-access-token/
		// )
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// result, err := weather.GetWeather()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// message := linebot.NewTextMessage(result)
	message := linebot.NewTextMessage("hello world")

	// NOTE: Send to all registered friends (https://developers.line.biz/ja/reference/messaging-api/#send-broadcast-message)
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
