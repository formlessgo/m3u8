package cmd

import (
	"os"
	"strings"

	"github.com/formlessgo/m3u8/src/config"
)

func Mod() {
	// init config
	config.InitConfig()
	// exec command
	if len(os.Args) == 2 {
		if os.Args[1] == "config" {
			_config()
		} else if os.Args[1] == "version" {
			version()
		} else if strings.HasSuffix(os.Args[1], ".m3u8") {
			main(os.Args[1])
		} else if strings.Contains(os.Args[1], ".m3u8") {
			main(os.Args[1])
		} else {
			help()
		}
	} else if len(os.Args) == 4 {
		if os.Args[1] == "set" {
			set()
		}
	}
}
