package telegram

import (
	"fmt"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shifty11/cosmos-gov/authz"
	"github.com/shifty11/cosmos-gov/log"
	"golang.org/x/exp/slices"
	"os"
	"strings"
)

var _api *tgbotapi.BotAPI = nil

func getApi() *tgbotapi.BotAPI {
	if _api == nil {
		telegramToken := os.Getenv("TELEGRAM_TOKEN")
		if telegramToken == "" {
			log.Sugar.Panic("you must provide a telegram token as env variable")
		}
		botApi, err := tgbotapi.NewBotAPI(telegramToken)
		if err != nil {
			log.Sugar.Panic(err)
		}
		_api = botApi
		//_api.Debug = os.Getenv("DEBUG") == "true"
	}
	return _api
}

type MessageCommand string

const (
	MessageCmdStart         MessageCommand = "start"
	MessageCmdSubscriptions MessageCommand = "subscriptions"
	MessageCmdProposals     MessageCommand = "proposals"
	MessageCmdHelp          MessageCommand = "help"
	MessageCmdSupport       MessageCommand = "support"

	MessageCmdStats     MessageCommand = "stats"     // admin command
	MessageCmdChains    MessageCommand = "chains"    // admin command
	MessageCmdBroadcast MessageCommand = "broadcast" // admin command
)

type CallbackCommand string

const (
	CallbackCmdShowSubscriptions CallbackCommand = "SHOW_SUBSCRIPTION"
	CallbackCmdShowProposals     CallbackCommand = "SHOW_PROPOSALS"
	CallbackCmdShowHelp          CallbackCommand = "SHOW_HELP"
	CallbackCmdShowSupport       CallbackCommand = "SHOW_SUPPORT"

	CallbackCmdStats        CallbackCommand = "STATS"         // admin command
	CallbackCmdEnableChains CallbackCommand = "ENABLE_CHAINS" // admin command
	CallbackCmdVote         CallbackCommand = "VOTE"
	//CallbackCmdBroadcast    CallbackCommand = "BROADCAST"     // admin command
)

type CallbackData struct {
	Command CallbackCommand
	Data    string
}

func (cd CallbackData) String() string {
	return fmt.Sprintf("%v:%v", cd.Command, cd.Data)
}

func ToCallbackData(str string) CallbackData {
	split := strings.Split(str, ":")
	if len(split) == 1 {
		return CallbackData{Command: CallbackCommand(split[0])}
	} else if len(split) >= 2 {
		return CallbackData{Command: CallbackCommand(split[0]), Data: strings.Join(split[1:], ":")}
	}
	log.Sugar.Errorf("Can not convert string to CallbackData: '%v'", str)
	return CallbackData{}
}

type Button struct {
	Text         string
	CallbackData *CallbackData
	LoginURL     *tgbotapi.LoginURL
}

func NewButton(text string, callbackData *CallbackData) Button {
	return Button{Text: text, CallbackData: callbackData}
}

func createKeyboard(buttons [][]Button) *tgbotapi.InlineKeyboardMarkup {
	var keyboard [][]tgbotapi.InlineKeyboardButton
	for _, row := range buttons {
		var keyboardRow []tgbotapi.InlineKeyboardButton
		for _, button := range row {
			btn := tgbotapi.InlineKeyboardButton{Text: button.Text, LoginURL: button.LoginURL}
			if button.CallbackData != nil {
				s := button.CallbackData.String()
				btn.CallbackData = &s
			}
			keyboardRow = append(keyboardRow, btn)
		}
		keyboard = append(keyboard, keyboardRow)
	}
	return &tgbotapi.InlineKeyboardMarkup{InlineKeyboard: keyboard}
}

func hasChatId(update *tgbotapi.Update) bool {
	if update.CallbackQuery != nil {
		return true
	}
	if update.Message != nil {
		return true
	}
	return false
}

func getChatIdX(update *tgbotapi.Update) int64 {
	if update.CallbackQuery != nil {
		return update.CallbackQuery.Message.Chat.ID
	}
	if update.Message != nil {
		return update.Message.Chat.ID
	}
	log.Sugar.Panic("getChatIdX: unreachable code reached!!!")
	return 0
}

func getChatName(update *tgbotapi.Update) string {
	if update.CallbackQuery != nil {
		if isGroupX(update) {
			return update.CallbackQuery.Message.Chat.Title
		}
		return update.CallbackQuery.Message.Chat.UserName
	}
	if update.Message != nil {
		if isGroupX(update) {
			return update.Message.Chat.Title
		}
		return update.Message.Chat.UserName
	}
	return ""
}

func isGroupX(update *tgbotapi.Update) bool {
	if update.CallbackQuery != nil {
		return !update.CallbackQuery.Message.Chat.IsPrivate()
	}
	if update.Message != nil {
		return !update.Message.Chat.IsPrivate()
	}
	log.Sugar.Panic("isGroupX: unreachable code reached!!!")
	return false
}

func getUserIdX(update *tgbotapi.Update) int64 {
	if update.CallbackQuery != nil {
		return update.CallbackQuery.From.ID
	}
	if update.Message != nil {
		return update.Message.From.ID
	}
	log.Sugar.Panic("getUserIdX: unreachable code reached!!!")
	return 0
}

func getUserName(update *tgbotapi.Update) string {
	if update.CallbackQuery != nil {
		return update.CallbackQuery.From.UserName
	}
	if update.Message != nil {
		return update.Message.From.UserName
	}
	return "<not found>"
}

