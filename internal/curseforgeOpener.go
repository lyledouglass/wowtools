package internal

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os/exec"
	"wowtools/pkg/utilities"
)

func OpenCurseforge() {
	curseforgeExe := viper.GetString("curseforge_exe")
	updatePrompt := utilities.AskForConfirmation("Do you want to launch Curseforge to update addons?")
	if updatePrompt {
		fmt.Println("Opening Cureseforge")
		cmd := exec.Command("powershell", "Start-Process",
			fmt.Sprintf("-Filepath '%s'", curseforgeExe))
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
	}
}
