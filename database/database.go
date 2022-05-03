package database

import (
	"context"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/shifty11/cosmos-gov/ent/migrate"
	"github.com/shifty11/cosmos-gov/ent/migrationinfo"
	"github.com/shifty11/cosmos-gov/ent/user"
	regen "github.com/zach-klippenstein/goregen"
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
	migrateUsers()
}

// TODO: remove after migration
// TODO: I don't have the user ids of telegram/discord... needs an additional migration step or empty user
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
		token, err := regen.Generate("[A-Za-z0-9]{32}")
		if err != nil {
			log.Sugar.Panicf("Failed migrating %v", err)
		}
		u.Update().SetName(u.UserName).SetLoginToken(token).SaveX(ctx)
		if u.Type == user.TypeTelegram {
			err = client.TelegramChat.
				Create().
				SetChatID(u.ChatID).
				SetName(u.ChatName).
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
				SetChannelID(u.ChatID).
				SetName(u.ChatName).
				SetIsGroup(true). // TODO: set this field properly
				SetUser(u).
				AddChains(chains...).
				Exec(ctx)
			if err != nil {
				log.Sugar.Panicf("Failed migrating %v", err)
			}
		}
	}

	for _, c := range client.Chain.Query().AllX(ctx) {
		c.Update().
			SetCreateTime(c.CreatedAt).
			SetUpdateTime(c.UpdatedAt).
			SaveX(ctx)
	}
	for _, c := range client.LensChainInfo.Query().AllX(ctx) {
		c.Update().
			SetCreateTime(c.CreatedAt).
			SetUpdateTime(c.UpdatedAt).
			SaveX(ctx)
	}
	for _, c := range client.Proposal.Query().AllX(ctx) {
		c.Update().
			SetCreateTime(c.CreatedAt).
			SetUpdateTime(c.UpdatedAt).
			SaveX(ctx)
	}
	for _, c := range client.RpcEndpoint.Query().AllX(ctx) {
		c.Update().
			SetCreateTime(c.CreatedAt).
			SetUpdateTime(c.UpdatedAt).
			SaveX(ctx)
	}
	for _, c := range client.User.Query().AllX(ctx) {
		c.Update().
			SetCreateTime(c.CreatedAt).
			SetUpdateTime(c.UpdatedAt).
			SaveX(ctx)
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

type DbManagers struct {
	ChainManager                *ChainManager
	UserManager                 *UserManager
	TelegramUserManager         *TypedUserManager
	DiscordUserManager          *TypedUserManager
	TelegramChatManager         *TelegramChatManager
	DiscordChannelManager       *DiscordChannelManager
	ProposalManager             *ProposalManager
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
		SubscriptionManager:         NewSubscriptionManager(client, ctx, userManager, chainManager, telegramChatManager, discordChannelManager),
		TelegramSubscriptionManager: NewTelegramSubscriptionManager(client, ctx, telegramUserManager, chainManager, telegramChatManager),
		DiscordSubscriptionManager:  NewDiscordSubscriptionManager(client, ctx, discordUserManager, chainManager, discordChannelManager),
		LensChainInfoManager:        NewLensChainInfoManager(client, ctx),
		StatsManager:                NewStatsManager(client, ctx),
		WalletManager:               NewWalletManager(client, ctx, chainManager),
	}
}
