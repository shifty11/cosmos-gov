package database

import (
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/user"
	regen "github.com/zach-klippenstein/goregen"
)

type UserManager struct {
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (server *UserManager) GetUser(chatId int64, userType user.Type, token string) (*ent.User, error) {
	client, ctx := connect()
	entUser, err := client.User.
		Query().
		Where(user.And(
			user.ChatIDEQ(chatId),
			user.TypeEQ(userType),
		)).
		Only(ctx)
	return entUser, err
}

func (server *UserManager) GenerateNewLoginToken(chatId int64, userType user.Type) error {
	token, err := regen.Generate("[A-Za-z0-9]{32}")
	if err != nil {
		return err
	}
	client, ctx := connect()
	_, err = client.User.
		Update().
		Where(user.And(
			user.ChatIDEQ(chatId),
			user.TypeEQ(userType),
		)).
		SetLogingToken(token).
		Save(ctx)
	return err
}
