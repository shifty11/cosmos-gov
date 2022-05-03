package datasource

import (
	"context"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
)

func (ds Datasource) Migrate(chainManager *database.ChainManager) {
	log.Sugar.Info("Update chains")
	for _, c := range chainManager.All() {
		if c.ChainID == "" {
			_, info, _, err := ds.getChainInfo(c.Name)
			if err != nil {
				return
			}
			c.Update().SetChainID(info.ChainID).SetAccountPrefix(c.AccountPrefix).SaveX(context.Background())
		}
	}
	log.Sugar.Info("All chains updated")
}
