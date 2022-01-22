package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GenMD5(v string) string {
	if len(v) == 0 {
		return ""
	}
	hash := md5.Sum([]byte(v))
	ret := hex.EncodeToString(hash[:])
	return ret
}
