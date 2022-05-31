package datasource

import (
	"context"
	"errors"
	"fmt"
	querytypes "github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/liamylian/jsontime"
	"github.com/microcosm-cc/bluemonday"
	"github.com/shifty11/cosmos-gov/api/discord"
	"github.com/shifty11/cosmos-gov/api/telegram"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/log"
	lens "github.com/strangelove-ventures/lens/client"
	registry "github.com/strangelove-ventures/lens/client/chain_registry"
	"os"
	"regexp"
	"strings"
)

var json = jsontime.ConfigWithCustomTimeFormat
var stripPolicy = bluemonday.StrictPolicy()

type State struct {
	fetchErrors                     map[int]int
	maxFetchErrorsUntilAttemptToFix int // max fetch errors until attempt to fix it will start
	maxFetchErrorsUntilReport       int // max fetch errors until fetching will be reported
}

type ProposalDatasource struct {
	ctx                   context.Context
	chainRegistry         registry.CosmosGithubRegistry
	chainManager          *database.ChainManager
	telegramChatManager   *database.TelegramChatManager
	discordChannelManager *database.DiscordChannelManager
	proposalManager       *database.ProposalManager
	state                 *State
	tgClient              *telegram.TelegramLightClient
	discordClient         *discord.DiscordLightClient
}

func NewProposalDatasource(
	ctx context.Context,
	managers database.DbManagers,
	chainRegistry registry.CosmosGithubRegistry,
	state *State,
	tgClient *telegram.TelegramLightClient,
	discordClient *discord.DiscordLightClient,
) *ProposalDatasource {
	if state == nil {
		state = &State{
			fetchErrors:                     make(map[int]int),
			maxFetchErrorsUntilAttemptToFix: 10,
			maxFetchErrorsUntilReport:       20,
		}
	}
	return &ProposalDatasource{
		ctx:                   ctx,
		chainRegistry:         chainRegistry,
		chainManager:          managers.ChainManager,
		proposalManager:       managers.ProposalManager,
		telegramChatManager:   managers.TelegramChatManager,
		discordChannelManager: managers.DiscordChannelManager,
		state:                 state,
		tgClient:              tgClient,
		discordClient:         discordClient,
	}
}

func sanitizeTitle(title string) string { // Removes first character if it is not a letter
	r := regexp.MustCompile("[a-zA-Z]+")
	result := r.FindString(string(title[0]))
	if result == "" {
		return title[1:]
	}
	return title
}

func extractContentByRegEx(value []byte) (*database.ProposalContent, error) {
	r := regexp.MustCompile("[ -~]+") // search for all printable characters
	result := r.FindAll(value[1:], -1)
	if len(result) >= 2 {
		description := strings.Replace(string(result[1]), "\\n", "\n", -1)
		description = stripPolicy.Sanitize(description)
		return &database.ProposalContent{
			Title:       sanitizeTitle(string(result[0])),
			Description: description,
		}, nil
	}
	return nil, errors.New(fmt.Sprintf("Length of regex result is %v", len(result)))
}

// This is a bit a hack. The reason for this is that lens doesn't support chain specific proposals
func extractContent(cl *lens.ChainClient, response types.QueryProposalsResponse, proposalId uint64) (*database.ProposalContent, error) {
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

	var proposals database.Proposals
	err = json.Unmarshal(proto, &proposals) // transform the json []byte to our proposal structure
	if err != nil {
		return nil, err
	}
	if len(proposals.Proposals) == 1 {
		description := strings.Replace(proposals.Proposals[0].Content.Description, "\\n", "\n", -1)
		description = stripPolicy.Sanitize(description)
		return &database.ProposalContent{ // We just need the content
			Title:       proposals.Proposals[0].Content.Title,
			Description: description,
		}, nil
	}
	return nil, errors.New(fmt.Sprintf("Length of proposals is %v. This should never happen!", len(proposals.Proposals)))
}

