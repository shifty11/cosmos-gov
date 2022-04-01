package datasource

import (
	"errors"
	"fmt"
	querytypes "github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/liamylian/jsontime"
	"github.com/microcosm-cc/bluemonday"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/discord"
	"github.com/shifty11/cosmos-gov/dtos"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
	"github.com/shifty11/cosmos-gov/telegram"
	"github.com/strangelove-ventures/lens/client"
	"github.com/strangelove-ventures/lens/cmd"
	"regexp"
	"strings"
)

var json = jsontime.ConfigWithCustomTimeFormat
var stripPolicy = bluemonday.StrictPolicy()

func extractContentByRegEx(value []byte) (*dtos.ProposalContent, error) {
	r := regexp.MustCompile("[ -~]+") // search for all printable characters
	result := r.FindAll(value[1:], -1)
	if len(result) >= 2 {
		description := strings.Replace(string(result[1]), "\\n", "\n", -1)
		description = stripPolicy.Sanitize(description)
		return &dtos.ProposalContent{
			Title:       string(result[0])[1:],
			Description: description,
		}, nil
	}
	return nil, errors.New(fmt.Sprintf("Length of regex result is %v", len(result)))
}

// This is a bit a hack. The reason for this is that lens doesn't support chain specific proposals
func extractContent(cl *client.ChainClient, response types.QueryProposalsResponse, proposalId uint64) (*dtos.ProposalContent, error) {
	// We want just the proposal with proposalId
	for _, prop := range response.Proposals {
		if prop.ProposalId == proposalId {
			response.Proposals = []types.Proposal{prop}
			break
		}
	}

	proto, err := cl.MarshalProto(&response) // this will use the correct type to produce json []byte
	if err != nil {                          // it will fail if there is a chain specific proposal
		log.Sugar.Debugf("extractContentByRegEx for proposal #%v on %v", proposalId, cl.Config.ChainID)
		return extractContentByRegEx(response.Proposals[0].Content.Value) // extract content by regex in this case
	}

	var proposals dtos.Proposals
	err = json.Unmarshal(proto, &proposals) // transform the json []byte to our proposal structure
	if err != nil {
		return nil, err
	}
	if len(proposals.Proposals) == 1 {
		description := strings.Replace(proposals.Proposals[0].Content.Description, "\\n", "\n", -1)
		description = stripPolicy.Sanitize(description)
		return &dtos.ProposalContent{ // We just need the content
			Title:       proposals.Proposals[0].Content.Title,
			Description: description,
		}, nil
	}
	return nil, errors.New(fmt.Sprintf("Length of proposals is %v. This should never happen!", len(proposals.Proposals)))
}

func fetchProposals(chainId string, proposalStatus types.ProposalStatus, pageReq *querytypes.PageRequest) (*dtos.Proposals, error) {
	config, err := cmd.GetConfig(false)
	if err != nil {
		log.Sugar.Panicf("Error while reading config %v", err)
	}
	cl := config.GetClient(chainId)
	if cl == nil {
		log.Sugar.Panicf("Chain client '%v' not found ", chainId)
	}

	log.Sugar.Debugf("QueryGovernanceProposals on %v --status %v", chainId, strings.ToLower(strings.Replace(proposalStatus.String(), "PROPOSAL_STATUS_", "", 1)))
	response, err := cl.QueryGovernanceProposals(proposalStatus, "", "", pageReq)
	if err != nil {
		log.Sugar.Debugf("Error while querying proposals on %v: %v", chainId, err)
		return nil, err
	}
	log.Sugar.Debugf("Got %v proposals", len(response.Proposals))

	var proposals dtos.Proposals
	for _, respProp := range response.Proposals {
		content, err := extractContent(cl, *response, respProp.ProposalId)
		if err != nil {
			log.Sugar.Error(err)
			continue
		}
		prop := dtos.Proposal{
			ProposalId:      respProp.ProposalId,
			Content:         *content,
			VotingStartTime: respProp.VotingStartTime,
			VotingEndTime:   respProp.VotingEndTime,
			Status:          respProp.Status.String(),
		}
		proposals.Proposals = append(proposals.Proposals, prop)
	}
	return &proposals, nil
}

func saveAndSendProposals(props *dtos.Proposals, entChain *ent.Chain) {
	for _, prop := range props.Proposals {
		entProp := database.CreateProposalIfNotExists(&prop, entChain)
		if entProp != nil && entChain.IsEnabled {
			errIds := telegram.SendProposals(entProp, entChain)
			if len(errIds) > 0 {
				database.DeleteUsers(errIds, user.TypeTelegram)
			}

			errIds = discord.SendProposals(entProp, entChain)
			if len(errIds) > 0 {
				database.DeleteUsers(errIds, user.TypeDiscord)
			}
		}
	}
}

const maxFetchErrorsUntilAttemptToFix = 10 // max fetch errors until attempt to fix it will start
const maxFetchErrorsUntilReport = 20       // max fetch errors until fetching will be reported

var fetchErrors = make(map[int]int) // map of chain and number of errors

func handleFetchError(chain *ent.Chain, err error) {
	if err != nil {
		fetchErrors[chain.ID] += 1
		if fetchErrors[chain.ID] >= maxFetchErrorsUntilAttemptToFix {
			addOrUpdateChainInLensConfig(chain.Name)
		}
		if fetchErrors[chain.ID] >= maxFetchErrorsUntilReport {
			log.Sugar.Errorf("Chain '%v' has %v errors", chain.DisplayName, fetchErrors[chain.ID])
		}
	} else {
		fetchErrors[chain.ID] = 0
	}
}

func updateProposal(entProp *ent.Proposal, status types.ProposalStatus) bool {
	pageRequest := querytypes.PageRequest{
		Key:        nil,
		Offset:     0,
		Limit:      100,
		CountTotal: false,
		Reverse:    true,
	}
	proposals, err := fetchProposals(entProp.Edges.Chain.Name, status, &pageRequest)
	handleFetchError(entProp.Edges.Chain, err)
	if err != nil {
		return false
	}
	for _, prop := range proposals.Proposals {
		if prop.ProposalId == entProp.ProposalID {
			database.CreateOrUpdateProposal(&prop, entProp.Edges.Chain)
			return false
		}
	}
	return true
}

// CheckForUpdates checks if proposal that are in voting period need to be updated
func CheckForUpdates() {
	votingProposals := database.GetFinishedProposalsInVotingPeriod()
	if len(votingProposals) == 0 { // do nothing if there is no finished votingProposal
		return
	}

	for _, entProp := range votingProposals {
		continueUpdating := updateProposal(entProp, types.StatusPassed)
		if continueUpdating {
			continueUpdating = updateProposal(entProp, types.StatusRejected)
			if continueUpdating {
				continueUpdating = updateProposal(entProp, types.StatusFailed)
				if continueUpdating {
					log.Sugar.Errorf("Status of proposal #%v on chain %v could not be updated", entProp.ProposalID, entProp.Edges.Chain.DisplayName)
				}
			}
		}
	}
}

func FetchProposals() {
	log.Sugar.Info("Fetch proposals")
	for _, chain := range database.GetChains() {
		proposals, err := fetchProposals(chain.Name, types.StatusVotingPeriod, nil)
		handleFetchError(chain, err)
		if err == nil {
			saveAndSendProposals(proposals, chain)
		}
	}
}
