package bot

import (
	"fmt"
	"golang-discord-bot/command"
	"golang-discord-bot/config"

	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running !")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	response, err := command.Resolver.Run(m)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
	} else {
		s.ChannelMessageSend(m.ChannelID, response)
	}
}
