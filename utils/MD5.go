package utils

import (
	"crypto/md5"
	"fmt"
)

func MD5(str string) string {
	has := md5.Sum([]byte(str))
	md5Str := fmt.Sprintf("%x", has)
	return md5Str
}
