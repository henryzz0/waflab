package utils

import (
	"strings"
)

const charSet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

// RandomString generate a random string with given length using character from charSet
func RandomString(length int) string {
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteByte(charSet[randomGenerator.rand.Intn(len(charSet))])
	}
	return builder.String()
}
