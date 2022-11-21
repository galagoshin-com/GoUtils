package encoding

import (
	"encoding/base64"
)

type Base64 string

func EncodeBase64(bytes []byte) Base64 {
	return Base64(base64.StdEncoding.EncodeToString(bytes))
}

func DecodeBase64(encoded Base64) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(encoded))
}
