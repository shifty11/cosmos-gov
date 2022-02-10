package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
)

func performUpdateNotification(update *tgbotapi.Update) {
	chatId := getChatIdX(update)
	chain := update.CallbackQuery.Data
	log.Sugar.Debugf("Toggle chain %v for user #%v", chain, chatId)
	err := database.AddOrRemoveChainForUser(chatId, chain)
	if err != nil {
		log.Sugar.Error("Error while toggle chain for user %v", chatId)
	}
}
