package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/shifty11/cosmos-gov/database"
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

func sendMenu(update *tgbotapi.Update) {
	chatId := getChatId(update)
	chains := database.GetChainsForUser(chatId)

	var buttons []Button
	for _, c := range chains {
		symbol := "❌ "
		if c.Notify {
			symbol = "✅ "
		}
		buttons = append(buttons, NewButton(c.ChainId, symbol+c.Name))
	}
	replyMarkup := createKeyboard([][]Button{buttons})

	if update.CallbackQuery == nil {
		msg := tgbotapi.NewMessage(chatId, menuInfoMsg)
		msg.ReplyMarkup = replyMarkup
		sendMessage(msg)
	} else {
		msg := tgbotapi.EditMessageTextConfig{
			BaseEdit: tgbotapi.BaseEdit{ChatID: chatId,
				MessageID:   update.CallbackQuery.Message.MessageID,
				ReplyMarkup: &replyMarkup,
			},
			Text: menuInfoMsg,
		}
		sendMessage(msg)
	}
}

func updateNotification(update *tgbotapi.Update) {
	if update.CallbackQuery != nil {
		chatId := getChatId(update)
		chainId := update.CallbackQuery.Data
		err := database.AddOrRemoveChainForUser(chatId, chainId)
		if err != nil {
			log.Sugar.Error("Error while toggle chain for user %v", chatId)
		}
	}
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
		if update.CallbackQuery != nil {
			updateNotification(&update)
			sendMenu(&update)
		} else {
			sendMenu(&update)
		}
	}
}
