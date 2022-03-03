package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
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

func sendSubscriptions(update *tgbotapi.Update) {
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
		text := subscriptionsMsg
		msg := tgbotapi.NewMessage(chatId, text)
		msg.ReplyMarkup = replyMarkup
		sendMessageX(msg)
	} else {
		msg := tgbotapi.EditMessageTextConfig{
			BaseEdit: tgbotapi.BaseEdit{ChatID: chatId,
				MessageID:   update.CallbackQuery.Message.MessageID,
				ReplyMarkup: &replyMarkup,
			},
			Text: subscriptionsMsg,
		}
		answerCallbackQuery(update)
		sendMessageX(msg)
	}
}

func sendCurrentProposals(update *tgbotapi.Update) {
	chatId := getChatIdX(update)
	text := proposalsMsg
	chains := database.GetProposalsInVotingPeriodForUser(chatId)
	if len(chains) == 0 {
		text = noSubscriptionsMsg
	} else {
		for _, chain := range chains {
			for _, prop := range chain.Edges.Proposals {
				text += fmt.Sprintf("<b>%v #%d</b> %v\n\n", chain.DisplayName, prop.ProposalID, prop.Title)
			}
		}
		if len(text) == len(proposalsMsg) {
			text = noProposalsMsg
		}
	}
	log.Sugar.Debugf("Send current proposals to user #%v", chatId)
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ParseMode = "html"
	sendMessageX(msg)
}

func sendHelp(update *tgbotapi.Update) {
	chatId := getChatIdX(update)
	log.Sugar.Debugf("Send help to user #%v", chatId)
	text := helpMsg
	if isBotAdmin(update) {
		text += "\n\n" + adminHelpMsg
	}
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ParseMode = "html"
	sendMessageX(msg)
}

func sendSupport(update *tgbotapi.Update) {
	chatId := getChatIdX(update)
	log.Sugar.Debugf("Send support message to user #%v", chatId)
	msg := tgbotapi.NewMessage(chatId, supportMsg)
	msg.DisableWebPagePreview = true
	msg.ParseMode = "html"
	sendMessageX(msg)
}
