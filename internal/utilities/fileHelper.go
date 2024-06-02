package utilities

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

func ZipSource(source, target string) error {
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Method = zip.Deflate
		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}

		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(headerWriter, f)
		return err
	})
}

func DownloadFiles(filename string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		Log.WithError(err).Errorf("Failed to get URL %s", url)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		Log.WithFields(logrus.Fields{
			"StatusCode": resp.StatusCode,
		}).Error("StatusCode not 200")
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		Log.WithError(err).Error("Failed to get Home dir")
	}

	fileOutput, err := os.Create(homeDir + "\\Downloads\\" + filename)
	if err != nil {
		Log.WithError(err).Errorf("Failed to create %s", homeDir+"\\Downloads\\"+filename)
	}
	defer fileOutput.Close()

	_, err = io.Copy(fileOutput, resp.Body)
	return err
}

func RemoveFolder(folderPath string) {
	err := os.RemoveAll(folderPath)
	if err != nil {
		Log.WithError(err).Errorf("Failed to remove %s", folderPath)
	}
}

func Unzip(src string, dest string) ([]string, error) {
	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		Log.WithError(err).Error("Failed to open zip reader")
	}
	defer r.Close()

	for _, f := range r.File {

		fpath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

func VerifyFolders(filepath string, wg *sync.WaitGroup) {
	defer wg.Done()
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		Log.Infof("%s did not exist. Creating it", filepath)
		os.Mkdir(filepath, os.ModePerm)
	}
}

func GetOldestFolder(filepath string) string {
	var oldestFile fs.FileInfo
	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		Log.WithError(err).Error("Error getting oldest folder")
	}
	oldestTime := time.Now()
	for _, file := range files {
		if file.Mode().IsRegular() && file.ModTime().Before(oldestTime) {
			oldestFile = file
			oldestTime = file.ModTime()
		}
	}
	if oldestFile == nil {
		err = os.ErrNotExist
	}
	return oldestFile.Name()
}

func GetFileCount(directory string) int {
	files, _ := ioutil.ReadDir(directory)
	count := len(files)
	return count
}
