package file

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"judger/pkg/config"
	"net/http"
	"os"
)

func GetFile(path, uri string) error {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	request, _ := http.NewRequest(http.MethodGet, uri, nil)

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	contentReader := io.LimitReader(response.Body, config.TestMaxFileSize)
	defer c(response.Body)

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer c(file)

	// Write the body to file
	if response.StatusCode == 200 {
		_, err = io.Copy(file, contentReader)
	} else {
		return errors.New(fmt.Sprintf("error when pulling file %s with status code %d", uri, response.StatusCode))
	}
	return err
}

func c(io io.Closer) {
	_ = io.Close()
}
