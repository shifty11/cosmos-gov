package database

import (
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
)

// PerformUpdateSubscription toggles the subscription for a chain
func PerformUpdateSubscription(userId int64, userType user.Type, chainName string) {
	if chainName == "" {
		return
	}
	log.Sugar.Debugf("Toggle subscription %v for %v #%v", chainName, userType, userId)
	err := AddOrRemoveChainForUser(userId, userType, chainName)
	if err != nil {
		log.Sugar.Errorf("Error while toggle subscription %v for %v %v", chainName, userType, userId)
	}
}
