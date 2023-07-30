package utilities

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// CurrentAppVersion Returns the local version of the application. Not the best
// way to handle the app version but works for both OS
func CurrentAppVersion() string {
	appVersion := "4.2.0"
	return appVersion
}

type githubApiData struct {
	AppVersion string `json:"tag_name"`
}

func GetPublishedAppVersion(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		Log.WithError(err).Errorf("Unable to GET %s", url)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Log.WithError(err).Error("Error reading response body")
	}
	var data githubApiData
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		Log.WithError(jsonErr).Error("Error unmarshalling json")
	}
	return strings.Trim(data.AppVersion, "{ v }")
}
