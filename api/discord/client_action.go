package discord

import (
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
)

// performUpdateSubscription toggles the subscription for a chain
func performUpdateSubscription(channelId int64, chainName string) {
	if chainName == "" {
		return
	}
	log.Sugar.Debugf("Toggle subscription %v for Telegram chat #%v", chainName, channelId)
	manager := database.NewDiscordChannelManager()
	_, err := manager.AddOrRemoveChain(channelId, chainName)
	if err != nil {
		log.Sugar.Errorf("Error while toggle subscription %v for Telegram chat #%v", chainName, channelId)
	}
}
