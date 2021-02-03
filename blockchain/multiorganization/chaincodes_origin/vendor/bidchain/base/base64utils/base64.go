package base64utils

import "encoding/base64"

func Base64Encode(buf []byte) string {
	s := base64.StdEncoding.EncodeToString(buf)
	return s
}

func Base64Decode(str string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	return data, err
}