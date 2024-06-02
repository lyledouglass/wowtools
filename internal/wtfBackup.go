package internal

import (
	"fmt"
	"os"
	"time"
	"wowtools/internal/utilities"

	"github.com/spf13/viper"
)

func WtfBackup() {
	wtfFolder := viper.GetString("wtf_dir")
	wtfBackupDir := viper.GetString("backup_dir") + "WTF\\"
	currentTime := time.Now()
	folderName := currentTime.Format("2006-01-02")
	removeOldestWtfZip()
	fmt.Println("Beginning backup of WTF folder. This may take a moment")
	utilities.Log.Info("Beginning backup of WTF folder. This may take a moment")
	if err := utilities.ZipSource(wtfFolder, wtfBackupDir+folderName+".zip"); err != nil {
		utilities.Log.WithError(err).Error("Failed to zip source")
	}
	utilities.Log.Info("Folder backup complete")
}

func removeOldestWtfZip() {
	retentionRate := viper.GetInt("retention_rate")
	wtfBackupDir := viper.GetString("backup_dir") + "WTF\\"
	fileCount := utilities.GetFileCount(wtfBackupDir)
	if fileCount > retentionRate {
		oldestFile := utilities.GetOldestFolder(wtfBackupDir)
		err := os.Chdir(wtfBackupDir)
		if err != nil {
			utilities.Log.WithError(err).Errorf("Failed to change directories to %s", wtfBackupDir)
			return
		}
		removeFile := os.Remove(oldestFile)
		if removeFile != nil {
			utilities.Log.WithError(removeFile).Errorf("Failed to remove %s", oldestFile)
		}
	}
}
