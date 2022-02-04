package telegram

import (
	"fmt"
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

const NbrOfButtonsPerRow = 3

func sendMenu(update *tgbotapi.Update) {
	chatId, err := getChatId(update)
	if err != nil {
		log.Sugar.Error(err)
		return
	}
	chains := database.GetChainsForUser(chatId)

	var buttons [][]Button
	var buttonRow []Button
	for ix, c := range chains {
		symbol := "❌ "
		if c.Notify {
			symbol = "✅ "
		}
		buttonRow = append(buttonRow, NewButton(c.Name, symbol+c.DisplayName))
		if (ix+1)%NbrOfButtonsPerRow == 0 || ix == len(chains)-1 {
			buttons = append(buttons, buttonRow)
			buttonRow = []Button{}
		}
	}
	replyMarkup := createKeyboard(buttons)

	if update.CallbackQuery == nil {
		msg := tgbotapi.NewMessage(chatId, menuInfoMsg)
		msg.ReplyMarkup = replyMarkup
		sendMessageX(msg)
	} else {
		msg := tgbotapi.EditMessageTextConfig{
			BaseEdit: tgbotapi.BaseEdit{ChatID: chatId,
				MessageID:   update.CallbackQuery.Message.MessageID,
				ReplyMarkup: &replyMarkup,
			},
			Text: menuInfoMsg,
		}
		sendMessageX(msg)
	}
}

func updateNotification(update *tgbotapi.Update) {
	if update.CallbackQuery != nil {
		chatId, err := getChatId(update)
		if err != nil {
			log.Sugar.Error(err)
			return
		}
		chain := update.CallbackQuery.Data
		log.Sugar.Debugf("Toggle chain %v for user #%v", chain, chatId)
		err = database.AddOrRemoveChainForUser(chatId, chain)
		if err != nil {
			log.Sugar.Error("Error while toggle chain for user %v", chatId)
		}
	}
}

func isNotificationCommand(update *tgbotapi.Update) bool {
	api := getApi()
	return update.Message != nil &&
		(update.Message.Text == "/start" ||
			update.Message.Text == "/notifications" ||
			update.Message.Text == fmt.Sprintf("/notifications@%v", api.Self.UserName))
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
		} else if isNotificationCommand(&update) {
			sendMenu(&update)
		}
	}
}

func SendProposal(proposalText string, chatIds []int) map[int]struct{} {
	errIds := make(map[int]struct{})
	var exists = struct{}{}
	for _, chatId := range chatIds {
		msg := tgbotapi.NewMessage(int64(chatId), proposalText)
		msg.ParseMode = "markdown"
		log.Sugar.Debugf("Send proposal to chat #%v", chatId)
		err := sendMessage(msg)
		if err != nil {
			errIds[chatId] = exists
		}
	}
	return errIds
}
