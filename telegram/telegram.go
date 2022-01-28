package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/shifty11/cosmos-gov/log"
	"os"
)

var api *tgbotapi.BotAPI = nil

func getApi() *tgbotapi.BotAPI {
	if api == nil {
		telegramToken := os.Getenv("TELEGRAM_TOKEN")
		if telegramToken == "" {
			log.Sugar.Panic("you must provide a telegram token as env variable")
		}
		botApi, err := tgbotapi.NewBotAPI(telegramToken)
		if err != nil {
			log.Sugar.Panic(err)
		}
		api = botApi
	}
	return api
}

func sendMainCommands(update tgbotapi.Update) {
	replyMarkup := createKeyboard([][]Button{{
		NewButton("Cosmos", "Cosmos"),
		NewButton("Osmosis", "Osmosis"),
	}},
	)

	chatId := getChatId(update)
	msg := tgbotapi.NewMessage(chatId, "Select the projects that you want to follow and receive notifications about new governance proposals")
	msg.ReplyMarkup = replyMarkup
	sendMessage(msg)
}

func Listen() {
	log.Sugar.Info("Start listening for commands")
	api := getApi()

	updateConfig := tgbotapi.NewUpdate(0)

	updates, err := api.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Sugar.Panic(err)
	}

	for update := range updates {
		sendMainCommands(update)
	}
}
