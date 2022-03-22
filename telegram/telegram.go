package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/shifty11/cosmos-gov/common"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
	"strings"
)

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
	switch MessageCommand(update.Message.Command()) { // Check for non admin commands
	case MessageCmdStart, MessageCmdSubscriptions:
		sendSubscriptions(update)
		setState(update, StateNil, nil)
	case MessageCmdProposals:
		sendCurrentProposals(update)
		setState(update, StateNil, nil)
	case MessageCmdHelp:
		sendHelp(update)
		setState(update, StateNil, nil)
	case MessageCmdSupport:
		sendSupport(update)
		setState(update, StateNil, nil)
	default:
		if isBotAdmin(update) { // Check for admin commands
			switch MessageCommand(update.Message.Command()) {
			case MessageCmdStats:
				sendUserStatistics(update)
				setState(update, StateNil, nil)
			case MessageCmdBroadcast:
				sendBroadcastStart(update)
				setState(update, StateStartBroadcast, nil)
			case MessageCmdChains:
				sendChains(update)
				setState(update, StateNil, nil)
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
		if common.Contains(yesOptions, strings.ToLower(update.Message.Text)) {
			data := getStateData(update)
			if data.BroadcastStateData == nil || data.BroadcastStateData.Message == "" {
				log.Sugar.Fatal("No message to broadcast. This should never happen!")
			}
			sendBroadcastMessage(data.BroadcastStateData.Message)
			sendBroadcastEndInfoMessage(update, true)
			setState(update, StateNil, nil)
		} else if common.Contains(abortOptions, strings.ToLower(update.Message.Text)) {
			sendBroadcastEndInfoMessage(update, false)
			setState(update, StateNil, nil)
		} else {
			sendBroadcastStart(update)
			setState(update, StateStartBroadcast, nil)
		}
	}
}

func handleCallbackQuery(update *tgbotapi.Update) {
	callbackData := ToCallbackData(update.CallbackQuery.Data)
	switch callbackData.Command {
	case CallbackCmdShowSubscriptions:
		database.PerformUpdateSubscription(getChatIdX(update), user.TypeTelegram, callbackData.Data)
		sendSubscriptions(update)
	case CallbackCmdShowProposals:
		sendCurrentProposals(update)
	case CallbackCmdShowHelp:
		sendHelp(update)
	case CallbackCmdShowSupport:
		sendSupport(update)
	default:
		if isBotAdmin(update) { // Check for admin callbacks
			switch callbackData.Command {
			case CallbackCmdStats:
				sendUserStatistics(update)
			case CallbackCmdEnableChains:
				performToggleChain(callbackData.Data)
				sendChains(update)
			default:
				sendError(update)
				sendHelp(update)
				setState(update, StateNil, nil)
			}
		} else {
			sendError(update)
			sendHelp(update)
			setState(update, StateNil, nil)
		}
	}
}

// groups -> just admins and creators can interact with the bot
// private -> everything is allowed
func isInteractionAllowed(update *tgbotapi.Update) bool {
	if isUpdateFromGroup(update) {
		return isUpdateFromCreatorOrAdministrator(update)
	}
	return true
}

// Handles updates for only 1 user in a serial way
func handleUpdates(channel chan tgbotapi.Update) {
	for update := range channel {
		chatId := getChatIdX(&update)
		if isInteractionAllowed(&update) {
			if update.Message != nil && update.Message.IsCommand() {
				handleCommand(&update)
			} else if update.Message != nil && isExpectingMessage(&update) {
				handleMessage(&update)
			} else if update.CallbackQuery != nil {
				handleCallbackQuery(&update)
			}
		} else {
			log.Sugar.Debugf("Interaction with bot for user #%v is not allowed", chatId)
			if update.CallbackQuery != nil {
				answerCallbackQuery(&update)
			}
		}
		updateCountChannel <- UpdateCount{ChatId: chatId, Updates: -1}
	}
}

type UpdateCount struct {
	ChatId  int64
	Updates int
}

// updateChannels contains one update channel for every user.
// This means the updates can be processed parallel for multiple users but serial for every single user
var updateChannels map[int64]chan tgbotapi.Update

// updateCountChannel is used to communicate to `manageUpdateChannels` from `handleUpdates`
var updateCountChannel chan UpdateCount

func hasChannel(channelId int64) bool {
	for key := range updateChannels {
		if key == channelId {
			return true
		}
	}
	return false
}

func sendToChannelAsync(chatId int64, update tgbotapi.Update) {
	updateCountChannel <- UpdateCount{ChatId: chatId, Updates: 1}
	updateChannels[chatId] <- update
}

func sendToChannel(update *tgbotapi.Update) {
	chatId := getChatIdX(update)
	if !hasChannel(chatId) {
		updateChannels[chatId] = make(chan tgbotapi.Update)
		go handleUpdates(updateChannels[chatId])
	}
	go sendToChannelAsync(chatId, *update)
}

// Keeps track of all the user channels and closes them if there are no more updates
func manageUpdateChannels() {
	updateCountChannel = make(chan UpdateCount)
	var count = make(map[int64]int)
	for msg := range updateCountChannel {
		count[msg.ChatId] += msg.Updates
		if count[msg.ChatId] == 0 {
			close(updateChannels[msg.ChatId])
			delete(updateChannels, msg.ChatId)
			delete(count, msg.ChatId)
		}
	}
}

func Start() {
	log.Sugar.Info("Start telegram bot")
	api := getApi()

	updateConfig := tgbotapi.NewUpdate(0)
	updates, err := api.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Sugar.Panic(err)
	}

	updateChannels = make(map[int64]chan tgbotapi.Update)
	go manageUpdateChannels()

	for update := range updates {
		if !hasChatId(&update) { // no chat id means there is something strange or the update is not for us
			continue
		}

		sendToChannel(&update)
	}
}
