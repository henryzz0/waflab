package utils

import (
	"strings"
)

const charSet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

// RandomStringFromSet generate a string by concating freq numbers of strings randomly draw from set.
// Notice that string draw from set may be repetitive.
func RandomStringFromSet(freq int, set []string) string {
	var builder strings.Builder
	for i := 0; i < freq; i++ {
		builder.WriteString(set[randomGenerator.rand.Intn(len(set))])
	}
	return builder.String()
}

// RandomString generate a random string with given length using character from charSet.
func RandomString(length int) string {
	return RandomStringFromSet(length, strings.Split(charSet, ""))
}
