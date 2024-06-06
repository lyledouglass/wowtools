package internal

import (
	"fmt"
	"os"
	"os/exec"
	"wowtools/internal/utilities"

	"github.com/spf13/viper"
)

var (
	retailAddonFolder string
	destAddonFolder   string
)

func CopyAddonData(wowVersion string, addonName string) {

	wowDir := viper.GetString("wow_dir")

	if addonName != "" {
		retailAddonFolder = wowDir + "\\_retail_\\Interface\\AddOns\\" + addonName
		destAddonFolder = wowDir + "\\" + wowVersion + "\\Interface\\AddOns\\" + addonName
	} else {
		retailAddonFolder = wowDir + "\\_retail_\\Interface\\AddOns"
		destAddonFolder = wowDir + "\\" + wowVersion + "\\Interface\\AddOns"
	}

	utilities.Log.Debug(retailAddonFolder)
	utilities.Log.Debug(destAddonFolder)

	removeDir(destAddonFolder)

	roboCmd := fmt.Sprintf("robocopy '%s' '%s' /mir", retailAddonFolder, destAddonFolder)
	utilities.Log.Debug(roboCmd)
	utilities.Log.Debugf("Copying AddOns to %s dir", wowVersion)
	cmd := exec.Command("powershell", roboCmd)
	err := cmd.Start()
	if err != nil {
		utilities.Log.WithError(err).Error("CopyAddonData - Failed to start command")
	}
	err = cmd.Wait()
	if err != nil {
		if err.Error() == "exit status 1" {
			return
		} else {
			utilities.Log.WithError(err).Error("CopyAddonData - Failed to copy AddOns")
		}
	}
	utilities.Log.Info("AddOn data copied successfully")
}

func removeDir(folder string) {
	utilities.Log.Debugf("Removing %s\n", folder)
	err := os.RemoveAll(folder)
	if err != nil {
		utilities.Log.WithError(err).Error("removeDir - failed to remove files")
	}
}
