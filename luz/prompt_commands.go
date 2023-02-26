package main

import (
	"fmt"
	"os"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/colors"
	"github.com/c-bata/go-prompt"
)

func getCommand() string {
	userPrompt("Choose a command", lastCommand())
	val := prompt.Input("> ", commandOptions)
	if val == "quit" {
		os.Exit(0)
	}
	if val == "" {
		val = lastCommand()
	}
	return val
}

func lastCommand() string {
	return "get"
}

func commandOptions(d prompt.Document) []prompt.Suggest {
	var s = []prompt.Suggest{
		{Text: "get", Description: "retrieve the latest value of the manifest hash"},
		{Text: "publish", Description: "publish the manifest hash"},
		{Text: "quit", Description: "quit this tool"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func userPrompt(p string, d string) {
	fmt.Printf("%s%s. (default = '%s', tab for options)%s\n", colors.Yellow, p, d, colors.Off)
}