func fetchProposals(chainId string, proposalStatus types.ProposalStatus, pageReq *querytypes.PageRequest, client *lens.ChainClient) (*database.Proposals, error) {
	log.Sugar.Debugf("QueryGovernanceProposals on %v --status %v", chainId, strings.ToLower(strings.Replace(proposalStatus.String(), "PROPOSAL_STATUS_", "", 1)))

	response, err := client.QueryGovProposals(context.Background(), proposalStatus, pageReq)
	if err != nil {
		log.Sugar.Debugf("Error while querying proposals on %v: %v", chainId, err)
		return nil, err
	}

	var proposals database.Proposals
	for _, respProp := range response.Proposals {
		content, err := extractContent(client, *response, respProp.ProposalId)
		if err != nil {
			log.Sugar.Error(err)
			continue
		}
		prop := database.Proposal{
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

func getChainInfo(chainName string, chainRegistry registry.CosmosGithubRegistry) (*lens.ChainClient, *registry.ChainInfo, []string, error) {
	chainInfo, err := chainRegistry.GetChain(context.Background(), chainName)
	if err != nil {
		log.Sugar.Debugf("Failed to get chain client on %v: %v \n", chainName, err)
		return nil, nil, nil, err
	}

	rpcs, err := chainInfo.GetRPCEndpoints(context.Background())
	if err != nil {
		log.Sugar.Errorf("Failed to get RPC endpoints on chain %s: %v \n", chainInfo.ChainID, err)
		return nil, nil, nil, err
	}

	if len(rpcs) <= 0 {
		log.Sugar.Debugf("Found no working RPC endpoints on chain %s: %v \n", chainInfo.ChainID, err)
		return nil, nil, nil, errors.New("found no working RPC endpoints")
	}

	pwd, _ := os.Getwd()
	key_dir := pwd + "/keys"

	chainConfig := lens.ChainClientConfig{
		Key:            "default",
		ChainID:        chainInfo.ChainID,
		RPCAddr:        rpcs[0],
		AccountPrefix:  chainInfo.Bech32Prefix,
		KeyringBackend: "test",
		Debug:          true,
		Timeout:        "20s",
		Modules:        lens.ModuleBasics,
	}

	// Creates client object to pull chain info
	chainClient, err := lens.NewChainClient(log.Sugar.Desugar(), &chainConfig, key_dir, os.Stdin, os.Stdout)
	if err != nil {
		log.Sugar.Fatalf("Failed to build new chain client for %s. Err: %v \n", chainInfo.ChainID, err)
	}
	return chainClient, &chainInfo, rpcs, nil
}

func (ds ProposalDatasource) saveAndSendProposals(props *database.Proposals, entChain *ent.Chain) {
	for _, prop := range props.Proposals {
		entProp := ds.proposalManager.CreateIfNotExists(&prop, entChain)
		if entProp != nil && entChain.IsEnabled {
			errIds := ds.tgClient.SendProposals(entProp, entChain)
			if len(errIds) > 0 {
				ds.telegramChatManager.DeleteMultiple(errIds)
			}

			errIds = ds.discordClient.SendProposals(entProp, entChain)
			if len(errIds) > 0 {
				ds.discordChannelManager.DeleteMultiple(errIds)
			}
		}
	}
}

func (ds ProposalDatasource) updateRpcs(chainName string) {
	_, _, rpcs, err := getChainInfo(chainName, ds.chainRegistry)
	if err != nil {
		log.Sugar.Errorf("Error getting RPC's for chain %v: %v", chainName, err)
	}
	if len(rpcs) == 0 {
		log.Sugar.Errorf("Found no RPC's for chain %v: %v", chainName, err)
		return
	}
	err = ds.chainManager.UpdateRpcs(chainName, rpcs)
	if err != nil {
		log.Sugar.Errorf("Error while updating RPC's for chain %v: %v", chainName, err)
	}
}

func (ds ProposalDatasource) handleFetchError(chain *ent.Chain, err error) {
	if err != nil {
		ds.state.fetchErrors[chain.ID] += 1
		if ds.state.fetchErrors[chain.ID] >= ds.state.maxFetchErrorsUntilAttemptToFix {
			ds.updateRpcs(chain.Name)
		}
		if ds.state.fetchErrors[chain.ID] >= ds.state.maxFetchErrorsUntilReport {
			log.Sugar.Errorf("Chain '%v' has %v errors", chain.DisplayName, ds.state.fetchErrors[chain.ID])
		}
	} else {
		ds.state.fetchErrors[chain.ID] = 0
	}
}

func (ds ProposalDatasource) updateProposal(entProp *ent.Proposal, status types.ProposalStatus) bool {
	pageRequest := querytypes.PageRequest{
		Key:        nil,
		Offset:     0,
		Limit:      1000,
		CountTotal: false,
		Reverse:    true,
	}
	client, err := ds.chainManager.BuildLensClient(entProp.Edges.Chain)
	if err != nil {
		log.Sugar.Fatalf("Could not get client for chain %v. It's probably not saved into the db.", chain.Name)
	}
	proposals, err := fetchProposals(entProp.Edges.Chain.Name, status, &pageRequest, client)
	ds.handleFetchError(entProp.Edges.Chain, err)
	if err != nil {
		return false
	}
	for _, prop := range proposals.Proposals {
		if prop.ProposalId == entProp.ProposalID {
			ds.proposalManager.CreateOrUpdateProposal(&prop, entProp.Edges.Chain)
			return false
		}
	}
	return true
}

// CheckForUpdates checks if proposal that are in voting period need to be updated
func (ds ProposalDatasource) CheckForUpdates() {
	votingProposals := ds.proposalManager.GetFinishedProposalsInVotingPeriod()
	if len(votingProposals) == 0 { // do nothing if there is no finished votingProposal
		return
	}

	for _, entProp := range votingProposals {
		continueUpdating := ds.updateProposal(entProp, types.StatusPassed)
		if continueUpdating {
			continueUpdating = ds.updateProposal(entProp, types.StatusRejected)
			if continueUpdating {
				continueUpdating = ds.updateProposal(entProp, types.StatusFailed)
				if continueUpdating {
					log.Sugar.Errorf("Status of proposal #%v on chain %v could not be updated", entProp.ProposalID, entProp.Edges.Chain.DisplayName)
				}
			}
		}
	}
}

func (ds ProposalDatasource) FetchProposals() {
	log.Sugar.Info("Fetch proposals")
	chains := ds.chainManager.All()
	for _, c := range chains {
		client, err := ds.chainManager.BuildLensClient(c)
		if err != nil {
			log.Sugar.Errorf("Could not get client for chain %v. It's probably not saved into the db.", c.Name)
			continue
		}
		proposals, err := fetchProposals(c.Name, types.StatusVotingPeriod, nil, client)
		ds.handleFetchError(c, err)
		if err == nil {
			ds.saveAndSendProposals(proposals, c)
		}
	}
}
