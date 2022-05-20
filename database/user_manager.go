package database

import (
	"context"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
	regen "github.com/zach-klippenstein/goregen"
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
