package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
func main() {
	token := "6106995947:AAH2u_s7odcUnECWwWJjMY9NtuUVeFThNiM"
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			if strings.Contains(update.Message.Text, "lahir") {
				messag, _ := os.ReadFile("answer/lahir.txt")
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(messag))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			} else if strings.Contains(update.Message.Text, "hobi") {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hobi Ardi Firmansyah adalah bermain Volly")
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			} else if strings.Contains(update.Message.Text, "start") {
				mess, _ := os.ReadFile("answer/listpertanyaan.txt")
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hallo "+update.Message.From.UserName+"\n"+string(mess))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			} else if strings.Contains(update.Message.Text, "minuman") {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Minuman kesukaan Ardi adalah Lemon Tea")
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			} else if strings.Contains(update.Message.Text, "github") {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "https://github.com/ardifirmansy4h")
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			} else if strings.Contains(update.Message.Text, "linkedin") {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "https://www.linkedin.com/in/ardi-firmansyah-29346a22b")
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			} else if strings.Contains(update.Message.Text, "makanan") {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Makanan kesukaan Ardi adalah Soto")
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			} else if strings.Contains(update.Message.Text, "foto") {
				foto := tgbotapi.FilePath("answer/ardi.jpg")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, foto)
				photo.ReplyToMessageID = update.Message.MessageID
				bot.Send(photo)
			} else if strings.Contains(update.Message.Text, "lagu") {

				audio := tgbotapi.FilePath("answer/sudah.mp3")
				msg := tgbotapi.NewAudio(update.Message.Chat.ID, audio)
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			} else if strings.Contains(update.Message.Text, "pendidikan") {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ardi Firmansyah menempuh pendidikan di Universitas Malikussaleh")
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			} else if strings.Contains(update.Message.Text, "bahasa") {
				file, _ := os.ReadFile("answer/lang.txt")
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(file))
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Silahkan Bertanya Pada "+bot.Self.UserName)
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}

		}
	}
}


