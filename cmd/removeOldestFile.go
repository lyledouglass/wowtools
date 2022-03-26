package cmd

import (
	"log"
	"os"
	"wowtools/utilities"

	"github.com/spf13/viper"
)

func RemoveOldestWtfZip() {
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
