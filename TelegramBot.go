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
		logistics := LogisticsGetOriginals("https://vsuwt.ru/files/abitur/lsitOfAbiturs/spo/23.02.01%20%D0%9E%D1%80%D0%B3%D0%B0%D0%BD%D0%B8%D0%B7%D0%B0%D1%86%D0%B8%D1%8F%20%D0%BF%D0%B5%D1%80%D0%B5%D0%B2%D0%BE%D0%B7%D0%BE%D0%BA%20%D0%B8%20%D1%83%D0%BF%D1%80%D0%B0%D0%B2%D0%BB%D0%B5%D0%BD%D0%B8%D0%B5%20%D0%BD%D0%B0%20%D1%82%D1%80%D0%B0%D0%BD%D1%81%D0%BF%D0%BE%D1%80%D1%82%D0%B5%20(%D0%BF%D0%BE%20%D0%B2%D0%B8%D0%B4%D0%B0%D0%BC)%209%20%D0%BA%D0%BB%D0%B0%D1%81%D1%81%D0%BE%D0%B2%20-%20%D0%9E%D1%87%D0%BD%D0%B0%D1%8F%20-%20%D0%B1%D1%8E%D0%B4%D0%B6%D0%B5%D1%82_%D0%91.html", "Логистика речное")
		cs := LogisticsGetOriginals("https://vsuwt.ru/files/abitur/lsitOfAbiturs/spo/09.02.07%20%D0%98%D0%BD%D1%84%D0%BE%D1%80%D0%BC%D0%B0%D1%86%D0%B8%D0%BE%D0%BD%D0%BD%D1%8B%D0%B5%20%D1%81%D0%B8%D1%81%D1%82%D0%B5%D0%BC%D1%8B%20%D0%B8%20%D0%BF%D1%80%D0%BE%D0%B3%D1%80%D0%B0%D0%BC%D0%BC%D0%B8%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5%209%20%D0%BA%D0%BB%D0%B0%D1%81%D1%81%D0%BE%D0%B2%20-%20%D0%9E%D1%87%D0%BD%D0%B0%D1%8F%20-%20%D0%B1%D1%8E%D0%B4%D0%B6%D0%B5%D1%82_%D0%91.html", "Информатика речное")
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%v\n%v\n%v", nntc, logistics, cs))

		bot.Send(msg)

	}
}
