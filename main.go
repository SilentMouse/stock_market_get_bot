package main

import (
	"flag"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"os"
)

var (
	// глобальная переменная в которой храним токен
	telegramBotToken string
)


var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("USD"),
		tgbotapi.NewKeyboardButton("EUR"),
		tgbotapi.NewKeyboardButton("JPY"),
	),
	//tgbotapi.NewKeyboardButtonRow(
	//	tgbotapi.NewKeyboardButton("4"),
	//	tgbotapi.NewKeyboardButton("5"),
	//	tgbotapi.NewKeyboardButton("6"),
	//),
)


func init() {
	// принимаем на входе флаг -telegrambottoken
	flag.StringVar(&telegramBotToken, "telegrambottoken", "", "Telegram Bot Token")
	flag.Parse()

	// без него не запускаемся
	if telegramBotToken == "" {
		log.Print("-telegrambottoken is required")
		os.Exit(1)
	}
}

func main() {
	// используя токен создаем новый инстанс бота
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// u - структура с конфигом для получения апдейтов
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// используя конфиг u создаем канал в который будут прилетать новые сообщения
	updates, err := bot.GetUpdatesChan(u)

	// в канал updates прилетают структуры типа Update
	// вычитываем их и обрабатываем
	for update := range updates {

		//reply := "Не знаю что сказать"
		if update.Message == nil {
			continue
		}

		// логируем от кого какое сообщение пришло
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.ReplyMarkup = numericKeyboard

		switch update.Message.Command() {
		case "start":
			msg.Text = "Hello. I am telegram bot"
		case "hello":
			msg.Text = "world"
		default:
			msg.Text = "Не знаю что сказать"
		}

		// создаем ответное сообщение
		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		//msg.ReplyMarkup = numericKeyboard

		switch update.Message.Text {
		case "open":
			msg.ReplyMarkup = numericKeyboard
		case "close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		case "USD":
			msg.Text = "57.78"
		case "EUR":
			msg.Text = "70.11"
		case "JPY":
			msg.Text = "53,84"
		}

		// отправляем
		bot.Send(msg)
	}
}