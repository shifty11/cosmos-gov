package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/shifty11/cosmos-gov/log"
	"os"
	"os/signal"
)

var s *discordgo.Session

func initDiscord() {
	log.Sugar.Info("Init discord bot")

	var err error
	s, err = discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Sugar.Fatalf("Invalid bot parameters: %v", err)
	}
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := cmdHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		case discordgo.InteractionMessageComponent:
			var action = getAction(i.MessageComponentData().CustomID)
			if h, ok := actionHandlers[action.Name]; ok {
				h(s, i, action.Data)
			}
		}
	})
	s.AddHandler(messageHandler)
}

func addCommands() {
	for _, v := range cmds {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Sugar.Panic("Cannot create '%v' command: %v", v.Name, err)
		}
	}
}
func removeCommands() {
	registeredCommands, err := s.ApplicationCommands(s.State.User.ID, "")
	if err != nil {
		log.Sugar.Fatalf("Could not fetch registered commands: %v", err)
	}

	for _, v := range registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			log.Sugar.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}

func Start() {
	initDiscord()
	log.Sugar.Info("Start discord bot")

	err := s.Open()
	if err != nil {
		log.Sugar.Fatalf("Cannot open the s: %v", err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer s.Close()

	removeCommands()
	addCommands()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}

//Definition of messageHandler function it takes two arguments first one is discordgo.Session which is s , second one is discordgo.MessageCreate which is m.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	//Bot musn't reply to it's own messages , to confirm it we perform this check.
	if m.Author.ID == s.State.User.ID {
		return
	}
	//If we message ping to our bot in our discord it will return us pong .
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "*pong*: "+m.ChannelID)
	}
}
