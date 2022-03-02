package datasource

import (
	"fmt"
	"github.com/PumpkinSeed/cage"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/shifty11/cosmos-gov/common"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
	"github.com/strangelove-ventures/lens/cmd"
	"strings"
)

func getChainsFromRegistry() ([]string, error) {
	c := cage.Start()
	query := fmt.Sprint("chains registry-list")
	rootCmd := cmd.NewRootCmd()
	rootCmd.SetArgs(strings.Fields(query))
	log.Sugar.Debug(query)
	err := rootCmd.Execute()
	cage.Stop(c)
	if err != nil {
		log.Sugar.Errorf("Error while querying '%v': %v", query, err)
		return nil, err
	}

	var chains []string
	dataBytes := []byte(strings.Join(c.Data, ""))
	err = json.Unmarshal(dataBytes, &chains)
	if err != nil {
		log.Sugar.Errorf("Error while decoding response for query '%v': %v", query, err)
		return nil, err
	}
	var filteredChains []string
	for _, chain := range chains {
		if !strings.Contains(chain, "/") {
			filteredChains = append(filteredChains, chain)
		}
	}
	return filteredChains, err
}

func getNewChains() []string {
	newChains, err := getChainsFromRegistry()
	if err != nil {
		return nil
	}

	chains := database.GetChains()
	var chainNames []string
	for _, chain := range chains {
		chainNames = append(chainNames, chain.Name)
	}

	var filteredChains []string
	for _, chain := range newChains {
		if !common.Contains(chainNames, chain) {
			filteredChains = append(filteredChains, chain)
		}
	}
	return filteredChains
}

func isChainInConfig(chainName string) bool {
	config, err := cmd.GetConfig(true)
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

func AddNewChains() {
	log.Sugar.Info("Add new chains")
	chains := getNewChains()
	for _, chain := range chains {
		addOrUpdateChainInLensConfig(chain)
		if isChainInConfig(chain) {
			proposals, err := fetchProposals(chain, types.StatusNil, nil)
			if err != nil {
				log.Sugar.Debugf("Chain '%v' has %v errors", chain, err)
				removeChainFromLensConfig(chain)
			} else {
				if len(proposals.Proposals) >= 1 {
					chainEnt := database.CreateChain(chain)
					for _, prop := range proposals.Proposals {
						database.CreateOrUpdateProposal(&prop, chainEnt)
					}
				} else {
					removeChainFromLensConfig(chain)
				}
			}
		}
	}
}
