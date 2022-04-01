package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/shifty11/cosmos-gov/api"
	"github.com/shifty11/cosmos-gov/common"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
	"strings"
)

const NbrOfButtonsPerRow = 5

func createKeyboard(chains *[]common.Chain) []discordgo.MessageComponent {
	var buttons []discordgo.MessageComponent
	var buttonRows []discordgo.MessageComponent
	for ix, c := range *chains {
		symbol := "❌ "
		if c.Notify {
			symbol = "✅ "
		}
		var button = discordgo.Button{
			Label:    symbol + " " + c.DisplayName,
			Style:    discordgo.SecondaryButton,
			Disabled: false,
			CustomID: common.SubscriptionCmd + ":" + c.Name,
		}
		buttons = append(buttons, button)
		if (ix+1)%NbrOfButtonsPerRow == 0 || ix == len(*chains)-1 {
			var buttonRow = discordgo.ActionsRow{Components: buttons}
			buttonRows = append(buttonRows, buttonRow)
			buttons = []discordgo.MessageComponent{}
		}
	}
	return buttonRows
}

func chunks(xs []common.Chain, chunkSize int) [][]common.Chain {
	if len(xs) == 0 {
		return nil
	}
	divided := make([][]common.Chain, (len(xs)+chunkSize-1)/chunkSize)
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

func getSpecificChunk(chunks *[][]common.Chain, name string) *[]common.Chain {
	for _, c1 := range *chunks {
		for _, c2 := range c1 {
			if c2.Name == name {
				return &c1
			}
		}
	}
	return nil
}

var (
	cmds = []*discordgo.ApplicationCommand{
		{
			Name:        common.SubscriptionCmd,
			Description: "Manage your Subscriptions",
		},
		{
			Name:        common.ProposalsCmd,
			Description: "Show ongoing proposals",
		},
		//{
		//	Name:        common.SupportCmd,
		//	Description: "Show ongoing proposals",
		//},
	}
	cmdHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		common.SubscriptionCmd: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if !canInteractWithBot(s, i) {
				sendEmptyResponse(s, i)
				return
			}

			channelId := getChannelId(i)

			chains := chunks(database.GetChainsForUser(channelId, user.TypeDiscord), 25)
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content:    sanitizeUrls(common.SubscriptionsMsg),
					Components: createKeyboard(&chains[0]),
				},
			})
			if err != nil {
				log.Sugar.Errorf("Error while sending subscriptions: %v", err)
			}
			if len(chains) > 1 {
				for _, chain := range chains[1:] {
					_, err := s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
						Content:    "",
						Components: createKeyboard(&chain),
					})
					if err != nil {
						log.Sugar.Errorf("Error while sending subscriptions: %v", err)
					}
				}
			}
		},
		common.ProposalsCmd: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if !canInteractWithBot(s, i) {
				sendEmptyResponse(s, i)
				return
			}

			channelId := getChannelId(i)

			text := api.GetOngoingProposalsText(channelId, user.TypeDiscord, api.MsgFormatMarkdown)

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
		common.SupportCmd: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if !canInteractWithBot(s, i) {
				sendEmptyResponse(s, i)
				return
			}

			text := fmt.Sprintf(common.SupportMsg, "[Rapha](https://discordapp.com/users/228978159440232453/)")
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
		common.SubscriptionCmd: func(s *discordgo.Session, i *discordgo.InteractionCreate, action string) {
			if !canInteractWithBot(s, i) {
				sendEmptyResponse(s, i)
				return
			}

			userId := getChannelId(i)

			database.PerformUpdateSubscription(userId, user.TypeDiscord, action)
			allChains := chunks(database.GetChainsForUser(userId, user.TypeDiscord), 25)
			chains := getSpecificChunk(&allChains, action)

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
