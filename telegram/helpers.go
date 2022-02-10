package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
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
	log.Sugar.Panic("unreachable code reached!!!")
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

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func handleError(chatId int, err error) {
	if err != nil {
		if err.Error() == "Forbidden: bot was blocked by the user" {
			log.Sugar.Debugf("Delete user #%v", chatId)
			database.DeleteUser(int64(chatId))
		} else {
			log.Sugar.Errorf("Error while sending message: %v", err)
		}
	}
}
