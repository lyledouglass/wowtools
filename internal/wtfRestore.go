package internal

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"wowtools/pkg/utilities"
)

func WtfRestore(file string) {
	fmt.Println("WARNING: This is a destructive action and will DELETE your current WTF directory. Are you sure " +
		"you want to proceed?")
	updatePrompt := utilities.AskForConfirmation("")
	if updatePrompt {
		fmt.Println("Continuing")
		fmt.Println("Removing current WTF folder")
		retailFolder := viper.GetString("retail_dir")
		wtfFolder := viper.GetString("wtf_dir")
		// Delete files recursively in order to delete the
		err := os.RemoveAll(wtfFolder)
		if err != nil {
			log.Fatal()
		}
		fmt.Printf("Restoring %s", file)
		restoreFile := viper.GetString("backup_dir") + "WTF\\" + file
		utilities.Unzip(restoreFile, retailFolder)
	} else {
		fmt.Println("Exiting")
		os.Exit(0)
	}

}
