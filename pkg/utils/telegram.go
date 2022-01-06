package utils

import (
	"kaznet-status/config"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendToTelegram(templateMessage string) {
	bot, err := tgbotapi.NewBotAPI(config.Config("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	chatId, err := strconv.ParseInt(config.Config("TELEGRAM_BOT_CHAT_ID"), 10, 64)
	if err != nil {
		log.Panic(err)
	}

	msg := tgbotapi.NewMessage(chatId, templateMessage)

	bot.Send(msg)

}
