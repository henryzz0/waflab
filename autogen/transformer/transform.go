package transformer

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"unicode"

	"github.com/waflab/waflab/autogen/utils"
)

// probability list for reversing transformation
const (
	randomStringLength        = 10
	reverseLowerCaseProb      = 0.5
	reverseCompressProb       = 0.5
	reverseCommentProb        = 0.10
	reverseCommentCharProb    = 0.10
	reverseNullProb           = 0.10
	reverseReplaceCommentProb = 0.5
	reverseReplaceNullProb    = 0.5
)

var whiteSpaceCharacters = []string{"\f", "\t", "\n", "\r", "\v"}

// randomStringsInsertion randomly insert string from reverse to str.
// At each rune of str, randomstringsInsertion will randomly pick a string from reserve and
// insert it between the rune with given probability
func randomStringsInsertion(str string, reserve []string, probability float64) string {
	var builder strings.Builder
	for _, r := range str {
		builder.WriteRune(r)
		if utils.RandomBiasedBool(probability) {
			builder.WriteString(reserve[utils.RandomIntWithRange(0, len(reserve))])
		}
	}
	return builder.String()
}

func reverseBase64Decode(variable string) string {
	return base64.StdEncoding.EncodeToString([]byte(variable))
}

// reverseCompressWhiteSpace assume that the only kinds of whitespace character
// variable contains is space and there is not any consecutive space.
func reverseCompressWhiteSpace(variable string) string {
	var builder strings.Builder

	for _, r := range variable {
		if unicode.IsSpace(r) {
			for p := true; p; p = utils.RandomBiasedBool(reverseCompressProb) {
				builder.WriteString(utils.PickRandomString(whiteSpaceCharacters))
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
	redundantPath := []string{"", ".", "foo/.."}
	parts := strings.Split(variable, "/")
	res := []string{}

	for index, part := range parts {
		res = append(res, part)
		if index < len(parts)-1 { // add redundant path in between the path
			res = append(res, utils.PickRandomString(redundantPath))
		}
	}

	return strings.Join(res, "/")
}

func reverseNormalizePathWin(variable string) string {
	return reverseNormalizePath(strings.ReplaceAll(variable, "/", "\\"))
}

func reverseLowercase(variable string) string {
	var builder strings.Builder
	for _, char := range variable {
		if utils.RandomBiasedBool(reverseLowerCaseProb) {
			builder.WriteRune(unicode.ToUpper(char))
		} else {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}

func reverseRemoveComments(variable string) string {
	res := randomStringsInsertion(variable,
		[]string{fmt.Sprintf("/*%s*/", utils.RandomString(randomStringLength))},
		reverseCommentProb)
	res += fmt.Sprintf("#%s", utils.RandomString(randomStringLength))
	return res
}

func reverseRemoveCommentsChar(variable string) string {
	return randomStringsInsertion(variable, []string{"/**/", "--", "#"}, reverseCommentCharProb)
}

func reverseReplaceComments(variable string) string {
	var builder strings.Builder
	for _, r := range variable {
		if unicode.IsSpace(r) && utils.RandomBiasedBool(reverseReplaceCommentProb) {
			builder.WriteString(fmt.Sprintf("/*%s*/", utils.RandomString(randomStringLength)))
		} else {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func reverseRemoveNulls(variable string) string {
	return randomStringsInsertion(variable, []string{"\x00"}, reverseNullProb)
}

func reverseReplaceNulls(variable string) string {
	var builder strings.Builder
	for _, r := range variable {
		builder.WriteRune(r)
		if unicode.IsSpace(r) && utils.RandomBiasedBool(reverseReplaceNullProb) {
			builder.WriteString("\000")
		} else {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func reverseTrim(variable string) string {
	return reverseTrimLeft(reverseTrimRight(variable))
}

func reverseTrimLeft(variable string) string {
	return fmt.Sprintf("%s%s", utils.RandomStringFromSet(randomStringLength, whiteSpaceCharacters), variable)
}

func reverseTrimRight(variable string) string {
	return fmt.Sprintf("%s%s", variable, utils.RandomStringFromSet(randomStringLength, whiteSpaceCharacters))
}

func reverseUrlDecode(variable string) string {
	return url.QueryEscape(variable)
}
