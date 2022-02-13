package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
	"strings"
)

type State int

const (
	StateNil              State = 0
	StateStartBroadcast   State = 1
	StateConfirmBroadcast State = 2
)

var MessageStates = [...]State{StateStartBroadcast, StateConfirmBroadcast}

var state = make(map[int64]State)
var stateData = make(map[int64]StateData)

type BroadcastStateData struct {
	Message string
}

type StateData struct {
	BroadcastStateData *BroadcastStateData
}

func sendMenu(update *tgbotapi.Update) {
	chatId := getChatIdX(update)
	chains := database.GetChainsForUser(chatId)

	var buttons [][]Button
	var buttonRow []Button
	for ix, c := range chains {
		symbol := "❌ "
		if c.Notify {
			symbol = "✅ "
		}
		buttonRow = append(buttonRow, NewButton(c.Name, symbol+c.DisplayName))
		if (ix+1)%NbrOfButtonsPerRow == 0 || ix == len(chains)-1 {
			buttons = append(buttons, buttonRow)
			buttonRow = []Button{}
		}
	}
	replyMarkup := createKeyboard(buttons)

	if update.CallbackQuery == nil {
		msg := tgbotapi.NewMessage(chatId, menuInfoMsg)
		msg.ReplyMarkup = replyMarkup
		sendMessageX(msg)
	} else {
		msg := tgbotapi.EditMessageTextConfig{
			BaseEdit: tgbotapi.BaseEdit{ChatID: chatId,
				MessageID:   update.CallbackQuery.Message.MessageID,
				ReplyMarkup: &replyMarkup,
			},
			Text: menuInfoMsg,
		}
		answerCallbackQuery(update)
		sendMessageX(msg)
	}
}

func sendUserStatistics(update *tgbotapi.Update) {
	chatId := getChatIdX(update)
	statistics, err := database.GetUserStatistics()
	if err != nil {
		log.Sugar.Error(err)
		return
	}
	chainStatistics, err := database.GetChainStatistics()
	if err != nil {
		log.Sugar.Error(err)
		return
	}

	userMsg := fmt.Sprintf("`"+userStatisticMsg+"`", statistics.CntUsers,
		statistics.CntUsersThisWeek, statistics.ChangeThisWeekInPercent,
		statistics.CntUsersSinceYesterday, statistics.ChangeSinceYesterdayInPercent)
	chainMsg := fmt.Sprintf("`" + chainStatisticHeaderMsg)
	sumUsers := 0
	sumChains := 0
	for _, chain := range *chainStatistics {
		chainMsg += fmt.Sprintf(chainStatisticRowMsg, strings.Title(chain.Name), chain.Notifications)
		sumUsers += chain.Notifications
		sumChains += 1
	}
	chainMsg += fmt.Sprintf(chainStatisticFooterMsg+"`", fmt.Sprintf("Total(%v)", sumChains), sumUsers)

	msg := tgbotapi.NewMessage(chatId, chainMsg+"\n\n"+userMsg)
	msg.ParseMode = "markdown"
	sendMessageX(msg)
}

func sendBroadcastStart(update *tgbotapi.Update) {
	chatId := getChatIdX(update)
	msg := tgbotapi.NewMessage(chatId, startBroadcastInfoMsg)
	msg.DisableWebPagePreview = true
	sendMessageX(msg)
}

func sendConfirmBroadcastMessage(update *tgbotapi.Update, text string) {
	chatId := getChatIdX(update)
	cntUsers := database.CountUsers()
	broadcastMsg := tgbotapi.NewMessage(chatId, text)
	broadcastMsg.DisableWebPagePreview = true
	broadcastMsg.ParseMode = "html"
	sendMessageX(broadcastMsg)
	msg := tgbotapi.NewMessage(chatId, fmt.Sprintf(confirmBroadcastMsg, cntUsers))
	msg.ParseMode = "markdown"
	sendMessageX(msg)
}

func sendBroadcastMessage(text string) {
	chatIds := database.GetAllUserChatIds()
	log.Sugar.Debugf("Broadcast message to %v users", len(chatIds))
	for _, chatId := range chatIds {
		broadcastMsg := tgbotapi.NewMessage(int64(chatId), text)
		broadcastMsg.DisableWebPagePreview = true
		broadcastMsg.ParseMode = "html"
		err := sendMessage(broadcastMsg)
		handleError(chatId, err)
	}
}

func sendBroadcastEndInfoMessage(update *tgbotapi.Update, success bool) {
	chatId := getChatIdX(update)
	text := abortBroadcastMsg
	if success {
		cntUsers := database.CountUsers()
		text = fmt.Sprintf(successBroadcastMsg, cntUsers)
	}
	msg := tgbotapi.NewMessage(chatId, text)
	sendMessageX(msg)
}

func sendCurrentProposals(update *tgbotapi.Update) {
	chatId := getChatIdX(update)
	text := showProposalMessage
	chains := database.GetProposalsInVotingPeriodForUser(chatId)
	for _, chain := range chains {
		for _, prop := range chain.Edges.Proposals {
			text += fmt.Sprintf("<b>%v #%d</b> %v\n\n", chain.DisplayName, prop.ProposalID, prop.Title)
		}
	}
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ParseMode = "html"
	sendMessageX(msg)
}
