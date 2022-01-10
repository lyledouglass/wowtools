package test

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/spf13/viper"
)

func TestCurseforgeOpener(t *testing.T) {
	curseforgeExe := viper.GetString("curseforge_exe")
	curseforgeArgs := viper.GetString("curseforge_args")

	fmt.Println("Opening Cureseforge")
	cmd := exec.Command("powershell", "Start-Process",
		fmt.Sprintf("-Filepath '%s'", curseforgeExe),
		fmt.Sprintf("-ArgumentList '%s'", curseforgeArgs),
	)
	fmt.Println(cmd.String())
}
