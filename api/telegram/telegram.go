package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
	"golang.org/x/exp/slices"
	"strings"
)

var mHack database.DbManagers // TODO: get rid of this hack

func (client TelegramClient) isExpectingMessage(update *tgbotapi.Update) bool {
	currentState := client.getState(update)
	for _, s := range MessageStates {
		if s == currentState {
			return true
		}
	}
	return false
}

func (client TelegramClient) setState(update *tgbotapi.Update, newState State, data *StateData) {
	chatId := getChatIdX(update)
	if newState == StateNil {
		delete(client.state, chatId)
		delete(client.stateData, chatId)
	} else {
		client.state[chatId] = newState
		if data != nil {
			client.stateData[chatId] = *data
		} else {
			delete(client.stateData, chatId)
		}
	}
}

func (client TelegramClient) getState(update *tgbotapi.Update) State {
	chatId := getChatIdX(update)
	if _, exists := client.state[chatId]; exists {
		return client.state[chatId]
	}
	return StateNil
}

func (client TelegramClient) getStateData(update *tgbotapi.Update) StateData {
	chatId := getChatIdX(update)
	if _, exists := client.stateData[chatId]; exists {
		return client.stateData[chatId]
	}
	return StateData{}
}

func (client TelegramClient) handleCommand(update *tgbotapi.Update) {
	switch MessageCommand(update.Message.Command()) { // Check for non admin commands
	case MessageCmdStart, MessageCmdSubscriptions:
		sendSubscriptions(update)
		client.setState(update, StateNil, nil)
	case MessageCmdProposals:
		sendCurrentProposals(update)
		client.setState(update, StateNil, nil)
	case MessageCmdHelp:
		sendHelp(update)
		client.setState(update, StateNil, nil)
	case MessageCmdSupport:
		sendSupport(update)
		client.setState(update, StateNil, nil)
	default:
		if isBotAdmin(update) { // Check for admin commands
			switch MessageCommand(update.Message.Command()) {
			case MessageCmdStats:
				sendUserStatistics(update)
				client.setState(update, StateNil, nil)
			case MessageCmdBroadcast:
				sendBroadcastStart(update)
				client.setState(update, StateStartBroadcast, nil)
			}
		}
	}
}

func (client TelegramClient) handleMessage(update *tgbotapi.Update) {
	switch client.getState(update) {
	case StateStartBroadcast:
		data := StateData{BroadcastStateData: &BroadcastStateData{Message: update.Message.Text}}
		sendConfirmBroadcastMessage(update, data.BroadcastStateData.Message)
		client.setState(update, StateConfirmBroadcast, &data)
	case StateConfirmBroadcast:
		yesOptions := []string{"yes", "y"}
		abortOptions := []string{"abort", "a"}
		if slices.Contains(yesOptions, strings.ToLower(update.Message.Text)) {
			data := client.getStateData(update)
			if data.BroadcastStateData == nil || data.BroadcastStateData.Message == "" {
				log.Sugar.Fatal("No message to broadcast. This should never happen!")
			}
			sendBroadcastMessage(data.BroadcastStateData.Message)
			sendBroadcastEndInfoMessage(update, true)
			client.setState(update, StateNil, nil)
		} else if slices.Contains(abortOptions, strings.ToLower(update.Message.Text)) {
			sendBroadcastEndInfoMessage(update, false)
			client.setState(update, StateNil, nil)
		} else {
			sendBroadcastStart(update)
			client.setState(update, StateStartBroadcast, nil)
		}
	}
}

func (client TelegramClient) handleCallbackQuery(update *tgbotapi.Update) {
	callbackData := ToCallbackData(update.CallbackQuery.Data)
	switch callbackData.Command {
	case CallbackCmdShowSubscriptions:
		performUpdateSubscription(update, callbackData.Data)
		sendSubscriptions(update)
	case CallbackCmdShowProposals:
		sendCurrentProposals(update)
	case CallbackCmdShowHelp:
		sendHelp(update)
	case CallbackCmdShowSupport:
		sendSupport(update)
	case CallbackCmdVote:
		performVote(update, callbackData.Data)
	default:
		if isBotAdmin(update) { // Check for admin callbacks
			switch callbackData.Command {
			case CallbackCmdStats:
				sendUserStatistics(update)
			default:
				sendError(update)
				sendHelp(update)
				client.setState(update, StateNil, nil)
			}
		} else {
			sendError(update)
			sendHelp(update)
			client.setState(update, StateNil, nil)
		}
	}
}

