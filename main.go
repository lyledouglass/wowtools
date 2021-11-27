package main

import (
	L "wowtools/cmd"
	// "wowtools/utilities"
)

func main() {
	L.InitConfig()
	L.WtfBackup()
	L.UpdateElvUI()
}
