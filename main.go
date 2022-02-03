package main

import (
	_ "github.com/lib/pq"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/datasource"
	"github.com/shifty11/cosmos-gov/log"
	"github.com/shifty11/cosmos-gov/telegram"
	"time"
)

func initDatabase() {
	database.MigrateDatabase()
	//database.DropChains()
	chains := datasource.ReadLensConfig("/home/rapha/.lens/config.yaml")
	database.CreateChains(chains)
	//database.DropProposals() // only testing
}

func main() {
	log.InitLogger()
	defer log.SyncLogger() // flushes buffer, if any

	defer database.Close()

	initDatabase()
	go datasource.FetchProposals()
	go telegram.Listen()
	time.Sleep(time.Duration(1<<63 - 1))
}
