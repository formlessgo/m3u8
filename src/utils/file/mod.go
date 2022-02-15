package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func Exist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func Create(path string) {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func Read(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func Write(data string, path string) {
	err := ioutil.WriteFile(path, []byte(data), 0666)
	if err != nil {
		panic(err)
	}
}
