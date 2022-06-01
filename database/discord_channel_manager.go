package database

import (
	"context"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/discordchannel"
	"github.com/shifty11/cosmos-gov/log"
)

type DiscordChannelManager struct {
	client       *ent.Client
	ctx          context.Context
	chainManager *ChainManager
}

func NewDiscordChannelManager(client *ent.Client, ctx context.Context, chainManager *ChainManager) *DiscordChannelManager {
	return &DiscordChannelManager{client: client, ctx: ctx, chainManager: chainManager}
}

func (manager *DiscordChannelManager) AddOrRemoveChain(dChannelId int64, chainName string) (bool, error) {
	dChannel, err := manager.client.DiscordChannel.
		Query().
		Where(discordchannel.ChannelID(dChannelId)).
		First(manager.ctx)
	if err != nil {
		return false, err
	}

	chainEnt, err := manager.chainManager.ByName(chainName)
	if err != nil {
		return false, err
	}

	exists, err := dChannel.
		QueryChains().
		Where(chain.IDEQ(chainEnt.ID)).
		Exist(manager.ctx)
	if err != nil {
		return false, err
	}
	if exists {
		_, err := dChannel.
			Update().
			RemoveChainIDs(chainEnt.ID).
			Save(manager.ctx)
		if err != nil {
			return false, err
		}
	} else {
		_, err := dChannel.
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
func setUserIfNotPresentD(channel *ent.DiscordChannel, manager *DiscordChannelManager, oldErr error, channelId int64, entUser *ent.User) (*ent.DiscordChannel, error) {
	if oldErr != nil {
		channel, err := manager.client.DiscordChannel.
			Query().
			Where(discordchannel.ChannelIDEQ(channelId)).
			First(manager.ctx)
		if err != nil {
			return nil, oldErr
		}
		channel, err = channel.
			Update().
			SetUser(entUser).
			Save(manager.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while updating telegram chat: %v", err)
		}
		return channel, nil
	}
	return channel, oldErr
}

func (manager *DiscordChannelManager) GetOrCreateDiscordChannel(entUser *ent.User, channelId int64, name string, isGroup bool) *ent.DiscordChannel {
	dChannel, err := entUser.
		QueryDiscordChannels().
		Where(discordchannel.ChannelIDEQ(channelId)).
		Only(manager.ctx)

	dChannel, err = setUserIfNotPresentD(dChannel, manager, err, channelId, entUser)

	if err != nil {
		dChannel, err = manager.client.DiscordChannel.
			Create().
			SetChannelID(channelId).
			SetName(name).
			SetIsGroup(isGroup).
			SetUser(entUser).
			Save(manager.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while creating discord channel: %v", err)
		}
	}
	return dChannel
}

func (manager *DiscordChannelManager) DeleteMultiple(channelIds []int64) {
	log.Sugar.Debugf("Delete %v Discord channel's", len(channelIds))
	_, err := manager.client.DiscordChannel.
		Delete().
		Where(discordchannel.ChannelIDIn(channelIds...)).
		Exec(manager.ctx)
	if err != nil {
		log.Sugar.Errorf("Error while deleting Discord channels: %v", err)
	}
}

func (manager *DiscordChannelManager) GetChannelIds(entChain *ent.Chain) []int {
	channelIds, err := entChain.
		QueryDiscordChannels().
		Select(discordchannel.FieldChannelID).
		Ints(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying Discord channelIds for chain %v: %v", entChain.Name, err)
	}
	return channelIds
}

func (manager *DiscordChannelManager) GetChannelIdsWithDraftPropsEnabled(entChain *ent.Chain) []int {
	channelIds, err := entChain.
		QueryDiscordChannels().
		Where(discordchannel.WantsDraftProposalsEQ(true)).
		Select(discordchannel.FieldChannelID).
		Ints(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying Discord channelIds for chain %v: %v", entChain.Name, err)
	}
	return channelIds
}
