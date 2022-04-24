package command

import (
	"golang-discord-bot/command/diceRoll"
	pingpong "golang-discord-bot/command/pingPong"

	"github.com/bwmarrin/discordgo"
)

var Resolver CommandResolver = CommandResolver{}

type CommandResolver struct {
}

type CommandInterface interface {
	Match(content string) (matched bool, err error)
	Doc() (documentation string)
	Run(content string) (response string, err error)
}

var commandResolverSlice []CommandInterface = []CommandInterface{
	&pingpong.PingPong{},
	&diceRoll.DiceRoll{},
}

func (c *CommandResolver) Run(m *discordgo.MessageCreate) (response string, err error) {
	var matched bool
	for _, command := range commandResolverSlice {
		matched, err = command.Match(m.Content)
		if err != nil {
			return
		}
		if matched {
			return command.Run(m.Content)
		}
	}
	return
}
