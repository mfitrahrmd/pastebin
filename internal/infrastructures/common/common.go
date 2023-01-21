package common

import (
	"math/rand"
	"os"
)

const (
	DEFAULT_CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func GenerateRandomString(length uint) string {
	var charset string
	envCharset := os.Getenv("CHARSET")

	if envCharset == "" {
		charset = DEFAULT_CHARSET
	} else {
		charset = envCharset
	}

	str := make([]byte, length)

	for i := 0; i < len(str); i++ {
		str[i] = charset[rand.Intn(len(charset))]
	}

	return string(str)
}
