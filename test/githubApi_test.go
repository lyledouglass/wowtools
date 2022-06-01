package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestGitHubApi(t *testing.T) {
	type githubApiData struct {
		AppVersion string `json:"tag_name"`
		Assets     []struct {
			Name               string `json:"name"`
			BrowserDownloadURL string `json:"browser_download_url"`
		}
	}
	uri := "https://api.github.com/repos/lyledouglass/wowtools/releases/latest"
	resp, err := http.Get(uri)
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
	fmt.Println("Version: " + data.AppVersion)
	for _, asset := range data.Assets {
		if asset.Name == "wowtools.exe" {
			fmt.Println("Download URI: " + asset.BrowserDownloadURL)
		}
	}
}
