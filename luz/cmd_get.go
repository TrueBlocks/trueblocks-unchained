package main

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/colors"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/rpcClient"
	"github.com/ethereum/go-ethereum"
)

func Get(theChain string) {
	msg := fmt.Sprintf("%s%s%s", colors.Yellow, "Getting the current value of the manifest hash...", colors.Off)
	logger.Log(logger.Info, msg)

	theAbi, err := getAbi()
	if err != nil {
		logger.Log(logger.Error, err.Error())
		return
	}

	method := theAbi.Methods["manifestHashMap"]
	encoding := hex.EncodeToString(method.ID[:4])
	publisher := getPublisher()
	chainStr := stringToHex(32*2, theChain)
	fmt.Printf("%s\n%s\n%s\n", encoding, publisher, chop(chainStr))
	input := "0x" + encoding + publisher + chainStr

	ctx := context.Background()
	theCall := ethereum.CallMsg{
		To:   &theUnchainedIndex,
		Data: rpcClient.DecodeHex(input),
	}

	ethClient := rpcClient.GetClient(config.GetRpcProvider("mainnet"))
	defer ethClient.Close()

	response, err := ethClient.CallContract(ctx, theCall, nil)
	if err != nil {
		fmt.Println(fmt.Errorf("while calling contract: %w", err))
		return
	}

	result, err := theAbi.Unpack("manifestHashMap", response)
	if err != nil {
		fmt.Println(fmt.Errorf("while unpacking value: %w", err))
		return
	}

	if len(result) == 0 {
		logger.Log(logger.Error, "contract returned empty data")
	} else {
		msg := fmt.Sprintf("%sUnchained Index for %s%s: %s%s%s\n", colors.Yellow, colors.BrightBlue, theChain, colors.BrightGreen, result[0].(string), colors.Off)
		logger.Log(logger.Info, msg)
	}
}
