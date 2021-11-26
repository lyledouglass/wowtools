package cmd

import (
	"fmt"
	"strings"
	"wowtools/utilities"
)

func UpdateElvUI() {
	currentVersion := utilities.GetCurrentVersion()
	latestVersion := utilities.GetLatestVersion()
	stringCurrentVersion := strings.Join(currentVersion, "")
	fmt.Println(stringCurrentVersion)
	fmt.Println(latestVersion)
	if latestVersion > stringCurrentVersion {
		fmt.Printf("A later version of ElvUI is available. Current version: %s; New version: %s\n", stringCurrentVersion, latestVersion)
	} else {
		fmt.Println("ElvUI is up to date")
	}
}

// Compare it with current released version at https://www.tukui.org/download.php?ui=elvui

// if version is newer, zip up old installation and unzip new one.
