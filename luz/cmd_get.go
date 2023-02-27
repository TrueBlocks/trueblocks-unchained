package main

import (
	"fmt"
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/colors"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/manifest"
)

func Get() {
	theDatabase := getDatabase()
	thePublisher := getPublisher()

	parts := strings.Split(theDatabase, "-")
	if len(parts) != 2 {
		parts = append(parts, "")
	}

	msg := fmt.Sprintf("%s%s%s", colors.Yellow, "Getting the current value of the manifest hash...", colors.Off)
	logger.Log(logger.Info, msg)

	cid, err := manifest.ReadUnchainIndex(parts[0], parts[1], thePublisher)
	if err != nil {
		logger.Log(logger.Error, err.Error())
		return
	}

	if len(cid) == 0 {
		logger.Log(logger.Error, "contract returned empty data")
	} else {
		db := colors.BrightBlue + theDatabase + colors.Off
		pub := colors.BrightBlue + thePublisher + colors.Off
		res := colors.BrightGreen + cid + colors.Off
		msg := fmt.Sprintf("%sUnchained Index published by %s for %s: %s\n", colors.Yellow, pub, db, res)
		logger.Log(logger.Info, msg)
	}
}
