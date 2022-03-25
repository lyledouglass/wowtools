package test

import (
	"log"
	"os"
	"testing"
	"wowtools/utilities"
)

func TestRemoveFile(t *testing.T) {
	os.Chdir("C:\\Temp\\wowtools")
	testFile, err := os.Create("Test.txt")
	if err != nil {
		log.Fatal()
	}
	t.Log(testFile)
	testFile.Close()

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
