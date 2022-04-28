package database

import (
	"context"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
)

type Subscription struct {
	Name        string
	DisplayName string
	Notify      bool
}

type ChatRoom struct {
	Id            int64
	Name          string
	Subscriptions []*Subscription
}

type SubscriptionManager struct {
	client                *ent.Client
	ctx                   context.Context
	userManager           *UserManager
	chainManager          *ChainManager
	telegramChatManager   *TelegramChatManager
	discordChannelManager *DiscordChannelManager
}

func NewSubscriptionManager(
	client *ent.Client,
	ctx context.Context,
	userManager *UserManager,
	chainManager *ChainManager,
	telegramChatManager *TelegramChatManager,
	discordChannelManager *DiscordChannelManager,
) *SubscriptionManager {
	return &SubscriptionManager{
		client:                client,
		ctx:                   ctx,
		userManager:           userManager,
		chainManager:          chainManager,
		telegramChatManager:   telegramChatManager,
		discordChannelManager: discordChannelManager,
	}
}

func getSubscriptions(chainManager *ChainManager, chainsOfUser []*ent.Chain) []*Subscription {
	allChains := chainManager.Enabled()
	var chains []*Subscription
	for _, c := range allChains {
		var chainEntry = Subscription{Name: c.Name, DisplayName: c.DisplayName, Notify: false}
		for _, nc := range chainsOfUser { // check if user gets notified for this chain (c)
			if nc.ID == c.ID {
				chainEntry.Notify = true
			}
		}
		chains = append(chains, &chainEntry)
	}
	return chains
}

func (manager *SubscriptionManager) ToggleSubscription(entUser *ent.User, chatRoomId int64, chainName string) (bool, error) {
	if entUser.Type == user.TypeTelegram {
		return manager.telegramChatManager.AddOrRemoveChain(chatRoomId, chainName)
	} else {
		return manager.discordChannelManager.AddOrRemoveChain(chatRoomId, chainName)
	}
}

func (manager *SubscriptionManager) GetSubscriptions(entUser *ent.User) []*ChatRoom {
	if entUser.Type == user.TypeTelegram {
		tgChats, err := entUser.
			QueryTelegramChats().
			WithChains().
			All(manager.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while querying telegram chats of user %v (%v): %v", entUser.Name, entUser.ID, err)
		}

		var chats []*ChatRoom
		for _, tgChat := range tgChats {
			chats = append(chats, &ChatRoom{
				Id:            tgChat.ID,
				Name:          tgChat.Name,
				Subscriptions: getSubscriptions(manager.chainManager, tgChat.Edges.Chains),
			})
		}
		return chats
	} else {
		dChannels, err := entUser.
			QueryDiscordChannels().
			WithChains().
			All(manager.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while querying discord channels of user %v (%v): %v", entUser.Name, entUser.ID, err)
		}

		var chats []*ChatRoom
		for _, dChannel := range dChannels {
			chats = append(chats, &ChatRoom{
				Id:            dChannel.ID,
				Name:          dChannel.Name,
				Subscriptions: getSubscriptions(manager.chainManager, dChannel.Edges.Chains),
			})
		}
		return chats
	}
}

type TelegramSubscriptionManager struct {
	client        *ent.Client
	ctx           context.Context
	userManager   *TypedUserManager
	chainManager  *ChainManager
	tgChatManager *TelegramChatManager
}

func NewTelegramSubscriptionManager(
	client *ent.Client,
	ctx context.Context,
	userManager *TypedUserManager,
	chainManager *ChainManager,
	tgChatManager *TelegramChatManager,
) *TelegramSubscriptionManager {
	return &TelegramSubscriptionManager{
		client:        client,
		ctx:           ctx,
		userManager:   userManager,
		chainManager:  chainManager,
		tgChatManager: tgChatManager,
	}
}

type DiscordSubscriptionManager struct {
	client                *ent.Client
	ctx                   context.Context
	userManager           *TypedUserManager
	chainManager          *ChainManager
	discordChannelManager *DiscordChannelManager
}

func NewDiscordSubscriptionManager(
	client *ent.Client,
	ctx context.Context,
	userManager *TypedUserManager,
	chainManager *ChainManager,
	discordChannelManager *DiscordChannelManager,
) *DiscordSubscriptionManager {
	return &DiscordSubscriptionManager{
		client:                client,
		ctx:                   ctx,
		userManager:           userManager,
		chainManager:          chainManager,
		discordChannelManager: discordChannelManager,
	}
}

func (manager *TelegramSubscriptionManager) GetOrCreateSubscriptions(userId int64, userName string, chatId int64, chatName string, isGroup bool) []*Subscription {
	entUser := manager.userManager.GetOrCreateUser(userId, userName)
	tgChat := manager.tgChatManager.GetOrCreateTelegramChat(entUser, chatId, chatName, isGroup)

	chainsOfUser, err := tgChat.QueryChains().All(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while fetching chains for chat %v (%v): %v", tgChat.Name, tgChat.ID, err)
	}
	return getSubscriptions(manager.chainManager, chainsOfUser)
}

func (manager *DiscordSubscriptionManager) GetOrCreateSubscriptions(userId int64, userName string, channelId int64, channelName string, isGroup bool) []*Subscription {
	entUser := manager.userManager.GetOrCreateUser(userId, userName)
	dChannel := manager.discordChannelManager.GetOrCreateDiscordChannel(entUser, channelId, channelName, isGroup)

	chainsOfUser, err := dChannel.QueryChains().All(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while fetching chains for chat %v (%v): %v", dChannel.Name, dChannel.ID, err)
	}
	return getSubscriptions(manager.chainManager, chainsOfUser)
}
