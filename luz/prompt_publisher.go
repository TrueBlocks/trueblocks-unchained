package main

import (
	"fmt"
	"os"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/colors"
	"github.com/c-bata/go-prompt"
)

func getPublisher() string {
	userPrompt("Choose a publisher", lastPublisher())
	optionsPtr = &theConfig.Publishers
	filterFunc = nil
	val := prompt.Input("> ", getOptions)
	if isQuit(val) {
		os.Exit(0)
	} else if isHelp(val) {
		msg := "Enter the name of the publisher you'd like to read from or write from."
		fmt.Printf("%s%s%s\n", colors.BrightGreen, msg, colors.Off)
		return getPublisher()
	} else if val == "" {
		val = lastPublisher()
	}

	NoteConfigValue("publishers", val)
	return val
}

func lastPublisher() string {
	return theConfig.Publishers.LastValue
}

var defaultPublishers = Options{
	LastValue: "0xf503017d7baf7fbc0fff7492b751025c6a78179b",
	CanAdd:    true,
	Values: []option{
		{Text: "0xf503017d7baf7fbc0fff7492b751025c6a78179b", Description: "trueblocks.eth"},
	},
}
