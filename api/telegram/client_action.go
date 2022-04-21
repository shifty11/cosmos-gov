package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
)

// Enables or disables a chain for all users. Can only be performed by botadmins.
func performToggleChain(chainName string) {
	if chainName == "" {
		return
	}
	log.Sugar.Debugf("Enable/disable chain %v", chainName)
	err := database.NewChainManager().EnableOrDisableChain(chainName)
	if err != nil {
		log.Sugar.Error("Error while toggle chain %v", chainName)
	}
}

// performUpdateSubscription toggles the subscription for a chain
func performUpdateSubscription(update *tgbotapi.Update, chainName string) {
	if chainName == "" {
		return
	}
	tgChatId := getChatIdX(update)
	chatName := getChatName(update)
	log.Sugar.Debugf("Toggle subscription %v for Telegram chat %v (%v)", chainName, chatName, tgChatId)
	manager := database.NewTelegramChatManager()
	_, err := manager.AddOrRemoveChain(tgChatId, chainName)
	if err != nil {
		log.Sugar.Errorf("Error while toggle subscription %v for Telegram chat %v (%v)", chainName, chatName, tgChatId)
	}
}
