package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/shifty11/cosmos-gov/common"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
	"math"
	"os"
	"strings"
)

var api *tgbotapi.BotAPI = nil

func getApi() *tgbotapi.BotAPI {
	if api == nil {
		telegramToken := os.Getenv("TELEGRAM_TOKEN")
		if telegramToken == "" {
			log.Sugar.Panic("you must provide a telegram token as env variable")
		}
		botApi, err := tgbotapi.NewBotAPI(telegramToken)
		if err != nil {
			log.Sugar.Panic(err)
		}
		api = botApi
	}
	return api
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
	} else if len(split) == 2 {
		return CallbackData{Command: CallbackCommand(split[0]), Data: split[1]}
	}
	log.Sugar.Errorf("Can not convert string to CallbackData: '%v'", str)
	return CallbackData{}
}

type Button struct {
	Text         string
	CallbackData CallbackData
}

func NewButton(text string, callbackData CallbackData) Button {
	return Button{Text: text, CallbackData: callbackData}
}

func createKeyboard(buttons [][]Button) tgbotapi.InlineKeyboardMarkup {
	var keyboard [][]tgbotapi.InlineKeyboardButton
	for _, row := range buttons {
		var keyboardRow []tgbotapi.InlineKeyboardButton
		for _, button := range row {
			data := button.CallbackData.String()
			btn := tgbotapi.InlineKeyboardButton{Text: button.Text, CallbackData: &data}
			keyboardRow = append(keyboardRow, btn)
		}
		keyboard = append(keyboard, keyboardRow)
	}
	return tgbotapi.InlineKeyboardMarkup{InlineKeyboard: keyboard}
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

func getUserIdX(update *tgbotapi.Update) int {
	if update.CallbackQuery != nil {
		return update.CallbackQuery.From.ID
	}
	if update.Message != nil {
		return update.Message.From.ID
	}
	log.Sugar.Panic("getUserIdX: unreachable code reached!!!")
	return 0
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
		_, err := api.AnswerCallbackQuery(callback)
		if err != nil {
			log.Sugar.Error(err)
		}
	}
}

var forbiddenErrors = []string{
	"Forbidden: bot was blocked by the user",
	"Forbidden: bot was kicked from the group chat",
	"Forbidden: bot was kicked from the supergroup chat",
}

func handleError(chatId int, err error) {
	if err != nil {
		if common.Contains(forbiddenErrors, err.Error()) {
			log.Sugar.Debugf("Delete user #%v", chatId)
			database.DeleteUser(int64(chatId), user.TypeTelegram)
		} else {
			log.Sugar.Errorf("Error while sending message to chat #%v: %v", chatId, err)
		}
	}
}

// ChatId is negative for groups
func isUpdateFromGroup(update *tgbotapi.Update) bool {
	chatId := getChatIdX(update)
	return math.Signbit(float64(chatId))
}

func isUpdateFromCreatorOrAdministrator(update *tgbotapi.Update) bool {
	api := getApi()
	chatId := getChatIdX(update)
	userId := getUserIdX(update)
	memberConfig := tgbotapi.ChatConfigWithUser{
		ChatID:             chatId,
		SuperGroupUsername: "",
		UserID:             userId,
	}
	member, err := api.GetChatMember(memberConfig)
	if err != nil {
		if common.Contains(forbiddenErrors, err.Error()) {
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
		buttonRow = append(buttonRow, NewButton("🔔 Subscriptions", CallbackData{Command: CallbackCmdShowSubscriptions}))
	}
	if config.ShowProposals {
		buttonRow = append(buttonRow, NewButton("🗳 Proposals", CallbackData{Command: CallbackCmdShowProposals}))
	}
	//if config.ShowHelp {
	//	buttonRow = append(buttonRow, NewButton("🆘 Help", CallbackData{Command: CallbackCmdShowHelp}))
	//}
	if config.ShowSupport {
		buttonRow = append(buttonRow, NewButton("💰 Support", CallbackData{Command: CallbackCmdShowSupport}))
	}
	return buttonRow
}

type BotAdminMenuButtonConfig struct {
	ShowStats  bool
	ShowChains bool
	//ShowBroadcast bool
}

func createBotAdminMenuButtonConfig() BotAdminMenuButtonConfig {
	return BotAdminMenuButtonConfig{ShowStats: true, ShowChains: true}
}

func getBotAdminMenuButtonRow(config BotAdminMenuButtonConfig) []Button {
	var buttonRow []Button
	if config.ShowStats {
		buttonRow = append(buttonRow, NewButton("📈 Stats", CallbackData{Command: CallbackCmdStats}))
	}
	if config.ShowChains {
		buttonRow = append(buttonRow, NewButton("🔗 Chains", CallbackData{Command: CallbackCmdEnableChains}))
	}
	//if config.ShowBroadcast {
	//	buttonRow = append(buttonRow, NewButton("🔊 Broadcast", CallbackData{Command: CallbackCmdBroadcast}))
	//}
	return buttonRow
}
