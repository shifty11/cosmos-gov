package database

import (
	"errors"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/shifty11/cosmos-gov/dtos"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/proposal"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
)

// CreateProposalIfNotExists creates a proposal if it does not exist. If it exists it doesn't do anything.
// returns new proposal or nil if it already exists.
func CreateProposalIfNotExists(prop *dtos.Proposal, chainDb *ent.Chain) *ent.Proposal {
	client, ctx := connect()
	exist, err := chainDb.QueryProposals().
		Where(proposal.And(proposal.ProposalIDEQ(prop.ProposalId))).
		Exist(ctx)
	if err != nil {
		log.Sugar.Panicf("Error while checking for proposal #%v: %v", prop.ProposalId, err)
	}
	if !exist {
		log.Sugar.Debugf("Save proposal #%v on chain %v", prop.ProposalId, chainDb.DisplayName)
		status, err := types.ProposalStatusFromString(prop.Status)
		if err != nil {
			log.Sugar.Panicf("Error while reading proposal status of proposal #%v: %v", prop.ProposalId, err)
		}
		propDb, err := client.Proposal.
			Create().
			SetProposalID(prop.ProposalId).
			SetTitle(prop.Content.Title).
			SetDescription(prop.Content.Description).
			SetVotingStartTime(prop.VotingStartTime).
			SetVotingEndTime(prop.VotingEndTime).
			SetChainID(chainDb.ID).
			SetStatus(proposal.Status(status.String())).
			Save(ctx)
		if err != nil {
			log.Sugar.Panicf("Error while creating proposal #%v: %v", prop.ProposalId, err)
		}
		return propDb
	}
	return nil
}

func CreateOrUpdateProposal(prop *dtos.Proposal, chainDb *ent.Chain) *ent.Proposal {
	status, err := types.ProposalStatusFromString(prop.Status)
	if err != nil {
		log.Sugar.Panicf("Error while reading proposal status of proposal #%v: %v", prop.ProposalId, err)
	}

	client, ctx := connect()
	propDb, err := chainDb.QueryProposals().
		Where(proposal.ProposalIDEQ(prop.ProposalId)).
		Only(ctx)
	notFoundError := &ent.NotFoundError{}
	if err != nil {
		if errors.As(err, &notFoundError) {
			log.Sugar.Debugf("Save proposal #%v on chain %v", prop.ProposalId, chainDb.DisplayName)

			propDb, err := client.Proposal.
				Create().
				SetProposalID(prop.ProposalId).
				SetTitle(prop.Content.Title).
				SetDescription(prop.Content.Description).
				SetVotingStartTime(prop.VotingStartTime).
				SetVotingEndTime(prop.VotingEndTime).
				SetChainID(chainDb.ID).
				SetStatus(proposal.Status(status.String())).
				Save(ctx)
			if err != nil {
				log.Sugar.Panicf("Error while creating proposal #%v: %v", prop.ProposalId, err)
			}
			return propDb
		} else {
			log.Sugar.Panicf("Error while checking for proposal #%v: %v", prop.ProposalId, err)
		}
	}

	log.Sugar.Debugf("Update proposal #%v on chain %v", prop.ProposalId, chainDb.DisplayName)
	newProp, err := propDb.
		Update().
		SetTitle(prop.Content.Title).
		SetDescription(prop.Content.Description).
		SetVotingStartTime(prop.VotingStartTime).
		SetVotingEndTime(prop.VotingEndTime).
		SetStatus(proposal.Status(status.String())).
		Save(ctx)
	if err != nil {
		log.Sugar.Panicf("Error while updating proposal #%v: %v", prop.ProposalId, err)
	}
	return newProp
}

func GetProposalsInVotingPeriod(chainName string) []*ent.Proposal {
	client, ctx := connect()
	props, err := client.Proposal.
		Query().
		Where(proposal.And(
			proposal.HasChainWith(chain.NameEQ(chainName)),
			proposal.StatusEQ(proposal.StatusPROPOSAL_STATUS_VOTING_PERIOD),
		)).
		All(ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying first/second proposal for chain %v: %v", chainName, err)
	}
	return props
}

func GetProposalsInVotingPeriodForUser(chatId int64) []*ent.Chain {
	client, ctx := connect()
	props, err := client.Chain.
		Query().
		Where(chain.And(
			chain.HasUsersWith(user.ChatIDEQ(chatId)),
		)).
		Order(ent.Asc(chain.FieldName)).
		WithProposals(func(q *ent.ProposalQuery) {
			q.Where(proposal.StatusEQ(proposal.StatusPROPOSAL_STATUS_VOTING_PERIOD))
		}).
		All(ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying proposals for user #%v: %v", chatId, err)
	}
	return props
}
