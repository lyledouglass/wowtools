package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/viper"
)

func zipSource(source, target string) error {
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Method = zip.Deflate
		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}

		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(headerWriter, f)
		return err
	})
}

func WtfBackup() {
	wtfFolder := viper.GetString("wtf_dir")
	wtfBackupDir := viper.GetString("wtf_backup_dir")
	currentTime := time.Now()
	folderName := currentTime.Format("2006-01-02")

	fmt.Println("Beginning backup of WTF folder")
	if err := zipSource(wtfFolder, wtfBackupDir+folderName+".zip"); err != nil {
		log.Fatal(err)
	}
	// Not really a true progress bar at the moment - more of a visual for the user - need to reseach better implementation, but works for now, as the zip process is fairly quick for the WTF folder
	bar := progressbar.Default(100)
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(20 * time.Millisecond)
	}
	fmt.Println("Folder backup complete")
}
