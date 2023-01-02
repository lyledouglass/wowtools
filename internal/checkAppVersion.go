package internal

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	utilities "wowtools/pkg/utilities"
)

const wowtoolsUri = "https://api.github.com/repos/lyledouglass/wowtools/releases/latest"

func compareAppVersioning(currentVersion string, latestVersion string) bool {
	var updateApp bool
	fmt.Println(currentVersion)
	if currentVersion < latestVersion {
		updateApp = true
	}
	return updateApp
}

func UpdateWowtools() {
	// Variable declaration
	utilities.LoadConfig(".")
	var latestVersion = utilities.GetPublishedAppVersion(wowtoolsUri)
	var currentVersion = viper.GetString("wowtools_version")

	updateApp := compareAppVersioning(currentVersion, latestVersion)
	if updateApp == true {
		fmt.Printf("You are running on an older version (%s) of this application. Would you like to download the latest version (%s)?", currentVersion, latestVersion)
		updatePrompt := utilities.AskForConfirmation("")
		if updatePrompt {
			fmt.Println("Downloading latest package...")
			downloadUri := utilities.GetReleaseAsset(wowtoolsUri, "wowtools.exe")
			err := utilities.DownloadFiles("wowtools.exe", downloadUri)
			if err != nil {
				log.Fatal("Download step failed")
			}
			homeDir, _ := os.UserHomeDir()
			fmt.Printf("Wowtools version %s hase been downloaded to %s. Please close this application and replace it with the new executable", latestVersion, homeDir+"\\Downloads\\")
			fmt.Println("")
			fmt.Println("Press 'Enter' to close the program to update")
			var input string
			fmt.Scanln(&input)
			os.Exit(0)
		}
	} else {
		fmt.Println("wowtools is up to date, continuing...")
	}
}
