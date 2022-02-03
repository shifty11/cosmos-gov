package datasource

import (
	"fmt"
	"github.com/PumpkinSeed/cage"
	"github.com/liamylian/jsontime"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/dtos"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/log"
	"github.com/shifty11/cosmos-gov/telegram"
	"github.com/strangelove-ventures/lens/cmd"
	"strings"
)

var json = jsontime.ConfigWithCustomTimeFormat

func fetchProposals(query string) (*dtos.Proposals, error) {
	c := cage.Start() // start capturing output from stdout

	rootCmd := cmd.NewRootCmd()
	rootCmd.SetArgs(strings.Fields(query))
	err := rootCmd.Execute()
	if err != nil {
		log.Sugar.Errorf("Error while querying '%v': %v", query, err)
		return nil, err
	}

	cage.Stop(c) // stop capturing output from stdout

	var proposals dtos.Proposals
	dataBytes := []byte(strings.Join(c.Data, ""))
	err = json.Unmarshal(dataBytes, &proposals)
	if err != nil {
		log.Sugar.Errorf("Error while decoding response for query '%v': %v", query, err)
		return nil, err
	}
	return &proposals, nil
}

func saveAndSendProposals(props *dtos.Proposals, chainDb *ent.Chain) {
	for _, prop := range props.Proposals {
		errIds := make(map[int]struct{})
		propDb := database.CreateProposalIfNotExists(&prop, chainDb)
		if propDb != nil {
			chatIds := database.GetUsers(chainDb)
			text := fmt.Sprintf("*%v*\n*#%v - %v*\n%v", chainDb.DisplayName, propDb.ProposalID, propDb.Title, propDb.Description)
			errIds = telegram.SendProposal(text, chatIds)
		}
		if len(errIds) != 0 {
			database.DeleteUsers(errIds)
		}
	}
}

const filter = "--limit 1"

//const filter = "--status voting_period"

func FetchProposals() {
	for _, chain := range database.GetChains() {
		query := fmt.Sprintf("query governance proposals %v --chain %v", filter, chain.Name)
		proposals, err := fetchProposals(query)
		if err == nil {
			saveAndSendProposals(proposals, chain)
		}
	}
}
