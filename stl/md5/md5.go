package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// 参考 https://wangbjun.site/2020/coding/golang/md5.html

func ToMD5(s string) string  {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}