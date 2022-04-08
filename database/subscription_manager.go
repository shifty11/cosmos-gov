package database

import (
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
)

type Subscription struct {
	Name        string
	DisplayName string
	Notify      bool
}

type SubscriptionManager struct {
	userManager *UserManager
}

func NewSubscriptionManager(userManager *UserManager) *SubscriptionManager {
	return &SubscriptionManager{userManager: userManager}
}

func (manager *SubscriptionManager) ToggleSubscription(chatId int64, userType user.Type, chainName string) (bool, error) {
	_, ctx := connect()
	userDto, err := manager.userManager.GetUser(chatId, userType)
	if err != nil {
		return false, err
	}
	chainDto, err := getChainByName(chainName)
	if err != nil {
		return false, err
	}
	exists, err := userDto.
		QueryChains().
		Where(chain.IDEQ(chainDto.ID)).
		Exist(ctx)
	if err != nil {
		return false, err
	}
	if exists {
		_, err = userDto.
			Update().
			RemoveChainIDs(chainDto.ID).
			Save(ctx)
	} else {
		_, err = userDto.
			Update().
			AddChainIDs(chainDto.ID).
			Save(ctx)
	}
	return !exists, err
}

func (manager *SubscriptionManager) GetSubscriptions(chatId int64, userType user.Type) []Subscription {
	client, ctx := connect()
	var userDto = getOrCreateUser(chatId, userType)
	chainsOfUser, err := client.Chain.
		Query().
		Where(chain.HasUsersWith(user.ID(userDto.ID))).
		All(ctx)
	if err != nil {
		log.Sugar.Panic("Error while querying chains for user %v: %v", userDto.ID, err)
	}
	allChains, err := client.Chain.
		Query().
		Where(chain.IsEnabledEQ(true)).
		Order(ent.Asc(chain.FieldDisplayName)).
		All(ctx)
	var chains []Subscription
	for _, c := range allChains {
		var chainEntry = Subscription{Name: c.Name, DisplayName: c.DisplayName, Notify: false}
		for _, nc := range chainsOfUser { // check if user gets notified for this chain (c)
			if nc.ID == c.ID {
				chainEntry.Notify = true
			}
		}
		chains = append(chains, chainEntry)
	}
	return chains
}
