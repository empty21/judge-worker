package util

import (
	"crypto/tls"
	"errors"
	"io"
	"judger/pkg/config"
	"net/http"
	"os"
)

func DownloadFile(path string, uri string) error {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	req.Header.Set("X-API-Key", config.Config.FileAPIKey)

	resp, err := client.Do(req)
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
	if resp.StatusCode == 200 {
		_, err = io.Copy(file, resp.Body)
	} else {
		return errors.New("error when pulling file " + uri)
	}
	return err
}
