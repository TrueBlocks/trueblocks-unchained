package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	cmd := getCommand()
	switch cmd {
	case "publish":
		Publish(getChain())
	case "get":
		Get(getChain())
	default:
		fmt.Println("Unknown command:", cmd)
	}
}

// func getSigs(theAbi *abi.ABI) []string {
// 	ret := []string{}
// 	ret = append(ret, "Quit")
// 	cnt := 1
// 	for _, method := range theAbi.Methods {
// 		f := types.FunctionFromAbiMethod(&method, "TrueBlocks")
// 		sig := fmt.Sprintf("\t%d. %s%s%s(", cnt, colors.BrightGreen, f.Name, colors.Off)
// 		for i, param := range f.Inputs {
// 			if i > 0 {
// 				sig += ", "
// 			}
// 			sig += param.Name + " " + param.ParameterType
// 		}
// 		sig += ")"
// 		ret = append(ret, sig)
// 		cnt++
// 	}
// 	return ret
// }

var theUnchainedIndex = common.HexToAddress("0x0c316b7042b419d07d343f2f4f5bd54ff731183d")

func getAbi() (*abi.ABI, error) {
	reader, err := os.Open(config.GetPathToRootConfig() + "abis/known-000/unchainedV2.json")
	if err != nil {
		return nil, fmt.Errorf("while reading contract ABI: %w", err)
	}
	defer reader.Close()

	theAbi, err := abi.JSON(reader)
	if err != nil {
		return nil, fmt.Errorf("while parsing contract ABI: %w", err)
	}

	return &theAbi, nil
}

func stringToHex(offset int, str string) string {
	off := fmt.Sprintf("%0.64x", offset) // offset in number of bytes
	strLen := fmt.Sprintf("%0.64d", len(str))
	val := fmt.Sprintf("%x%s", str, strings.Repeat("0", (32-(len(str)%32))*2))
	return off + strLen + val
}

func chop(str string) string {
	ret := []string{}
	for {
		if len(str) >= 64 {
			ret = append(ret, str[:64])
		} else {
			ret = append(ret, str)
			break
		}
		str = str[64:]
	}
	return strings.Trim(strings.Join(ret, "\n"), "\n")
}
