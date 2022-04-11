package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"wowtools/utilities"

	"github.com/spf13/viper"
)

func UpdateElvUI() {
	currentVersion := utilities.GetCurrentVersion()
	latestVersion := utilities.GetLatestVersion()
	stringCurrentVersion := strings.Join(currentVersion, "")
	filename := "elvui-" + latestVersion + ".zip"
	homeDir, err := os.UserHomeDir()
	downloadUri := "https://www.tukui.org/downloads/" + filename
	zipFile := homeDir + "\\Downloads\\" + filename

	if latestVersion > stringCurrentVersion {
		fmt.Printf("A later version of ElvUI is available. Current version: %s; New version: %s\n", stringCurrentVersion, latestVersion)
		updatePrompt := utilities.AskForConfirmation("Do you want to install the lastest version of ElvUI?")
		if updatePrompt {
			removeOldElvuiZip()
			ZipElvUI()
			fmt.Printf("Downloading ElvUI %s\n", latestVersion)
			utilities.DownloadFiles(filename, downloadUri)
			utilities.RemoveFolder(viper.GetString("elvui_dir"))
			utilities.RemoveFolder(viper.GetString("elvui_options_dir"))
			if err != nil {
				log.Fatal(err)
			}
			utilities.Unzip(zipFile, viper.GetString("addons_dir"))
			// if version is newer, zip up old installation and unzip new one.
		}
	} else {
		fmt.Println("ElvUI is up to date")
	}
}

func ZipElvUI() {
	elvuiFolder := viper.GetString("elvui_dir")
	backupFolder := viper.GetString("backup_dir") + "ElvUI\\"
	currentTime := time.Now()
	folderName := currentTime.Format("2006-01-02")

	fmt.Println("Beginning backup of ElvUI folder. This may take a moment")
	if err := utilities.ZipSource(elvuiFolder, backupFolder+folderName+".zip"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Folder backup complete")
}

func removeOldElvuiZip() {
	backupFolder := viper.GetString("backup_dir") + "ElvUI\\"
	fileCount := utilities.GetFileCount(backupFolder)
	fmt.Println(fileCount)
	if fileCount > 2 {
		oldestFile := utilities.GetOldestFolder(backupFolder)
		os.Chdir(backupFolder)
		removeFile := os.Remove(oldestFile)
		if removeFile != nil {
			log.Fatal()
		}
	}
}
