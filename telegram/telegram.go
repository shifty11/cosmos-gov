package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/shifty11/cosmos-gov/log"
	"os"
	"strconv"
	"strings"
)

func isAdmin(fromId int) bool {
	admins := strings.Split(strings.Trim(os.Getenv("ADMIN_IDS"), " "), ",")
	return contains(admins, strconv.Itoa(fromId))
}

func isExpectingMessage(update *tgbotapi.Update) bool {
	currentState := getState(update)
	for _, s := range MessageStates {
		if s == currentState {
			return true
		}
	}
	return false
}

func setState(update *tgbotapi.Update, newState State, data *StateData) {
	chatId := getChatIdX(update)
	if newState == StateNil {
		delete(state, chatId)
		delete(stateData, chatId)
	} else {
		state[chatId] = newState
		if data != nil {
			stateData[chatId] = *data
		} else {
			delete(stateData, chatId)
		}
	}
}

func getState(update *tgbotapi.Update) State {
	chatId := getChatIdX(update)
	if _, exists := state[chatId]; exists {
		return state[chatId]
	}
	return StateNil
}

func getStateData(update *tgbotapi.Update) StateData {
	chatId := getChatIdX(update)
	if _, exists := stateData[chatId]; exists {
		return stateData[chatId]
	}
	return StateData{DataString: ""}
}

func handleCommand(update *tgbotapi.Update) {
	switch update.Message.Command() { // Check for non admin commands
	case "start", "notifications":
		sendMenu(update)
		setState(update, StateNil, nil)
	case "proposals":
		sendCurrentProposals(update)
		setState(update, StateNil, nil)
	default:
		if isAdmin(update.Message.From.ID) { // Check for admin commands
			switch update.Message.Command() {
			case "stats":
				sendUserStatistics(update)
				setState(update, StateNil, nil)
			case "broadcast":
				if sendBroadcastStart(update) {
					setState(update, StateStartBroadcast, nil)
				} else {
					setState(update, StateNil, nil)
				}
			}
		}
	}
}

func handleMessage(update *tgbotapi.Update) {
	switch getState(update) {
	case StateStartBroadcast:
		data := StateData{
			DataInt:   update.Message.MessageID,
			DataInt64: update.Message.Chat.ID,
		}
		sendConfirmBroadcastMessage(update, data.DataInt64, data.DataInt)
		setState(update, StateConfirmBroadcast, &data)
	case StateConfirmBroadcast:
		yesOptions := []string{"yes", "y"}
		if contains(yesOptions, strings.ToLower(update.Message.Text)) {
			data := getStateData(update)
			if data.DataInt == 0 || data.DataInt64 == 0 {
				log.Sugar.Fatal("No message to broadcast. This should never happen!")
			}
			sendBroadcastMessage(update, data.DataInt64, data.DataInt)
			sendBroadcastEndInfoMessage(update, true)
		} else {
			sendBroadcastEndInfoMessage(update, false)
		}
		setState(update, StateNil, nil)
	}
}

func handleCallbackQuery(update *tgbotapi.Update) {
	performUpdateNotification(update)
	sendMenu(update)
}

func Listen() {
	log.Sugar.Info("Start listening for commands")
	api := getApi()

	updateConfig := tgbotapi.NewUpdate(0)

	updates, err := api.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Sugar.Panic(err)
	}

	for update := range updates {
		if !hasChatId(&update) { // no chat id means there is something strange or the update is not for us
			continue
		}

		if update.Message != nil && update.Message.IsCommand() { // handle commands
			handleCommand(&update)
		} else if update.Message != nil && isExpectingMessage(&update) {
			//handleMessage(&update)
		} else if update.CallbackQuery != nil {
			handleCallbackQuery(&update)
		}
	}
}

func SendProposal(proposalText string, chatIds []int) {
	for _, chatId := range chatIds {
		msg := tgbotapi.NewMessage(int64(chatId), proposalText)
		msg.ParseMode = "html"
		msg.DisableWebPagePreview = true
		log.Sugar.Debugf("Send proposal to chat #%v", chatId)
		err := sendMessage(msg)
		handleError(chatId, err)
	}
}
