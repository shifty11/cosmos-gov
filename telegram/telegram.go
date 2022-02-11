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
	return StateData{}
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
				sendBroadcastStart(update)
				setState(update, StateStartBroadcast, nil)
			}
		}
	}
}

func handleMessage(update *tgbotapi.Update) {
	switch getState(update) {
	case StateStartBroadcast:
		data := StateData{BroadcastStateData: &BroadcastStateData{Message: update.Message.Text}}
		sendConfirmBroadcastMessage(update, data.BroadcastStateData.Message)
		setState(update, StateConfirmBroadcast, &data)
	case StateConfirmBroadcast:
		yesOptions := []string{"yes", "y"}
		abortOptions := []string{"abort", "a"}
		if contains(yesOptions, strings.ToLower(update.Message.Text)) {
			data := getStateData(update)
			if data.BroadcastStateData == nil || data.BroadcastStateData.Message == "" {
				log.Sugar.Fatal("No message to broadcast. This should never happen!")
			}
			sendBroadcastMessage(data.BroadcastStateData.Message)
			sendBroadcastEndInfoMessage(update, true)
			setState(update, StateNil, nil)
		} else if contains(abortOptions, strings.ToLower(update.Message.Text)) {
			sendBroadcastEndInfoMessage(update, false)
			setState(update, StateNil, nil)
		} else {
			sendBroadcastStart(update)
			setState(update, StateStartBroadcast, nil)
		}
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
			handleMessage(&update)
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
