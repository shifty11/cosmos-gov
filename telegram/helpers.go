package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/shifty11/cosmos-gov/log"
)

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

func getChatId(update *tgbotapi.Update) int64 {
	if update.CallbackQuery != nil {
		return update.CallbackQuery.Message.Chat.ID
	}
	return update.Message.Chat.ID
}

func sendMessage(message tgbotapi.Chattable) {
	api := getApi()
	_, err := api.Send(message)
	if err != nil {
		log.Sugar.Panic(err)
	}
}
