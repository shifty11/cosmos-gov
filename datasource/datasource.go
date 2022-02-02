package datasource

import (
	"github.com/PumpkinSeed/cage"
	"github.com/liamylian/jsontime"
	"github.com/shifty11/cosmos-gov/log"
	"github.com/strangelove-ventures/lens/cmd"
	"strings"
	"time"
)

type ProposalContent struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Proposal struct {
	ProposalId      int             `json:"proposal_id,string"`
	Content         ProposalContent `json:"content"`
	VotingStartTime time.Time       `json:"voting_start_time"`
	VotingEndTime   time.Time       `json:"voting_end_time"`
	Status          string          `json:"status"`
}

type Proposals struct {
	Proposals []Proposal `json:"proposals"`
}

var json = jsontime.ConfigWithCustomTimeFormat

func fetchProposals(query string) (*Proposals, error) {
	c := cage.Start() // start capturing output from stdout

	rootCmd := cmd.NewRootCmd()
	rootCmd.SetArgs(strings.Fields(query))
	err := rootCmd.Execute()
	if err != nil {
		log.Sugar.Errorf("Error while querying '%v': %v", query, err)
		return nil, err
	}

	cage.Stop(c) // stop capturing output from stdout

	var proposals Proposals
	dataBytes := []byte(strings.Join(c.Data, ""))
	err = json.Unmarshal(dataBytes, &proposals)
	if err != nil {
		log.Sugar.Errorf("Error while decoding response for query '%v': %v", query, err)
		return nil, err
	}
	return &proposals, nil
}

func saveProposals(*Proposals) {

}

var queries = [...]string{
	//"query governance proposals --chain cosmoshub --status voting_period",
	//"query governance proposals --chain osmosis --status voting_period",
	"query governance proposals --chain juno",
}

func FetchProposals() {
	for _, query := range queries {
		proposals, err := fetchProposals(query)
		if err == nil {
			saveProposals(proposals)
		}
	}
}
