package main

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
	"github.com/shifty11/cosmos-gov/api/discord"
	"github.com/shifty11/cosmos-gov/api/grpc"
	"github.com/shifty11/cosmos-gov/api/telegram"
	"github.com/shifty11/cosmos-gov/authz"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/datasource"
	"github.com/shifty11/cosmos-gov/log"
	registry "github.com/strangelove-ventures/lens/client/chain_registry"
	"os"
	"time"
)

func initDatabase(cd *datasource.ChainDatasource, chainManager *database.ChainManager) {
	database.MigrateDatabase()

	if len(chainManager.All()) == 0 { // Add chains after DB has been newly created
		cd.AddNewChains()
	}
}

func startProposalFetching(ds *datasource.ProposalDatasource) {
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

func startDraftProposalFetching(dc *datasource.DiscourseCrawler) {
	go func() {
		dc.FetchDraftProposals() // start immediately and then every 15 minutes
		c := cron.New()
		_, err := c.AddFunc("@every 15m", func() { dc.FetchDraftProposals() })
		if err != nil {
			log.Sugar.Errorf("while executing 'datasource.FetchDraftProposals()' via cron: %v", err)
		}
		c.Start()
	}()
}

func startNewChainFetching(cd *datasource.ChainDatasource) {
	c := cron.New()
	_, err := c.AddFunc("0 10 * * *", func() { cd.AddNewChains() }) // execute every day at 10.00
	if err != nil {
		log.Sugar.Errorf("while executing 'datasource.AddNewChains()' via cron: %v", err)
	}
	c.Start()
}

func startProposalUpdating(ds *datasource.ProposalDatasource) {
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

func startTelegramServer(tgClient *telegram.TelegramClient) {
	go tgClient.Start()
}

func startDiscordServer(discordClient *discord.DiscordClient) {
	go discordClient.Start()
}

func startGrpcServer() {
	go grpc.Start()
}

func main() {
	defer log.SyncLogger() // flushes buffer, if any

	defer database.Close()

	managers := database.NewDefaultDbManagers()
	authzClient := authz.NewAuthzClient(managers.ChainManager, managers.WalletManager)
	tgClient := telegram.NewTelegramClient(managers, authzClient)
	tgLightClient := telegram.NewTelegramLightClient(managers)
	discordClient := discord.NewDiscordClient(managers)
	discordLightClient := discord.NewDiscordLightClient(managers)
	reg := registry.NewCosmosGithubRegistry(log.Sugar.Desugar())
	cd := datasource.NewChainDatasource(context.Background(), managers, reg, tgLightClient, discordLightClient)
	ds := datasource.NewProposalDatasource(context.Background(), managers, reg, nil, tgLightClient, discordLightClient)
	dc := datasource.NewDiscourseCrawler(context.Background(), managers, tgLightClient, discordLightClient)

	args := os.Args[1:]
	if len(args) > 0 && args[0] == "fetching" {
		initDatabase(cd, managers.ChainManager)
		startProposalFetching(ds)
		startDraftProposalFetching(dc)
		startNewChainFetching(cd)
		startProposalUpdating(ds)
	} else if len(args) > 0 && args[0] == "telegram" {
		startTelegramServer(tgClient)
	} else if len(args) > 0 && args[0] == "discord" {
		startDiscordServer(discordClient)
	} else if len(args) > 0 && args[0] == "grpc" {
		startGrpcServer()
	} else {
		initDatabase(cd, managers.ChainManager)
		startProposalFetching(ds)
		startDraftProposalFetching(dc)
		startNewChainFetching(cd)
		startProposalUpdating(ds)
		startTelegramServer(tgClient)
		startDiscordServer(discordClient)
	}

	time.Sleep(time.Duration(1<<63 - 1))
}
