package transformer

import (
	"encoding/base64"
	"encoding/hex"
	"net/url"
	"strconv"
	"strings"
	"unicode"

	"github.com/waflab/waflab/autogen/utils"
)

// probability list for reversing transformation
const (
	reverseLowerCaseProb = 0.5
	reverseCompressProb  = 0.5
)

var whiteSpaceCharacters = []string{"\f", "\t", "\n", "\r", "\v"}

func reverseBase64Decode(variable string) string {
	return base64.StdEncoding.EncodeToString([]byte(variable))
}

// reverseCompressWhiteSpace assume that the only kinds of whitespace character
// variable contains is space and there is not any consecutive space.
func reverseCompressWhiteSpace(variable string) string {
	var builder strings.Builder

	for _, r := range variable {
		if unicode.IsSpace(r) {
			for p := utils.RandomFloat32(); p < reverseCompressProb; p = utils.RandomFloat32() {
				builder.WriteString(whiteSpaceCharacters[utils.RandomIntWithRange(0, len(whiteSpaceCharacters))])
			}
		}
		builder.WriteRune(r)
	}

	return builder.String()
}

func reverseHexDecode(variable string) string {
	return hex.EncodeToString([]byte(variable))
}

func reverseLength(variable string) string {
	length, err := strconv.Atoi(variable)
	if err != nil {
		panic("Length must be an integer")
	}
	return utils.RandomString(length)
}

// reverseNormalizePath assume that the variable string is a normalized path,
// otherwise the function may return illegal path
func reverseNormalizePath(variable string) string {
	redundantPath := []string{"//", "/./", "foo/../"}
	parts := strings.Split(variable, "/")
	var builder strings.Builder

	for index, part := range parts {
		builder.WriteString(part)
		if index < len(parts)-1 { // add redundant path in between the path
			builder.WriteString(redundantPath[utils.RandomIntWithRange(0, len(redundantPath))])
		}
	}

	return builder.String()
}

func reverseNormalizePathWin(variable string) string {
	return reverseNormalizePath(strings.ReplaceAll(variable, "/", "\\"))
}

func reverseLowercase(variable string) string {
	var builder strings.Builder
	for _, char := range variable {
		if utils.RandomFloat32() < reverseLowerCaseProb {
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
