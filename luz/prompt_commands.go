package main

import (
	"fmt"
	"os"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/colors"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/c-bata/go-prompt"
)

func getCommand() string {
	val := ""
	if len(os.Args) < 2 {
		userPrompt("Choose a command", lastCommand())
		optionsPtr = &theConfig.Commands
		filterFunc = nil
		val = prompt.Input("> ", getOptions)
		if !validateCommand(val) {
			logger.Log(logger.Error, "You may not add new values to the commands option.")
			os.Exit(0)
		}
	} else {
		val = os.Args[1]
		if !validateCommand(val) {
			logger.Log(logger.Error, "Invalid command", val)
			os.Exit(0)
		}
	}

	if isQuit(val) {
		os.Exit(0)
	} else if isHelp(val) {
		msg := "Enter 'get' to retrieve the latest value of the manifest hash. Enter 'publish' to publish the manifest hash."
		fmt.Printf("%s%s%s\n", colors.BrightGreen, msg, colors.Off)
		os.Exit(0)
		// return getCommand()
	} else if val == "" {
		val = lastCommand()
	}

	NoteConfigValue("commands", val)
	return val
}

func lastCommand() string {
	return theConfig.Commands.LastValue
}

func validateCommand(val string) bool {
	if len(val) == 0 || isHelp(val) || isQuit(val) {
		return true
	}
	for _, opts := range theConfig.Commands.Values {
		if opts.Text == val {
			return true
		}
	}
	return false
}

var defaultCommands = Options{
	LastValue: "get",
	CanAdd:    false,
	Values: []option{
		{Text: "get", Description: "retrieve the latest value of the manifest hash"},
		{Text: "publish", Description: "publish the manifest hash"},
		{Text: "add-name", Description: "add a name to the names database"},
	},
}

func userPrompt(p string, d string) {
	if d == "" {
		fmt.Printf("%s%s. (tab for options)%s\n", colors.Yellow, p, colors.Off)
	} else {
		fmt.Printf("%s%s. (default = '%s', tab for options)%s\n", colors.Yellow, p, d, colors.Off)
	}
}

func isQuit(val string) bool {
	return val == "quit" || val == "q" || val == "Q"
}

func isHelp(val string) bool {
	return val == "help" || val == "h" || val == "H"
}
