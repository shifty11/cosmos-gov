package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/log"
	"os"
	"strconv"
	"strings"
)

//goland:noinspection GoNameStartsWithPackageName
type DiscordLightClient struct {
	DiscordChannelManager *database.DiscordChannelManager
}

func NewDiscordLightClient(managers database.DbManagers) *DiscordLightClient {
	return &DiscordLightClient{
		DiscordChannelManager: managers.DiscordChannelManager,
	}
}

func (dc DiscordLightClient) startSession() *discordgo.Session {
	var err error
	session, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Sugar.Fatalf("Invalid bot parameters: %v", err)
	}

	err = session.Open()
	if err != nil {
		log.Sugar.Fatalf("Cannot open the s: %v", err)
	}
	return session
}

func (dc DiscordLightClient) closeSession(session *discordgo.Session) {
	err := session.Close()
	if err != nil {
		log.Sugar.Errorf("Error while closing discord s: %v", err)
	}
}

func (dc DiscordLightClient) shouldDeleteUser(err error) bool {
	if restErr, ok := err.(*discordgo.RESTError); ok {
		return restErr.Response.StatusCode == 403 || restErr.Response.StatusCode == 404
	} else {
		return false
	}
}

func (dc DiscordLightClient) SendProposals(entProp *ent.Proposal, entChain *ent.Chain) []int64 {
	session := dc.startSession()
	defer dc.closeSession(session)

	// Remove bold text inside of description
	description := strings.Replace(entProp.Description, "*", "", -1)
	description = sanitizeUrls(description)

	text := fmt.Sprintf("🎉  **%v - Proposal %v\n\n%v**\n\n%v", entChain.DisplayName, entProp.ProposalID, entProp.Title, description)
	if len(text) > 2000 {
		text = text[:1997] + "..."
	}

	var errIds []int64
	channelIds := dc.DiscordChannelManager.GetChannelIds(entChain)
	for _, channelId := range channelIds {
		log.Sugar.Debugf("Send proposal #%v on %v to discord chat #%v", entProp.ProposalID, entChain.DisplayName, channelId)
		var _, err = session.ChannelMessageSendComplex(strconv.Itoa(channelId), &discordgo.MessageSend{
			Content: text,
		})
		if err != nil {
			if dc.shouldDeleteUser(err) {
				errIds = append(errIds, int64(channelId))
			} else {
				log.Sugar.Errorf("Error while sending proposal to discord chat #%v: %v", channelId, err)
			}
		}
	}
	return errIds
}

func (dc DiscordLightClient) SendDraftProposals(entProp *ent.DraftProposal, entChain *ent.Chain) []int64 {
	session := dc.startSession()
	defer dc.closeSession(session)

	text := fmt.Sprintf("💬  **%v - New pre-vote proposal\n\n%v**\n<%v>", entChain.DisplayName, entProp.Title, entProp.URL)

	var errIds []int64
	channelIds := dc.DiscordChannelManager.GetChannelIds(entChain)
	for _, channelId := range channelIds {
		log.Sugar.Debugf("Send draft proposal #%v on %v to discord chat #%v", entProp.DraftProposalID, entChain.DisplayName, channelId)
		var _, err = session.ChannelMessageSendComplex(strconv.Itoa(channelId), &discordgo.MessageSend{
			Content: text,
		})
		if err != nil {
			if dc.shouldDeleteUser(err) {
				errIds = append(errIds, int64(channelId))
			} else {
				log.Sugar.Errorf("Error while sending proposal to discord chat #%v: %v", channelId, err)
			}
		}
	}
	return errIds
}
