package internal

import (
	"fmt"
	"os"
	"wowtools/pkg/utilities"
)

const wowtoolsUri = "https://api.github.com/repos/lyledouglass/wowtools/releases/latest"

func compareAppVersioning(currentVersion string, latestVersion string) bool {
	var updateApp bool
	utilities.Log.Debug(currentVersion)
	utilities.Log.Debug(latestVersion)
	if currentVersion < latestVersion {
		updateApp = true
	}
	return updateApp
}

func UpdateWowtools() {
	// Variable declaration
	var latestVersion = utilities.GetPublishedAppVersion(wowtoolsUri)
	var currentVersion = utilities.CurrentAppVersion()

	updateApp := compareAppVersioning(currentVersion, latestVersion)
	if updateApp {
		utilities.Log.Infof("You are running on an older version (%s) of this application. Would you like to download the latest version (%s)?", currentVersion, latestVersion)
		updatePrompt := utilities.AskForConfirmation("")
		if updatePrompt {
			utilities.Log.Debug("Downloading latest package")
			downloadUri := utilities.GetReleaseAsset(wowtoolsUri, "wowtools_client.exe")
			err := utilities.DownloadFiles("wowtools.exe", downloadUri)
			if err != nil {
				utilities.Log.WithError(err).Error("Download of new Wowtools failed")
			}
			homeDir, _ := os.UserHomeDir()
			utilities.Log.Infof("Wowtools version %s hase been downloaded to %s. Please close this application and replace it with the new executable", latestVersion, homeDir+"\\Downloads\\")
			fmt.Println("")
			fmt.Println("Press 'Enter' to close the program to update")
			var input string
			fmt.Scanln(&input)
			os.Exit(0)
		}
	} else {
		utilities.Log.Debug("Wowtools is up to date")
	}
}
