package test

import (
	"log"
	"os"
	"testing"
	"wowtools/pkg/utilities"
)

func TestRemoveFile(t *testing.T) {
	t.Skip()
	err := os.Chdir("C:\\Temp\\wowtools")
	if err != nil {
		log.Fatal(err)
	}
	testFile, err := os.Create("Test.txt")
	if err != nil {
		log.Fatal()
	}
	t.Log(testFile)
	err = testFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	retentionRate := 1
	fileCount := utilities.GetFileCount("C:\\Temp\\wowtools\\")
	t.Log(fileCount)
	if fileCount > retentionRate {
		oldestFile := utilities.GetOldestFolder("C:\\Temp\\wowtools")
		t.Log(oldestFile)
		removeFile := os.Remove(oldestFile)
		if removeFile != nil {
			log.Fatal()
		}
	}
}
