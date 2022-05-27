package datasource

import (
	"context"
	"github.com/shifty11/cosmos-gov/database"
	registry "github.com/strangelove-ventures/lens/client/chain_registry"
)

type State struct {
	fetchErrors                     map[int]int
	maxFetchErrorsUntilAttemptToFix int // max fetch errors until attempt to fix it will start
	maxFetchErrorsUntilReport       int // max fetch errors until fetching will be reported
}

type Datasource struct {
	ctx                   context.Context
	chainRegistry         registry.CosmosGithubRegistry
	chainManager          *database.ChainManager
	telegramChatManager   *database.TelegramChatManager
	discordChannelManager *database.DiscordChannelManager
	proposalManager       *database.ProposalManager
	draftProposalManager  *database.DraftProposalManager
	lensChainInfoManager  *database.LensChainInfoManager
	state                 *State
}

func NewDatasource(
	ctx context.Context,
	managers database.DbManagers,
	chainRegistry registry.CosmosGithubRegistry,
	state *State,
) *Datasource {
	if state == nil {
		state = &State{
			fetchErrors:                     make(map[int]int),
			maxFetchErrorsUntilAttemptToFix: 10,
			maxFetchErrorsUntilReport:       20,
		}
	}
	return &Datasource{ctx: ctx,
		chainRegistry:         chainRegistry,
		chainManager:          managers.ChainManager,
		proposalManager:       managers.ProposalManager,
		draftProposalManager:  managers.DraftProposalManager,
		lensChainInfoManager:  managers.LensChainInfoManager,
		telegramChatManager:   managers.TelegramChatManager,
		discordChannelManager: managers.DiscordChannelManager,
		state:                 state,
	}
}
