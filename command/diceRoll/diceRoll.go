package diceRoll

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type DiceRoll struct {
}

func (d *DiceRoll) Match(content string) (matched bool, err error) {
	matched, err = regexp.MatchString("^[0-9]+[d][0-9]+$", content)
	return
}

func (d *DiceRoll) Doc() (documentation string) {
	//TODO: Implementar
	return
}

func (d *DiceRoll) Run(content string) (response string, err error) {
	//TODO: Refatorar
	//Get parameters
	//Generate dices
	//prepare response
	ara := strings.Split(content, "d")
	if len(ara) != 2 {
		return "", errors.New("ERROR")
	}
	qtd, _ := strconv.Atoi(ara[0])
	max, _ := strconv.ParseInt(ara[1], 10, 64)
	if qtd < 1 || qtd > 50 {
		return "", errors.New("qtd error")
	}
	if max < 2 || max > 100 {
		return "", errors.New("max error")
	}

	dices := []int{}

	for i := 0; i < qtd; i++ {
		RandomCrypto, _ := rand.Int(rand.Reader, big.NewInt(max))
		dices = append(dices, int(RandomCrypto.Int64())+1)
	}
	sort.Slice(dices, func(i, j int) bool {
		return dices[i] > dices[j]
	})
	var diceSum int = 0
	for _, dice := range dices {
		diceSum += dice
	}
	return fmt.Sprint(diceSum) + " <= " + strings.Replace(fmt.Sprint(dices), " ", ", ", -1), nil
}
