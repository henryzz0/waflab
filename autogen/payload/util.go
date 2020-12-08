package payload

import (
	"math/rand"
	"strings"
)

const charSet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(42))
}

func randomString(length int) string {
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteByte(charSet[random.Intn(len(charSet))])
	}
	return builder.String()
}
