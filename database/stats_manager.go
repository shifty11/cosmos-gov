package database

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"errors"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/discordchannel"
	"github.com/shifty11/cosmos-gov/ent/proposal"
	"github.com/shifty11/cosmos-gov/ent/telegramchat"
	"github.com/shifty11/cosmos-gov/ent/user"
	"time"
)

type StatsManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewStatsManager() *StatsManager {
	client, ctx := connect()
	return &StatsManager{client: client, ctx: ctx}
}

type UserStatistic struct {
	CntUsers                      int
	CntUsersSinceYesterday        int
	CntUsersThisWeek              int
	ChangeSinceYesterdayInPercent float64
	ChangeThisWeekInPercent       float64
}

type ChainStatistic struct {
	DisplayName   string `json:"display_name"`
	Proposals     int
	Subscriptions int
	IsEnabled     bool `json:"is_enabled"`
}

func (manager *StatsManager) getChainStats(userType user.Type) ([]*ChainStatistic, error) {
	client, ctx := connect()
	var chainsWithNotifications []ChainStatistic
	err := client.Chain.Query().
		Order(ent.Desc(chain.FieldIsEnabled), ent.Asc(chain.FieldDisplayName)).
		GroupBy(chain.FieldIsEnabled, chain.FieldDisplayName).
		Aggregate(
			func(s *sql.Selector) string {
				if userType == user.TypeTelegram {
					t := sql.Table(chain.TelegramChatsTable)
					s.Join(t).On(s.C(chain.FieldID), t.C(telegramchat.ChainsPrimaryKey[1]))
					return sql.As(sql.Count(t.C(telegramchat.ChainsPrimaryKey[1])), "subscriptions")
				}
				t := sql.Table(chain.DiscordChannelsTable)
				s.Join(t).On(s.C(chain.FieldID), t.C(discordchannel.ChainsPrimaryKey[1]))
				return sql.As(sql.Count(t.C(discordchannel.ChainsPrimaryKey[1])), "subscriptions")
			},
		).
		Scan(ctx, &chainsWithNotifications)
	if err != nil {
		return nil, err
	}
	var chainsWithProposals []ChainStatistic
	err = client.Chain.Query().
		Order(ent.Desc(chain.FieldIsEnabled), ent.Asc(chain.FieldDisplayName)).
		GroupBy(chain.FieldIsEnabled, chain.FieldDisplayName).
		Aggregate(
			func(s *sql.Selector) string {
				t := sql.Table(chain.ProposalsTable)
				s.Join(t).On(s.C(chain.FieldID), t.C(proposal.ChainColumn))
				return sql.As(sql.Count(t.C(proposal.FieldID)), "proposals")
			},
		).
		Scan(ctx, &chainsWithProposals)
	if err != nil {
		return nil, err
	}
	var stats []*ChainStatistic
	for _, cp := range chainsWithProposals {
		found := false
		for _, cn := range chainsWithNotifications {
			if cp.DisplayName == cn.DisplayName {
				stats = append(stats, &ChainStatistic{
					DisplayName:   cp.DisplayName,
					Proposals:     cp.Proposals,
					Subscriptions: cn.Subscriptions,
				})
				found = true
			}
		}
		if !found {
			stats = append(stats, &ChainStatistic{
				DisplayName:   cp.DisplayName,
				Proposals:     cp.Proposals,
				Subscriptions: 0,
			})
		}
	}
	return stats, err
}

func (manager *StatsManager) GetChainStats() ([]*ChainStatistic, error) {
	tgStats, err := manager.getChainStats(user.TypeTelegram)
	if err != nil {
		return nil, err
	}

	dStats, err := manager.getChainStats(user.TypeDiscord)
	if err != nil {
		return nil, err
	}
	for i, stats := range tgStats {
		stats.Subscriptions += dStats[i].Subscriptions
	}
	return tgStats, nil
}

func (manager *StatsManager) GetUserStatistics(userType user.Type) (*UserStatistic, error) {
	client, ctx := connect()
	cntAll, err := client.User.
		Query().
		Where(user.TypeEQ(userType)).
		Count(ctx)
	if err != nil {
		return nil, err
	}
	cntSinceYesterday, err := client.User.
		Query().
		Where(user.And(
			user.CreateTimeGTE(time.Now().AddDate(0, 0, -1)),
			user.TypeEQ(userType),
		)).
		Count(ctx)
	if err != nil {
		return nil, err
	}
	cntSinceSevenDays, err := client.User.
		Query().
		Where(user.And(
			user.CreateTimeGTE(time.Now().AddDate(0, 0, -7)),
			user.TypeEQ(userType),
		)).
		Count(ctx)
	if err != nil {
		return nil, err
	}
	if cntAll == 0 {
		return nil, errors.New("no users -> division with 0 not allowed")
	}
	changeSinceYesterdayInPercent := float64(cntSinceYesterday) / float64(cntAll) * 100
	changeThisWeekInPercent := float64(cntSinceSevenDays) / float64(cntAll) * 100
	return &UserStatistic{
		CntUsers:                      cntAll,
		CntUsersSinceYesterday:        cntSinceYesterday,
		CntUsersThisWeek:              cntSinceSevenDays,
		ChangeSinceYesterdayInPercent: changeSinceYesterdayInPercent,
		ChangeThisWeekInPercent:       changeThisWeekInPercent,
	}, nil
}
