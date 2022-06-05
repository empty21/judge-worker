package util

import (
	"archive/zip"
	"io"
	"judger/pkg/telegram"
	_ "judger/pkg/telegram"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ZipAndBackup(workDir string) error {
	targetFile := path.Join(path.Dir(workDir), strings.ReplaceAll(workDir, path.Dir(workDir), "")+".zip")
	defer os.RemoveAll(targetFile)
	err := zipFile(workDir, targetFile)
	if err != nil {
		return err
	}
	telegram.SendFile(targetFile)
	return nil
}

func zipFile(workDir, target string) error {
	// 1. Create a ZIP file and zip.Writer
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()
	writer := zip.NewWriter(f)
	defer writer.Close()
	// 2. Go through all the files of the source
	return filepath.Walk(workDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 3. Create a local file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// set compression
		header.Method = zip.Deflate

		// 4. Set relative path of a file as the header name
		header.Name, err = filepath.Rel(filepath.Dir(workDir), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}

		// 5. Create writer for the file header and save content of the file
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
