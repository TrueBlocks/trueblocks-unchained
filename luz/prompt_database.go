package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/colors"
	"github.com/c-bata/go-prompt"
)

func getDatabase() string {
	userPrompt("Database name", lastDatabase())
	optionsPtr = &theConfig.Databases
	filterFunc = nil
	val := prompt.Input("> ", getOptions)
	if isQuit(val) {
		os.Exit(0)
	} else if isHelp(val) {
		msg := "Enter the name of the database you wish to publish or retrieve. Any value is valid."
		fmt.Printf("%s%s%s\n", colors.BrightGreen, msg, colors.Off)
		return getDatabase()
	} else if val == "" {
		val = lastDatabase()
	}

	if len(val) == 0 {
		log.Panic("Should not happen in getDatabase")
	}

	NoteConfigValue("databases", val)
	return val
}

func lastDatabase() string {
	return theConfig.Databases.LastValue
}

var defaultDatabases = Options{
	LastValue: "",
	CanAdd:    true,
	Values: []option{
		{Text: "mainnet", Description: "IPFS hash for mainnet manifest"},
		{Text: "mainnet-ts", Description: "IPFS hash for mainnet timestamps database"},
		{Text: "sepolia", Description: "IPFS hash for sepolia manifest"},
		{Text: "sepolia-ts", Description: "IPFS hash for sepolia timestamps database"},
	},
}
