package cmd

import (
	"fmt"

	"github.com/formlessgo/m3u8/src/constants"
)

func version() {
	fmt.Printf("%s version: %s\n", constants.AppName, constants.Version)
}
