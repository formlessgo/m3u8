package download

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/formlessgo/m3u8/src/config"
	"github.com/formlessgo/m3u8/src/logger"
	"github.com/formlessgo/m3u8/src/parse"
	"github.com/formlessgo/m3u8/src/utils"
	"github.com/formlessgo/m3u8/src/utils/base64"
)

const (
	tsExt            = ".ts"
	tempFolderName   = "temp"
	mergeTSFilename  = "main.ts"
	tsTempFileSuffix = "_tmp"
	progressWidth    = 40
)

type Downloader struct {
	lock     sync.Mutex
	queue    []int
	folder   string
	tsFolder string
	finish   int32
	segLen   int

	result *parse.Result
}

// NewTask returns a Task instance
func NewTask(url string) (*Downloader, error) {
	result, err := parse.FromURL(url)
	if err != nil {
		return nil, err
	}
	folder := config.GetConfig().DownloadPath
	if folder == "_" {
		folder = config.DefaultConfig.DownloadPath
	}
	err = os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		logger.Error("create storage folder failed: %s", err)
	}
	tempFolder := filepath.Join(folder, tempFolderName)
	err = os.MkdirAll(tempFolder, os.ModePerm)
	if err != nil {
		logger.Error("create ts folder '[%s]' failed: %s", tempFolder, err.Error())
	}
	d := &Downloader{
		folder:   folder,
		tsFolder: tempFolder,
		result:   result,
	}
	d.segLen = len(result.M3u8.Segments)
	d.queue = genSlice(d.segLen)
	return d, nil
}

// Start runs downloader
func (d *Downloader) Start(concurrency int, url string) error {
	var wg sync.WaitGroup
	// struct{} zero size
	limitChan := make(chan struct{}, concurrency)
	for {
		tsIdx, end, err := d.next()
		if err != nil {
			if end {
				break
			}
			continue
		}
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			if err := d.download(idx); err != nil {
				// Back into the queue, retry request
				fmt.Printf("[failed] %s\n", err.Error())
				if err := d.back(idx); err != nil {
					logger.Error("%s", err)
				}
			}
			<-limitChan
		}(tsIdx)
		limitChan <- struct{}{}
	}
	wg.Wait()

	downloadName := config.GetConfig().DownloadName
	if downloadName == "" || downloadName == "_" {
		downloadName = base64.Base64(url)
	}
	err := d.merge(downloadName)

	if err != nil {
		logger.Error("%s", err)
	}
	if config.GetConfig().Transcode {
		d.transcode(downloadName)
	}
	return nil
}

func (d *Downloader) next() (segIndex int, end bool, err error) {
	d.lock.Lock()
	defer d.lock.Unlock()
	if len(d.queue) == 0 {
		err = fmt.Errorf("queue empty")
		if d.finish == int32(d.segLen) {
			end = true
			return
		}
		// Some segment indexes are still running.
		end = false
		return
	}
	segIndex = d.queue[0]
	d.queue = d.queue[1:]
	return
}

func (d *Downloader) back(segIndex int) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	if sf := d.result.M3u8.Segments[segIndex]; sf == nil {
		return fmt.Errorf("invalid segment index: %d", segIndex)
	}
	d.queue = append(d.queue, segIndex)
	return nil
}

func (d *Downloader) tsURL(segIndex int) string {
	seg := d.result.M3u8.Segments[segIndex]
	return utils.ResolveURL(d.result.URL, seg.URI)
}

func tsFilename(ts int) string {
	return strconv.Itoa(ts) + tsExt
}

func genSlice(len int) []int {
	s := make([]int, 0)
	for i := 0; i < len; i++ {
		s = append(s, i)
	}
	return s
}
