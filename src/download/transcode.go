package download

import (
	"os/exec"
	"path/filepath"

	"github.com/formlessgo/m3u8/src/config"
	"github.com/formlessgo/m3u8/src/logger"
)

func (d *Downloader) transcode(name string) {
	logger.Info("Start transcode %s.ts", name)
	_, err := exec.Command("ffmpeg", "-hide_banner", "-loglevel", "panic", "-i", filepath.Join(config.DefaultConfig.DownloadPath, name+".ts"), "-f", "mp4", "-c", "copy", filepath.Join(config.DefaultConfig.DownloadPath, name+".mp4")).Output()
	if err != nil {
		logger.Error("Transcode failed: %s", err)
	} else {
		logger.Info("Transcode success")
		logger.Info("Output file: %s", filepath.Join(config.DefaultConfig.DownloadPath, name+".mp4"))
	}
}
