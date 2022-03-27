package main

import (
	"sync"
	"wowtools/cmd"
	"wowtools/utilities"

	"github.com/spf13/viper"
)

func main() {
	cmd.InitConfig()

	// WaitGroup for creating missing folders.
	var wg sync.WaitGroup
	wg.Add(3)
	go utilities.VerifyFolders(viper.GetString("backup_dir"), &wg)
	go utilities.VerifyFolders(viper.GetString("backup_dir")+"ElvUI", &wg)
	go utilities.VerifyFolders(viper.GetString("backup_dir")+"WTF", &wg)
	wg.Wait()

	cmd.WtfBackup()
	cmd.UpdateElvUI()
	utilities.OpenCurseforge()
}
