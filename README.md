# The Unchained Index

The UnchainedIndex is a smart contract that accomplishes a very simple task: it allows anyone to publish 
a pointer to any data and records the address of that publisher.

That's all it does. That's it.

While the previous sentence sounds sort of silly, behind it there's an interesting idea.

## IPFS hashes are pointers into a world-wide shared global memory

A pointer in a programming language is able to point to an arbitrarily large piece of memory. In the same
way, an IPFS hash can be seen to do the same thing. Pointers can point to something as simple as an integer
or something as complex as an multi-Gigabyte in-memory database. Pointers to Memory-mapped files can point
to Terabytes of hard drive space.

No matter what the pointer points to, it's the same size. In the case of memory pointer, the pointer
is a 64-bit unsigned integer. In the case of IPFS, the pointer is a 32-byte IPFS hash.

The UnchainedIndex, by providing a place to record the publication of pointers, can be seen as granting
access to a global memory space. Unlike a computer's memory or a database, however, the memory pointed
to by an IPFS hash is immutable and repeatable. Furthermore, because we publish the IPFS to a smart
contract, the fact that we did so can never disappear. Once we publish, everyone can see the pointer
for the rest of time.

## Unchained Index is permissionless and immutable

We've chosen to make Unchained Index permissionless, in the sense that anyone may read from it, but more importantly, anyone
may write to it. Each time someone calls `publishHash`, the contract records the sender's address. We call this address 
the `publisher` address. Anyone may then later query for data produced by that `publisher`. The Unchained Index doesn't
care.

There is a Preferred publisher (called `owner`) which is the first deployer of the contract (us!). In this way, if a user wishes
to retrieve data published by TrueBlocks, they may do so by querying the pointer provided by us. However, anyone else may also
publish, and users who wish to may choose to query against that publisher. The Unchained Index doesn't care.

In this sense, the Unchained Index is a form of 'oracle by reputation'. If we can convince our users that the memory we
point to is better than anyone else's they will use it. If someone else comes along and convinces people to use thier
pointer to memory, Unchained Index doesn't care. It's not on the end user's decision, it's his/her responsibility.

If the "US Academy of Accounting Professionals" (a made-up organization) one day decides to publish an IPFS hash
pointing to the world's best Ethereum mainnet index, so be it. We can't (and don't want to) stop our users from
switching over. This is all by design. We're purposefully giving up our ability to capture our users.

The contract is also immutable (i.e. non-upgradable). This is also by design. We are also specifically relinquishing
our ability to change the rules out from under our users. We're creating an immutable history to the data that's
being produced by our system.

## Pre-requisites for building

You must have `foundry` installed. [Follow these instructions](https://book.getfoundry.sh/getting-started/installation.html).

## Building the Unchained Index

```[bash]
git clone https://github.com/TrueBlocks/trueblocks-unchained
cd trueblocks-unchained
forge build
```

## Testing

```[bash]
cd trueblocks-unchained
forge test
```

## Reading the IPFS hash of the manifest

### Go code

```[go]
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
```

If you run this code with

```
go run main.go
```

it prints

```
QmcBzCmvdcY5s3qt8fLz8hcYxS8QR2K7KoBuC2qi2NuaTx
```

You may access the manifest with 

```[bash]
curl "https://gateway.pinata.cloud/ipfs/QmcBzCmvdcY5s3qt8fLz8hcYxS8QR2K7KoBuC2qi2NuaTx"
```
