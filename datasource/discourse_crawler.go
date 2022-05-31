package datasource

import (
	"context"
	"fmt"
	"github.com/FrenchBen/godisco"
	_ "github.com/lib/pq"
	"github.com/shifty11/cosmos-gov/api/discord"
	"github.com/shifty11/cosmos-gov/api/telegram"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/log"
	"golang.org/x/exp/slices"
)

type DiscourseCrawler struct {
	ctx                   context.Context
	chainManager          *database.ChainManager
	telegramChatManager   *database.TelegramChatManager
	discordChannelManager *database.DiscordChannelManager
	proposalManager       *database.ProposalManager
	draftProposalManager  *database.DraftProposalManager
	lensChainInfoManager  *database.LensChainInfoManager
	tgClient              *telegram.TelegramLightClient
	discordClient         *discord.DiscordLightClient
}

func NewDiscourseCrawler(ctx context.Context, managers database.DbManagers, tgClient *telegram.TelegramLightClient, discordClient *discord.DiscordLightClient) *DiscourseCrawler {
	return &DiscourseCrawler{
		ctx:                   ctx,
		chainManager:          managers.ChainManager,
		proposalManager:       managers.ProposalManager,
		draftProposalManager:  managers.DraftProposalManager,
		lensChainInfoManager:  managers.LensChainInfoManager,
		telegramChatManager:   managers.TelegramChatManager,
		discordChannelManager: managers.DiscordChannelManager,
		tgClient:              tgClient,
		discordClient:         discordClient,
	}
}

type Topic struct {
	Id    int      `json:"id"`
	Title string   `json:"title"`
	Slug  string   `json:"slug"`
	Tags  []string `json:"tags"`
}

func (t Topic) url(baseUrl string) string {
	return fmt.Sprintf("%v/t/%v", baseUrl, t.Id)
}

type TopicList struct {
	Topics []Topic `json:"topics"`
}

type TopicResponse struct {
	TopicList TopicList `json:"topic_list"`
}

func (dc DiscourseCrawler) getDraftProposals(url string) ([]Topic, error) {
	discourseClient, err := godisco.NewClient(url, "", "")
	if err != nil {
		log.Sugar.Fatal(err)
	}

	body, _, err := discourseClient.Get("/latest.json")
	if err != nil {
		return nil, err
	}
	response := TopicResponse{}
	err = json.Unmarshal(body, &response)

	log.Sugar.Info(response)

	var topics []Topic
	for _, topic := range response.TopicList.Topics {
		if slices.Contains(topic.Tags, "draft") {
			topics = append(topics, topic)
			log.Sugar.Infof("%v/t/%v/%v", url, topic.Slug, topic.Id)
		}
	}
	return topics, nil
}

func (dc DiscourseCrawler) saveAndSendDraftProposal(topic Topic, entChain *ent.Chain, url string) {
	prop, err := dc.draftProposalManager.Create(entChain, int64(topic.Id), topic.Title, topic.url(url))
	if err != nil {
		log.Sugar.Errorf("while creating draft proposal: %v", err)
		return
	}
	if entChain.IsEnabled {
		errIds := dc.tgClient.SendDraftProposals(prop, entChain)
		if len(errIds) > 0 {
			dc.telegramChatManager.DeleteMultiple(errIds)
		}

		errIds = dc.discordClient.SendDraftProposals(prop, entChain)
		if len(errIds) > 0 {
			dc.telegramChatManager.DeleteMultiple(errIds)
		}
	}
}

func (dc DiscourseCrawler) FetchDraftProposals() {
	log.Sugar.Info("Fetch draft proposals")
	chains := dc.chainManager.All()
	for _, c := range chains {
		if c.Name == "cosmoshub" {
			//database.DeleteAllDrafts()

			url := "https://forum.cosmos.network"
			topics, err := dc.getDraftProposals(url)
			if err != nil {
				log.Sugar.Errorf("while fetching draft proposals from %v", c.DisplayName)
			}

			props, err := dc.draftProposalManager.ByChain(c.Name)
			if err != nil {
				log.Sugar.Errorf("while querying draft proposals from %v", c.DisplayName)
			}

			for _, t := range topics {
				var found = false
				for _, prop := range props {
					if prop.ID == t.Id {
						found = true
					}
				}
				if !found {
					dc.saveAndSendDraftProposal(t, c, url)
				}
			}
		}
	}
}
