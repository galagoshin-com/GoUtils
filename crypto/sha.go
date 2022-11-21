package crypto

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func Sha1(bytes []byte) string {
	h := sha1.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}

func Sha256(bytes []byte) string {
	h := sha256.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}

func Sha512(bytes []byte) string {
	h := sha512.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}
