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
	"github.com/shifty11/cosmos-gov/common"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
	lens "github.com/strangelove-ventures/lens/client"
	registry "github.com/strangelove-ventures/lens/client/chain_registry"
	"os"
	"regexp"
	"strings"
)

var json = jsontime.ConfigWithCustomTimeFormat
var stripPolicy = bluemonday.StrictPolicy()

func extractContentByRegEx(value []byte) (*common.ProposalContent, error) {
	r := regexp.MustCompile("[ -~]+") // search for all printable characters
	result := r.FindAll(value[1:], -1)
	if len(result) >= 2 {
		description := strings.Replace(string(result[1]), "\\n", "\n", -1)
		description = stripPolicy.Sanitize(description)
		return &common.ProposalContent{
			Title:       string(result[0])[1:],
			Description: description,
		}, nil
	}
	return nil, errors.New(fmt.Sprintf("Length of regex result is %v", len(result)))
}

// This is a bit a hack. The reason for this is that lens doesn't support chain specific proposals
func extractContent(cl *lens.ChainClient, response types.QueryProposalsResponse, proposalId uint64) (*common.ProposalContent, error) {
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

	var proposals common.Proposals
	err = json.Unmarshal(proto, &proposals) // transform the json []byte to our proposal structure
	if err != nil {
		return nil, err
	}
	if len(proposals.Proposals) == 1 {
		description := strings.Replace(proposals.Proposals[0].Content.Description, "\\n", "\n", -1)
		description = stripPolicy.Sanitize(description)
		return &common.ProposalContent{ // We just need the content
			Title:       proposals.Proposals[0].Content.Title,
			Description: description,
		}, nil
	}
	return nil, errors.New(fmt.Sprintf("Length of proposals is %v. This should never happen!", len(proposals.Proposals)))
}

func fetchProposals(chainId string, proposalStatus types.ProposalStatus, pageReq *querytypes.PageRequest, client *lens.ChainClient) (*common.Proposals, error) {
	log.Sugar.Debugf("QueryGovernanceProposals on %v --status %v", chainId, strings.ToLower(strings.Replace(proposalStatus.String(), "PROPOSAL_STATUS_", "", 1)))

	response, err := client.QueryGovProposals(context.Background(), proposalStatus, pageReq)
	if err != nil {
		log.Sugar.Debugf("Error while querying proposals on %v: %v", chainId, err)
		return nil, err
	}

	var proposals common.Proposals
	for _, respProp := range response.Proposals {
		content, err := extractContent(client, *response, respProp.ProposalId)
		if err != nil {
			log.Sugar.Error(err)
			continue
		}
		prop := common.Proposal{
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

func getChainClient(chainName string) (*lens.ChainClient, error) {
	chainInfo, err := registry.DefaultChainRegistry(log.Sugar.Desugar()).GetChain(context.Background(), chainName)
	if err != nil {
		log.Sugar.Errorf("Failed to get chain client on %v: %v \n", chainName, err)
		return nil, err
	}

	//	Use Chain info to select random endpoint
	rpc, err := chainInfo.GetRandomRPCEndpoint(context.Background())
	if err != nil {
		log.Sugar.Errorf("Failed to get random RPC endpoint on chain %s: %v \n", chainInfo.ChainID, err)
		return nil, err
	}

	// For this example, lets place the key directory in your PWD.
	pwd, _ := os.Getwd()
	key_dir := pwd + "/keys"

	// Build chain config
	chainConfig := lens.ChainClientConfig{
		Key:     "default",
		ChainID: chainInfo.ChainID,
		RPCAddr: rpc,
		// GRPCAddr       string,
		//AccountPrefix:  chainInfo.Bech32Prefix,
		KeyringBackend: "test",
		//GasAdjustment:  1.2,
		//GasPrices:      "0.01uosmo",
		//KeyDirectory:   key_dir,
		Debug:   true,
		Timeout: "20s",
		//OutputFormat:   "json",
		//SignModeStr:    "direct",
		Modules: lens.ModuleBasics,
	}

	// Creates client object to pull chain info
	chainClient, err := lens.NewChainClient(log.Sugar.Desugar(), &chainConfig, key_dir, os.Stdin, os.Stdout)
	if err != nil {
		log.Sugar.Fatalf("Failed to build new chain client for %s. Err: %v \n", chainInfo.ChainID, err)
	}
	return chainClient, nil
}

func getChainClientFromDb(entChain *ent.Chain) (*lens.ChainClient, error) {
	rpc, err := entChain.
		QueryRPCEndpoints().
		First(context.Background())
	if err != nil {
		log.Sugar.Fatalf("Failed to build new chain client for %s. Err: %v \n", entChain.DisplayName, err)
	}

	pwd, _ := os.Getwd()
	key_dir := pwd + "/keys"

	chainConfig := lens.ChainClientConfig{
		Key:            "default",
		ChainID:        entChain.Name,
		RPCAddr:        rpc.Endpoint,
		KeyringBackend: "test",
		Debug:          true,
		Timeout:        "20s",
		Modules:        lens.ModuleBasics,
	}

	chainClient, err := lens.NewChainClient(log.Sugar.Desugar(), &chainConfig, key_dir, os.Stdin, os.Stdout)
	if err != nil {
		log.Sugar.Fatalf("Failed to build new chain client for %s. Err: %v \n", entChain.DisplayName, err)
	}
	return chainClient, nil
}

func saveAndSendProposals(props *common.Proposals, entChain *ent.Chain) {
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
	client, err := getChainClientFromDb(entProp.Edges.Chain)
	if err != nil {
		log.Sugar.Fatalf("Could not get client for chain %v. It's probably not saved into the db.", chain.Name)
	}
	proposals, err := fetchProposals(entProp.Edges.Chain.Name, status, &pageRequest, client)
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
		client, err := getChainClientFromDb(chain)
		if err != nil {
			log.Sugar.Errorf("Could not get client for chain %v. It's probably not saved into the db.", chain.Name)
			continue
		}
		proposals, err := fetchProposals(chain.Name, types.StatusVotingPeriod, nil, client)
		handleFetchError(chain, err)
		if err == nil {
			saveAndSendProposals(proposals, chain)
		}
	}
}
