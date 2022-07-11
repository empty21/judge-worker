package util

import (
	"errors"
	"os"
	"time"
)

func FileNotExisted(path string, ttl time.Duration) bool {
	stat, err := os.Stat(path)
	if err == nil {
		return time.Now().After(stat.ModTime().Add(ttl))

	}
	return errors.Is(err, os.ErrNotExist)
}
