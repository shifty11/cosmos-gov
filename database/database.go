package database

import (
	"context"
	"github.com/shifty11/cosmos-gov/ent/migrate"
	registry "github.com/strangelove-ventures/lens/client/chain_registry"
	"os"

	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/log"
)

var dbClient *ent.Client
var localhost = "host=localhost user=cosmosgov password=cosmosgov dbname=cosmosgov port=5432 sslmode=disable TimeZone=Europe/Zurich"

func connect() (*ent.Client, context.Context) {
	if dbClient == nil {
		dsn := os.Getenv("DATABASE_URL")
		if dsn == "" {
			dsn = localhost
		}
		newClient, err := ent.Open("postgres", dsn)
		if err != nil {
			log.Sugar.Panic("failed to connect to server ", err)
		}
		dbClient = newClient
	}
	return dbClient, context.Background()
}

func Close() {
	if dbClient != nil {
		err := dbClient.Close()
		if err != nil {
			log.Sugar.Error(err)
		}
	}
}

func MigrateDatabase() {
	log.Sugar.Info("Migrate database")
	client, ctx := connect()
	err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Sugar.Panic("Failed creating schema resources: %v", err)
	}
	migrateChains()
}

func migrateChains() {
	client, ctx := connect()

	chains := client.Chain.Query().AllX(ctx)
	reg := registry.DefaultChainRegistry(log.Sugar.Desugar())

	for _, c := range chains {
		exist, err := c.QueryRPCEndpoints().Exist(ctx)
		if err != nil {
			log.Sugar.Errorf("Failed to query rpcs on %v: %v \n", c.Name, err)
			continue
		}
		if exist {
			log.Sugar.Debugf("Skip %s:", c.Name)
			continue
		}

		chainInfo, err := reg.GetChain(context.Background(), c.Name)
		if err != nil {
			log.Sugar.Errorf("Failed to get chain client on %v: %v \n", c.Name, err)
			continue
		}

		rpcs, err := chainInfo.GetRPCEndpoints(context.Background())
		if err != nil {
			log.Sugar.Errorf("Failed to get RPC endpoints on chain %s: %v \n", chainInfo.ChainID, err)
			continue
		}
		for _, rpc := range rpcs {
			_, err := client.RpcEndpoint.
				Create().
				SetEndpoint(rpc).
				SetChain(c).
				Save(ctx)
			if err != nil {
				log.Sugar.Errorf("Failed to save RPC endpoint on chain %s: %v \n", chainInfo.ChainID, err)
				continue
			}
		}
		log.Sugar.Infof("Added %v RPC's for chain %v", len(rpcs), c.DisplayName)
	}
}
