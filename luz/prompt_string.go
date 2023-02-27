package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/colors"
	"github.com/c-bata/go-prompt"
)

func getString(pmpt, def, help string) string {
	userPrompt(pmpt, def)
	optionsPtr = &Options{
		CanAdd: true,
		Values: []option{},
	}
	filterFunc = nil
	val := prompt.Input("> ", getOptions)
	if isQuit(val) {
		os.Exit(0)
	} else if isHelp(val) {
		fmt.Printf("%s%s%s\n", colors.BrightGreen, help, colors.Off)
		return getString(pmpt, def, help)
	}
	return val
}

func getValue(pos int, field, def, prompt, help string, validate func(string) (bool, string)) string {
	val := ""
	if len(os.Args) > (pos + 1) {
		val = os.Args[pos+1]

	} else {
		val = os.Getenv("TB_NAME_" + strings.ToUpper(field))
		if val == "" {
			val = getString(strings.Replace(prompt, "{0}", field, -1), def, help)
			if def != "" && val == "" {
				val = def
			}
		}
	}

	if validate != nil {
		if ok, reason := validate(val); !ok {
			msg := "The value '" + val + "' is not valid for " + field + ": " + reason
			fmt.Printf("%s%s%s\n", colors.BrightRed, msg, colors.Off)
			return getValue(pos, field, def, prompt, help, validate)
		}
	}

	return val
}

func mustParseUint(input any) (result uint64) {
	result, _ = strconv.ParseUint(fmt.Sprint(input), 0, 64)
	return
}
