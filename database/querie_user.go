package database

import (
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
)

func getOrCreateUser(chatId int64) *ent.User {
	client, ctx := connect()
	var userDto *ent.User
	var err error
	userDto, err = client.User.
		Query().
		Where(user.ChatIDEQ(chatId)).
		Only(ctx)
	if err != nil {
		userDto, err = client.User.
			Create().
			SetChatID(chatId).
			Save(ctx)
		if err != nil {
			log.Sugar.Panic("Error while creating user: %v", err)
		}
	}
	return userDto
}

func DeleteUsers(chatIds map[int]struct{}) {
	var chatIds64 []int64
	for chatId := range chatIds {
		chatIds64 = append(chatIds64, int64(chatId))
	}
	client, ctx := connect()
	_, err := client.User.
		Delete().
		Where(user.ChatIDIn(chatIds64...)).
		Exec(ctx)
	if err != nil {
		log.Sugar.Errorf("Error while deleting user: %v", err)
	}
}
