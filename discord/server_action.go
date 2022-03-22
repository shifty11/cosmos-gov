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

func startSession() *discordgo.Session {
	var err error
	session, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Sugar.Fatalf("Invalid bot parameters: %v", err)
	}
	log.Sugar.Info("Start discord bot")

	err = session.Open()
	if err != nil {
		log.Sugar.Fatalf("Cannot open the s: %v", err)
	}
	return session
}

func closeSession(session *discordgo.Session) {
	err := session.Close()
	if err != nil {
		log.Sugar.Errorf("Error while closing discord s: %v", err)
	}
}

func shouldDeleteUser(err error) bool {
	if restErr, ok := err.(*discordgo.RESTError); ok {
		return restErr.Response.StatusCode == 403
	} else {
		return false
	}
}

func SendProposals(entProp *ent.Proposal, entChain *ent.Chain) []int64 {
	session := startSession()
	defer closeSession(session)

	description := strings.Replace(entProp.Description, "*", "", -1)

	text := fmt.Sprintf("🎉  **%v - Proposal %v\n\n%v**\n\n%v", entChain.DisplayName, entProp.ProposalID, entProp.Title, description)
	if len(text) > 2000 {
		text = text[:1997] + "..."
	}

	var errIds []int64
	channelIds := database.GetDiscordChatIds(entChain)
	for _, channelId := range channelIds {
		log.Sugar.Debugf("Send proposal #%v on %v to discord chat #%v", entProp.ProposalID, entChain.DisplayName, channelId)
		var _, err = session.ChannelMessageSend(strconv.Itoa(channelId), text)
		if err != nil {
			if shouldDeleteUser(err) {
				errIds = append(errIds, int64(channelId))
			} else {
				log.Sugar.Errorf("Error while sending proposal to discord chat #%v: %v", channelId, err)
			}
		}
	}
	return errIds
}
