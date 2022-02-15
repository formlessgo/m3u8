package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"

	"github.com/formlessgo/m3u8/src/config"
	"github.com/formlessgo/m3u8/src/utils/file"
)

func set() {
	data := file.Read(config.ConfigPath)
	// convert toml to struct
	var _config config.Config
	err := toml.Unmarshal([]byte(data), &_config)
	if err != nil {
		panic(err)
	}
	// convert struct to json
	jsonData, err := json.Marshal(_config)
	if err != nil {
		fmt.Println(err.Error())
	}
	// convert json to map
	var mapData map[string]interface{}
	err = json.Unmarshal(jsonData, &mapData)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		return
	}
	mapData[os.Args[2]] = os.Args[3]
	// convert map to json
	byteData, err := json.Marshal(mapData)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}
	// convert json to stuct
	var _config2 config.Config
	json.Unmarshal(byteData, &_config2)
	// convert struct to toml
	data2, err := toml.Marshal(_config2)
	if err != nil {
		panic(err)
	}
	file.Write(string(data2), config.ConfigPath)
}
