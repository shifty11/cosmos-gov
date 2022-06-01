package datasource

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/shifty11/cosmos-gov/api/discord"
	"github.com/shifty11/cosmos-gov/api/telegram"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
	registry "github.com/strangelove-ventures/lens/client/chain_registry"
	"golang.org/x/exp/slices"
	"sort"
	"strings"
)

type ChainDatasource struct {
	ctx                  context.Context
	chainRegistry        registry.CosmosGithubRegistry
	chainManager         *database.ChainManager
	proposalManager      *database.ProposalManager
	lensChainInfoManager *database.LensChainInfoManager
	tgClient             *telegram.TelegramLightClient
	discordClient        *discord.DiscordLightClient
}

func NewChainDatasource(
	ctx context.Context,
	managers *database.DbManagers,
	chainRegistry registry.CosmosGithubRegistry,
	tgClient *telegram.TelegramLightClient,
	discordClient *discord.DiscordLightClient,
) *ChainDatasource {
	return &ChainDatasource{
		ctx:                  ctx,
		chainRegistry:        chainRegistry,
		chainManager:         managers.ChainManager,
		proposalManager:      managers.ProposalManager,
		lensChainInfoManager: managers.LensChainInfoManager,
		tgClient:             tgClient,
		discordClient:        discordClient,
	}
}

func (cd ChainDatasource) getChainsFromRegistry() ([]string, error) {
	chains, err := cd.chainRegistry.ListChains(cd.ctx)
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

func (cd ChainDatasource) orderChainsByErrorCnt(chains []string) []string {
	chainInfo := cd.lensChainInfoManager.GetLensChainInfos()
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

func (cd ChainDatasource) getNewChains() []string {
	chainsInRegistry, err := cd.getChainsFromRegistry()
	if err != nil {
		return nil
	}

	chains := cd.chainManager.All()
	var chainNames []string
	for _, chain := range chains {
		chainNames = append(chainNames, chain.Name)
	}

	var newChains []string
	for _, chain := range chainsInRegistry {
		if !slices.Contains(chainNames, chain) {
			newChains = append(newChains, chain)
		}
	}
	return cd.orderChainsByErrorCnt(newChains)
}

func (cd ChainDatasource) AddNewChains() {
	log.Sugar.Info("Add new chains")
	chains := cd.getNewChains()
	message := ""
	chainManager := cd.chainManager
	propManager := cd.proposalManager
	lensChainManager := cd.lensChainInfoManager

	for _, chainName := range chains {
		client, chainInfo, rpcs, err := getChainInfo(chainName, cd.chainRegistry)
		if err != nil {
			log.Sugar.Debugf("Chain '%v' has %v errors", chainName, err)
			lensChainManager.AddErrorToLensChainInfo(chainName)
		} else {
			proposals, err := fetchProposals(chainName, types.StatusNil, nil, client)
			if err != nil {
				log.Sugar.Debugf("Chain '%v' has %v errors", chainName, err)
				lensChainManager.AddErrorToLensChainInfo(chainName)
			} else {
				chainEnt := chainManager.Create(chainInfo.ChainID, chainName, chainInfo.Bech32Prefix, rpcs)
				for _, prop := range proposals.Proposals {
					propManager.CreateOrUpdateProposal(&prop, chainEnt)
				}
				lensChainManager.DeleteLensChainInfo(chainName)
				message += fmt.Sprintf("Added chain '%v' including %v proposals\n", chainName, len(proposals.Proposals))
			}
		}
	}
	if message != "" {
		intro := "<b>New chain update info</b>\n"
		cd.tgClient.SendMessageToBotAdmins(intro + message)
	}
}
