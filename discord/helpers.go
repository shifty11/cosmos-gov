package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/shifty11/cosmos-gov/log"
	"strconv"
	"strings"
)

func getChannelId(i *discordgo.InteractionCreate) int64 {
	channelId, err := strconv.ParseInt(i.ChannelID, 10, 64)
	if err != nil {
		log.Sugar.Panicf("Error while converting user ID to int: %v", err)
	}
	return channelId
}

type Action struct {
	Name string
	Data string
}

func getAction(dataStr string) Action {
	var data = strings.Split(dataStr, ":")
	if len(data) > 1 {
		return Action{Name: data[0], Data: data[1]}
	}
	return Action{Name: data[0], Data: ""}
}
