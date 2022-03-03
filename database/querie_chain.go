package database

import (
	"entgo.io/ent/dialect/sql"
	"github.com/shifty11/cosmos-gov/dtos"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
	"strings"
)

func getChainByName(name string) (*ent.Chain, error) {
	client, ctx := connect()
	return client.Chain.
		Query().
		Where(chain.NameEQ(name)).
		Only(ctx)
}

func AddOrRemoveChainForUser(chatId int64, chainName string) error {
	_, ctx := connect()
	var userDto = getOrCreateUser(chatId)
	chainDto, err := getChainByName(chainName)
	if err != nil {
		return err
	}
	exists, err := userDto.
		QueryChains().
		Where(chain.IDEQ(chainDto.ID)).
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

func GetChainsForUser(chatId int64) []dtos.Chain {
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
		Order(ent.Asc(chain.FieldDisplayName)).
		All(ctx)
	var chains []dtos.Chain
	for _, c := range allChains {
		var chainEntry = dtos.Chain{Name: c.Name, DisplayName: c.DisplayName, Notify: false}
		for _, nc := range chainsOfUser { // check if user gets notified for this chain (c)
			if nc.ID == c.ID {
				chainEntry.Notify = true
			}
		}
		chains = append(chains, chainEntry)
	}
	return chains
}

func CreateChain(chainName string) *ent.Chain {
	client, ctx := connect()
	c, err := client.Chain.
		Query().
		Where(chain.NameEQ(chainName)).
		Only(ctx)
	if err != nil {
		log.Sugar.Infof("Create new chain: %v", chainName)
		c, err = client.Chain.
			Create().
			SetName(chainName).
			SetDisplayName(strings.Title(chainName)).
			Save(ctx)
		if err != nil {
			log.Sugar.Panic("Error while creating chain: %v", err)
		}
	}
	return c
}

func GetChains() []*ent.Chain {
	client, ctx := connect()
	chains, err := client.Chain.
		Query().
		All(ctx)
	if err != nil {
		log.Sugar.Panic("Error while querying chains: %v", err)
	}
	return chains
}

func GetChainStatistics() (*[]dtos.ChainStatistic, error) {
	client, ctx := connect()
	var chains []dtos.ChainStatistic
	err := client.Chain.Query().
		Order(ent.Asc(chain.FieldDisplayName)).
		GroupBy(chain.FieldDisplayName).
		Aggregate(func(s *sql.Selector) string {
			t := sql.Table(chain.UsersTable)
			s.Join(t).On(s.C(chain.FieldID), t.C(user.ChainsPrimaryKey[1]))
			return sql.As(sql.Count(t.C(user.ChainsPrimaryKey[1])), "notifications")
		}).
		Scan(ctx, &chains)
	if err != nil {
		return nil, err
	}
	return &chains, err
}
