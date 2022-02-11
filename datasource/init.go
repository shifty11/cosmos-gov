package datasource

import (
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
)

func InitChains() {
	for _, chain := range database.GetChains() {
		if !database.HasFirstOrSecondProposal(chain.Name) {
			proposals, err := fetchProposals(chain.Name, types.StatusNil)
			if err != nil {
				log.Sugar.Errorf("Chain '%v' has %v errors", chain.DisplayName, err)
			} else {
				for _, prop := range proposals.Proposals {
					database.CreateOrUpdateProposal(&prop, chain)
				}
			}
		}
	}
}
