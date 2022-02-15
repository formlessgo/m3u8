package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/powerforus/color"

	"github.com/formlessgo/m3u8/src/config"
)

var date = time.Now().UTC().Format("2006-01-02 15:04:05")

func Debug(format string, v ...interface{}) {
	message := fmt.Sprintf(date+" [DEBUG] "+format, v...)
	colorMessage := fmt.Sprintf(date+color.Blue(" [DEBUG] ")+format, v...)
	writeLog(message + "\n")
	fmt.Fprintf(os.Stderr, colorMessage+"\n")
}

func Info(format string, v ...interface{}) {
	message := fmt.Sprintf(date+" [INFO] "+format, v...)
	colorMessage := fmt.Sprintf(date+color.Green(" [INFO] ")+format, v...)
	writeLog(message + "\n")
	fmt.Fprintf(os.Stderr, colorMessage+"\n")
}

func Error(format string, v ...interface{}) {
	message := fmt.Sprintf(date+" [ERROR] "+format, v...)
	colorMessage := fmt.Sprintf(date+color.Red(" [ERROR] ")+format, v...)
	writeLog(message + "\n")
	fmt.Fprintf(os.Stderr, colorMessage+"\n")
}

func Fatal(format string, v ...interface{}) {
	message := fmt.Sprintf(date+" [FATAL] "+format, v...)
	colorMessage := fmt.Sprintf(date+color.Yellow(" [FATAL] ")+format, v...)
	writeLog(message + "\n")
	fmt.Fprintf(os.Stderr, colorMessage+"\n")
	os.Exit(1)
}

func writeLog(words string) {
	filePath := config.DefaultConfig.LogPath
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	if _, err := file.Write([]byte(words)); err != nil {
		file.Close()
		fmt.Println(err)
	}
	if err := file.Close(); err != nil {
		fmt.Println(err)
	}
}
