package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"wowtools/utilities"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/viper"
)

func WtfBackup() {
	wtfFolder := viper.GetString("wtf_dir")
	wtfBackupDir := viper.GetString("backup_dir") + "WTF\\"
	currentTime := time.Now()
	folderName := currentTime.Format("2006-01-02")

	removeOldElvuiZip()

	fmt.Println("Beginning backup of WTF folder")
	if err := utilities.ZipSource(wtfFolder, wtfBackupDir+folderName+".zip"); err != nil {
		log.Fatal(err)
	}
	// Not really a true progress bar at the moment - more of a visual for the user - need to reseach better implementation, but works for now, as the zip process is fairly quick for the WTF folder
	bar := progressbar.Default(100)
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(20 * time.Millisecond)
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
