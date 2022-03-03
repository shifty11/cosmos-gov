package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/shifty11/cosmos-gov/common"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
	"os"
	"strconv"
	"strings"
)

func isBotAdmin(update *tgbotapi.Update) bool {
	var fromId int
	if update.Message != nil {
		fromId = update.Message.From.ID
	} else if update.CallbackQuery != nil {
		fromId = update.CallbackQuery.From.ID
	} else {
		return false
	}
	admins := strings.Split(strings.Trim(os.Getenv("ADMIN_IDS"), " "), ",")
	return common.Contains(admins, strconv.Itoa(fromId))
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
		chainMsg += fmt.Sprintf(chainStatisticRowMsg, strings.Title(chain.DisplayName), chain.Notifications)
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

func sendChains(update *tgbotapi.Update) {
	chatId := getChatIdX(update)
	chains := database.GetChains()

	var buttons [][]Button
	var buttonRow []Button
	for ix, c := range chains {
		symbol := "❌ "
		if c.IsEnabled {
			symbol = "✅ "
		}
		callbackData := CallbackData{Command: CallbackCommandENABLE_CHAINS, Data: c.Name}
		buttonRow = append(buttonRow, NewButton(symbol+c.DisplayName, callbackData))
		if (ix+1)%NbrOfButtonsPerRow == 0 || ix == len(chains)-1 {
			buttons = append(buttons, buttonRow)
			buttonRow = []Button{}
		}
	}
	replyMarkup := createKeyboard(buttons)

	if update.CallbackQuery == nil {
		text := newChainsMsg
		msg := tgbotapi.NewMessage(chatId, text)
		msg.ReplyMarkup = replyMarkup
		err := sendMessage(msg)
		if err != nil {
			log.Sugar.Errorf("Error while sendChains for user #%v: %v", chatId, err)
		}
	} else {
		msg := tgbotapi.EditMessageTextConfig{
			BaseEdit: tgbotapi.BaseEdit{ChatID: chatId,
				MessageID:   update.CallbackQuery.Message.MessageID,
				ReplyMarkup: &replyMarkup,
			},
			Text: newChainsMsg,
		}
		answerCallbackQuery(update)
		err := sendMessage(msg)
		if err != nil {
			log.Sugar.Errorf("Error while sendChains for user #%v: %v", chatId, err)
		}
	}
}

func SendMessageToBotAdmins(message string) {
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
		handleError(chatId, err)
	}
}
