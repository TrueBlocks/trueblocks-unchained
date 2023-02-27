package main

import (
	"fmt"
	"os"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/colors"
	"github.com/c-bata/go-prompt"
)

func yesNo(msg string) {
	userPrompt(msg, "no")
	optionsPtr = &Options{
		LastValue: "no",
		CanAdd:    false,
		Values: []option{
			{Text: "yes", Description: "yes, complete the operation"},
			{Text: "no", Description: "no, do not complete the operation"},
		},
	}
	filterFunc = nil
	val := prompt.Input("> ", getOptions)
	if isQuit(val) {
		os.Exit(0)
	} else if isHelp(val) {
		msg := "Enter 'yes' or 'no'."
		fmt.Printf("%s%s%s\n", colors.BrightGreen, msg, colors.Off)
		yesNo(msg)
		return
	} else if val == "" {
		val = "no"
	}

	if val != "yes" && val != "y" && val != "Y" {
		fmt.Println("Canceled.")
		os.Exit(0)
	}
}
