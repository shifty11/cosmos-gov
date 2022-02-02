package main

import (
	_ "github.com/lib/pq"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/datasource"
	"github.com/shifty11/cosmos-gov/log"
)

func initDatabase() {
	var chains = []database.ChainBase{
		{ChainId: "cosmos-1", Name: "Cosmos"},
		{ChainId: "osmosis-1", Name: "Osmosis"},
	}
	database.MigrateDatabase()
	database.CreateChains(chains)
}

func main() {
	log.InitLogger()
	defer log.SyncLogger() // flushes buffer, if any

	defer database.Close()

	initDatabase()
	datasource.FetchProposals()
	//go datasource.PerformLensQuery()
	//go telegram.Listen()
	//time.Sleep(time.Duration(1<<63 - 1))
}
