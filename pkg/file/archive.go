package file

import (
	"archive/zip"
	"io"
	"os"
)

func Compress(destination string, sources ...string) error {
	zipFile, err := os.Create(destination)
	sourceFiles := make([]*os.File, 0)
	if err != nil {
		return err
	}
	zipWriter := zip.NewWriter(zipFile)
	defer func() {
		c(zipFile)
		c(zipWriter)
		for _, file := range sourceFiles {
			c(file)
		}
	}()
	for _, source := range sources {
		file, err := os.Open(source)
		if err != nil {
			return err
		}
		sourceFiles = append(sourceFiles, file)

	}
	for _, file := range sourceFiles {
		info, err := file.Stat()
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = file.Name()
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		if err != nil {
			return err
		}
	}

	return nil
}
