package utilities

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetReleaseAsset(uri string, assetName string) string {
	var downloadUri string
	type githubApiData struct {
		AppVersion string `json:"tag_name"`
		Assets     []struct {
			Name               string `json:"name"`
			BrowserDownloadURL string `json:"browser_download_url"`
		}
	}
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data githubApiData
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	for _, asset := range data.Assets {
		if asset.Name == assetName {
			downloadUri = asset.BrowserDownloadURL
		}
	}
	return downloadUri
}
