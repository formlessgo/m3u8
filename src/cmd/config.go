package cmd

import (
	"fmt"

	"github.com/formlessgo/m3u8/src/config"
	"github.com/formlessgo/m3u8/src/constants"
	"github.com/formlessgo/m3u8/src/utils/file"
)

func _config() {
	fmt.Printf("Config file path: %s\n", config.ConfigPath)
	data := file.Read(config.ConfigPath)
	fmt.Printf("%s config:\n", constants.AppName)
	fmt.Println(data)
}
