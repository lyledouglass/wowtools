package test

import (
	"encoding/json"
	"fmt"
	"github.com/gonutz/w32/v2"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestGetAppVersion(t *testing.T) {
	t.Skip()
	t.Log("Test")
	executablePath := "D:\\Development\\wowtools\\wowtools.exe"
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
	versionString, err := fmt.Printf("%d.%d.%d.%d\n",
		fileVersion&0xFFFF000000000000>>48,
		fileVersion&0x0000FFFF00000000>>32,
		fileVersion&0x00000000FFFF0000>>16,
		fileVersion&0x000000000000FFFF>>0)
	if err != nil {
		log.Fatalln("Failed to assign version to string")
	}
	t.Log(versionString)
}

type githubApiData struct {
	AppVersion string `json:"tag_name"`
}

func TestGetLatestVersion(t *testing.T) {
	t.Skip()
	url := "https://api.github.com/repos/lyledouglass/wowtools/releases/latest"
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
	t.Log(data)
	trimed := strings.Trim(data.AppVersion, "{ v }")
	t.Log(trimed)
}
