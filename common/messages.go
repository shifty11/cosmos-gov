package common

import (
	"fmt"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent/user"
	"strings"
)

type MsgFormat string

const (
	MsgFormatMarkdown MsgFormat = "markdown"
	MsgFormatHtml     MsgFormat = "html"
)

func (m MsgFormat) String() string {
	return string(m)
}

func GetOngoingProposalsText(chatId int64, userType user.Type, format MsgFormat) string {
	text := ProposalsMsg
	chains := database.GetProposalsInVotingPeriodForUser(chatId, userType)
	if len(chains) == 0 {
		text = NoSubscriptionsMsg
	} else {
		for _, chain := range chains {
			for _, prop := range chain.Edges.Proposals {
				if format == MsgFormatMarkdown {
					title := strings.Replace(prop.Title, "_", "\\_", -1)
					title = strings.Replace(title, "*", "\\*", -1)
					text += fmt.Sprintf("**%v #%d** _%v_\n\n", chain.DisplayName, prop.ProposalID, title)
				} else {
					text += fmt.Sprintf("<b>%v #%d</b> <i>%v</i>\n\n", chain.DisplayName, prop.ProposalID, prop.Title)
				}
			}
		}
		if len(text) == len(ProposalsMsg) {
			text = NoProposalsMsg
		}
	}
	return text
}
