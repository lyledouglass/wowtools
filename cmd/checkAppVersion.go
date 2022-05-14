package cmd

import (
	"wowtools/utilities"
)

func compareAppVersioning() {
	currentVersion := utilities.GetCurrentAppVersion()
	latestVersion := utilities.GetPublishedAppVersion()
	if currentVersion > latestVersion {
		updatePrompt := utilities.AskForConfirmation("You are running on an older version of this application. Would you like to download the latest version?")
		if updatePrompt {

		}
	}
}
