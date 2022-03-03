package database

import (
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/lenschaininfo"
	"github.com/shifty11/cosmos-gov/log"
)

func AddErrorToLensChainInfo(chainName string) {
	client, ctx := connect()
	entity, err := client.LensChainInfo.
		Query().
		Where(lenschaininfo.NameEQ(chainName)).
		Only(ctx)
	if err != nil {
		log.Sugar.Infof("Create new LensChainInfo for %v", chainName)
		entity, err = client.LensChainInfo.
			Create().
			SetName(chainName).
			SetCntErrors(1).
			Save(ctx)
		if err != nil {
			log.Sugar.Panic("Error while creating LensChainInfo: %v", err)
		}
	} else {
		log.Sugar.Infof("Update LensChainInfo for %v; Errors: %v", chainName, entity.CntErrors+1)
		entity, err = entity.
			Update().
			SetCntErrors(entity.CntErrors + 1).
			Save(ctx)
		if err != nil {
			log.Sugar.Panic("Error while creating LensChainInfo: %v", err)
		}
	}
}

func GetLensChainInfos() []*ent.LensChainInfo {
	client, ctx := connect()
	chains, err := client.LensChainInfo.
		Query().
		All(ctx)
	if err != nil {
		log.Sugar.Panic("Error while querying chains: %v", err)
	}
	return chains
}

func DeleteLensChainInfo(chainName string) {
	client, ctx := connect()
	_, err := client.LensChainInfo.
		Delete().
		Where(lenschaininfo.NameEQ(chainName)).
		Exec(ctx)
	if err != nil {
		log.Sugar.Errorf("Error while deleting LensChainInfo: %v", err)
	}
}