func sendMessageX(message tgbotapi.Chattable) {
	api := getApi()
	_, err := api.Send(message)
	if err != nil {
		log.Sugar.Panic(err)
	}
}

func sendMessage(message tgbotapi.Chattable) error {
	api := getApi()
	_, err := api.Send(message)
	return err
}

func answerCallbackQuery(update *tgbotapi.Update) {
	if update.CallbackQuery != nil {
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
		api := getApi()
		_, err := api.Request(callback)
		if err != nil {
			log.Sugar.Error(err)
		}
	}
}

var forbiddenErrors = []string{
	"Forbidden: bot was blocked by the user",
	"Forbidden: bot was kicked from the group chat",
	"Forbidden: bot was kicked from the supergroup chat",
	"Forbidden: bot is not a member of the supergroup chat",
	"Forbidden: user is deactivated",
	"Bad Request: chat not found",
}

func handleError(chatId int, err error) {
	if err != nil {
		if slices.Contains(forbiddenErrors, err.Error()) {
			log.Sugar.Debugf("Delete user #%v", chatId)
			mHack.TelegramChatManager.Delete(int64(chatId))
		} else {
			log.Sugar.Errorf("Error while sending message to chat #%v: %v", chatId, err)
		}
	}
}

func isUpdateFromCreatorOrAdministrator(update *tgbotapi.Update) bool {
	api := getApi()
	chatId := getChatIdX(update)
	userId := getUserIdX(update)
	memberConfig := tgbotapi.GetChatMemberConfig{
		ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
			ChatID:             chatId,
			SuperGroupUsername: "",
			UserID:             userId,
		},
	}
	member, err := api.GetChatMember(memberConfig)
	if err != nil {
		if slices.Contains(forbiddenErrors, err.Error()) {
			log.Sugar.Debugf("Error while getting member (ChatID: %v; UserID: %v): %v", chatId, userId, err)
			return false
		}
		log.Sugar.Errorf("Error while getting member (ChatID: %v; UserID: %v): %v", chatId, userId, err)
		return false
	}
	return member.Status == "creator" || member.Status == "administrator"
}

type MenuButtonConfig struct {
	ShowSubscriptions bool
	ShowProposals     bool
	ShowHelp          bool
	ShowSupport       bool
}

func createMenuButtonConfig() MenuButtonConfig {
	return MenuButtonConfig{ShowSubscriptions: true, ShowProposals: true, ShowHelp: true, ShowSupport: true}
}

func getMenuButtonRow(config MenuButtonConfig) []Button {
	var buttonRow []Button
	if config.ShowSubscriptions {
		buttonRow = append(buttonRow, NewButton("🔔 Subscriptions", &CallbackData{Command: CallbackCmdShowSubscriptions}))
	}
	if config.ShowProposals {
		buttonRow = append(buttonRow, NewButton("🗳 Proposals", &CallbackData{Command: CallbackCmdShowProposals}))
	}
	//if config.ShowHelp {
	//	buttonRow = append(buttonRow, NewButton("🆘 Help", CallbackData{Command: CallbackCmdShowHelp}))
	//}
	//if config.ShowSupport {
	//	buttonRow = append(buttonRow, NewButton("💰 Support", CallbackData{Command: CallbackCmdShowSupport}))
	//}
	return buttonRow
}

type BotAdminMenuButtonConfig struct {
	ShowStats  bool
	ShowChains bool
	//ShowBroadcast bool
	ShowWebAppLogin bool
}

func createBotAdminMenuButtonConfig() BotAdminMenuButtonConfig {
	return BotAdminMenuButtonConfig{ShowStats: true, ShowChains: true, ShowWebAppLogin: true}
}

func getBotAdminMenuButtonRow(config BotAdminMenuButtonConfig) []Button {
	var buttonRow []Button
	if config.ShowStats {
		buttonRow = append(buttonRow, NewButton("📈 Stats", &CallbackData{Command: CallbackCmdStats}))
	}
	if config.ShowChains {
		buttonRow = append(buttonRow, NewButton("🔗 Chains", &CallbackData{Command: CallbackCmdEnableChains}))
	}
	//if config.ShowBroadcast {
	//	buttonRow = append(buttonRow, NewButton("🔊 Broadcast", CallbackData{Command: CallbackCmdBroadcast}))
	//}
	if config.ShowWebAppLogin {
		button := NewButton("🌎 Web app", nil)
		button.LoginURL = &tgbotapi.LoginURL{URL: "test.mydomain.com:40001"}
		buttonRow = append(buttonRow, button)
	}
	return buttonRow
}

func getVoteButtons(vd *authz.VoteData) [][]Button {
	var buttons [][]Button
	var buttonRow []Button
	var options = []govtypes.VoteOption{govtypes.OptionYes, govtypes.OptionNo, govtypes.OptionAbstain, govtypes.OptionNoWithVeto}
	for i, option := range options {
		s := authz.NotVoted
		if option == vd.Vote {
			s = vd.State
		}
		voteData := authz.ToVoteData(vd.ChainName, vd.ProposalId, option, s)
		callbackData := &CallbackData{Command: CallbackCmdVote, Data: voteData.ToString()}
		buttonRow = append(buttonRow, NewButton(voteData.ButtonText(), callbackData))
		if (i+1)%2 == 0 {
			buttons = append(buttons, buttonRow)
			buttonRow = []Button{}
		}
	}
	return buttons
}
