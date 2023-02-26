package main

import (
	"fmt"
	"os"

	"github.com/c-bata/go-prompt"
)

func getPublisher() string {
	userPrompt("Choose a publisher", lastPublisher())
	val := prompt.Input("> ", publisherOptions)
	if val == "quit" {
		os.Exit(0)
	}
	if val == "" {
		val = lastPublisher()
	}
	return fmt.Sprintf("%064s", val)
}

func lastPublisher() string {
	return "f503017d7baf7fbc0fff7492b751025c6a78179b"
}

func publisherOptions(d prompt.Document) []prompt.Suggest {
	var s = []prompt.Suggest{
		{Text: "f503017d7baf7fbc0fff7492b751025c6a78179b", Description: "trueblocks.eth"},
		{Text: "quit", Description: "quit this tool"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
