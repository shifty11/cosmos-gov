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
	database.MigrateDatabase()
}

func startProposalFetching() {
	go func() {
		datasource.FetchProposals()
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
	_, err := c.AddFunc("0 10 * * *", func() { datasource.AddNewChains() }) // execute every Monday at 10.00
	if err != nil {
		log.Sugar.Errorf("while executing 'datasource.AddNewChains()' via cron: %v", err)
	}
	c.Start()
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
		startNewChainFetching()
	} else if len(args) > 0 && args[0] == "telegram" {
		startTelegramServer()
	} else {
		initDatabase()
		startProposalFetching()
		startNewChainFetching()
		startTelegramServer()
	}

	time.Sleep(time.Duration(1<<63 - 1))
}
