package datasource

import (
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
)

func updateRpcs(chainName string) {
	_, rpcs, err := getChainInfo(chainName)
	if err != nil {
		log.Sugar.Errorf("Error getting RPC's for chain %v: %v", chainName, err)
	}
	if len(rpcs) == 0 {
		log.Sugar.Errorf("Found no RPC's for chain %v: %v", chainName, err)
		return
	}
	err = database.NewChainManager().UpdateRpcs(chainName, rpcs)
	if err != nil {
		log.Sugar.Errorf("Error while updating RPC's for chain %v: %v", chainName, err)
	}
}
