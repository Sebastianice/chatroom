package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(pw string, b ...byte) string {
	hash := md5.New()
	hash.Write([]byte(pw))
	return hex.EncodeToString(hash.Sum(b))
}
