package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"wowtools/utilities"

	"github.com/spf13/viper"
)

func WtfBackup() {
	wtfFolder := viper.GetString("wtf_dir")
	wtfBackupDir := viper.GetString("backup_dir") + "WTF\\"
	currentTime := time.Now()
	folderName := currentTime.Format("2006-01-02")
	removeOldElvuiZip()
	fmt.Println("Beginning backup of WTF folder. This may take a moment")
	if err := utilities.ZipSource(wtfFolder, wtfBackupDir+folderName+".zip"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Folder backup complete")
}

func removeOldestWtfZip() {
	retentionRate := viper.GetInt("retention_rate")
	wtfBackupDir := viper.GetString("backup_dir") + "WTF\\"
	fileCount := utilities.GetFileCount(wtfBackupDir)
	if fileCount > retentionRate {
		oldestFile := utilities.GetOldestFolder(wtfBackupDir)
		os.Chdir(wtfBackupDir)
		removeFile := os.Remove(oldestFile)
		if removeFile != nil {
			log.Fatal()
		}
	}
}
