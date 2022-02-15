package config

import (
	"os"
	"path/filepath"

	"github.com/formlessgo/m3u8/src/constants"
)

var homeDir, _ = os.UserHomeDir()
var configPath = filepath.Join(homeDir, ".config", constants.BinName, "config.toml")
var downloadPath = filepath.Join(homeDir, ".config", constants.BinName, "downloads")
var logPath = filepath.Join(homeDir, ".config", constants.BinName, constants.BinName+".log")
