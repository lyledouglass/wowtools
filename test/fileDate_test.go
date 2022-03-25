package test

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

func TestCleanFolders(t *testing.T) {
	filepath := "C:\\Program Files (x86)\\World of Warcraft\\_retail_\\Backups\\WTF"
	var oldestFile fs.FileInfo
	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		log.Fatal(err)
	}
	oldestTime := time.Now()
	for _, file := range files {
		if file.Mode().IsRegular() && file.ModTime().Before(oldestTime) {
			oldestFile = file
			oldestTime = file.ModTime()
		}
	}
	if oldestFile == nil {
		err = os.ErrNotExist
	}
	fmt.Println(oldestFile.Name())
}
