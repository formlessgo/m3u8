package cmd

import (
	"fmt"

	"github.com/formlessgo/m3u8/src/constants"
)

var helpCollection = map[string]string{
	"config":  "show the config",
	"set":     "set config value",
	"help":    "show this help",
	"version": "show the version",
}

func help() {
	fmt.Printf("Usage: %s [command]\n\n", constants.BinName)
	// show the help of each command
	for k, v := range helpCollection {
		fmt.Printf("  %s: %s\n", k, v)
		fmt.Printf("\n")
	}
}
