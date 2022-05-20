package datasource

import (
	"github.com/FrenchBen/godisco"
	_ "github.com/lib/pq"
	"github.com/shifty11/cosmos-gov/log"
	"golang.org/x/exp/slices"
)

type Topic struct {
	Id    int      `json:"id"`
	Title string   `json:"title"`
	Slug  string   `json:"slug"`
	Tags  []string `json:"tags"`
}

type TopicList struct {
	Topics []Topic `json:"topics"`
}

type TopicResponse struct {
	TopicList TopicList `json:"topic_list"`
}

func GetDraftProposals() {
	url := "https://forum.cosmos.network"
	discourseClient, err := godisco.NewClient(url, "", "")
	if err != nil {
		log.Sugar.Fatal(err)
	}

	body, _, err := discourseClient.Get("/latest.json")
	if err != nil {
		//return nil, err
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
	log.Sugar.Info(topics)
}
