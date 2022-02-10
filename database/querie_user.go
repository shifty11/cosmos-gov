package database

import (
	"errors"
	"github.com/shifty11/cosmos-gov/dtos"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
	"time"
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

func DeleteUser(chatId int64) {
	client, ctx := connect()
	_, err := client.User.
		Delete().
		Where(user.ChatIDEQ(chatId)).
		Exec(ctx)
	if err != nil {
		log.Sugar.Errorf("Error while deleting user: %v", err)
	}
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

func GetUserStatistics() (*dtos.UserStatistic, error) {
	client, ctx := connect()
	cntAll, err := client.User.
		Query().
		Count(ctx)
	if err != nil {
		return nil, err
	}
	cntSinceYesterday, err := client.User.
		Query().
		Where(user.CreatedAtGTE(time.Now().AddDate(0, 0, -1))).
		Count(ctx)
	if err != nil {
		return nil, err
	}
	cntSinceSevenDays, err := client.User.
		Query().
		Where(user.CreatedAtGTE(time.Now().AddDate(0, 0, -7))).
		Count(ctx)
	if err != nil {
		return nil, err
	}
	if cntAll == 0 {
		return nil, errors.New("no users -> division with 0 not allowed")
	}
	changeSinceYesterdayInPercent := float64(cntSinceYesterday) / float64(cntAll) * 100
	changeThisWeekInPercent := float64(cntSinceSevenDays) / float64(cntAll) * 100
	return &dtos.UserStatistic{
		CntUsers:                      cntAll,
		CntUsersSinceYesterday:        cntSinceYesterday,
		CntUsersThisWeek:              cntSinceSevenDays,
		ChangeSinceYesterdayInPercent: changeSinceYesterdayInPercent,
		ChangeThisWeekInPercent:       changeThisWeekInPercent,
	}, nil
}

func GetAllUserChatIds() []int {
	client, ctx := connect()
	chatIds, err := client.User.
		Query().
		Select(user.FieldChatID).
		Ints(ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying chatIds of all users: %v", err)
	}
	return chatIds
}

func CountUsers() int {
	client, ctx := connect()
	cnt, err := client.User.
		Query().
		Count(ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying chatIds of all users: %v", err)
	}
	return cnt
}
