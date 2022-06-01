package utilities

import (
	"encoding/json"
	"fmt"
	"github.com/gonutz/w32/v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

func GetCurrentVersion() []string {
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

func GetLatestVersion() string {
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

type githubApiData struct {
	AppVersion string `json:"tag_name"`
}

func GetPublishedAppVersion(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data githubApiData
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return strings.Trim(data.AppVersion, "{ v }")
}

func GetCurrentAppVersion() string {
	executablePath, err := os.Executable()
	if err != nil {
		log.Fatalln("Failed to get path of executable")
	}
	size := w32.GetFileVersionInfoSize(executablePath)
	if size <= 0 {
		log.Fatalln("GetFileVersionInfoSize failed")
	}
	info := make([]byte, size)
	getFileInfo := w32.GetFileVersionInfo(executablePath, info)
	if !getFileInfo {
		log.Fatalln("GetFileVersionInfo failed")
	}
	fixed, getFileInfo := w32.VerQueryValueRoot(info)
	if !getFileInfo {
		log.Fatalln("VerQueryValueRoot failed")
	}
	fileVersion := fixed.FileVersion()
	versionString := fmt.Sprintf("%d.%d.%d\n",
		fileVersion&0xFFFF000000000000>>48,
		fileVersion&0x0000FFFF00000000>>32,
		fileVersion&0x00000000FFFF0000>>16)

	return versionString
}
