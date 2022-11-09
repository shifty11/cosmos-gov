package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
	regen "github.com/zach-klippenstein/goregen"
	"os"
)

type UserManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewUserManager(client *ent.Client, ctx context.Context) *UserManager {
	return &UserManager{client: client, ctx: ctx}
}

func (manager *UserManager) Get(userId int64, userType user.Type) (*ent.User, error) {
	return manager.client.User.
		Query().
		Where(user.And(
			user.UserIDEQ(userId),
			user.TypeEQ(userType),
		)).
		Only(manager.ctx)
}

func (manager *UserManager) ByToken(userId int64, userType user.Type, token string) (*ent.User, error) {
	return manager.client.User.
		Query().
		Where(user.And(
			user.UserIDEQ(userId),
			user.TypeEQ(userType),
			user.LoginTokenEQ(token),
		)).
		Only(manager.ctx)
}

func (manager *UserManager) GenerateNewLoginToken(userId int64, userType user.Type) error {
	token, err := regen.Generate("[A-Za-z0-9]{32}")
	if err != nil {
		return err
	}
	_, err = manager.client.User.
		Update().
		Where(user.And(
			user.UserIDEQ(userId),
			user.TypeEQ(userType),
		)).
		SetLoginToken(token).
		Save(manager.ctx)
	return err
}

func (manager *UserManager) SetName(entUser *ent.User, name string) (*ent.User, error) {
	return entUser.Update().
		SetName(name).
		Save(manager.ctx)
}

type TypedUserManager struct {
	client   *ent.Client
	ctx      context.Context
	userType user.Type
}

func NewTypedUserManager(client *ent.Client, ctx context.Context, userType user.Type) *TypedUserManager {
	return &TypedUserManager{client: client, ctx: ctx, userType: userType}
}

func (manager *TypedUserManager) Exists(id int64, userType user.Type) bool {
	exists, err := manager.client.User.
		Query().
		Where(user.And(
			user.UserIDEQ(id),
			user.TypeEQ(userType),
		)).
		Exist(manager.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while checking for existence of user #%v (%v)", id, userType)
	}
	return exists
}

func (manager *TypedUserManager) GetOrCreateUser(id int64, name string) *ent.User {
	userEnt, err := manager.client.User.
		Query().
		Where(
			user.And(
				user.UserIDEQ(id),
				user.TypeEQ(manager.userType),
			)).
		Only(manager.ctx)
	if err != nil {
		token, err := regen.Generate("[A-Za-z0-9]{32}")
		if err != nil {
			log.Sugar.Errorf("Could not generate new login token: %v", err)
		}
		userEnt, err = manager.client.User.
			Create().
			SetUserID(id).
			SetName(name).
			SetType(manager.userType).
			SetLoginToken(token).
			Save(manager.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while creating user: %v", err)
		}
	}
	return userEnt
}

type TelegramUser struct {
	UserId   int64  `json:"user_id"`
	Name     string `json:"name"`
	ChatId   int64  `json:"chat_id"`
	ChatName string `json:"chat_name"`
	IsGroup  bool   `json:"is_group"`
	ChainId  string `json:"chain_id"`
}

