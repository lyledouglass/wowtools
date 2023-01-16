package utilities

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// CurrentAppVersion Returns the local version of the application. Not the best
// way to handle the app version but works for both OS
func CurrentAppVersion() string {
	appVersion := "4.0.0"
	return appVersion
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
