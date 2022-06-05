package util

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(path string, uri string) error {
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the body to file
	_, err = io.Copy(file, resp.Body)
	return err
}
