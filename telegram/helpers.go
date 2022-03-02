package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/shifty11/cosmos-gov/common"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
	"math"
	"os"
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

type Button struct {
	Command string
	Text    string
}

func NewButton(command string, text string) Button {
	return Button{Command: command, Text: text}
}

func createKeyboard(buttons [][]Button) tgbotapi.InlineKeyboardMarkup {
	var keyboard [][]tgbotapi.InlineKeyboardButton
	for _, row := range buttons {
		var keyboardRow []tgbotapi.InlineKeyboardButton
		for _, button := range row {
			data := button.Command
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
			database.DeleteUser(int64(chatId))
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
