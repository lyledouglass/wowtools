package utilities

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

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
