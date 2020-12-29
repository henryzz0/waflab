package utils

import (
	"strings"
)

var charset = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r",
	"s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
	"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "_"}

// RandomStringFromSet generate a string by concating freq numbers of strings randomly draw from set.
// Notice that string draw from set may be repetitive.
func RandomStringFromSet(freq int, set []string) string {
	var builder strings.Builder
	for i := 0; i < freq; i++ {
		builder.WriteString(PickRandomString(set))
	}
	return builder.String()
}

// RandomString generate a random string with given length using character from charSet.
func RandomString(length int) string {
	return RandomStringFromSet(length, charset)
}

// PickRandomString picks a random string from the given string slice
func PickRandomString(set []string) string {
	return set[randomGenerator.rand.Intn(len(set))]
}
