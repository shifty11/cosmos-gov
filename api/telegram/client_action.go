package telegram

import (
	"fmt"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shifty11/cosmos-gov/authz"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
)

// performUpdateSubscription toggles the subscription for a chain
func performUpdateSubscription(update *tgbotapi.Update, chainName string) {
	if chainName == "" {
		return
	}
	tgChatId := getChatIdX(update)
	chatName := getChatName(update)
	log.Sugar.Debugf("Toggle subscription %v for Telegram chat %v (%v)", chainName, chatName, tgChatId)

	_, err := mHack.TelegramChatManager.AddOrRemoveChain(tgChatId, chainName)
	if err != nil {
		log.Sugar.Errorf("Error while toggle subscription %v for Telegram chat %v (%v)", chainName, chatName, tgChatId)
	}
}

func executeVote(update *tgbotapi.Update, entUser *ent.User, voteData *authz.VoteData) {
	authzClient := authz.NewAuthzClient(mHack.ChainManager, mHack.WalletManager)
	err := authzClient.ExecAuthzVote(entUser, voteData)
	if err != nil {
		log.Sugar.Errorf("while getting executing vote per authz %v", err)
		errMsg := err.Error()
		vote, err := authzClient.GetVoteStatus(entUser, voteData.ChainName, voteData.ProposalId)
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
		editSentProposal(update, voteData)

		text := fmt.Sprintf("🤯 <b>Error</b>\n\nThere was an error while performing your vote. Please try again or vote with another tool.\n\nThe error was: %v", errMsg)
		msg := tgbotapi.NewMessage(getChatIdX(update), text)
		msg.ParseMode = "html"
		err = sendMessage(msg)
		if err != nil {
			log.Sugar.Errorf("while sending error message to user %v (%v): %v", entUser.Name, entUser.UserID, err)
		}
	} else {
		voteData.State = authz.Voted
		editSentProposal(update, voteData)
	}
}

func performVote(update *tgbotapi.Update, data string) {
	voteData, err := unpackVoteData(data)
	if err != nil {
		log.Sugar.Errorf("while unpacking VoteData: %v", err)
	}
	entUser, err := mHack.UserManager.Get(getUserIdX(update), user.TypeTelegram)
	if err != nil {
		log.Sugar.Errorf("while getting user %v", err)
	}
	voteData.State = authz.Voting

	go executeVote(update, entUser, voteData)
	editSentProposal(update, voteData)
}
