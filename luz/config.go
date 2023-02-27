package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/c-bata/go-prompt"
)

type Options struct {
	LastValue string   `json:"lastValue"`
	CanAdd    bool     `json:"canAdd"`
	Values    []option `json:"values"`
}

type option struct {
	Text        string `json:"text"`
	Description string `json:"description"`
}

type Config struct {
	Commands   Options `json:"commands"`
	Chains     Options `json:"chains"`
	Publishers Options `json:"publishers"`
	Databases  Options `json:"databases"`
}

var theConfig Config

func init() {
	if file.FileExists(".config.json") {
		jsonFile, err := os.Open(".config.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &theConfig)
	} else {
		theConfig = Config{
			Chains:     defaultChains,
			Commands:   defaultCommands,
			Publishers: defaultPublishers,
			Databases:  defaultDatabases,
		}
		WriteConfig(theConfig)
	}
}

func NoteConfigValue(name string, val string) {
	switch name {
	case "chain":
		theConfig.Chains.LastValue = val
		theConfig.Chains = AddOption(theConfig.Chains, val)
	case "commands":
		theConfig.Commands.LastValue = val
		theConfig.Commands = AddOption(theConfig.Commands, val)
	case "databases":
		theConfig.Databases.LastValue = val
		theConfig.Databases = AddOption(theConfig.Databases, val)
	case "publishers":
		theConfig.Publishers.LastValue = val
		theConfig.Publishers = AddOption(theConfig.Publishers, val)
	}
	WriteConfig(theConfig)
}

func AddOption(opts Options, val string) Options {
	for _, opt := range opts.Values {
		if opt.Text == val {
			return opts
		}
	}

	yesNo("Are you sure you want to add a new value '" + val + "' to the options?")

	opts.Values = append(opts.Values, option{Text: val, Description: "user-entered value"})
	return opts
}

func WriteConfig(config Config) {
	file, _ := json.MarshalIndent(config, "", " ")
	_ = ioutil.WriteFile(".config.json", file, 0644)
}

var optionsPtr *Options = nil
var filterFunc func(string) bool = nil

func getOptions(d prompt.Document) []prompt.Suggest {
	var s = []prompt.Suggest{}
	for _, opts := range optionsPtr.Values {
		if filterFunc == nil || filterFunc(opts.Text) {
			s = append(s, prompt.Suggest{Text: opts.Text, Description: opts.Description})
		}
	}
	s = append(s, prompt.Suggest{Text: "help", Description: "get help on these options"})
	s = append(s, prompt.Suggest{Text: "quit", Description: "quit this tool"})

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
