package database

import "github.com/shifty11/cosmos-gov/ent/proposal"

func DropChains() {
	client, ctx := connect()
	client.Chain.
		Delete().
		ExecX(ctx)
}

func DropProposals() {
	client, ctx := connect()
	client.Proposal.
		Delete().
		//Where(proposal.And(
		//	proposal.HasChainWith(chain.NameEQ("osmosis")),
		//	proposal.ProposalIDEQ(144),
		//)).
		Where(proposal.StatusEQ(proposal.StatusPROPOSAL_STATUS_VOTING_PERIOD)).
		ExecX(ctx)
}
