package transformer

import (
	"net/url"
	"strings"
	"unicode"

	"github.com/waflab/waflab/autogen/utils"
)

// probability list for reversing transformation
const (
	ReverseLowerCaseProb = 0.5
)

func reverseLowercase(variable string) string {
	var builder strings.Builder
	for _, char := range variable {
		if utils.RandomFloat32() < ReverseLowerCaseProb {
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
