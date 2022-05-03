package main

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
	"github.com/shifty11/cosmos-gov/api/discord"
	"github.com/shifty11/cosmos-gov/api/grpc"
	"github.com/shifty11/cosmos-gov/api/telegram"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/datasource"
	"github.com/shifty11/cosmos-gov/log"
	registry "github.com/strangelove-ventures/lens/client/chain_registry"
	"os"
	"time"
)

func initDatabase(ds *datasource.Datasource, chainManager *database.ChainManager) {
	database.MigrateDatabase()
	ds.Migrate(chainManager)

	if len(chainManager.All()) == 0 { // Add chains after DB has been newly created
		ds.AddNewChains()
	}
}

func startProposalFetching(ds *datasource.Datasource) {
	go func() {
		ds.FetchProposals() // start immediately and then every 5 minutes
		c := cron.New()
		_, err := c.AddFunc("@every 5m", func() { ds.FetchProposals() })
		if err != nil {
			log.Sugar.Errorf("while executing 'datasource.FetchProposals()' via cron: %v", err)
		}
		c.Start()
	}()
}

func startNewChainFetching(ds *datasource.Datasource) {
	c := cron.New()
	_, err := c.AddFunc("0 10 * * *", func() { ds.AddNewChains() }) // execute every day at 10.00
	if err != nil {
		log.Sugar.Errorf("while executing 'datasource.AddNewChains()' via cron: %v", err)
	}
	c.Start()
}

func startProposalUpdating(ds *datasource.Datasource) {
	go func() {
		ds.CheckForUpdates() // start immediately and then every hour
		c := cron.New()
		_, err := c.AddFunc("@every 1h", func() { ds.CheckForUpdates() }) // execute every hour
		if err != nil {
			log.Sugar.Errorf("while executing 'datasource.CheckForUpdates()' via cron: %v", err)
		}
		c.Start()
	}()
}

func startTelegramServer() {
	go telegram.Start()
}

func startDiscordServer() {
	go discord.Start()
}

func startGrpcServer() {
	go grpc.Start()
}

func main() {
	defer log.SyncLogger() // flushes buffer, if any

	defer database.Close()

	managers := database.NewDefaultDbManagers()
	reg := registry.NewCosmosGithubRegistry(log.Sugar.Desugar())
	ds := datasource.NewDatasource(context.Background(), managers, reg, nil)

	args := os.Args[1:]
	if len(args) > 0 && args[0] == "fetching" {
		initDatabase(ds, managers.ChainManager)
		startProposalFetching(ds)
		startNewChainFetching(ds)
		startProposalUpdating(ds)
	} else if len(args) > 0 && args[0] == "telegram" {
		startTelegramServer()
	} else if len(args) > 0 && args[0] == "discord" {
		startDiscordServer()
	} else if len(args) > 0 && args[0] == "grpc" {
		startGrpcServer()
	} else {
		initDatabase(nil, nil)
		startProposalFetching(ds)
		startNewChainFetching(ds)
		startProposalUpdating(ds)
		startTelegramServer()
		startDiscordServer()
	}

	time.Sleep(time.Duration(1<<63 - 1))
}