// groups -> just admins and creators can interact with the bot
// private -> everything is allowed
func (client TelegramClient) isInteractionAllowed(update *tgbotapi.Update) bool {
	if isGroupX(update) {
		return isUpdateFromCreatorOrAdministrator(update)
	}
	return true
}

// Handles updates for only 1 user in a serial way
func (client TelegramClient) handleUpdates(channel chan tgbotapi.Update) {
	for update := range channel {
		chatId := getChatIdX(&update)
		if client.isInteractionAllowed(&update) {
			if update.Message != nil && update.Message.IsCommand() {
				client.handleCommand(&update)
			} else if update.Message != nil && client.isExpectingMessage(&update) {
				client.handleMessage(&update)
			} else if update.CallbackQuery != nil {
				client.handleCallbackQuery(&update)
			}
		} else {
			log.Sugar.Debugf("Interaction with bot for user #%v is not allowed", chatId)
			if update.CallbackQuery != nil {
				answerCallbackQuery(&update)
			}
		}
		client.updateCountChannel <- UpdateCount{ChatId: chatId, Updates: -1}
	}
}

type UpdateCount struct {
	ChatId  int64
	Updates int
}

func (client TelegramClient) hasChannel(channelId int64) bool {
	for key := range client.updateChannels {
		if key == channelId {
			return true
		}
	}
	return false
}

func (client TelegramClient) sendToChannelAsync(chatId int64, update tgbotapi.Update) {
	client.updateCountChannel <- UpdateCount{ChatId: chatId, Updates: 1}
	client.updateChannels[chatId] <- update
}

func (client TelegramClient) sendToChannel(update *tgbotapi.Update) {
	chatId := getChatIdX(update)
	if !client.hasChannel(chatId) {
		client.updateChannels[chatId] = make(chan tgbotapi.Update)
		go client.handleUpdates(client.updateChannels[chatId])
	}
	go client.sendToChannelAsync(chatId, *update)
}

// Keeps track of all the user channels and closes them if there are no more updates
func (client TelegramClient) manageUpdateChannels() {
	var count = make(map[int64]int)
	for msg := range client.updateCountChannel {
		count[msg.ChatId] += msg.Updates
		if count[msg.ChatId] == 0 {
			close(client.updateChannels[msg.ChatId])
			delete(client.updateChannels, msg.ChatId)
			delete(count, msg.ChatId)
		}
	}
}

//goland:noinspection GoNameStartsWithPackageName
type TelegramClient struct {
	//userManager *database.UserManager
	api *tgbotapi.BotAPI

	// updateChannels contains one update channel for every user.
	// This means the updates can be processed parallel for multiple users but serial for every single user
	updateChannels map[int64]chan tgbotapi.Update

	// updateCountChannel is used to communicate to `manageUpdateChannels` from `handleUpdates`
	updateCountChannel chan UpdateCount
	state              map[int64]State
	stateData          map[int64]StateData
}

func NewTelegramClient() *TelegramClient {
	return &TelegramClient{
		api:                getApi(),
		updateChannels:     make(map[int64]chan tgbotapi.Update),
		updateCountChannel: make(chan UpdateCount),
		state:              make(map[int64]State),
		stateData:          make(map[int64]StateData),
	}
}

func (client TelegramClient) Start() {
	log.Sugar.Info("Start telegram bot")

	mHack = database.NewDefaultDbManagers()

	updateConfig := tgbotapi.NewUpdate(0)
	updates := client.api.GetUpdatesChan(updateConfig)

	go client.manageUpdateChannels()

	for update := range updates {
		if !hasChatId(&update) { // no chat id means there is something strange or the update is not for us
			continue
		}

		client.sendToChannel(&update)
	}
}
