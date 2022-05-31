package telegram

import (
	"errors"
	"fmt"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shifty11/cosmos-gov/authz"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
	"strconv"
	"strings"
)

// performUpdateSubscription toggles the subscription for a chain
func (client TelegramClient) performUpdateSubscription(update *tgbotapi.Update, chainName string) {
	if chainName == "" {
		return
	}
	tgChatId := getChatIdX(update)
	chatName := getChatName(update)
	log.Sugar.Debugf("Toggle subscription %v for Telegram chat %v (%v)", chainName, chatName, tgChatId)

	_, err := client.TelegramChatManager.AddOrRemoveChain(tgChatId, chainName)
	if err != nil {
		log.Sugar.Errorf("Error while toggle subscription %v for Telegram chat %v (%v)", chainName, chatName, tgChatId)
	}
}

func (client TelegramClient) executeVote(update *tgbotapi.Update, entUser *ent.User, voteData *authz.VoteData) {
	err := client.AuthzClient.ExecAuthzVote(entUser, voteData)
	if err != nil {
		log.Sugar.Errorf("while getting executing vote per authz %v", err)
		errMsg := err.Error()
		vote, err := client.AuthzClient.GetVoteStatus(entUser, voteData.ChainName, voteData.ProposalId)
		if err != nil {
			log.Sugar.Errorf("Could not get vote of proposal %v on %v for user %v (%v)",
				voteData.ProposalId, voteData.ChainName, entUser.Name, entUser.UserID)
		}
		if vote == govtypes.OptionEmpty {
			voteData.State = authz.NotVoted
		} else {
			voteData.State = authz.Voted
			voteData.Vote = vote
		}
		client.editSentProposal(update, voteData)

		text := fmt.Sprintf("🤯 <b>Error</b>\n\nThere was an error while performing your vote. Please try again or vote with another tool.\n\nThe error was: %v", errMsg)
		msg := tgbotapi.NewMessage(getChatIdX(update), text)
		msg.ParseMode = "html"
		err = sendMessage(msg)
		if err != nil {
			log.Sugar.Errorf("while sending error message to user %v (%v): %v", entUser.Name, entUser.UserID, err)
		}
	} else {
		voteData.State = authz.Voted
		client.editSentProposal(update, voteData)
	}
}

func (client TelegramClient) performVote(update *tgbotapi.Update, data string) {
	voteData, err := client.unpackVoteData(data)
	if err != nil {
		log.Sugar.Errorf("while unpacking VoteData: %v", err)
	}
	entUser, err := client.UserManager.Get(getUserIdX(update), user.TypeTelegram)
	if err != nil {
		log.Sugar.Errorf("while getting user %v", err)
	}
	voteData.State = authz.Voting

	go client.executeVote(update, entUser, voteData)
	client.editSentProposal(update, voteData)
}

func (client TelegramClient) unpackVoteData(voteDataStr string) (*authz.VoteData, error) {
	var parts = strings.Split(voteDataStr, ":")
	if len(parts) != 4 {
		return nil, errors.New(fmt.Sprintf("expected 3 parts but got %v", len(parts)))
	}
	proposalId, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%d of type %T", proposalId, proposalId))
	}
	voteOption, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%d of type %T", voteOption, voteOption))
	}
	voteState, err := strconv.ParseInt(parts[3], 10, 64)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%d of type %T", voteState, voteState))
	}
	return &authz.VoteData{
		ChainName:  parts[0],
		ProposalId: proposalId,
		Vote:       govtypes.VoteOption(voteOption),
		State:      authz.VoteState(voteState),
	}, nil
}
