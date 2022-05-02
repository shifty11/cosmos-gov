package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/shifty11/cosmos-gov/api"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
	"strings"
)

const NbrOfButtonsPerRow = 5

func createKeyboard(chains []*database.Subscription) []discordgo.MessageComponent {
	var buttons []discordgo.MessageComponent
	var buttonRows []discordgo.MessageComponent
	for ix, c := range chains {
		symbol := "❌ "
		if c.Notify {
			symbol = "✅ "
		}
		var button = discordgo.Button{
			Label:    symbol + " " + c.DisplayName,
			Style:    discordgo.SecondaryButton,
			Disabled: false,
			CustomID: messages.SubscriptionCmd + ":" + c.Name,
		}
		buttons = append(buttons, button)
		if (ix+1)%NbrOfButtonsPerRow == 0 || ix == len(chains)-1 {
			var buttonRow = discordgo.ActionsRow{Components: buttons}
			buttonRows = append(buttonRows, buttonRow)
			buttons = []discordgo.MessageComponent{}
		}
	}
	return buttonRows
}

func chunks(xs []*database.Subscription, chunkSize int) [][]*database.Subscription {
	if len(xs) == 0 {
		return nil
	}
	divided := make([][]*database.Subscription, (len(xs)+chunkSize-1)/chunkSize)
	prev := 0
	i := 0
	till := len(xs) - chunkSize
	for prev < till {
		next := prev + chunkSize
		divided[i] = xs[prev:next]
		prev = next
		i++
	}
	divided[i] = xs[prev:]
	return divided
}

func getSpecificChunk(chunks [][]*database.Subscription, name string) []*database.Subscription {
	for _, c1 := range chunks {
		for _, c2 := range c1 {
			if c2.Name == name {
				return c1
			}
		}
	}
	return nil
}

func getOngoingProposalsText(chatId int64) string {
	text := messages.ProposalsMsg
	chains := mHack.ProposalManager.GetProposalsInVotingPeriod(chatId, user.TypeDiscord)
	if len(chains) == 0 {
		text = messages.NoSubscriptionsMsg
	} else {
		for _, chain := range chains {
			for _, prop := range chain.Edges.Proposals {
				title := strings.Replace(prop.Title, "_", "\\_", -1)
				title = strings.Replace(title, "*", "\\*", -1)
				text += fmt.Sprintf("**%v #%d** _%v_\n\n", chain.DisplayName, prop.ProposalID, title)
			}
		}
		if len(text) == len(messages.ProposalsMsg) {
			text = messages.NoProposalsMsg
		}
	}
	return text
}

var (
	cmds = []*discordgo.ApplicationCommand{
		{
			Name:        messages.SubscriptionCmd,
			Description: "Manage your Subscriptions",
		},
		{
			Name:        messages.ProposalsCmd,
			Description: "Show ongoing proposals",
		},
		//{
		//	Name:        messages.SupportCmd,
		//	Description: "Show ongoing proposals",
		//},
	}
	cmdHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		messages.SubscriptionCmd: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if !canInteractWithBot(s, i) {
				sendEmptyResponse(s, i)
				return
			}

			userId := getUserIdX(i)
			userName := getUserName(i)
			channelId := getChannelId(i)
			channelName := getChannelName(i)
			isGroup := isGroup(i)

			subs := mHack.DiscordSubscriptionManager.GetOrCreateSubscriptions(userId, userName, channelId, channelName, isGroup)

			chains := chunks(subs, 25)
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content:    sanitizeUrls(messages.SubscriptionsMsg),
					Components: createKeyboard(chains[0]),
				},
			})
			if err != nil {
				log.Sugar.Errorf("Error while sending subscriptions: %v", err)
			}
			if len(chains) > 1 {
				for _, chain := range chains[1:] {
					_, err := s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
						Content:    "",
						Components: createKeyboard(chain),
					})
					if err != nil {
						log.Sugar.Errorf("Error while sending subscriptions: %v", err)
					}
				}
			}
		},
		messages.ProposalsCmd: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if !canInteractWithBot(s, i) {
				sendEmptyResponse(s, i)
				return
			}

			channelId := getChannelId(i)

			text := getOngoingProposalsText(channelId)

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: text,
				},
			})
			if err != nil {
				log.Sugar.Errorf("Error while sending subscriptions: %v", err)
			}
		},
		messages.SupportCmd: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if !canInteractWithBot(s, i) {
				sendEmptyResponse(s, i)
				return
			}

			text := fmt.Sprintf(messages.SupportMsg, "[Rapha](https://discordapp.com/users/228978159440232453/)")
			text = strings.Replace(text, "*", "**", -1)

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: text,
				},
			})
			if err != nil {
				log.Sugar.Errorf("Error while sending subscriptions: %v", err)
			}
		},
	}
)

var (
	actionHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate, action string){
		messages.SubscriptionCmd: func(s *discordgo.Session, i *discordgo.InteractionCreate, action string) {
			if !canInteractWithBot(s, i) {
				sendEmptyResponse(s, i)
				return
			}

			userId := getUserIdX(i)
			userName := getUserName(i)
			channelId := getChannelId(i)
			channelName := getChannelName(i)
			isGroup := isGroup(i)

			subs := mHack.DiscordSubscriptionManager.GetOrCreateSubscriptions(userId, userName, channelId, channelName, isGroup)

			allChains := chunks(subs, 25)
			chains := getSpecificChunk(allChains, action)

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Content:    i.Message.Content,
					Components: createKeyboard(chains),
				},
			})
			if err != nil {
				log.Sugar.Errorf("Error while changing subscriptions: %v", err)
			}
		},
	}
)
