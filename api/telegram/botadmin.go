package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
	"golang.org/x/exp/slices"
	"os"
	"strconv"
	"strings"
)

func isBotAdmin(update *tgbotapi.Update) bool {
	var fromId int64
	if update.Message != nil {
		fromId = update.Message.From.ID
	} else if update.CallbackQuery != nil {
		fromId = update.CallbackQuery.From.ID
	} else {
		return false
	}
	admins := strings.Split(strings.Trim(os.Getenv("ADMIN_IDS"), " "), ",")
	return slices.Contains(admins, strconv.FormatInt(fromId, 10))
}

func handleError(chatId int, err error, tgChatManager *database.TelegramChatManager) {
	if err != nil {
		if slices.Contains(forbiddenErrors, err.Error()) {
			log.Sugar.Debugf("Delete user #%v", chatId)
			tgChatManager.Delete(int64(chatId))
		} else {
			log.Sugar.Errorf("Error while sending message to chat #%v: %v", chatId, err)
		}
	}
}

func (client TelegramClient) sendUserStatistics(update *tgbotapi.Update) {
	chatId := getChatIdX(update)
	statsManager := client.StatsManager
	chainStatistics, err := statsManager.GetChainStats()
	if err != nil {
		log.Sugar.Error(err)
		return
	}

	telegramStats, err := statsManager.GetUserStatistics(user.TypeTelegram)
	if err != nil {
		log.Sugar.Error(err)
		return
	}

	discordStats, err := statsManager.GetUserStatistics(user.TypeDiscord)
	if err != nil {
		log.Sugar.Error(err)
		return
	}

	sumSubscriptions := 0
	sumProposals := 0
	sumChains := 0
	chainMsg := fmt.Sprintf("`" + chainStatisticHeaderMsg)
	for _, chain := range chainStatistics {
		chainMsg += fmt.Sprintf(chainStatisticRowMsg, chain.DisplayName, chain.Proposals, chain.Subscriptions)
		sumSubscriptions += chain.Subscriptions
		sumProposals += chain.Proposals
		sumChains += 1
	}
	chainMsg += fmt.Sprintf(chainStatisticFooterMsg+"`", fmt.Sprintf("Total(%v)", sumChains), sumProposals, sumSubscriptions)

	telegramMsg := fmt.Sprintf("`"+userStatisticMsg+"`", user.TypeTelegram, telegramStats.CntUsers,
		telegramStats.CntUsersThisWeek, telegramStats.ChangeThisWeekInPercent,
		telegramStats.CntUsersSinceYesterday, telegramStats.ChangeSinceYesterdayInPercent)

	discordMsg := fmt.Sprintf("`"+userStatisticMsg+"`", user.TypeDiscord, discordStats.CntUsers,
		discordStats.CntUsersThisWeek, discordStats.ChangeThisWeekInPercent,
		discordStats.CntUsersSinceYesterday, discordStats.ChangeSinceYesterdayInPercent)

	text := chainMsg + "\n\n" + telegramMsg + "\n\n" + discordMsg

	config := createMenuButtonConfig(update)
	buttons := [][]Button{getMenuButtonRow(config, update)}
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
			log.Sugar.Errorf("Error while sendUserStatistics for user #%v: %v", chatId, err)
		}
	} else {
		msg := tgbotapi.NewEditMessageText(chatId, update.CallbackQuery.Message.MessageID, text)
		msg.ReplyMarkup = replyMarkup
		msg.ParseMode = "markdown"
		answerCallbackQuery(update)
		err := sendMessage(msg)
		if err != nil {
			log.Sugar.Debugf("Error while sendUserStatistics for user #%v: %v", chatId, err)
		}
	}
}

func (client TelegramClient) sendBroadcastStart(update *tgbotapi.Update) {
	chatId := getChatIdX(update)
	msg := tgbotapi.NewMessage(chatId, startBroadcastInfoMsg)
	msg.DisableWebPagePreview = true
	sendMessageX(msg)
}

func (client TelegramClient) sendConfirmBroadcastMessage(update *tgbotapi.Update, text string) {
	chatId := getChatIdX(update)
	cntUsers := client.TelegramChatManager.CountChats()
	broadcastMsg := tgbotapi.NewMessage(chatId, text)
	broadcastMsg.DisableWebPagePreview = true
	broadcastMsg.ParseMode = "html"
	sendMessageX(broadcastMsg)
	msg := tgbotapi.NewMessage(chatId, fmt.Sprintf(confirmBroadcastMsg, cntUsers))
	msg.ParseMode = "markdown"
	sendMessageX(msg)
}

func (client TelegramClient) sendBroadcastMessage(text string) {
	chatIds := client.TelegramChatManager.GetAllChatIds()
	log.Sugar.Debugf("Broadcast message to %v users", len(chatIds))
	for _, chatId := range chatIds {
		broadcastMsg := tgbotapi.NewMessage(int64(chatId), text)
		broadcastMsg.DisableWebPagePreview = true
		broadcastMsg.ParseMode = "html"
		err := sendMessage(broadcastMsg)
		handleError(chatId, err, client.TelegramChatManager)
	}
}

func (client TelegramClient) sendBroadcastEndInfoMessage(update *tgbotapi.Update, success bool) {
	chatId := getChatIdX(update)
	text := abortBroadcastMsg
	if success {
		cntUsers := client.TelegramChatManager.CountChats()
		text = fmt.Sprintf(successBroadcastMsg, cntUsers)
	}
	msg := tgbotapi.NewMessage(chatId, text)
	sendMessageX(msg)
}
