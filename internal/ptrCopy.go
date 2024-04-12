package internal

import (
	"fmt"
	"os"
	"os/exec"
	"wowtools/pkg/utilities"

	"github.com/spf13/viper"
)

func CopyPtrData() {
	ptrFolder := viper.GetString("wow_dir") + "\\_ptr_\\"
	retailFolder := viper.GetString("wow_dir") + "\\_retail_\\"

	var folders [2]string
	folders[0] = "Interface"
	folders[1] = "WTF"

	removePtrSubDirs(folders, ptrFolder)

	for _, element := range folders {
		roboCmd := fmt.Sprintf("robocopy \"%s\" \"%s\" /s", retailFolder+element, ptrFolder+element)
		utilities.Log.Debug(roboCmd)
		utilities.Log.Debugf("Copying %s to PTR dir", element)
		cmd := exec.Command("powershell", roboCmd)

		err := cmd.Start()
		if err != nil {
			utilities.Log.WithError(err).Error("CopyPtrData - Failed to start command")
		}
	}
	utilities.Log.Info("PTR data copied successfully")
}

func removePtrSubDirs(folderArray [2]string, dstFolder string) {
	for _, element := range folderArray {
		utilities.Log.Debug("Removing %s\n", element)
		err := os.RemoveAll(dstFolder + element)
		if err != nil {
			utilities.Log.WithError(err).Error("removePtrSubDirs - failed to remove files")
		}
	}
}
