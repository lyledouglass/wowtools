package utilities

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/viper"
)

func OpenCurseforge() {
	curseforgeExe := viper.GetString("curseforge_exe")
	curseforgeArgs := viper.GetString("curseforge_args")
	updatePrompt := AskForConfirmation("Do you want to launch Curseforge to update addons?")
	if updatePrompt {
		fmt.Println("Opening Cureseforge")
		cmd := exec.Command("powershell", "Start-Process",
			fmt.Sprintf("-Filepath '%s'", curseforgeExe),
			fmt.Sprintf("-ArgumentList '%s'", curseforgeArgs))
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
	}
}
