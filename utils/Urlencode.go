package utils

import "net/url"

func UrlEncode(param string) string {
	encodeValue := url.QueryEscape(param)
	return encodeValue
}

func UrlDecode(param string) string {
	decodeValue, _ := url.QueryUnescape(param)
	return decodeValue
}
