package cmd

import (
	"fmt"
	"log"
	"os"
	"wowtools/utilities"
)

const wowtoolsUri = "https://api.github.com/repos/lyledouglass/wowtools/releases/latest"

var latestVersion = utilities.GetPublishedAppVersion(wowtoolsUri)

func compareAppVersioning() bool {
	var updateApp bool
	currentVersion := utilities.GetCurrentAppVersion()
	if currentVersion > latestVersion {
		updateApp = true
	}
	return updateApp
}

func UpdateWowtools() {
	updateApp := compareAppVersioning()
	if updateApp == true {
		updatePrompt := utilities.AskForConfirmation("You are running on an older version of this application. Would you like to download the latest version?")
		if updatePrompt {
			fmt.Println("Downloading latest package...")
			downloadUri := utilities.GetReleaseAsset(wowtoolsUri, "wowtools.exe")
			err := utilities.DownloadFiles("wowtools.exe", downloadUri)
			if err != nil {
				log.Fatal("Download step failed")
			}
			homeDir, _ := os.UserHomeDir()
			fmt.Printf("Wowtools version %s hase been downloaded to %s. Please close this application and replace it with the new executable", latestVersion, homeDir+"\\Downloads\\")
			fmt.Println("Press 'Enter' to close the program to update")
			var input string
			fmt.Scanln(&input)
			os.Exit(0)
		}
	} else {
		fmt.Println("wowtools is up to date, continuing...")
	}
}
