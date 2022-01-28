package database

import (
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/migrate"
	"github.com/shifty11/cosmos-gov/log"
)

type Chain struct {
	Name    string
	ChainId string
}

func CreateChains(chains []Chain) {
	client, ctx := connect()
	for _, c := range chains {
		_, err := client.Chain.
			Query().
			Where(chain.ChainIDEQ(c.ChainId)).
			Only(ctx)
		if err != nil {
			_, err = client.Chain.
				Create().
				SetName(c.Name).
				SetChainID(c.ChainId).
				Save(ctx)
			if err != nil {
				log.Sugar.Panic(err)
			}
		}
	}
}

func MigrateDatabase() {
	client, ctx := connect()
	err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Sugar.Panic("failed creating schema resources: %v", err)
	}
}
