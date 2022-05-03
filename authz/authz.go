package authz

import (
	"context"
	"errors"
	"fmt"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent"
)

//goland:noinspection GoNameStartsWithPackageName
type AuthzClient struct {
	ctx           context.Context
	chainManager  *database.ChainManager
	walletManager *database.WalletManager
	userManager   *database.UserManager
}

func NewAuthzClient(chainManager *database.ChainManager, walletManager *database.WalletManager) *AuthzClient {
	return &AuthzClient{ctx: context.Background(), chainManager: chainManager, walletManager: walletManager}
}

type VoteState int

var (
	NotVoted VoteState = 0
	Voting   VoteState = 1
	Voted    VoteState = 2
)

type VoteData struct {
	ChainName  string
	ProposalId uint64
	Vote       govtypes.VoteOption
	State      VoteState
}

func (vd VoteData) ToString() string {
	return fmt.Sprintf("%v:%v:%v:%v", vd.ChainName, vd.ProposalId, vd.Vote, vd.State)
}

func (vd VoteData) ButtonText() string {
	text := ""
	if vd.State == Voting {
		text = "❗ " + text
	}
	switch vd.Vote {
	case govtypes.OptionYes:
		text += "Yes"
		if vd.State == Voted {
			text += " 👍🏽"
		}
		return text
	case govtypes.OptionNo:
		text += "No"
		if vd.State == Voted {
			text += " 👎🏽"
		}
		return text
	case govtypes.OptionAbstain:
		text += "Abstain"
		if vd.State == Voted {
			text += " 😴"
		}
		return text
	case govtypes.OptionNoWithVeto:
		text += "No with veto"
		if vd.State == Voted {
			text += " 😡"
		}
		return text
	}
	return ""
}

func ToVoteData(chainName string, proposalId uint64, vote govtypes.VoteOption, state VoteState) VoteData {
	return VoteData{
		ChainName:  chainName,
		ProposalId: proposalId,
		Vote:       vote,
		State:      state,
	}
}

func (client AuthzClient) GetVoteStatus(entUser *ent.User, chainName string, proposalId uint64) (govtypes.VoteOption, error) {
	entChain, err := client.chainManager.ByName(chainName)
	if err != nil {
		return govtypes.OptionEmpty, err
	}

	lensClient, err := client.chainManager.BuildLensClient(entChain)
	if err != nil {
		return govtypes.OptionEmpty, err
	}

	wallets, err := client.walletManager.ByUserAndChain(entUser, entChain)
	if err != nil {
		return govtypes.OptionEmpty, err
	}

	if len(wallets) == 0 {
		return govtypes.OptionEmpty, errors.New("no wallet found")
	}
	if len(wallets) > 1 {
		return govtypes.OptionEmpty, nil
	}

	vote, err := lensClient.QueryGovVote(client.ctx, proposalId, wallets[0].Address)
	if err != nil {
		return govtypes.OptionEmpty, err
	}

	return vote.Options[0].Option, nil
}

func (client AuthzClient) ExecAuthzVote(entUser *ent.User, voteData *VoteData) error {
	entChain, err := client.chainManager.ByName(voteData.ChainName)
	if err != nil {
		return err
	}

	lensClient, err := client.chainManager.BuildLensClient(entChain)
	if err != nil {
		return err
	}

	wallets, err := client.walletManager.ByUserAndChain(entUser, entChain)
	if err != nil {
		return err
	}
	if len(wallets) == 0 {
		return errors.New("no wallet was found")
	}

	var grantee string
	var addresses []string
	for _, w := range wallets {
		if len(w.Edges.Grants) == 0 {
			continue
		}
		grantee = w.Edges.Grants[0].Grantee
		addresses = append(addresses, w.Address)
	}
	if len(addresses) == 0 {
		return nil
	}

	result, err := lensClient.ExecAuthzVote(client.ctx, addresses, grantee, voteData.ProposalId, voteData.Vote, 200000)
	if err != nil {
		return err
	}
	if result.Code != 0 {
		return errors.New(result.Info)
	}
	return nil
}

func (client AuthzClient) GetGrant(chainName string, granter string, grantee string) (*database.GrantData, error) {
	entChain, err := client.chainManager.ByName(chainName)
	if err != nil {
		return nil, err
	}
	msgType := "/cosmos.gov.v1beta1.MsgVote"
	lensClient, err := client.chainManager.BuildLensClient(entChain)
	if err != nil {
		return nil, err
	}

	response, err := lensClient.QueryAuthzGrants(client.ctx, granter, grantee, msgType, nil)
	if err != nil {
		return nil, err
	}

	if len(response) == 0 {
		return nil, nil
	}
	if len(response) > 1 {
		return nil, errors.New(fmt.Sprintf("Found %v grants. There should only be one.", len(response)))
	}

	return &database.GrantData{
		Granter:   granter,
		Grantee:   grantee,
		Type:      msgType,
		ExpiresAt: response[0].Expiration,
	}, nil
}
