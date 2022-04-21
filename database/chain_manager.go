package database

import (
	"context"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/log"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var caser = cases.Title(language.English)

type ChainManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewChainManager() *ChainManager {
	client, ctx := connect()
	return &ChainManager{client: client, ctx: ctx}
}

func (manager *ChainManager) ByName(name string) (*ent.Chain, error) {
	return manager.client.Chain.
		Query().
		Where(chain.NameEQ(name)).
		Only(manager.ctx)
}

func (manager *ChainManager) Enabled() []*ent.Chain {
	allChains, err := manager.client.Chain.
		Query().
		Where(chain.IsEnabledEQ(true)).
		Order(ent.Asc(chain.FieldDisplayName)).
		All(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying enabled chains: %v", err)
	}
	return allChains
}

func (manager *ChainManager) All() []*ent.Chain {
	chains, err := manager.client.Chain.
		Query().
		Order(ent.Asc(chain.FieldDisplayName)).
		All(manager.ctx)
	if err != nil {
		log.Sugar.Panic("Error while querying chains: %v", err)
	}
	return chains
}

func (manager *ChainManager) EnableOrDisableChain(chainName string) error {
	chainDto, err := manager.ByName(chainName)
	if err != nil {
		return err
	}
	_, err = chainDto.
		Update().
		SetIsEnabled(!chainDto.IsEnabled).
		Save(manager.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (manager *ChainManager) Create(chainName string) *ent.Chain {
	c, err := manager.client.Chain.
		Query().
		Where(chain.NameEQ(chainName)).
		Only(manager.ctx)
	if err != nil {
		log.Sugar.Infof("Create new chain: %v", chainName)
		c, err = manager.client.Chain.
			Create().
			SetName(chainName).
			SetDisplayName(caser.String(chainName)).
			SetIsEnabled(false).
			Save(manager.ctx)
		if err != nil {
			log.Sugar.Panic("Error while creating chain: %v", err)
		}
	}
	return c
}
