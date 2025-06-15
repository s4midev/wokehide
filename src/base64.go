package main

import "encoding/base64"

func EncodeBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func DecodeBase64(s string) string {
	b, _ := base64.StdEncoding.DecodeString(s)
	return string(b)
}

func IsBase64(s string) bool {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return false
	}
	return base64.StdEncoding.EncodeToString(b) == s
}
