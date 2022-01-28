package database

import (
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/migrate"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
)

func getOrCreateUser(chatId int64) *ent.User {
	client, ctx := connect()
	var userDto *ent.User
	var err error
	userDto, err = client.User.
		Query().
		Where(user.ChatIDEQ(chatId)).
		Only(ctx)
	if err != nil {
		userDto, err = client.User.
			Create().
			SetChatID(chatId).
			Save(ctx)
		if err != nil {
			log.Sugar.Panic("Error while creating user: %v", err)
		}
	}
	return userDto
}

func getChainByChainId(chainId string) (*ent.Chain, error) {
	client, ctx := connect()
	return client.Chain.
		Query().
		Where(chain.ChainID(chainId)).
		Only(ctx)
}

func AddOrRemoveChainForUser(chatId int64, chainId string) error {
	_, ctx := connect()
	var userDto = getOrCreateUser(chatId)
	chainDto, err := getChainByChainId(chainId)
	if err != nil {
		return err
	}
	exists, err := userDto.
		QueryChains().
		Where(chain.ID(chainDto.ID)).
		Exist(ctx)
	if err != nil {
		return err
	}
	if exists {
		_, err := userDto.
			Update().
			RemoveChainIDs(chainDto.ID).
			Save(ctx)
		if err != nil {
			return err
		}
	} else {
		_, err := userDto.
			Update().
			AddChainIDs(chainDto.ID).
			Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

type Chain struct {
	Name    string
	ChainId string
	Notify  bool
}

func GetChainsForUser(chatId int64) []Chain {
	client, ctx := connect()
	var userDto = getOrCreateUser(chatId)
	chainsOfUser, err := client.Chain.
		Query().
		Where(chain.HasUsersWith(user.ID(userDto.ID))).
		All(ctx)
	if err != nil {
		log.Sugar.Panic("Error while fetching chains for user %v: %v", userDto.ID, err)
	}
	allChains, err := client.Chain.
		Query().
		All(ctx)
	var chains []Chain
	for _, c := range allChains {
		var chainEntry = Chain{Name: c.Name, ChainId: c.ChainID, Notify: false}
		for _, nc := range chainsOfUser { // check if user gets notified for this chain (c)
			if nc.ID == c.ID {
				chainEntry.Notify = true
			}
		}
		chains = append(chains, chainEntry)
	}
	return chains
}

type ChainBase struct {
	Name    string
	ChainId string
}

func CreateChains(chains []ChainBase) {
	client, ctx := connect()
	for _, c := range chains {
		_, err := client.Chain.
			Query().
			Where(chain.ChainIDEQ(c.ChainId)).
			Only(ctx)
		if err != nil {
			_, err = client.Chain.
				Create().
				SetName(c.Name).
				SetChainID(c.ChainId).
				Save(ctx)
			if err != nil {
				log.Sugar.Panic("Error while creating chains: %v", err)
			}
		}
	}
}

func MigrateDatabase() {
	client, ctx := connect()
	err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Sugar.Panic("Failed creating schema resources: %v", err)
	}
}
