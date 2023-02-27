package main

import (
	"fmt"
	"os"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/colors"
	"github.com/c-bata/go-prompt"
)

func getChain() string {
	userPrompt("Choose a chain", lastChain())
	optionsPtr = &theConfig.Chains
	filterFunc = nil
	val := prompt.Input("> ", getOptions)
	if isQuit(val) {
		os.Exit(0)
	} else if isHelp(val) {
		msg := "Type the name of the chain you want to use."
		fmt.Printf("%s%s%s\n", colors.BrightGreen, msg, colors.Off)
		return getChain()
	} else if val == "" {
		val = lastChain()
	}

	NoteConfigValue("chain", val)
	return val
}

func lastChain() string {
	return theConfig.Chains.LastValue
}

var defaultChains = Options{
	LastValue: "mainnet",
	CanAdd:    true,
	Values: []option{
		{Text: "mainnet", Description: "Ethereum mainnet"},
		{Text: "sepolia", Description: "the Sepolia testnet"},
	},
}
