package database

import (
	"context"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/draftproposal"
	"github.com/shifty11/cosmos-gov/log"
)

type DraftProposalManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewDraftProposalManager(client *ent.Client, ctx context.Context) *DraftProposalManager {
	return &DraftProposalManager{client: client, ctx: ctx}
}

func (manager *DraftProposalManager) ByChain(chainName string) ([]*ent.DraftProposal, error) {
	return manager.client.DraftProposal.
		Query().
		Where(draftproposal.HasChainWith(chain.NameEQ(chainName))).
		All(manager.ctx)
}

func (manager *DraftProposalManager) CreateIfNotExists(entChain *ent.Chain, draftProposalId int64, title string, url string) *ent.DraftProposal {
	exists, err := entChain.
		QueryDraftProposals().
		Where(draftproposal.DraftProposalID(draftProposalId)).
		Exist(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while checking for draft proposal #%v: %v", draftProposalId, err)
	}
	if exists {
		return nil
	}
	entProp, err := manager.client.DraftProposal.
		Create().
		SetChain(entChain).
		SetDraftProposalID(draftProposalId).
		SetTitle(title).
		SetURL(url).
		Save(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while creating draft proposal #%v: %v", draftProposalId, err)
	}
	return entProp
}
