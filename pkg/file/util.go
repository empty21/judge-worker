package file

import (
	"os"
)

func Read(path string) (string, error) {
	content, err := os.ReadFile(path)
	return string(content), err
}

func Write(path, data string) error {
	err := os.WriteFile(path, []byte(data), 0644)

	return err
}

func Copy(source, destination string) error {
	input, err := os.ReadFile(source)
	if err != nil {
		return err
	}
	err = os.WriteFile(destination, input, 0644)

	return err
}

func Remove(path string) error {
	return os.RemoveAll(path)
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
