package main

import (
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/datasource"
	"github.com/shifty11/cosmos-gov/log"
	"github.com/shifty11/cosmos-gov/telegram"
	"os"
	"time"
)

func initDatabase() {
	lensConfig := os.Getenv("LENS_CONFIG")
	if lensConfig == "" {
		log.Sugar.Panicf("LENS_CONFIG is not set. Please provide the path to the lens config.yaml.")
	}

	log.Sugar.Info("Migrate database")
	database.MigrateDatabase()
	chains := datasource.ReadLensConfig(lensConfig)
	log.Sugar.Info("Create chains")
	database.CreateChains(chains)
	log.Sugar.Info("Init chains")
	datasource.InitChains()
}

func startProposalFetching() {
	c := cron.New()
	_, err := c.AddFunc("@every 5m", func() { datasource.FetchProposals() })
	if err != nil {
		log.Sugar.Errorf("while executing 'datasource.FetchProposals()' via cron: %v", err)
	}
	log.Sugar.Info("Start proposal fetching")
	c.Start()
	//go datasource.FetchProposals()
}

func startTelegramServer() {
	go telegram.Listen()
}

func main() {
	log.InitLogger()
	defer log.SyncLogger() // flushes buffer, if any

	defer database.Close()

	args := os.Args[1:]
	if len(args) > 0 && args[0] == "fetching" {
		initDatabase()
		startProposalFetching()
	} else if len(args) > 0 && args[0] == "telegram" {
		startTelegramServer()
	} else {
		initDatabase()
		startProposalFetching()
		startTelegramServer()
	}

	time.Sleep(time.Duration(1<<63 - 1))
}
