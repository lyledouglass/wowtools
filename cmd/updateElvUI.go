package cmd

import (
	"fmt"
	"log"
	"strings"
	"time"
	"wowtools/utilities"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/viper"
)

func UpdateElvUI() {
	currentVersion := utilities.GetCurrentVersion()
	latestVersion := utilities.GetLatestVersion()
	stringCurrentVersion := strings.Join(currentVersion, "")
	fmt.Println(stringCurrentVersion)
	fmt.Println(latestVersion)
	if latestVersion > stringCurrentVersion {
		fmt.Printf("A later version of ElvUI is available. Current version: %s; New version: %s\n", stringCurrentVersion, latestVersion)
		updatePrompt := utilities.AskForConfirmation("Do you want to install the lastest version of ElvUI?")
		if updatePrompt {
			fmt.Println("Zipping current ElvUI folder")
			ZipElvUI()
			fmt.Printf("Downloading ElvUI %s\n", latestVersion)
			DownloadElvUI(latestVersion)
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

	fmt.Println("Beginning backup of WTF folder")
	if err := utilities.ZipSource(elvuiFolder, backupFolder+folderName+".zip"); err != nil {
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

func DownloadElvUI(latestVersion string) {
	filename := "elvui-" + latestVersion + ".zip"
	downloadUri := "https://www.tukui.org/downloads/" + filename
	utilities.DownloadFiles(filename, downloadUri)
}
