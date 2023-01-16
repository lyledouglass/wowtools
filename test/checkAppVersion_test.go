package test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestGetAppVersion(t *testing.T) {
	t.Skip()
	t.Log("Test")
	currentVersion := "3.0.0"
	t.Log(currentVersion)
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
