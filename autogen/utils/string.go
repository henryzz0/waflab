package utils

import (
	"math/rand"
	"strings"
)

const charSet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

var RandomGenerator *rand.Rand

func init() {
	RandomGenerator = rand.New(rand.NewSource(42))
}

func RandomString(length int) string {
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteByte(charSet[RandomGenerator.Intn(len(charSet))])
	}
	return builder.String()
}
