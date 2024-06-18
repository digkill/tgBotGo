package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	fmt.Printf("TELEGRAM BOT TOKEN: %s", os.Getenv("TELEGRAM_BOT_TOKEN"))

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			handleCommand(bot, update.Message)
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello! Choose an action:")
		msg.ReplyMarkup = getMainKeyboard()
		bot.Send(msg)
	}

	fmt.Println("Bot ready!")
}

func handleCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		msg := tgbotapi.NewMessage(message.Chat.ID, "Welcome!.")
		msg.ReplyMarkup = getMainKeyboard()
		bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "I don't know that command.")
		bot.Send(msg)
	}
}

func getMainKeyboard() tgbotapi.ReplyKeyboardMarkup {
	button1 := tgbotapi.NewKeyboardButton("Download")
	row1 := tgbotapi.NewKeyboardButtonRow(button1)

	keyboard := tgbotapi.NewReplyKeyboard(row1)
	return keyboard
}
