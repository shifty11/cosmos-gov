package database

import (
	"context"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/telegramchat"
	"github.com/shifty11/cosmos-gov/log"
)

type TelegramChatManager struct {
	client       *ent.Client
	ctx          context.Context
	chainManager *ChainManager
}

func NewTelegramChatManager(client *ent.Client, ctx context.Context, chainManager *ChainManager) *TelegramChatManager {
	return &TelegramChatManager{client: client, ctx: ctx, chainManager: chainManager}
}

func (manager *TelegramChatManager) AddOrRemoveChain(tgChatId int64, chainName string) (bool, error) {
	tgChat, err := manager.client.TelegramChat.
		Query().
		Where(telegramchat.ChatIDEQ(tgChatId)).
		First(manager.ctx)
	if err != nil {
		return false, err
	}

	chainEnt, err := manager.chainManager.ByName(chainName)
	if err != nil {
		return false, err
	}

	exists, err := tgChat.
		QueryChains().
		Where(chain.IDEQ(chainEnt.ID)).
		Exist(manager.ctx)
	if err != nil {
		return false, err
	}
	if exists {
		_, err := tgChat.
			Update().
			RemoveChainIDs(chainEnt.ID).
			Save(manager.ctx)
		if err != nil {
			return false, err
		}
	} else {
		_, err := tgChat.
			Update().
			AddChainIDs(chainEnt.ID).
			Save(manager.ctx)
		if err != nil {
			return false, err
		}
	}
	return !exists, nil
}

// TODO: remove after full migration
func setUserIfNotPresent(tgChat *ent.TelegramChat, manager *TelegramChatManager, oldErr error, tgChatId int64, entUser *ent.User) (*ent.TelegramChat, error) {
	if oldErr != nil {
		tgChat, err := manager.client.TelegramChat.
			Query().
			Where(telegramchat.ChatIDEQ(tgChatId)).
			First(manager.ctx)
		if err != nil {
			return nil, oldErr
		}
		tgChat, err = tgChat.
			Update().
			SetUser(entUser).
			Save(manager.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while updating telegram chat: %v", err)
		}
		return tgChat, nil
	}
	return tgChat, oldErr
}

func (manager *TelegramChatManager) GetOrCreateTelegramChat(entUser *ent.User, tgChatId int64, name string, isGroup bool) *ent.TelegramChat {
	tgChat, err := entUser.
		QueryTelegramChats().
		Where(telegramchat.ChatIDEQ(tgChatId)).
		Only(manager.ctx)

	tgChat, err = setUserIfNotPresent(tgChat, manager, err, tgChatId, entUser)

	if err != nil {
		tgChat, err = manager.client.TelegramChat.
			Create().
			SetChatID(tgChatId).
			SetName(name).
			SetIsGroup(isGroup).
			SetUser(entUser).
			Save(manager.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while creating telegram chat: %v", err)
		}
	}
	return tgChat
}

func (manager *TelegramChatManager) Delete(chatId int64) {
	log.Sugar.Debugf("Delete Telegram chat %v", chatId)
	_, err := manager.client.TelegramChat.
		Delete().
		Where(telegramchat.ChatIDEQ(chatId)).
		Exec(manager.ctx)
	if err != nil {
		log.Sugar.Errorf("Error while deleting telegram chat: %v", err)
	}
}

func (manager *TelegramChatManager) DeleteMultiple(chatIds []int64) {
	log.Sugar.Debugf("Delete %v Telegram chat's", len(chatIds))
	_, err := manager.client.TelegramChat.
		Delete().
		Where(telegramchat.ChatIDIn(chatIds...)).
		Exec(manager.ctx)
	if err != nil {
		log.Sugar.Errorf("Error while deleting Telegram channels: %v", err)
	}
}

func (manager *TelegramChatManager) GetChatIds(entChain *ent.Chain) []int {
	chatIds, err := entChain.
		QueryTelegramChats().
		Select(telegramchat.FieldID).
		Ints(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying Telegram chatIds for chain %v: %v", entChain.Name, err)
	}
	return chatIds
}

func (manager *TelegramChatManager) GetAllChatIds() []int {
	chatIds, err := manager.client.TelegramChat.
		Query().
		Select(telegramchat.FieldID).
		Ints(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying all Telegram chatIds: %v", err)
	}
	return chatIds
}

func (manager *TelegramChatManager) CountChats() int {
	cnt, err := manager.client.TelegramChat.
		Query().
		Count(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while counting all Telegram chatIds: %v", err)
	}
	return cnt
}
