package common

import (
	"fmt"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent/user"
)

func GetOngoingProposalsText(chatId int64, userType user.Type) string {
	text := ProposalsMsg
	chains := database.GetProposalsInVotingPeriodForUser(chatId, userType)
	if len(chains) == 0 {
		text = NoSubscriptionsMsg
	} else {
		for _, chain := range chains {
			for _, prop := range chain.Edges.Proposals {
				text += fmt.Sprintf("*%v #%d* _%v_\n\n", chain.DisplayName, prop.ProposalID, prop.Title)
			}
		}
		if len(text) == len(ProposalsMsg) {
			text = NoProposalsMsg
		}
	}
	return text
}
