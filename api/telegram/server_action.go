package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shifty11/cosmos-gov/common"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/log"
)

func shouldDeleteUser(err error) bool {
	if err != nil {
		return common.Contains(forbiddenErrors, err.Error())
	}
	return false
}

func SendProposals(entProp *ent.Proposal, entChain *ent.Chain) []int64 {
	text := fmt.Sprintf("🎉  <b>%v - Proposal %v\n\n%v</b>\n\n<i>%v</i>", entChain.DisplayName, entProp.ProposalID, entProp.Title, entProp.Description)
	if len(text) > 4096 {
		text = text[:4088] + "</i> ..."
	}

	var errIds []int64
	chatIds := database.NewTelegramChatManager().GetChatIds(entChain)
	for _, chatId := range chatIds {
		msg := tgbotapi.NewMessage(int64(chatId), text)
		msg.ParseMode = "html"
		msg.DisableWebPagePreview = true
		log.Sugar.Debugf("Send proposal #%v on %v to telegram chat #%v", entProp.ProposalID, entChain.DisplayName, chatId)
		err := sendMessage(msg)
		if err != nil {
			if shouldDeleteUser(err) {
				errIds = append(errIds, int64(chatId))
			} else {
				log.Sugar.Errorf("Error while sending message to telegram chat #%v: %v", chatId, err)
			}
		}
	}
	return errIds
}
