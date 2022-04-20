package database

import (
	"context"
	"github.com/shifty11/cosmos-gov/ent/migrate"
	"github.com/shifty11/cosmos-gov/ent/migrationinfo"
	"github.com/shifty11/cosmos-gov/ent/user"
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
		log.Sugar.Panicf("Failed creating schema resources: %v", err)
	}
	migrateUsers()
}

// TODO: remove after migration
func migrateUsers() {
	client, ctx := connect()
	doesExist, err := client.MigrationInfo.
		Query().
		Where(migrationinfo.IsMigratedEQ(true)).
		Exist(ctx)
	if err != nil {
		log.Sugar.Panicf("Failed migrating %v", err)
	}
	if doesExist {
		return
	}

	users, err := client.User.
		Query().
		WithChains().
		All(ctx)
	if err != nil {
		log.Sugar.Panicf("Failed migrating %v", err)
	}

	_, err = client.TelegramChat.
		Delete().
		Exec(ctx)
	if err != nil {
		log.Sugar.Panicf("Failed migrating %v", err)
	}
	_, err = client.DiscordChannel.
		Delete().
		Exec(ctx)
	if err != nil {
		log.Sugar.Panicf("Failed migrating %v", err)
	}

	for _, u := range users {
		chains, err := u.QueryChains().All(ctx)
		if err != nil {
			log.Sugar.Panicf("Failed migrating %v", err)
		}
		if u.Type == user.TypeTelegram {
			err = client.TelegramChat.
				Create().
				SetID(u.ChatID).
				SetName("<not set>"). // TODO: set this field properly
				SetIsGroup(u.ChatID < 0).
				SetUser(u).
				AddChains(chains...).
				Exec(ctx)
			if err != nil {
				log.Sugar.Panicf("Failed migrating %v", err)
			}
		} else {
			err = client.DiscordChannel.
				Create().
				SetID(u.ChatID).
				SetName("<not set>"). // TODO: set this field properly
				SetIsGroup(true).     // TODO: set this field properly
				SetUser(u).
				AddChains(chains...).
				Exec(ctx)
			if err != nil {
				log.Sugar.Panicf("Failed migrating %v", err)
			}
		}
	}

	err = client.MigrationInfo.
		Create().
		SetIsMigrated(true).
		Exec(ctx)
	if err != nil {
		log.Sugar.Panicf("Failed migrating %v", err)
	}
	log.Sugar.Info("User migration successful")
}
