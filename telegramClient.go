package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
)

func startTelegramBot(botToken string) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	updates, err := bot.GetUpdatesChan(ucfg)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		typing := tgbotapi.NewChatAction(update.Message.Chat.ID, "typing")
		bot.Send(typing)
		if update.Message.Command() == "info" {
			var msgtext string
			message := getInfoFromMiner()
			for _, msg := range message {
				msgtext += msg + "\n"
			}
			//msgtext = "info"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgtext)
			msg.ParseMode = "HTML"
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
		if update.Message.Command() == "start" {
			msgtxt := "Привет. Отправь мне команду /go и получи список фильмов в прокате города Могилева. "
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgtxt)
			bot.Send(msg)
		}
		if update.Message.Command() == "help" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Помоги себе сам")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
		//bot.Send(msg)
	}

}
