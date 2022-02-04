package main

import (
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/datasource"
	"github.com/shifty11/cosmos-gov/log"
	"github.com/shifty11/cosmos-gov/telegram"
	"time"
)

func initDatabase() {
	database.MigrateDatabase()
	chains := datasource.ReadLensConfig("/home/rapha/.lens/config.yaml")
	database.CreateChains(chains)
}

func startProposalFetching() {
	c := cron.New()
	_, err := c.AddFunc("@every 5m", func() { datasource.FetchProposals() })
	if err != nil {
		log.Sugar.Errorf("while executing 'datasource.FetchProposals()' via cron: %v", err)
	}
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

	initDatabase()
	startProposalFetching()
	startTelegramServer()
	time.Sleep(time.Duration(1<<63 - 1))
}
