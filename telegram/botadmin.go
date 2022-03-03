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
	sumSubscriptions := 0
	sumProposals := 0
	sumChains := 0
	for _, chain := range *chainStatistics {
		chainMsg += fmt.Sprintf(chainStatisticRowMsg, strings.Title(chain.DisplayName), chain.Proposals, chain.Subscriptions)
		sumSubscriptions += chain.Subscriptions
		sumProposals += chain.Proposals
		sumChains += 1
	}
	chainMsg += fmt.Sprintf(chainStatisticFooterMsg+"`", fmt.Sprintf("Total(%v)", sumChains), sumProposals, sumSubscriptions)

	text := chainMsg + "\n\n" + userMsg

	config := createMenuButtonConfig()
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
			log.Sugar.Errorf("Error while sendSubscriptions for user #%v: %v", chatId, err)
		}
	} else {
		msg := tgbotapi.NewEditMessageText(chatId, update.CallbackQuery.Message.MessageID, text)
		msg.ReplyMarkup = &replyMarkup
		msg.ParseMode = "markdown"
		answerCallbackQuery(update)
		err := sendMessage(msg)
		if err != nil {
			log.Sugar.Debugf("Error while sendSubscriptions for user #%v: %v", chatId, err)
		}
	}
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
		symbol := "🔴️ "
		if c.IsEnabled {
			symbol = "\U0001F7E2 "
		}
		callbackData := CallbackData{Command: CallbackCmdEnableChains, Data: c.Name}
		buttonRow = append(buttonRow, NewButton(symbol+c.DisplayName, callbackData))
		if (ix+1)%NbrOfButtonsPerRow == 0 || ix == len(chains)-1 {
			buttons = append(buttons, buttonRow)
			buttonRow = []Button{}
		}
	}

	config := createMenuButtonConfig()
	buttons = append(buttons, getMenuButtonRow(config))
	if isBotAdmin(update) {
		botAdminConfig := createBotAdminMenuButtonConfig()
		botAdminConfig.ShowChains = false
		buttons = append(buttons, getBotAdminMenuButtonRow(botAdminConfig))
	}
	replyMarkup := createKeyboard(buttons)

	if update.CallbackQuery == nil {
		msg := tgbotapi.NewMessage(chatId, newChainsMsg)
		msg.ReplyMarkup = replyMarkup
		err := sendMessage(msg)
		if err != nil {
			log.Sugar.Errorf("Error while sendChains for user #%v: %v", chatId, err)
		}
	} else {
		msg := tgbotapi.NewEditMessageText(chatId, update.CallbackQuery.Message.MessageID, newChainsMsg)
		msg.ReplyMarkup = &replyMarkup
		answerCallbackQuery(update)
		err := sendMessage(msg)
		if err != nil {
			log.Sugar.Debugf("Error while sendChains for user #%v: %v", chatId, err)
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
