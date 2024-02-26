package util

import "crypto/sha256"

func HashUri(uri string) string {
	hash := sha256.New()
	hash.Write([]byte(uri))
	return string(hash.Sum(nil))
}
