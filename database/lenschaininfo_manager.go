package database

import (
	"context"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/lenschaininfo"
	"github.com/shifty11/cosmos-gov/log"
)

type LensChainInfoManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewLensChainInfoManager(client *ent.Client, ctx context.Context) *LensChainInfoManager {
	return &LensChainInfoManager{client: client, ctx: ctx}
}

func (manager *LensChainInfoManager) AddErrorToLensChainInfo(chainName string) {
	entity, err := manager.client.LensChainInfo.
		Query().
		Where(lenschaininfo.NameEQ(chainName)).
		Only(manager.ctx)
	if err != nil {
		log.Sugar.Infof("Create new LensChainInfo for %v", chainName)
		entity, err = manager.client.LensChainInfo.
			Create().
			SetName(chainName).
			SetCntErrors(1).
			Save(manager.ctx)
		if err != nil {
			log.Sugar.Panic("Error while creating LensChainInfo: %v", err)
		}
	} else {
		log.Sugar.Infof("Update LensChainInfo for %v; Errors: %v", chainName, entity.CntErrors+1)
		entity, err = entity.
			Update().
			SetCntErrors(entity.CntErrors + 1).
			Save(manager.ctx)
		if err != nil {
			log.Sugar.Panic("Error while creating LensChainInfo: %v", err)
		}
	}
}

func (manager *LensChainInfoManager) GetLensChainInfos() []*ent.LensChainInfo {
	chains, err := manager.client.LensChainInfo.
		Query().
		All(manager.ctx)
	if err != nil {
		log.Sugar.Panic("Error while querying chains: %v", err)
	}
	return chains
}

func (manager *LensChainInfoManager) DeleteLensChainInfo(chainName string) {
	_, err := manager.client.LensChainInfo.
		Delete().
		Where(lenschaininfo.NameEQ(chainName)).
		Exec(manager.ctx)
	if err != nil {
		log.Sugar.Errorf("Error while deleting LensChainInfo: %v", err)
	}
}
