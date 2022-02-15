package config

import (
	"github.com/pelletier/go-toml/v2"

	"github.com/formlessgo/m3u8/src/utils/file"
)

type Config struct {
	DownloadPath string
	DownloadName string
	LogPath      string
	Transcode    bool
}

var DefaultConfig = Config{
	DownloadPath: downloadPath,
	DownloadName: "",
	LogPath:      logPath,
	Transcode:    false,
}

var ConfigPath = configPath

func InitConfig() {
	if !file.Exist(ConfigPath) {
		data, err := toml.Marshal(DefaultConfig)
		if err != nil {
			panic(err)
		}
		file.Create(ConfigPath)
		file.Write(string(data), ConfigPath)
	}
}

func GetConfig() Config {
	data := file.Read(ConfigPath)
	var _config Config
	err := toml.Unmarshal([]byte(data), &_config)
	if err != nil {
		panic(err)
	}
	return _config
}
