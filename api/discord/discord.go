package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
	"os"
	"os/signal"
)

//goland:noinspection GoNameStartsWithPackageName
type DiscordClient struct {
	s                          *discordgo.Session
	DiscordChannelManager      *database.DiscordChannelManager
	DiscordSubscriptionManager *database.DiscordSubscriptionManager
	ProposalManager            *database.ProposalManager
}

func NewDiscordClient(managers *database.DbManagers) *DiscordClient {
	return &DiscordClient{
		DiscordChannelManager:      managers.DiscordChannelManager,
		DiscordSubscriptionManager: managers.DiscordSubscriptionManager,
		ProposalManager:            managers.ProposalManager,
	}
}

func (dc DiscordClient) initDiscord() {
	log.Sugar.Info("Init discord bot")

	var err error
	dc.s, err = discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Sugar.Fatalf("Invalid bot parameters: %v", err)
	}
	dc.s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		//TODO: make this multithreaded (see Telegram bot)
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := cmdHandlers[i.ApplicationCommandData().Name]; ok {
				h(dc, s, i)
			}
		case discordgo.InteractionMessageComponent:
			var action = getAction(i.MessageComponentData().CustomID)
			if h, ok := actionHandlers[action.Name]; ok {
				h(dc, s, i, action.Data)
			}
		}
	})
}

func (dc DiscordClient) addCommands() {
	for _, v := range cmds {
		_, err := dc.s.ApplicationCommandCreate(dc.s.State.User.ID, "", v)
		if err != nil {
			log.Sugar.Panic("Cannot create '%v' command: %v", v.Name, err)
		}
	}
}
func (dc DiscordClient) removeCommands() {
	registeredCommands, err := dc.s.ApplicationCommands(dc.s.State.User.ID, "")
	if err != nil {
		log.Sugar.Fatalf("Could not fetch registered commands: %v", err)
	}

	for _, v := range registeredCommands {
		err := dc.s.ApplicationCommandDelete(dc.s.State.User.ID, "", v.ID)
		if err != nil {
			log.Sugar.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}

func (dc DiscordClient) Start() {
	dc.initDiscord()
	log.Sugar.Info("Start discord bot")

	err := dc.s.Open()
	if err != nil {
		log.Sugar.Fatalf("Cannot open the s: %v", err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer dc.s.Close()

	dc.removeCommands()
	dc.addCommands()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
