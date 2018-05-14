package lib

import (
	"crypto/md5"
	"encoding/hex"
	"path/filepath"
)

// GetMD5Hash return the md5 hash for text
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	hash := hex.EncodeToString(hasher.Sum(nil))

	return hash
}

// HashName use md5 hash of string to build a pseudo random path
func HashName(path string) string {
	hash := GetMD5Hash(path)

	return hash[0:2] + "/" + hash[len(hash)-2:]
}

// HashDir use HashName to build a pseudo random path
func HashDir(prefix, path string) string {
	hash := HashName(path)
	basedir := prefix + "/" + hash
	basename := filepath.Base(path)

	return basedir + "/" + basename
}
