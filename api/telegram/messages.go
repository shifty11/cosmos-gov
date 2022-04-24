package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shifty11/cosmos-gov/api"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent/user"
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
	if update.Message != nil && update.Message.Chat != nil && update.Message.Chat.Type == "group" {
		log.Sugar.Debugf("Send subscriptions to group '%v' #%v", update.Message.Chat.Title, chatId)
	} else {
		log.Sugar.Debugf("Send subscriptions to user #%v", chatId)
	}
	userId := getUserIdX(update)
	subsManager := database.NewTelegramSubscriptionManager()
	chains := subsManager.GetOrCreateSubscriptions(userId, getUserName(update), chatId, getChatName(update), isGroupX(update))

	var buttons [][]Button
	var buttonRow []Button
	for ix, c := range chains {
		symbol := "❌ "
		if c.Notify {
			symbol = "✅ "
		}
		callbackData := CallbackData{Command: CallbackCmdShowSubscriptions, Data: c.Name}
		buttonRow = append(buttonRow, NewButton(symbol+c.DisplayName, callbackData))
		if (ix+1)%NbrOfButtonsPerRow == 0 || ix == len(chains)-1 {
			buttons = append(buttons, buttonRow)
			buttonRow = []Button{}
		}
	}
	config := createMenuButtonConfig()
	config.ShowSubscriptions = false
	buttons = append(buttons, getMenuButtonRow(config))
	if isBotAdmin(update) {
		botAdminConfig := createBotAdminMenuButtonConfig()
		buttons = append(buttons, getBotAdminMenuButtonRow(botAdminConfig))
	}
	replyMarkup := createKeyboard(buttons)

	if update.CallbackQuery == nil {
		msg := tgbotapi.NewMessage(chatId, messages.SubscriptionsMsg)
		msg.ReplyMarkup = replyMarkup
		msg.ParseMode = "markdown"
		msg.DisableWebPagePreview = true
		err := sendMessage(msg)
		if err != nil {
			log.Sugar.Errorf("Error while sendSubscriptions for user #%v: %v", chatId, err)
		}
	} else {
		msg := tgbotapi.NewEditMessageText(chatId, update.CallbackQuery.Message.MessageID, messages.SubscriptionsMsg)
		msg.ReplyMarkup = &replyMarkup
		msg.ParseMode = "markdown"
		msg.DisableWebPagePreview = true
		answerCallbackQuery(update)
		err := sendMessage(msg)
		if err != nil {
			log.Sugar.Errorf("Error while sendSubscriptions for user #%v: %v", chatId, err)
		}
	}
}

func getOngoingProposalsText(chatId int64) string {
	text := messages.ProposalsMsg
	chains := database.NewProposalManager().GetProposalsInVotingPeriod(chatId, user.TypeTelegram)
	if len(chains) == 0 {
		text = messages.NoSubscriptionsMsg
	} else {
		for _, chain := range chains {
			for _, prop := range chain.Edges.Proposals {
				text += fmt.Sprintf("<b>%v #%d</b> <i>%v</i>\n\n", chain.DisplayName, prop.ProposalID, prop.Title)
			}
		}
		if len(text) == len(messages.ProposalsMsg) {
			text = messages.NoProposalsMsg
		}
	}
	return text
}

func sendCurrentProposals(update *tgbotapi.Update) {
	chatId := getChatIdX(update)
	log.Sugar.Debugf("Send current proposals to user #%v", chatId)

	text := getOngoingProposalsText(chatId)

	config := createMenuButtonConfig()
	config.ShowProposals = false
	buttons := [][]Button{getMenuButtonRow(config)}
	if isBotAdmin(update) {
		botAdminConfig := createBotAdminMenuButtonConfig()
		buttons = append(buttons, getBotAdminMenuButtonRow(botAdminConfig))
	}
	replyMarkup := createKeyboard(buttons)

	if update.CallbackQuery == nil {
		msg := tgbotapi.NewMessage(chatId, text)
		msg.ReplyMarkup = replyMarkup
		msg.ParseMode = "markdown"
		err := sendMessage(msg)
		if err != nil {
			log.Sugar.Errorf("Error while sendCurrentProposals for user #%v: %v", chatId, err)
		}
	} else {
		msg := tgbotapi.NewEditMessageText(chatId, update.CallbackQuery.Message.MessageID, text)
		msg.ReplyMarkup = &replyMarkup
		msg.ParseMode = "markdown"
		answerCallbackQuery(update)
		err := sendMessage(msg)
		if err != nil {
			log.Sugar.Errorf("Error while sendCurrentProposals for user #%v: %v", chatId, err)
		}
	}
}

