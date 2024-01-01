package internal

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	util "wowtools/pkg/utilities"

	"github.com/google/go-github/v56/github"
)

// Get the latest version of the application from the github api
func getPublishedAppVersion(githubClient *github.Client) (string, error) {
	opt := &github.ListOptions{Page: 1, PerPage: 1}
	releases, rsp, err := githubClient.Repositories.ListReleases(context.Background(), "lyledouglass", "wowtools", opt)
	if err != nil {
		util.Log.Fatalf("Error getting releases: %v", err)
	}
	util.Log.Debug(*releases[0].TagName)
	util.Log.Tracef("\n%+v\n", rsp)
	tageName := *releases[0].TagName
	versionRegex := regexp.MustCompile(`^v\d+\.\d+\.\d+$`)
	if !versionRegex.MatchString(tageName) {
		return "", fmt.Errorf("invalid tag name: %s", tageName)
	}
	version := strings.TrimPrefix(tageName, "v")
	return version, nil
}

// Download the latest version of the application from the github api
func downloadApp(githubClient *github.Client, version string, folderPath string) error {
	release, _, err := githubClient.Repositories.GetReleaseByTag(context.Background(), "lyledouglass", "wowtools", "v"+version)
	if err != nil {
		return fmt.Errorf("error getting release: %s", err)
	}
	for _, asset := range release.Assets {
		util.Log.Debug(*asset.Name)
		httpClient := &http.Client{Timeout: 60 * time.Second}
		resp, _, err := githubClient.Repositories.DownloadReleaseAsset(context.Background(), "lyledouglass", "wowtools", *release.Assets[0].ID, httpClient)
		if err != nil {
			return fmt.Errorf("error downloading asset: %s", err)
		}
		defer resp.Close()
		tmpfile, err := os.CreateTemp("", *asset.Name)
		if err != nil {
			return fmt.Errorf("error creating temp file: %s", err)
		}
		//defer tmpfile.Close()
		reader := bufio.NewReader(resp)
		writeFile, err := io.Copy(tmpfile, reader)
		if err != nil {
			return fmt.Errorf("error copying file: %s", err)
		}
		tmpFileName := tmpfile.Name()
		tmpfile.Close()
		util.Log.Debugf("Wrote %d bytes to %s", writeFile, tmpFileName)

		// Move the downloaded file to the specified location
		destPath := filepath.Join(folderPath, filepath.Base(*asset.Name))
		util.Log.Infof("Moving %s to %s", tmpFileName, destPath)
		err = os.Rename(tmpFileName, destPath)
		if err != nil {
			fmt.Print(err)
			return fmt.Errorf("error moving file: %s", err)
		}
		// update permissions on the new filepath
		err = os.Chmod(destPath, 0777)
		if err != nil {
			return fmt.Errorf("error updating permissions: %s", err)
		}
	}
	return nil
}

func UpdateApp(appVersion string) {
	// Wowtools is a public repo, so we won't need to authenticate to access and can pass nil to the client
	client := github.NewClient(nil)
	latestVersion, err := getPublishedAppVersion(client)
	if err != nil {
		util.Log.Fatalf("Error getting latest version: %s", err)
		//return
	}

	// Assuming that the version returned from getPublishedAppVersion is valid, we
	// can compare the two versions and see if we need to download a new version
	if latestVersion > appVersion {
		util.Log.Info("New version available, downloading...")
		user, err := user.Current()
		if err != nil {
			util.Log.Fatalf("Error getting current user: %s", err)
		}
		homeDir := user.HomeDir
		downloadApp(client, latestVersion, homeDir+"\\Downloads")
		util.Log.Info("New version downloaded to Downloads folder. Please overwrite the existing executable with the new one!")
	} else {
		util.Log.Info("You are running the latest version")
	}
}
