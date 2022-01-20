package utils

import "encoding/base64"

func Base64Encode(buf []byte) string {
	encodeValue := base64.StdEncoding.EncodeToString(buf)
	return encodeValue
}

func Base64Decode(enStr string) ([]byte, error) {
	encodeValue, err := base64.StdEncoding.DecodeString(enStr)
	if err != nil {
		return nil, err
	}
	return encodeValue, nil
}
