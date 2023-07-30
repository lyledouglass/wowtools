package utilities

import (
	"encoding/json"
	"io"
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
		Log.WithError(err).Errorf("GetReleaseAsset - failed to GET %s", uri)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			Log.WithError(err).Error("Error closing Body")
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		Log.WithError(err).Error("Failed to Read resp.Body")
	}
	var data githubApiData
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		Log.WithError(jsonErr).Error("Failed to unmarshal json")
	}
	for _, asset := range data.Assets {
		if asset.Name == assetName {
			downloadUri = asset.BrowserDownloadURL
		}
	}
	Log.Debug("Download URI - " + downloadUri)
	return downloadUri
}
