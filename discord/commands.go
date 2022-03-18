package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/dtos"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/log"
)

//const cmdSubscriptions = "subscriptions"
const cmdSubscriptions = "subs"
const subscriptionsMsg = `Select the projects that you want to follow. You will receive notifications about new governance proposals once they enter the voting period.`

const NbrOfButtonsPerRow = 5

func createKeyboard(chains *[]dtos.Chain) []discordgo.MessageComponent {
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
			CustomID: cmdSubscriptions + ":" + c.Name,
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

func chunks(xs []dtos.Chain, chunkSize int) [][]dtos.Chain {
	if len(xs) == 0 {
		return nil
	}
	divided := make([][]dtos.Chain, (len(xs)+chunkSize-1)/chunkSize)
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

func getSpecificChunk(chunks *[][]dtos.Chain, name string) *[]dtos.Chain {
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
			Name:        cmdSubscriptions,
			Description: "Edit your Subscriptions",
		},
	}
	cmdHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		cmdSubscriptions: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			userId := getChannelId(i)

			chains := chunks(database.GetChainsForUser(userId, user.TypeDiscord), 25)
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content:    subscriptionsMsg,
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
	}
)

var (
	actionHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate, action string){
		cmdSubscriptions: func(s *discordgo.Session, i *discordgo.InteractionCreate, action string) {
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
				log.Sugar.Errorf("Error while sending subscriptions: %v", err)
			}
		},
	}
)
