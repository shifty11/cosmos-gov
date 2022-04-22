package main

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
	"github.com/shifty11/cosmos-gov/api/discord"
	"github.com/shifty11/cosmos-gov/api/telegram"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/datasource"
	"github.com/shifty11/cosmos-gov/log"
	registry "github.com/strangelove-ventures/lens/client/chain_registry"
	"os"
	"time"
)

func initDatabase() {
	database.MigrateDatabase()
}

func startProposalFetching() {
	go func() {
		datasource.FetchProposals() // start immediately and then every 5 minutes
		c := cron.New()
		_, err := c.AddFunc("@every 5m", func() { datasource.FetchProposals() })
		if err != nil {
			log.Sugar.Errorf("while executing 'datasource.FetchProposals()' via cron: %v", err)
		}
		c.Start()
	}()
}

func startNewChainFetching() {
	c := cron.New()
	ds := datasource.NewDatasource(context.Background(), registry.NewCosmosGithubRegistry(log.Sugar.Desugar()))
	_, err := c.AddFunc("06 21 * * *", func() { ds.AddNewChains() }) // execute every day at 10.00
	if err != nil {
		log.Sugar.Errorf("while executing 'datasource.AddNewChains()' via cron: %v", err)
	}
	c.Start()
}

func startProposalUpdating() {
	go func() {
		datasource.CheckForUpdates() // start immediately and then every hour
		c := cron.New()
		_, err := c.AddFunc("@every 1h", func() { datasource.CheckForUpdates() }) // execute every hour
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

func main() {
	defer log.SyncLogger() // flushes buffer, if any

	defer database.Close()

	args := os.Args[1:]
	if len(args) > 0 && args[0] == "fetching" {
		initDatabase()
		startProposalFetching()
		startNewChainFetching()
		startProposalUpdating()
	} else if len(args) > 0 && args[0] == "telegram" {
		startTelegramServer()
	} else if len(args) > 0 && args[0] == "discord" {
		startDiscordServer()
	} else {
		initDatabase()
		startProposalFetching()
		startNewChainFetching()
		startProposalUpdating()
		startTelegramServer()
		startDiscordServer()
	}

	time.Sleep(time.Duration(1<<63 - 1))
}
