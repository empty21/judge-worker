package util

import "os"

func WriteToFile(data string, path string) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(file)

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}
