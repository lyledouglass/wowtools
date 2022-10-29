package internal

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
	utilities "wowtools/pkg/utilities"
)

func getCurrentElvuiVersion() []string {
	elvuiDir := viper.GetString("elvui_dir")
	fileOpen, err := ioutil.ReadFile(elvuiDir + "ElvUI_Mainline.toc")
	if err != nil {
		log.Fatal(err)
	}
	str := string(fileOpen)
	re := regexp.MustCompile(`[0-9]+\.[0-9]+`)
	v := re.FindStringSubmatch(str)

	return v
}

type apiData struct {
	Version string `json:"version"`
}

func getLatestElvuiVersion() string {
	url := "https://www.tukui.org/api.php?ui=elvui"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data apiData
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return (data.Version)
}

func UpdateElvUI() {
	currentVersion := getCurrentElvuiVersion()
	latestVersion := getLatestElvuiVersion()
	stringCurrentVersion := strings.Join(currentVersion, "")
	filename := "elvui-" + latestVersion + ".zip"
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	downloadUri := "https://www.tukui.org/downloads/" + filename
	zipFile := homeDir + "\\Downloads\\" + filename

	// if version is newer, zip up old installation and unzip new one.
	if latestVersion > stringCurrentVersion {
		fmt.Printf("A later version of ElvUI is available. Current version: %s; New version: %s\n", stringCurrentVersion, latestVersion)
		updatePrompt := utilities.AskForConfirmation("Do you want to install the lastest version of ElvUI?")
		if updatePrompt {
			removeOldElvuiZip()
			ZipElvUI()
			fmt.Printf("Downloading ElvUI %s\n", latestVersion)
			err := utilities.DownloadFiles(filename, downloadUri)
			if err != nil {
				log.Fatal(err)
			}
			// Defer removing the zip file that's downloaded
			defer os.Remove(zipFile)
			utilities.RemoveFolder(viper.GetString("elvui_dir"))
			utilities.RemoveFolder(viper.GetString("elvui_options_dir"))
			if err != nil {
				log.Fatal(err)
			}
			utilities.Unzip(zipFile, viper.GetString("addons_dir"))
		}
	} else {
		fmt.Println("ElvUI is up to date")
	}
}

func ZipElvUI() {
	elvuiFolder := viper.GetString("elvui_dir")
	backupFolder := viper.GetString("backup_dir") + "ElvUI\\"
	currentTime := time.Now()
	folderName := currentTime.Format("2006-01-02")

	fmt.Println("Beginning backup of ElvUI folder. This may take a moment")
	if err := utilities.ZipSource(elvuiFolder, backupFolder+folderName+".zip"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Folder backup complete")
}

func removeOldElvuiZip() {
	backupFolder := viper.GetString("backup_dir") + "ElvUI\\"
	fileCount := utilities.GetFileCount(backupFolder)
	fmt.Println(fileCount)
	if fileCount > 2 {
		oldestFile := utilities.GetOldestFolder(backupFolder)
		err := os.Chdir(backupFolder)
		if err != nil {
			log.Fatal(err)
		}
		removeFile := os.Remove(oldestFile)
		if removeFile != nil {
			log.Fatal()
		}
	}
}
