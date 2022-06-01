package telegram

import (
	"fmt"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shifty11/cosmos-gov/authz"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/log"
	"golang.org/x/exp/slices"
	"os"
	"strconv"
	"strings"
)

//goland:noinspection GoNameStartsWithPackageName
type TelegramLightClient struct {
	api                 *tgbotapi.BotAPI
	TelegramChatManager *database.TelegramChatManager
}

func NewTelegramLightClient(managers *database.DbManagers) *TelegramLightClient {
	return &TelegramLightClient{
		api: getApi(),

		TelegramChatManager: managers.TelegramChatManager,
	}
}

func (client TelegramLightClient) shouldDeleteUser(err error) bool {
	if err != nil {
		return slices.Contains(forbiddenErrors, err.Error())
	}
	return false
}

func (client TelegramLightClient) toSlice(res []database.TgChatQueryResult) []int64 {
	var ints []int64
	for _, entry := range res {
		ints = append(ints, entry.ChatId)
	}
	return ints
}

func (client TelegramLightClient) SendProposals(entProp *ent.Proposal, entChain *ent.Chain) []int64 {
	text := fmt.Sprintf("🎉  <b>%v - Proposal %v\n\n%v</b>\n\n<i>%v</i>", entChain.DisplayName, entProp.ProposalID, entProp.Title, entProp.Description)
	if len(text) > 4096 {
		text = text[:4088] + "</i> ..."
	}

	var errIds []int64
	allTgChats := client.TelegramChatManager.GetChatIds(entChain)
	tgChatsWithGrants := client.TelegramChatManager.GetChatIdsWithGrants(entChain)
	for _, chat := range allTgChats {
		log.Sugar.Debugf("Send proposal #%v on %v to telegram chat #%v", entProp.ProposalID, entChain.DisplayName, chat.ChatId)

		msg := tgbotapi.NewMessage(chat.ChatId, text)
		msg.ParseMode = "html"
		msg.DisableWebPagePreview = true

		if slices.Contains(client.toSlice(tgChatsWithGrants), chat.ChatId) {
			voteData := authz.ToVoteData(entChain.Name, entProp.ProposalID, govtypes.OptionEmpty, authz.NotVoted)
			msg.ReplyMarkup = createKeyboard(getVoteButtons(&voteData))
		}

		err := sendMessage(msg)
		if err != nil {
			if client.shouldDeleteUser(err) {
				errIds = append(errIds, chat.ChatId)
			} else {
				log.Sugar.Errorf("Error while sending message to telegram chat #%v: %v", chat, err)
			}
		}
	}
	return errIds
}

func (client TelegramLightClient) SendDraftProposals(entProp *ent.DraftProposal, entChain *ent.Chain) []int64 {
	text := fmt.Sprintf("💬  <b>%v - New pre-vote proposal\n\n%v</b>\n%v", entChain.DisplayName, entProp.Title, entProp.URL)

	var errIds []int64
	tgChats := client.TelegramChatManager.GetChatIdsWithDraftPropsEnabled(entChain)
	for _, chat := range tgChats {
		log.Sugar.Debugf("Send draft proposal #%v on %v to telegram chat #%v", entProp.DraftProposalID, entChain.DisplayName, chat.ChatId)

		msg := tgbotapi.NewMessage(chat.ChatId, text)
		msg.ParseMode = "html"
		msg.DisableWebPagePreview = true

		err := sendMessage(msg)
		if err != nil {
			if client.shouldDeleteUser(err) {
				errIds = append(errIds, chat.ChatId)
			} else {
				log.Sugar.Errorf("Error while sending message to telegram chat #%v: %v", chat, err)
			}
		}
	}
	return errIds
}

func (client TelegramLightClient) SendMessageToBotAdmins(message string) {
	admins := strings.Split(strings.Trim(os.Getenv("ADMIN_IDS"), " "), ",")
	for _, chatIdStr := range admins {
		chatId, err := strconv.Atoi(chatIdStr)
		if err != nil {
			log.Sugar.Error(err)
		}
		msg := tgbotapi.NewMessage(int64(chatId), message)
		msg.ParseMode = "html"
		msg.DisableWebPagePreview = true
		err = sendMessage(msg)
		handleError(chatId, err, client.TelegramChatManager)
	}
}
