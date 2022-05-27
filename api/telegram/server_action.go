package telegram

import (
	"errors"
	"fmt"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shifty11/cosmos-gov/authz"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/log"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

func shouldDeleteUser(err error) bool {
	if err != nil {
		return slices.Contains(forbiddenErrors, err.Error())
	}
	return false
}

func toSlice(res []database.TgChatQueryResult) []int64 {
	var ints []int64
	for _, entry := range res {
		ints = append(ints, entry.ChatId)
	}
	return ints
}

func unpackVoteData(voteDataStr string) (*authz.VoteData, error) {
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

func SendProposals(entProp *ent.Proposal, entChain *ent.Chain) []int64 {
	text := fmt.Sprintf("🎉  <b>%v - Proposal %v\n\n%v</b>\n\n<i>%v</i>", entChain.DisplayName, entProp.ProposalID, entProp.Title, entProp.Description)
	if len(text) > 4096 {
		text = text[:4088] + "</i> ..."
	}

	var errIds []int64
	mHack = database.NewDefaultDbManagers() //TODO: remove

	allTgChats := mHack.TelegramChatManager.GetChatIds(entChain)
	tgChatsWithGrants := mHack.TelegramChatManager.GetChatIdsWithGrants(entChain)
	for _, chat := range allTgChats {
		log.Sugar.Debugf("Send proposal #%v on %v to telegram chat #%v", entProp.ProposalID, entChain.DisplayName, chat.ChatId)

		msg := tgbotapi.NewMessage(chat.ChatId, text)
		msg.ParseMode = "html"
		msg.DisableWebPagePreview = true

		if slices.Contains(toSlice(tgChatsWithGrants), chat.ChatId) {
			voteData := authz.ToVoteData(entChain.Name, entProp.ProposalID, govtypes.OptionEmpty, authz.NotVoted)
			msg.ReplyMarkup = createKeyboard(getVoteButtons(&voteData))
		}

		err := sendMessage(msg)
		if err != nil {
			if shouldDeleteUser(err) {
				errIds = append(errIds, chat.ChatId)
			} else {
				log.Sugar.Errorf("Error while sending message to telegram chat #%v: %v", chat, err)
			}
		}
	}
	return errIds
}

func SendDraftProposals(entProp *ent.DraftProposal, entChain *ent.Chain) []int64 {
	text := fmt.Sprintf("💬  <b>%v - New pre-vote proposal\n\n%v</b>\n%v", entChain.DisplayName, entProp.Title, entProp.URL)

	var errIds []int64
	mHack = database.NewDefaultDbManagers() //TODO: remove

	allTgChats := mHack.TelegramChatManager.GetChatIds(entChain)
	for _, chat := range allTgChats {
		log.Sugar.Debugf("Send draft proposal #%v on %v to telegram chat #%v", entProp.DraftProposalID, entChain.DisplayName, chat.ChatId)

		msg := tgbotapi.NewMessage(chat.ChatId, text)
		msg.ParseMode = "html"
		msg.DisableWebPagePreview = true

		err := sendMessage(msg)
		if err != nil {
			if shouldDeleteUser(err) {
				errIds = append(errIds, chat.ChatId)
			} else {
				log.Sugar.Errorf("Error while sending message to telegram chat #%v: %v", chat, err)
			}
		}
	}
	return errIds
}
