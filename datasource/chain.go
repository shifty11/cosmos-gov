package datasource

import (
	"fmt"
	"github.com/PumpkinSeed/cage"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/shifty11/cosmos-gov/common"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
	"github.com/shifty11/cosmos-gov/telegram"
	"github.com/strangelove-ventures/lens/cmd"
	"sort"
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

func orderChainsByErrorCnt(chains []string) []string {
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

func getNewChains() []string {
	chainsInRegistry, err := getChainsFromRegistry()
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
	return orderChainsByErrorCnt(newChains)
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
	message := ""
	for _, chainName := range chains {
		addOrUpdateChainInLensConfig(chainName)
		if isChainInConfig(chainName) {
			proposals, err := fetchProposals(chainName, types.StatusNil, nil)
			if err != nil {
				log.Sugar.Debugf("Chain '%v' has %v errors", chainName, err)
				removeChainFromLensConfig(chainName)
				database.AddErrorToLensChainInfo(chainName)
			} else {
				if len(proposals.Proposals) >= 1 {
					chainEnt := database.CreateChain(chainName)
					for _, prop := range proposals.Proposals {
						database.CreateOrUpdateProposal(&prop, chainEnt)
					}
					database.DeleteLensChainInfo(chainName)
					message += fmt.Sprintf("Added chain '%v' including %v proposals\n", chainName, len(proposals.Proposals))
				} else {
					removeChainFromLensConfig(chainName)
					database.AddErrorToLensChainInfo(chainName)
				}
			}
		} else {
			database.AddErrorToLensChainInfo(chainName)
		}
	}
	if message != "" {
		intro := "<b>New chain update info</b>\n"
		telegram.SendMessageToBotAdmins(intro + message)
	}
}
