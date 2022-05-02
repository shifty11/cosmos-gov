package database

import (
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
)

// PerformUpdateSubscription toggles the subscription for a chain
func PerformUpdateSubscription(chatId int64, userType user.Type, chainName string, userId int64, userName string, chatName string, isGroup bool) {
	if chainName == "" {
		return
	}
	log.Sugar.Debugf("Toggle subscription %v for %v #%v", chainName, userType, userId)
	err := AddOrRemoveChainForUser(chatId, userType, chainName, userId, userName, chatName, isGroup)
	if err != nil {
		log.Sugar.Errorf("Error while toggle subscription %v for %v %v", chainName, userType, userId)
	}
}
