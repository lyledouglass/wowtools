package main

import (
	"wowtools/cmd"
	"wowtools/utilities"

	"github.com/spf13/viper"
)

func main() {
	cmd.InitConfig()
	utilities.VerifyFolders(viper.GetString("backup_dir"))
	go utilities.VerifyFolders(viper.GetString("backup_dir") + "ElvUI")
	go utilities.VerifyFolders(viper.GetString("backup_dir") + "WTF")
	cmd.WtfBackup()
	cmd.UpdateElvUI()
	utilities.OpenCurseforge()
}
