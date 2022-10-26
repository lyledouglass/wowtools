package internal

import (
	"github.com/spf13/viper"
	"wowtools/pkg/utilities"
)

func NewCharacter(serverName string, characterName string, accountName string) {
	// Copy the template zip to dst
	fileDst := viper.GetString("wtf_dir") + "\\Account\\" + accountName + serverName
	fileSrc := viper.GetString("char_template_path")
	utilities.CopyFolder(fileSrc, fileDst)
	utilities.Unzip()
}

func updateCharacterInFiles() {
	// Logic to find and replace the template name across all files

}
