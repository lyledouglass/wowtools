package utilities

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

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