type DiscordUser struct {
	UserId      int64  `json:"user_id"`
	Name        string `json:"name"`
	ChannelId   int64  `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	IsGroup     bool   `json:"is_group"`
	ChainId     string `json:"chain_id"`
}

func (manager *UserManager) DumpDb() {
	manager.client.Chain.Update().SetIsEnabled(true).ExecX(manager.ctx)

	users, err := manager.client.User.
		Query().
		WithTelegramChats().
		WithDiscordChannels().
		All(manager.ctx)
	if err != nil {
		log.Sugar.Errorf("Error while dumping db: %v", err)
	}
	var tgUsers []TelegramUser
	var zeroIdTgUsers []TelegramUser
	var discordUsers []DiscordUser
	var zeroIdDiscordUsers []DiscordUser
	for _, u := range users {
		if u.Type == user.TypeTelegram {
			if u.UserID == 0 {
				zeroIdTgUsers = append(zeroIdTgUsers, TelegramUser{
					UserId: u.UserID,
					Name:   u.Name,
				})
				for _, c := range u.Edges.TelegramChats {
					zeroIdTgUsers = append(zeroIdTgUsers, TelegramUser{
						UserId:   u.UserID,
						Name:     u.Name,
						ChatId:   c.ChatID,
						ChatName: c.Name,
						IsGroup:  c.IsGroup,
					})
					for _, sub := range c.QueryChains().AllX(manager.ctx) {
						zeroIdTgUsers = append(zeroIdTgUsers, TelegramUser{
							UserId:   u.UserID,
							Name:     u.Name,
							ChatId:   c.ChatID,
							ChatName: c.Name,
							IsGroup:  c.IsGroup,
							ChainId:  sub.ChainID,
						})
					}
				}
			} else {
				tgUsers = append(tgUsers, TelegramUser{
					UserId: u.UserID,
					Name:   u.Name,
				})
				for _, c := range u.Edges.TelegramChats {
					if c.ChatID == 0 {
						panic("ChatID is 0")
					}
					tgUsers = append(tgUsers, TelegramUser{
						UserId:   u.UserID,
						Name:     u.Name,
						ChatId:   c.ChatID,
						ChatName: c.Name,
						IsGroup:  c.IsGroup,
					})
					for _, sub := range c.QueryChains().AllX(manager.ctx) {
						tgUsers = append(tgUsers, TelegramUser{
							UserId:   u.UserID,
							Name:     u.Name,
							ChatId:   c.ChatID,
							ChatName: c.Name,
							IsGroup:  c.IsGroup,
							ChainId:  sub.ChainID,
						})
					}
				}
			}
		} else if u.Type == user.TypeDiscord {
			if u.UserID == 0 {
				zeroIdDiscordUsers = append(zeroIdDiscordUsers, DiscordUser{
					UserId: u.UserID,
					Name:   u.Name,
				})
				for _, c := range u.Edges.DiscordChannels {
					zeroIdDiscordUsers = append(zeroIdDiscordUsers, DiscordUser{
						UserId:      u.UserID,
						Name:        u.Name,
						ChannelId:   c.ChannelID,
						ChannelName: c.Name,
						IsGroup:     c.IsGroup,
					})
					for _, sub := range c.QueryChains().AllX(manager.ctx) {
						zeroIdDiscordUsers = append(zeroIdDiscordUsers, DiscordUser{
							UserId:      u.UserID,
							Name:        u.Name,
							ChannelId:   c.ChannelID,
							ChannelName: c.Name,
							IsGroup:     c.IsGroup,
							ChainId:     sub.ChainID,
						})
					}
				}
			} else {
				discordUsers = append(discordUsers, DiscordUser{
					UserId: u.UserID,
					Name:   u.Name,
				})
				for _, c := range u.Edges.DiscordChannels {
					if c.ChannelID == 0 {
						panic("ChannelID is 0")
					}
					discordUsers = append(discordUsers, DiscordUser{
						UserId:      u.UserID,
						Name:        u.Name,
						ChannelId:   c.ChannelID,
						ChannelName: c.Name,
						IsGroup:     c.IsGroup,
					})
					for _, sub := range c.QueryChains().AllX(manager.ctx) {
						discordUsers = append(discordUsers, DiscordUser{
							UserId:      u.UserID,
							Name:        u.Name,
							ChannelId:   c.ChannelID,
							ChannelName: c.Name,
							IsGroup:     c.IsGroup,
							ChainId:     sub.ChainID,
						})
					}
				}
			}
		}
	}

	if len(tgUsers) > 0 {
		writeToFile(tgUsers, "telegram_users.json")
	}
	if len(discordUsers) > 0 {
		writeToFile(discordUsers, "discord_users.json")
	}
	if len(zeroIdTgUsers) > 0 {
		writeToFile(zeroIdTgUsers, "telegram_users_zero_id.json")
	}
	if len(zeroIdDiscordUsers) > 0 {
		writeToFile(zeroIdDiscordUsers, "discord_users_zero_id.json")
	}
}

func writeToFile(thing any, fileName string) {
	if thing == nil {
		return
	}

	content, err := json.Marshal(thing)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	file, err := os.Create(fileName)
	if err != nil {
		log.Sugar.Errorf("Error creating file: %s", err)
		return
	}
	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		log.Sugar.Errorf("Error writing to file: %s", err)
		return
	}
	log.Sugar.Infof("Wrote %s", fileName)
}
