package database

import (
	"context"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/migrate"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
	"os"
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

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			//goland:noinspection GoUnhandledErrorResult
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
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
}

type DbManagers struct {
	ChainManager                *ChainManager
	UserManager                 *UserManager
	TelegramUserManager         *TypedUserManager
	DiscordUserManager          *TypedUserManager
	TelegramChatManager         *TelegramChatManager
	DiscordChannelManager       *DiscordChannelManager
	ProposalManager             *ProposalManager
	DraftProposalManager        *DraftProposalManager
	SubscriptionManager         *SubscriptionManager
	TelegramSubscriptionManager *TelegramSubscriptionManager
	DiscordSubscriptionManager  *DiscordSubscriptionManager
	LensChainInfoManager        *LensChainInfoManager
	StatsManager                *StatsManager
	WalletManager               *WalletManager
}

func NewDefaultDbManagers() DbManagers {
	client, ctx := connect()
	return NewCustomDbManagers(client, ctx)
}

func NewCustomDbManagers(client *ent.Client, ctx context.Context) DbManagers {
	chainManager := NewChainManager(client, ctx)
	userManager := NewUserManager(client, ctx)
	telegramChatManager := NewTelegramChatManager(client, ctx, chainManager)
	discordChannelManager := NewDiscordChannelManager(client, ctx, chainManager)
	telegramUserManager := NewTypedUserManager(client, ctx, user.TypeTelegram)
	discordUserManager := NewTypedUserManager(client, ctx, user.TypeDiscord)
	return DbManagers{
		ChainManager:                chainManager,
		UserManager:                 userManager,
		TelegramUserManager:         telegramUserManager,
		DiscordUserManager:          discordUserManager,
		TelegramChatManager:         telegramChatManager,
		DiscordChannelManager:       discordChannelManager,
		ProposalManager:             NewProposalManager(client, ctx),
		DraftProposalManager:        NewDraftProposalManager(client, ctx),
		SubscriptionManager:         NewSubscriptionManager(client, ctx, userManager, chainManager, telegramChatManager, discordChannelManager),
		TelegramSubscriptionManager: NewTelegramSubscriptionManager(client, ctx, telegramUserManager, chainManager, telegramChatManager),
		DiscordSubscriptionManager:  NewDiscordSubscriptionManager(client, ctx, discordUserManager, chainManager, discordChannelManager),
		LensChainInfoManager:        NewLensChainInfoManager(client, ctx),
		StatsManager:                NewStatsManager(client, ctx),
		WalletManager:               NewWalletManager(client, ctx, chainManager),
	}
}
