package main

import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/colors"
)

func Publish() {
	// theChain := getChain()
	theDatabase := colors.BrightBlue + getDatabase() + colors.Off
	thePublisher := colors.BrightBlue + getPublisher() + colors.Off
	fmt.Println("Publishing", theDatabase, "database from", thePublisher, "publisher")
	// ethClient := rpcClient.GetClient(config.GetRpcProvider("mainnet"))
	// defer ethClient.Close()

	// address := common.HexToAddress("0x0c316b7042b419d07d343f2f4f5bd54ff731183d")
	// chain := "mainnet"

	// encoding := "7087e4bd"
	// publisher := "000000000000000000000000f503017d7baf7fbc0fff7492b751025c6a78179b"
}
