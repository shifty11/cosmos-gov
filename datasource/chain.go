package datasource

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/shifty11/cosmos-gov/api/telegram"
	"github.com/shifty11/cosmos-gov/common"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
	registry "github.com/strangelove-ventures/lens/client/chain_registry"
	"github.com/strangelove-ventures/lens/cmd"
	"sort"
	"strings"
)

type Datasource struct {
	ctx           context.Context
	chainRegistry registry.CosmosGithubRegistry
}

func NewDatasource(ctx context.Context, chainRegistry registry.CosmosGithubRegistry) *Datasource {
	return &Datasource{ctx: ctx, chainRegistry: chainRegistry}
}

func (ds Datasource) getChainsFromRegistry() ([]string, error) {
	chains, err := ds.chainRegistry.ListChains(ds.ctx)
	if err != nil {
		log.Sugar.Errorf("Error calling reg.ListChains: %v", err)
		return nil, err
	}
	var filteredChains []string
	for _, chain := range chains {
		if !strings.Contains(chain, "/") {
			filteredChains = append(filteredChains, chain)
		}
	}
	return filteredChains, nil
}

func (ds Datasource) orderChainsByErrorCnt(chains []string) []string {
	chainInfo := database.GetLensChainInfos()
	var chainsWithErrors = make(map[int][]string)
	chainsWithErrors[0] = []string{}
	for _, chainName := range chains {
		var found = false
		for _, lcInfo := range chainInfo {
			if chainName == lcInfo.Name {
				if chainsWithErrors[lcInfo.CntErrors] != nil {
					chainsWithErrors[lcInfo.CntErrors] = append(chainsWithErrors[lcInfo.CntErrors], lcInfo.Name)
				} else {
					chainsWithErrors[lcInfo.CntErrors] = []string{lcInfo.Name}
				}
				found = true
			}
		}
		if !found {
			chainsWithErrors[0] = append(chainsWithErrors[0], chainName)
		}
	}

	keys := make([]int, len(chainsWithErrors))
	i := 0
	for k := range chainsWithErrors {
		keys[i] = k
		i++
	}

	var orderedChains []string
	sort.Ints(keys)
	for _, k := range keys {
		for _, chain := range chainsWithErrors[k] {
			orderedChains = append(orderedChains, chain)
		}
	}
	return orderedChains
}

func (ds Datasource) getNewChains() []string {
	chainsInRegistry, err := ds.getChainsFromRegistry()
	if err != nil {
		return nil
	}

	chains := database.GetChains()
	var chainNames []string
	for _, chain := range chains {
		chainNames = append(chainNames, chain.Name)
	}

	var newChains []string
	for _, chain := range chainsInRegistry {
		if !common.Contains(chainNames, chain) {
			newChains = append(newChains, chain)
		}
	}
	return ds.orderChainsByErrorCnt(newChains)
}

func isChainInConfig(chainName string) bool {
	config, err := cmd.GetConfig(true, log.Sugar.Desugar())
	if err != nil {
		log.Sugar.Errorf("while getting config: %v", err)
		return false
	}
	for key := range config.Chains {
		if key == chainName {
			return true
		}
	}
	return false
}

func (ds Datasource) AddNewChains() {
	log.Sugar.Info("Add new chains")
	chains := ds.getNewChains()
	message := ""

	for _, chainName := range chains {
		client, err := getChainClient(chainName)
		if err != nil {
			log.Sugar.Debugf("Chain '%v' has %v errors", chainName, err)
			database.AddErrorToLensChainInfo(chainName)
		} else {
			proposals, err := fetchProposals(chainName, types.StatusNil, nil, client)
			if err != nil {
				log.Sugar.Debugf("Chain '%v' has %v errors", chainName, err)
				database.AddErrorToLensChainInfo(chainName)
			} else {
				if len(proposals.Proposals) >= 1 {
					addOrUpdateChainInLensConfig(chainName)
					if isChainInConfig(chainName) {
						chainEnt := database.CreateChain(chainName)
						for _, prop := range proposals.Proposals {
							database.CreateOrUpdateProposal(&prop, chainEnt)
						}
						database.DeleteLensChainInfo(chainName)
						message += fmt.Sprintf("Added chain '%v' including %v proposals\n", chainName, len(proposals.Proposals))
					} else {
						log.Sugar.Errorf("Chain '%v' is not in lens config", chainName)
						database.AddErrorToLensChainInfo(chainName)
					}
				} else {
					database.AddErrorToLensChainInfo(chainName)
				}
			}
		}
	}
	if message != "" {
		intro := "<b>New chain update info</b>\n"
		telegram.SendMessageToBotAdmins(intro + message)
	}
}
