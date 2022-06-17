// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/rpcClient"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// We need an ABI for the Unchained Index
	reader, err := os.Open(config.GetPathToRootConfig() + "abis/known-000/unchainedV2.json")
	if err != nil {
		fmt.Println(fmt.Errorf("while reading contract ABI: %w", err))
		return
	}
	defer reader.Close()
	theAbi, err := abi.JSON(reader)
	if err != nil {
		fmt.Println(fmt.Errorf("while parsing contract ABI: %w", err))
		return
	}

	// We need access to the Ethereum client
	ethClient := rpcClient.GetClient(config.GetRpcProvider("mainnet"))
	defer ethClient.Close()

	// We need the address of the smart contract
	address := common.HexToAddress("0x0c316b7042b419d07d343f2f4f5bd54ff731183d")

	// And here we make the call
	response, err := ethClient.CallContract(
		context.Background(),
		ethereum.CallMsg{
			To:   &address,
			Data: rpcClient.DecodeHex("0x7087e4bd00000000000000000000000002f2b09b33fdbd406ead954a31f98bd29a2a3492000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000076d61696e6e657400000000000000000000000000000000000000000000000000"),
		},
		nil,
	)
	if err != nil {
		fmt.Println(fmt.Errorf("while calling contract: %w", err))
		return
	}

	// Unpack the result...
	result, err := theAbi.Unpack("manifestHashMap", response)
	if err != nil {
		fmt.Println(fmt.Errorf("while unpacking value: %w", err))
		return
	}

	// And print it if we got anything
	if len(result) == 0 {
		fmt.Println(errors.New("contract returned empty data"))
	} else {
		fmt.Println(result[0].(string))
	}
}
