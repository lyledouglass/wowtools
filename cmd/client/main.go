package main

import (
	"flag"
	"sync"
	"wowtools/internal"
	"wowtools/pkg/utilities"

	"github.com/spf13/viper"
)

func main() {

	utilities.LoadConfig(".")

	var (
		copyPtr   bool
		backup    bool
		noUpdates bool
		restore   bool
		wtfzip    string
	)
	// Configure Logging via Logrus using utilities.SetupLogger for a global config
	utilities.SetupLogger(viper.GetString("log_level"))

	// Flags
	flag.BoolVar(&copyPtr, "copy-ptr", false, "only performs copy of PTR folders from Retail")
	flag.BoolVar(&backup, "backup", false, "perfomrs only backup of wtf folder")
	flag.BoolVar(&noUpdates, "no-updates", false, "skips checking updates for wowtools")
	flag.BoolVar(&restore, "restore", false, "restores a WTF backup")
	flag.StringVar(&wtfzip, "wtfzip", "", "File name of WTF Zip")
	flag.Parse()

	// Check for updates to the application
	if !noUpdates {
		utilities.Log.Debug("Beginning UpdateWowtools function")
		internal.UpdateWowtools()
	}

	if copyPtr {
		utilities.Log.Debug("Beginning CopyPtrData function")
		internal.CopyPtrData()
	}

	if backup {
		// WaitGroup for creating missing folders.
		utilities.Log.Debug("Creating WaitGroup if folders are missing")
		var wg sync.WaitGroup
		wg.Add(2)
		go utilities.VerifyFolders(viper.GetString("backup_dir"), &wg)
		// go utilities.VerifyFolders(viper.GetString("backup_dir")+"ElvUI", &wg)
		go utilities.VerifyFolders(viper.GetString("backup_dir")+"WTF", &wg)
		wg.Wait()

		utilities.Log.Debug("Beginning WtfBackup function")
		internal.WtfBackup()
	}

	if restore && wtfzip != "" {
		utilities.Log.Debug("Beginning WtfRestore function")
		internal.WtfRestore(wtfzip)
	}

	// Deprecated functionality with new WowUp CF program!

	/*
		if !copyPtr && !backupOnly {
			internal.WtfBackup()
			if viper.GetString("elvui_dir") != "" {
				internal.UpdateElvUI()
			}
			internal.OpenCurseforge()
		}
	*/
}
