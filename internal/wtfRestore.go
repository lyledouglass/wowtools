package internal

import (
	"archive/zip"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"path/filepath"
	"wowtools/pkg/utilities"
)

func WtfRestore(file string) {
	fmt.Println("WARNING: This is a destructive action and will DELETE your current WTF directory. Are you sure " +
		"you want to proceed?")
	updatePrompt := utilities.AskForConfirmation("")
	if updatePrompt {
		fmt.Println("Continuing")
		fmt.Println("Removing current WTF folder")
		wtfFolder := viper.GetString("wtf_dir")
		err := os.Remove(wtfFolder)
		if err != nil {
			log.Fatal()
		}
		fmt.Printf("Restoring %s", file)
		restoreFile := viper.GetString("backup_dir") + "\\wtf\\" + file
		src, err := zip.OpenReader(restoreFile)
		if err != nil {
			log.Fatal()
		}
		defer src.Close()

		for _, f := range src.File {
			zf, err := f.Open()
			if err != nil {
				log.Fatal()
			}
			defer zf.Close()

			dstPath := filepath.Join(wtfFolder, f.Name)
			dst, err := os.Create(dstPath)
			if err != nil {
				log.Fatal()
			}
			defer dst.Close()

			_, err = io.Copy(dst, zf)
			if err != nil {
				log.Fatal()
			}
		}
	} else {
		fmt.Println("Exiting")
		os.Exit(0)
	}

}