func sendHelp(update *tgbotapi.Update) {
	chatId := getChatIdX(update)
	log.Sugar.Debugf("Send help to user #%v", chatId)
	text := helpMsg
	if isBotAdmin(update) {
		text += "\n\n" + adminHelpMsg
	}

	config := createMenuButtonConfig()
	config.ShowHelp = false
	buttons := [][]Button{getMenuButtonRow(config)}
	if isBotAdmin(update) {
		botAdminConfig := createBotAdminMenuButtonConfig()
		buttons = append(buttons, getBotAdminMenuButtonRow(botAdminConfig))
	}
	replyMarkup := createKeyboard(buttons)

	if update.CallbackQuery == nil {
		msg := tgbotapi.NewMessage(chatId, text)
		msg.ReplyMarkup = replyMarkup
		msg.ParseMode = "html"
		err := sendMessage(msg)
		if err != nil {
			log.Sugar.Errorf("Error while sendHelp for user #%v: %v", chatId, err)
		}
	} else {
		msg := tgbotapi.NewEditMessageText(chatId, update.CallbackQuery.Message.MessageID, text)
		msg.ReplyMarkup = &replyMarkup
		msg.ParseMode = "html"
		answerCallbackQuery(update)
		err := sendMessage(msg)
		if err != nil {
			log.Sugar.Errorf("Error while sendHelp for user #%v: %v", chatId, err)
		}
	}
}

func sendSupport(update *tgbotapi.Update) {
	chatId := getChatIdX(update)
	log.Sugar.Debugf("Send support message to user #%v", chatId)

	config := createMenuButtonConfig()
	config.ShowSupport = false
	buttons := [][]Button{getMenuButtonRow(config)}
	if isBotAdmin(update) {
		botAdminConfig := createBotAdminMenuButtonConfig()
		buttons = append(buttons, getBotAdminMenuButtonRow(botAdminConfig))
	}
	replyMarkup := createKeyboard(buttons)

	text := fmt.Sprintf(messages.SupportMsg, "@rapha\\_decrypto")
	if update.CallbackQuery == nil {
		msg := tgbotapi.NewMessage(chatId, text)
		msg.ReplyMarkup = replyMarkup
		msg.ParseMode = "markdown"
		msg.DisableWebPagePreview = true
		err := sendMessage(msg)
		if err != nil {
			log.Sugar.Errorf("Error while sendSupport for user #%v: %v", chatId, err)
		}
	} else {
		msg := tgbotapi.NewEditMessageText(chatId, update.CallbackQuery.Message.MessageID, text)
		msg.ReplyMarkup = &replyMarkup
		msg.ParseMode = "markdown"
		msg.DisableWebPagePreview = true
		answerCallbackQuery(update)
		err := sendMessage(msg)
		if err != nil {
			log.Sugar.Errorf("Error while sendSupport for user #%v: %v", chatId, err)
		}
	}
}

func sendError(update *tgbotapi.Update) {
	chatId := getChatIdX(update)
	log.Sugar.Debugf("Send error msg to user #%v", chatId)
	text := errMsg
	msg := tgbotapi.NewMessage(chatId, text)
	if update.CallbackQuery != nil {
		answerCallbackQuery(update)
	}
	err := sendMessage(msg)
	if err != nil {
		log.Sugar.Errorf("Error while sendError for user #%v: %v", chatId, err)
	}
}
