package main

import (
	"os"

	"github.com/c-bata/go-prompt"
)

func getChain() string {
	userPrompt("Choose a chain", lastChain())
	val := prompt.Input("> ", chainOptions)
	if val == "quit" {
		os.Exit(0)
	}
	if val == "" {
		val = lastChain()
	}
	return val
}

func lastChain() string {
	return "mainnet"
}

func chainOptions(d prompt.Document) []prompt.Suggest {
	var s = []prompt.Suggest{
		{Text: "mainnet", Description: "Ethereum mainnet"},
		{Text: "sepolia", Description: "the Sepolia testnet"},
		{Text: "quit", Description: "quit this tool"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
