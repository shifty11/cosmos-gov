package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
)

// Toggles the subscription for a chain
func performUpdateSubscription(update *tgbotapi.Update, chainName string) {
	if chainName == "" {
		return
	}
	chatId := getChatIdX(update)
	log.Sugar.Debugf("Toggle subscription %v for user #%v", chainName, chatId)
	err := database.AddOrRemoveChainForUser(chatId, chainName)
	if err != nil {
		log.Sugar.Error("Error while toggle subscription %v for user %v", chainName, chatId)
	}
}

// Enables or disables a chain for all users. Can only be performed by botadmins.
func performToggleChain(chainName string) {
	if chainName == "" {
		return
	}
	log.Sugar.Debugf("Enable/disable chain %v", chainName)
	err := database.EnableOrDisableChain(chainName)
	if err != nil {
		log.Sugar.Error("Error while toggle chain %v", chainName)
	}
}
