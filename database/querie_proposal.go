package database

import (
	"github.com/shifty11/cosmos-gov/dtos"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/proposal"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
)

func CreateProposalIfNotExists(prop *dtos.Proposal, chainDb *ent.Chain) *ent.Proposal {
	client, ctx := connect()
	exist, err := chainDb.QueryProposals().
		Where(proposal.ProposalIDEQ(prop.ProposalId)).
		Exist(ctx)
	if err != nil {
		log.Sugar.Panicf("Error while checking for proposal #%v: %v", prop.ProposalId, err)
	}
	if !exist {
		propDb, err := client.Proposal.
			Create().
			SetProposalID(prop.ProposalId).
			SetTitle(prop.Content.Title).
			SetDescription(prop.Content.Description).
			SetVotingStartTime(prop.VotingStartTime).
			SetVotingEndTime(prop.VotingEndTime).
			SetChainID(chainDb.ID).
			SetStatus(prop.Status).
			Save(ctx)
		if err != nil {
			log.Sugar.Panicf("Error while creating proposal #%v: %v", prop.ProposalId, err)
		}
		return propDb
	}
	return nil
}

func GetUsers(chainDb *ent.Chain) []int {
	_, ctx := connect()
	chatIds, err := chainDb.
		QueryUsers().
		Select(user.FieldChatID).
		Ints(ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying chatIds for chain %v: %v", chainDb.Name, err)
	}
	return chatIds
}

func DropProposals() {
	client, ctx := connect()
	client.Proposal.
		Delete().
		ExecX(ctx)
}
