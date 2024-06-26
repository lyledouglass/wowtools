package internal

import (
	"fmt"
	"log"
	"os/exec"
	"wowtools/internal/utilities"

	"github.com/spf13/viper"
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
