package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(bytes []byte) string {
	h := md5.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}
