package transformer

import (
	"math/rand"
	"net/url"
	"strings"
	"unicode"
)

var random *rand.Rand

// probability list for reversing transformation
const (
	ReverseLowerCaseProb = 0.5
)

func init() {
	random = rand.New(rand.NewSource(42))
}

func reverseLowercase(variable string) string {
	var builder strings.Builder
	for _, char := range variable {
		if random.Float32() < ReverseLowerCaseProb {
			builder.WriteRune(unicode.ToUpper(char))
		} else {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}

func reverseUrlDecode(variable string) string {
	return url.QueryEscape(variable)
}
