package telegram

import (
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
)

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
