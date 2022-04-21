package database

import (
	"context"
	"errors"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/shifty11/cosmos-gov/common"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/discordchannel"
	"github.com/shifty11/cosmos-gov/ent/proposal"
	"github.com/shifty11/cosmos-gov/ent/telegramchat"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
	"time"
)

type ProposalManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewProposalManager() *ProposalManager {
	client, ctx := connect()
	return &ProposalManager{client: client, ctx: ctx}
}

// CreateIfNotExists creates a proposal if it does not exist. If it exists it doesn't do anything.
// returns new proposal or nil if it already exists.
func (manager *ProposalManager) CreateIfNotExists(prop *common.Proposal, entChain *ent.Chain) *ent.Proposal {
	exist, err := entChain.QueryProposals().
		Where(proposal.And(proposal.ProposalIDEQ(prop.ProposalId))).
		Exist(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while checking for proposal #%v: %v", prop.ProposalId, err)
	}
	if !exist {
		log.Sugar.Debugf("Save proposal #%v on chain %v", prop.ProposalId, entChain.DisplayName)
		status, err := types.ProposalStatusFromString(prop.Status)
		if err != nil {
			log.Sugar.Panicf("Error while reading proposal status of proposal #%v: %v", prop.ProposalId, err)
		}
		propDb, err := manager.client.Proposal.
			Create().
			SetProposalID(prop.ProposalId).
			SetTitle(prop.Content.Title).
			SetDescription(prop.Content.Description).
			SetVotingStartTime(prop.VotingStartTime).
			SetVotingEndTime(prop.VotingEndTime).
			SetChainID(entChain.ID).
			SetStatus(proposal.Status(status.String())).
			Save(manager.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while creating proposal #%v: %v", prop.ProposalId, err)
		}
		return propDb
	}
	return nil
}

func (manager *ProposalManager) CreateOrUpdateProposal(prop *common.Proposal, chainDb *ent.Chain) *ent.Proposal {
	status, err := types.ProposalStatusFromString(prop.Status)
	if err != nil {
		log.Sugar.Panicf("Error while reading proposal status of proposal #%v: %v", prop.ProposalId, err)
	}

	entProp, err := chainDb.QueryProposals().
		Where(proposal.ProposalIDEQ(prop.ProposalId)).
		Only(manager.ctx)
	notFoundError := &ent.NotFoundError{}
	if err != nil {
		if errors.As(err, &notFoundError) {
			log.Sugar.Debugf("Save proposal #%v on chain %v", prop.ProposalId, chainDb.DisplayName)

			entProp, err := manager.client.Proposal.
				Create().
				SetProposalID(prop.ProposalId).
				SetTitle(prop.Content.Title).
				SetDescription(prop.Content.Description).
				SetVotingStartTime(prop.VotingStartTime).
				SetVotingEndTime(prop.VotingEndTime).
				SetChainID(chainDb.ID).
				SetStatus(proposal.Status(status.String())).
				Save(manager.ctx)
			if err != nil {
				log.Sugar.Panicf("Error while creating proposal #%v: %v", prop.ProposalId, err)
			}
			return entProp
		} else {
			log.Sugar.Panicf("Error while checking for proposal #%v: %v", prop.ProposalId, err)
		}
	}

	log.Sugar.Debugf("Update proposal #%v on chain %v", prop.ProposalId, chainDb.DisplayName)
	newProp, err := entProp.
		Update().
		SetTitle(prop.Content.Title).
		SetDescription(prop.Content.Description).
		SetVotingStartTime(prop.VotingStartTime).
		SetVotingEndTime(prop.VotingEndTime).
		SetStatus(proposal.Status(status.String())).
		Save(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while updating proposal #%v: %v", prop.ProposalId, err)
	}
	return newProp
}

func (manager *ProposalManager) GetFinishedProposalsInVotingPeriod() []*ent.Proposal {
	props, err := manager.client.Proposal.
		Query().
		Where(proposal.And(
			proposal.StatusEQ(proposal.StatusPROPOSAL_STATUS_VOTING_PERIOD),
			proposal.VotingEndTimeLT(time.Now()),
		)).
		WithChain().
		All(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying finished proposals in voting period: %v", err)
	}
	return props
}

func (manager *ProposalManager) GetProposalsInVotingPeriod(chatOrChannelId int64, userType user.Type) []*ent.Chain {
	if userType == user.TypeTelegram {
		props, err := manager.client.Chain.
			Query().
			Where(chain.HasTelegramChatsWith(telegramchat.IDEQ(chatOrChannelId))).
			Order(ent.Asc(chain.FieldName)).
			WithProposals(func(q *ent.ProposalQuery) {
				q.Where(proposal.StatusEQ(proposal.StatusPROPOSAL_STATUS_VOTING_PERIOD))
			}).
			All(manager.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while querying proposals for Telegram chat #%v: %v", chatOrChannelId, err)
		}
		return props
	} else {
		props, err := manager.client.Chain.
			Query().
			Where(chain.HasDiscordChannelsWith(discordchannel.IDEQ(chatOrChannelId))).
			Order(ent.Asc(chain.FieldName)).
			WithProposals(func(q *ent.ProposalQuery) {
				q.Where(proposal.StatusEQ(proposal.StatusPROPOSAL_STATUS_VOTING_PERIOD))
			}).
			All(manager.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while querying proposals for Discrod channel #%v: %v", chatOrChannelId, err)
		}
		return props
	}
}
