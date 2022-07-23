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
	internal.InitConfig()
	// Check for updates to the application
	internal.UpdateWowtools()

	var copyPtr bool
	flag.BoolVar(&copyPtr, "copy-ptr", false, "only performs copy of PTR folders from Retail")
	flag.Parse()
	if copyPtr {
		internal.CopyPtrData()
	} else {
		// WaitGroup for creating missing folders.
		var wg sync.WaitGroup
		wg.Add(3)
		go utilities.VerifyFolders(viper.GetString("backup_dir"), &wg)
		go utilities.VerifyFolders(viper.GetString("backup_dir")+"ElvUI", &wg)
		go utilities.VerifyFolders(viper.GetString("backup_dir")+"WTF", &wg)
		wg.Wait()

		internal.WtfBackup()
		internal.UpdateElvUI()
		internal.OpenCurseforge()
	}
}
