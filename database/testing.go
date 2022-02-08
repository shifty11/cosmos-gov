package database

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
		//Where(proposal.HasChainWith(chain.NameEQ("osmosis"))).
		ExecX(ctx)
}
