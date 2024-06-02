package test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"wowtools/internal/utilities"
)

func TestRemoveFile(t *testing.T) {
	t.Skip()
	retentionRate := 15
	t.Log(retentionRate)
	wtfBackupDir := "C:\\Program Files (x86)\\World of Warcraft\\_retail_\\Backups\\WTF"
	t.Log(wtfBackupDir)
	fileCount := getFileCountTest(wtfBackupDir)
	t.Log(fileCount)
	if fileCount > retentionRate {
		oldestFile := utilities.GetOldestFolder(wtfBackupDir)
		err := os.Chdir(wtfBackupDir)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(oldestFile)
	}
}

func getFileCountTest(directory string) int {
	files, _ := os.ReadDir(directory)
	count := len(files)
	fmt.Println(count)
	return count
}
