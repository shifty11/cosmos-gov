package datasource

import (
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

func getDraftProposals(url string) ([]Topic, error) {
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

func (ds Datasource) saveAndSendDraftProposal(topic Topic, entChain *ent.Chain, url string) {
	prop, err := ds.draftProposalManager.Create(entChain, int64(topic.Id), topic.Title, topic.url(url))
	if err != nil {
		log.Sugar.Errorf("while creating draft proposal: %v", err)
		return
	}
	if entChain.IsEnabled {
		errIds := telegram.SendDraftProposals(prop, entChain)
		if len(errIds) > 0 {
			ds.telegramChatManager.DeleteMultiple(errIds)
		}

		errIds = discord.SendDraftProposals(prop, entChain)
		if len(errIds) > 0 {
			ds.telegramChatManager.DeleteMultiple(errIds)
		}
	}
}

func (ds Datasource) FetchDraftProposals() {
	log.Sugar.Info("Fetch draft proposals")
	chains := ds.chainManager.All()
	for _, c := range chains {
		if c.Name == "cosmoshub" {
			database.DeleteAllDrafts()

			url := "https://forum.cosmos.network"
			topics, err := getDraftProposals(url)
			if err != nil {
				log.Sugar.Errorf("while fetching draft proposals from %v", c.DisplayName)
			}

			props, err := ds.draftProposalManager.ByChain(c.Name)
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
					ds.saveAndSendDraftProposal(t, c, url)
				}
			}
		}
	}
}
