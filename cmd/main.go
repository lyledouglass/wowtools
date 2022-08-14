//go:generate goversioninfo

package main

import (
	"flag"
	"github.com/spf13/viper"
	"sync"
	"wowtools/internal"
	"wowtools/pkg/utilities"
)

func main() {
	var (
		copyPtr    bool
		backupOnly bool
		noUpdates  bool
	)

	// Flags
	flag.BoolVar(&copyPtr, "copy-ptr", false, "only performs copy of PTR folders from Retail")
	flag.BoolVar(&backupOnly, "backup-only", false, "perfomrs only backup of wtf folder")
	flag.BoolVar(&noUpdates, "no-updtaes", false, "skips checking updates for wowtools")
	flag.Parse()

	internal.InitConfig()
	// Check for updates to the application
	if !noUpdates {
		internal.UpdateWowtools()
	}

	// WaitGroup for creating missing folders.
	var wg sync.WaitGroup
	wg.Add(3)
	go utilities.VerifyFolders(viper.GetString("backup_dir"), &wg)
	go utilities.VerifyFolders(viper.GetString("backup_dir")+"ElvUI", &wg)
	go utilities.VerifyFolders(viper.GetString("backup_dir")+"WTF", &wg)
	wg.Wait()

	if copyPtr {
		internal.CopyPtrData()
	}

	if backupOnly {
		internal.WtfBackup()
	}

	if !copyPtr && !backupOnly {
		internal.WtfBackup()
		if viper.GetString("elvui_dir") != "" {
			internal.UpdateElvUI()
		}
		internal.OpenCurseforge()
	}
}
