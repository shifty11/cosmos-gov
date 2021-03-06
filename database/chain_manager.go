package database

import (
	"context"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/rpcendpoint"
	"github.com/shifty11/cosmos-gov/log"
	lens "github.com/strangelove-ventures/lens/client"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
)

var caser = cases.Title(language.English)

type ChainManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewChainManager(client *ent.Client, ctx context.Context) *ChainManager {
	return &ChainManager{client: client, ctx: ctx}
}

type ChainQueryOptions struct {
	//WithWallets      bool
	//WithGrants       bool
	WithRpcAddresses bool
}

func (manager *ChainManager) ByName(name string) (*ent.Chain, error) {
	return manager.client.Chain.
		Query().
		Where(chain.NameEQ(name)).
		Only(manager.ctx)
}

func (manager *ChainManager) Enabled(options *ChainQueryOptions) []*ent.Chain {
	query := manager.client.Chain.
		Query().
		Where(chain.IsEnabledEQ(true)).
		Order(ent.Asc(chain.FieldDisplayName))
	if options != nil && options.WithRpcAddresses {
		query = query.WithRPCEndpoints()
	}
	allChains, err := query.All(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying enabled chains: %v", err)
	}
	return allChains
}

func (manager *ChainManager) CanVote() []*ent.Chain {
	chains, err := manager.client.Chain.
		Query().
		Where(chain.And(
			chain.IsEnabledEQ(true),
			chain.IsVotingEnabled(true),
		)).
		WithRPCEndpoints().
		Order(ent.Asc(chain.FieldDisplayName)).
		All(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying enabled chains: %v", err)
	}
	return chains
}

func (manager *ChainManager) All() []*ent.Chain {
	chains, err := manager.client.Chain.
		Query().
		Order(ent.Asc(chain.FieldDisplayName)).
		All(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying chains: %v", err)
	}
	return chains
}

func (manager *ChainManager) EnableOrDisableChain(chainName string) (*ent.Chain, error) {
	chainDto, err := manager.ByName(chainName)
	if err != nil {
		return nil, err
	}
	return chainDto.
		Update().
		SetIsEnabled(!chainDto.IsEnabled).
		Save(manager.ctx)
}

func (manager *ChainManager) Update(name string, isEnabled bool, isVotingEnabled bool, isFeegrantUsed bool) (*ent.Chain, error) {
	entChain, err := manager.ByName(name)
	if err != nil {
		return nil, err
	}
	return manager.client.Chain.
		UpdateOne(entChain).
		SetIsEnabled(isEnabled).
		SetIsVotingEnabled(isVotingEnabled).
		SetIsFeegrantUsed(isFeegrantUsed).
		Save(manager.ctx)
}

func (manager *ChainManager) Create(chainId string, chainName string, accountPrefix string, rpcs []string) *ent.Chain {
	log.Sugar.Infof("Create new chain: %v", chainName)
	c, err := manager.client.Chain.
		Create().
		SetChainID(chainId).
		SetAccountPrefix(accountPrefix).
		SetName(chainName).
		SetDisplayName(caser.String(chainName)).
		SetIsEnabled(false).
		Save(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while creating chain: %v", err)
	}
	for _, rpc := range rpcs {
		_, err := manager.client.RpcEndpoint.
			Create().
			SetEndpoint(rpc).
			SetChain(c).
			Save(manager.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while creating rpc endpoint: %v", err)
		}
	}
	return c
}

func (manager *ChainManager) UpdateRpcs(chainName string, rpcs []string) error {
	c, err := manager.client.Chain.
		Query().
		Where(chain.NameEQ(chainName)).
		WithRPCEndpoints().
		Only(manager.ctx)
	if err != nil {
		return err
	}
	_, err = manager.client.RpcEndpoint.
		Delete().
		Where(rpcendpoint.HasChainWith(chain.IDEQ(c.ID))).
		Exec(manager.ctx)
	if err != nil {
		return err
	}
	for _, rpc := range rpcs {
		_, err := manager.client.RpcEndpoint.
			Create().
			SetEndpoint(rpc).
			SetChain(c).
			Save(manager.ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (manager *ChainManager) GetFirstRpc(entChain *ent.Chain) *ent.RpcEndpoint {
	rpc, err := entChain.
		QueryRPCEndpoints().
		First(context.Background())
	if err != nil {
		log.Sugar.Panicf("Error while getting firs rpc of chain %v: %v", entChain.Name, err)
	}
	return rpc
}

func (manager *ChainManager) BuildLensClient(entChain *ent.Chain) (*lens.ChainClient, error) {
	rpc := manager.GetFirstRpc(entChain)

	key_dir := os.Getenv("LENS_PATH")
	if key_dir == "" {
		log.Sugar.Fatalf("LENS_PATH env var must be set")
	}

	chainConfig := lens.ChainClientConfig{
		Key:            "default",
		ChainID:        entChain.ChainID,
		RPCAddr:        rpc.Endpoint,
		AccountPrefix:  entChain.AccountPrefix,
		KeyringBackend: "test",
		Debug:          true,
		Timeout:        "20s",
		Modules:        lens.ModuleBasics,
	}

	chainClient, err := lens.NewChainClient(log.Sugar.Desugar(), &chainConfig, key_dir, os.Stdin, os.Stdout)
	if err != nil {
		log.Sugar.Fatalf("Failed to build new chain client for %s. Err: %v \n", entChain.DisplayName, err)
	}
	return chainClient, nil
}
