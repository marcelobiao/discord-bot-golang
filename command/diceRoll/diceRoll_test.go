package diceRoll_test

import (
	"fmt"
	"golang-discord-bot/command/diceRoll"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	diceRoll := diceRoll.DiceRoll{}

	response, err := diceRoll.Run("14d10")
	fmt.Println(response, err)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
}
