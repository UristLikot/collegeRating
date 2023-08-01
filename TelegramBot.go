package main

import (
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func telegramBot() {
	bot, err := tgbotapi.NewBotAPI("6596621401:AAGzae_kFiZQUdGrQMU9T_TiMzssZfid-kk")
	if err != nil {
		panic(err)
	}

	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Получаем обновления от бота
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		nntc := NntcHandler()
		logistics := LogisticsGetOriginals("https://tinyurl.com/4fnb7jka", "Логистика речное")
		cs := LogisticsGetOriginals("https://tinyurl.com/mr2kmznn", "Информатика речное")
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%v\n%v\n%v", nntc, logistics, cs))

		bot.Send(msg)

	}
}
