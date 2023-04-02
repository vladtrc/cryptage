package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type TgBot struct {
	bot *tgbotapi.BotAPI
}

func (b *TgBot) Broadcast(message string) (err error) {
	chatIds, err := GetChatIds()
	if err != nil {
		return
	}
	for _, chatId := range chatIds {
		r := tgbotapi.NewMessage(chatId, message)
		_, err := b.bot.Send(r)
		if err != nil {
			log.Printf("could not send message: %s", err)
			continue
		}
	}
	return
}
func (b *TgBot) Updates() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := b.bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.Text == "/start" {
			err := AddUserId(update.Message.Chat.ID)
			if err != nil {
				log.Printf("unable to add user: %s", err)
			}
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		if _, err := b.bot.Send(msg); err != nil {
			panic(err)
		}
	}
}

func NewTgBot() (res TgBot) {
	var err error
	res.bot, err = tgbotapi.NewBotAPI(config.token)
	if err != nil {
		panic(err)
	}
	res.bot.Debug = true
	return
}
