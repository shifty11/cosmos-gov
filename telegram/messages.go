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

type StateData struct {
	DataString string
	DataInt    int
	DataInt64  int64
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

func sendBroadcastStart(update *tgbotapi.Update) bool {
	chatId := getChatIdX(update)
	msg := tgbotapi.NewMessage(chatId, startBroadcastInfoMsg)
	err := sendMessage(msg)
	if err != nil {
		log.Sugar.Error(err)
	}
	return err == nil
}

func sendConfirmBroadcastMessage(update *tgbotapi.Update, fromChatId int64, messageId int) {
	chatId := getChatIdX(update)
	cntUsers := database.CountUsers()
	forwardMsg := tgbotapi.NewForward(chatId, fromChatId, messageId)
	sendMessageX(forwardMsg)
	msg := tgbotapi.NewMessage(chatId, fmt.Sprintf(confirmBroadcastMsg, cntUsers))
	sendMessageX(msg)
}

func sendBroadcastMessage(update *tgbotapi.Update, fromChatId int64, messageId int) {
	chatIds := database.GetAllUserChatIds()
	log.Sugar.Debugf("Broadcast message to %v users", len(chatIds))
	for _, chatId := range chatIds {
		forwardMsg := tgbotapi.NewForward(int64(chatId), fromChatId, messageId)
		err := sendMessage(forwardMsg)
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
	msg := tgbotapi.NewMessage(chatId, "Send current proposals -> not yet implemented")
	sendMessageX(msg)
}
