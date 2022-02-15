package cmd

import (
	"fmt"
	"os"

	"github.com/formlessgo/m3u8/src/logger"

	"github.com/formlessgo/m3u8/src/download"
)

func main(url string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[error]", r)
			os.Exit(-1)
		}
	}()

	downloader, err := download.NewTask(url)
	if err != nil {
		logger.Fatal("%s", err)
	}
	err = downloader.Start(25, url)
	if err != nil {
		logger.Fatal("%s", err)
	}
	logger.Info("Done!")
}

// test URL: http://1257120875.vod2.myqcloud.com/0ef121cdvodtransgzp1257120875/3055695e5285890780828799271/v.f230.m3u8
