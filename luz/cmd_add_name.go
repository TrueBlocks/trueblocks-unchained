package main

import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/rpcClient/ens"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/validate"
)

func AddName() {
	n := types.SimpleName{}
	n.Address = types.HexToAddress(getValue(1, "address", "", "Enter an {0}", "An address is a 20-byte hexidecimal string.", validAddress))
	fmt.Println(n.Address)
	n.Name = getValue(2, "name", "", "Enter a {0} for this address", "A name is a human-readable label for an address.", notEmpty)
	fmt.Println(n.Name)
	n.Tags = getValue(3, "tag", "30-Contracts", "Enter a {0} for this address", "A tag is a human-readable label for an address.", notEmpty)
	fmt.Println(n.Tags)
	n.Source = getValue(3, "source", "EtherScan.io", "Enter a {0} for this address", "A source is a human-readable label for the data source for this address.", notEmpty)
	fmt.Println(n.Source)
	n.Symbol = getValue(4, "symbol", "", "Enter a {0} for this address", "A symbol is a human-readable label for the data source for this address.", nil)
	fmt.Println(n.Symbol)
	n.Decimals = mustParseUint(getValue(5, "decimals", "18", "Enter a {0} for this address", "A decimal is a human-readable label for the data source for this address.", nil))
	fmt.Println(n.Decimals)
	output.Format(&n, "csv")
}

func notEmpty(input string) (ok bool, reason string) {
	return input != "", "string is empty"
}

func validAddress(address string) (ok bool, reason string) {
	a := ens.ConvertOneEns("mainnet", address)
	return validate.IsValidAddress(a), fmt.Sprintf("'%s' is not a valid address", a)
}
