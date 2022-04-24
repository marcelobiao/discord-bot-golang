package command

import (
	"golang-discord-bot/command/diceRoll"
	pingpong "golang-discord-bot/command/pingPong"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

var Resolver CommandResolver = CommandResolver{
	PingPong: &pingpong.PingPong{},
	DiceRoll: &diceRoll.DiceRoll{},
}

type CommandResolver struct {
	PingPong CommandInterface
	DiceRoll CommandInterface
}

func GetCommandRecolver() CommandResolver {
	return Resolver
}

type CommandInterface interface {
	Run(content string) (response string, err error)
}

func (c *CommandResolver) Run(m *discordgo.MessageCreate) (response string, err error) {

	if m.Content == "ping" {
		return c.PingPong.Run(m.Content)
	}

	match, err := regexp.MatchString("^[0-9]+[d][0-9]+$", m.Content)
	if match {
		return c.DiceRoll.Run(m.Content)
	}
	return
}
