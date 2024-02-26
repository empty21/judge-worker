package file

import (
	"judger/pkg/config"
	"judger/pkg/log"
	"judger/pkg/util"
	"os"
	"path"
	"sort"
	"time"
)

var accessedTime = make(map[string]int64)

func Cache(f func(path, uri string) error) func(path, uri string) error {
	return func(p, u string) error {
		hashedUri := util.HashUri(u)
		cachePath := path.Join(config.CacheFolder, hashedUri)

		// If the file exists in the cache, copy it to the destination
		if Exists(cachePath) {
			err := Copy(cachePath, p)
			if err != nil {
				return err
			}
			accessedTime[hashedUri] = time.Now().Unix()
		}

		// If the file does not exist or the copy failed, download the file
		err := f(p, u)
		if err == nil {
			// If the download was successful, cache the file
			defer func() {
				err = Copy(p, cachePath)
				if err != nil {
					accessedTime[hashedUri] = time.Now().Unix()
				}
			}()
		}

		return err
	}
}

func CleanUpCache() {
	log.Info("Start cleaning up cache")
	files, err := os.ReadDir(config.CacheFolder)
	if err != nil {
		log.Error("Clean cache folder error: %v", err)
		return
	}
	if len(files) > config.FileCachedCount {
		// Sort files by last access time
		sort.Slice(files, func(i, j int) bool {
			iAccessedTime := accessedTime[files[i].Name()]
			jAccessedTime := accessedTime[files[j].Name()]
			return iAccessedTime < jAccessedTime
		})

		// Remove the oldest files
		for i := 0; i < len(files)-config.FileCachedCount; i++ {
			err = os.RemoveAll(path.Join(config.CacheFolder, files[i].Name()))
			delete(accessedTime, files[i].Name())
		}
	}
	log.Info("Finish cleaning up cache")
}

func init() {
	_ = os.MkdirAll(config.CacheFolder, os.ModeDir)
}
