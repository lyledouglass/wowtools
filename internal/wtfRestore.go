package internal

import (
	"fmt"
	"os"
	"wowtools/pkg/utilities"

	"github.com/spf13/viper"
)

func WtfRestore(file string) {
	fmt.Println("WARNING: This is a destructive action and will DELETE your current WTF directory. Are you sure " +
		"you want to proceed?")
	updatePrompt := utilities.AskForConfirmation("")
	if updatePrompt {
		utilities.Log.Info("Continuing destructive action")
		utilities.Log.Debug("Removing current WTF folder")
		retailFolder := viper.GetString("retail_dir")
		wtfFolder := viper.GetString("wtf_dir")
		// Delete files recursively in order to delete the
		err := os.RemoveAll(wtfFolder)
		if err != nil {
			utilities.Log.WithError(err).Errorf("Failed to remove %s", wtfFolder)
		}
		utilities.Log.Infof("Restoring %s", file)
		restoreFile := viper.GetString("backup_dir") + "WTF\\" + file
		utilities.Unzip(restoreFile, retailFolder)
	} else {
		utilities.Log.Debug("Exiting")
		os.Exit(0)
	}

}
