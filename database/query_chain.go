package database

import (
	"entgo.io/ent/dialect/sql"
	"github.com/shifty11/cosmos-gov/common"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/proposal"
	"github.com/shifty11/cosmos-gov/ent/rpcendpoint"
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

func AddOrRemoveChainForUser(chatId int64, userType user.Type, chainName string) error {
	_, ctx := connect()
	var userDto = getOrCreateUser(chatId, userType)
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

func GetChainsForUser(chatId int64, userType user.Type) []common.Chain {
	client, ctx := connect()
	var userDto = getOrCreateUser(chatId, userType)
	chainsOfUser, err := client.Chain.
		Query().
		Where(chain.HasUsersWith(user.ID(userDto.ID))).
		All(ctx)
	if err != nil {
		log.Sugar.Panic("Error while fetching chains for user %v: %v", userDto.ID, err)
	}
	allChains, err := client.Chain.
		Query().
		Where(chain.IsEnabledEQ(true)).
		Order(ent.Asc(chain.FieldDisplayName)).
		All(ctx)
	var chains []common.Chain
	for _, c := range allChains {
		var chainEntry = common.Chain{Name: c.Name, DisplayName: c.DisplayName, Notify: false}
		for _, nc := range chainsOfUser { // check if user gets notified for this chain (c)
			if nc.ID == c.ID {
				chainEntry.Notify = true
			}
		}
		chains = append(chains, chainEntry)
	}
	return chains
}

func CreateChain(chainName string, rpcs []string) *ent.Chain {
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
			SetIsEnabled(false).
			Save(ctx)
		if err != nil {
			log.Sugar.Panic("Error while creating chain: %v", err)
		}
		for _, rpc := range rpcs {
			_, err := client.RpcEndpoint.
				Create().
				SetEndpoint(rpc).
				SetChain(c).
				Save(ctx)
			if err != nil {
				log.Sugar.Panic("Error while creating chain: %v", err)
			}
		}
	}
	return c
}

func UpdateRpcs(chainName string, rpcs []string) error {
	client, ctx := connect()
	c, err := client.Chain.
		Query().
		Where(chain.NameEQ(chainName)).
		WithRPCEndpoints().
		Only(ctx)
	if err != nil {
		return err
	}
	_, err = client.RpcEndpoint.
		Delete().
		Where(rpcendpoint.HasChainWith(chain.IDEQ(c.ID))).
		Exec(ctx)
	if err != nil {
		return err
	}
	for _, rpc := range rpcs {
		_, err := client.RpcEndpoint.
			Create().
			SetEndpoint(rpc).
			SetChain(c).
			Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetChains() []*ent.Chain {
	client, ctx := connect()
	chains, err := client.Chain.
		Query().
		Order(ent.Asc(chain.FieldDisplayName)).
		All(ctx)
	if err != nil {
		log.Sugar.Panic("Error while querying chains: %v", err)
	}
	return chains
}

func GetChainStatistics() (*[]common.ChainStatistic, error) {
	client, ctx := connect()
	var chainsWithNotifications []common.ChainStatistic
	err := client.Chain.Query().
		Order(ent.Desc(chain.FieldIsEnabled), ent.Asc(chain.FieldDisplayName)).
		GroupBy(chain.FieldIsEnabled, chain.FieldDisplayName).
		Aggregate(
			func(s *sql.Selector) string {
				t := sql.Table(chain.UsersTable)
				s.Join(t).On(s.C(chain.FieldID), t.C(user.ChainsPrimaryKey[1]))
				return sql.As(sql.Count(t.C(user.ChainsPrimaryKey[1])), "subscriptions")
			},
		).
		Scan(ctx, &chainsWithNotifications)
	if err != nil {
		return nil, err
	}
	var chainsWithProposals []common.ChainStatistic
	err = client.Chain.Query().
		Order(ent.Desc(chain.FieldIsEnabled), ent.Asc(chain.FieldDisplayName)).
		GroupBy(chain.FieldIsEnabled, chain.FieldDisplayName).
		Aggregate(
			func(s *sql.Selector) string {
				t := sql.Table(chain.ProposalsTable)
				s.Join(t).On(s.C(chain.FieldID), t.C(proposal.ChainColumn))
				return sql.As(sql.Count(t.C(proposal.FieldID)), "proposals")
			},
		).
		Scan(ctx, &chainsWithProposals)
	if err != nil {
		return nil, err
	}
	var stats []common.ChainStatistic
	for _, cp := range chainsWithProposals {
		found := false
		for _, cn := range chainsWithNotifications {
			if cp.DisplayName == cn.DisplayName {
				stats = append(stats, common.ChainStatistic{
					DisplayName:   cp.DisplayName,
					Proposals:     cp.Proposals,
					Subscriptions: cn.Subscriptions,
				})
				found = true
			}
		}
		if !found {
			stats = append(stats, common.ChainStatistic{
				DisplayName:   cp.DisplayName,
				Proposals:     cp.Proposals,
				Subscriptions: 0,
			})
		}
	}
	return &stats, err
}

func EnableOrDisableChain(chainName string) error {
	_, ctx := connect()
	chainDto, err := getChainByName(chainName)
	if err != nil {
		return err
	}
	_, err = chainDto.
		Update().
		SetIsEnabled(!chainDto.IsEnabled).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
