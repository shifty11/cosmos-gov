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
	client       *ent.Client
	ctx          context.Context
	userManager  *UserManager
	chainManager *ChainManager
}

func NewSubscriptionManager(userManager *UserManager, chainManager *ChainManager) *SubscriptionManager {
	client, ctx := connect()
	return &SubscriptionManager{client: client, ctx: ctx, userManager: userManager, chainManager: chainManager}
}

func getSubscriptions(chainsOfUser []*ent.Chain) []*Subscription {
	allChains := NewChainManager().Enabled()
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
		return NewTelegramChatManager().AddOrRemoveChain(chatRoomId, chainName)
	} else {
		return NewDiscordChannelManager().AddOrRemoveChain(chatRoomId, chainName)
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
				Subscriptions: getSubscriptions(tgChat.Edges.Chains),
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
				Subscriptions: getSubscriptions(dChannel.Edges.Chains),
			})
		}
		return chats
	}
}

type TelegramSubscriptionManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewTelegramSubscriptionManager() *TelegramSubscriptionManager {
	client, ctx := connect()
	return &TelegramSubscriptionManager{client: client, ctx: ctx}
}

type DiscordSubscriptionManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewDiscordSubscriptionManager() *DiscordSubscriptionManager {
	client, ctx := connect()
	return &DiscordSubscriptionManager{client: client, ctx: ctx}
}

func (manager *TelegramSubscriptionManager) GetOrCreateSubscriptions(userId int64, userName string, chatId int64, chatName string, isGroup bool) []*Subscription {
	userManager := NewTypedUserManager(user.TypeTelegram)
	entUser := userManager.GetOrCreateUser(userId, userName)
	tgChatManager := NewTelegramChatManager()
	tgChat := tgChatManager.GetOrCreateTelegramChat(entUser, chatId, chatName, isGroup)

	chainsOfUser, err := tgChat.QueryChains().All(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while fetching chains for chat %v (%v): %v", tgChat.Name, tgChat.ID, err)
	}
	return getSubscriptions(chainsOfUser)
}

func (manager *DiscordSubscriptionManager) GetOrCreateSubscriptions(userId int64, userName string, channelId int64, channelName string, isGroup bool) []*Subscription {
	userManager := NewTypedUserManager(user.TypeDiscord)
	entUser := userManager.GetOrCreateUser(userId, userName)
	dChannelManager := NewDiscordChannelManager()
	dChannel := dChannelManager.GetOrCreateDiscordChannel(entUser, channelId, channelName, isGroup)

	chainsOfUser, err := dChannel.QueryChains().All(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while fetching chains for chat %v (%v): %v", dChannel.Name, dChannel.ID, err)
	}
	return getSubscriptions(chainsOfUser)
}
