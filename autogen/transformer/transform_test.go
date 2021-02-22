// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
package transformer

import (
	"testing"

	"github.com/waflab/waflab/autogen/utils"
)

func TestReverseBase64Decode(t *testing.T) {
	transformationTestHelper(t, reverseBase64Decode, "nihao", "bmloYW8=")
}

func TestReverseCompressWhiteSpace(t *testing.T) {
	transformationTestHelper(t, reverseCompressWhiteSpace, "cat /etc/passwd", "cat\f\r\r\n\r /etc/passwd")
	transformationTestHelper(t, reverseCompressWhiteSpace, " ", "\f\r\r\n\r ")
	transformationTestHelper(t, reverseCompressWhiteSpace, "web application firewall", "web\f\r\r\n\r application\v\f\n\n firewall")
}

func TestReverseHexDecode(t *testing.T) {
	transformationTestHelper(t, reverseHexDecode, "nihao", "6e6968616f")
}

func TestReverseLength(t *testing.T) {
	transformationTestHelper(t, reverseLength, "10", "ASG5Ohg1l0")
	transformationTestHelper(t, reverseLength, "15", "ASG5Ohg1l0CMEef")
}

func TestReverseNormalizePath(t *testing.T) {
	transformationTestHelper(t, reverseNormalizePath, "/dev/stdout", "/foo/../dev/foo/../stdout")
	transformationTestHelper(t, reverseNormalizePath, "/", "/foo/../")
	transformationTestHelper(t, reverseNormalizePath, "", "")
}

func TestReverseNormalizePathWin(t *testing.T) {
	transformationTestHelper(t, reverseNormalizePathWin, "/dev/stdout", `\foo\..\dev\foo\..\stdout`)
	transformationTestHelper(t, reverseNormalizePathWin, "/", `\foo\..\`)
	transformationTestHelper(t, reverseNormalizePathWin, "", "")
}

func TestReverseLowercase(t *testing.T) {
	transformationTestHelper(t, reverseLowercase, "nihao", "NIhAO")
	transformationTestHelper(t, reverseLowercase, "Web Application Firewall", "WEb APpLIcaTIOn FireWALL")
}

func TestReverseRemoveComments(t *testing.T) {
	transformationTestHelper(t, reverseRemoveComments, "function hello() {\n alert(\"hello\")\n}",
		"function hello()/*ASG5Ohg1l0*/ {\n aler/*ASG5Ohg1l0*/t(\"hello/*ASG5Ohg1l0*/\")\n}#QzOrW_BTgE")
	transformationTestHelper(t, reverseRemoveComments, "", "#CMEefBrPrV")
}

func TestReverseRemoveCommentsChar(t *testing.T) {
	transformationTestHelper(t, reverseRemoveCommentsChar, "strconv.Quote()", "st#rc--onv.Quote()")
	transformationTestHelper(t, reverseRemoveCommentsChar, "", "")
}

func TestReverseReplaceComments(t *testing.T) {
	transformationTestHelper(t, reverseReplaceComments, "Transcript show: 100 factorial.",
		"Transcript/*SG5Ohg1l0C*/show:/*EefBrPrV9Q*/100/*azJoL6uaxF*/factorial.")
	transformationTestHelper(t, reverseReplaceComments, " ", "/*SG5Ohg1l0C*/")
}

func TestReverseRemoveNulls(t *testing.T) {
	transformationTestHelper(t, reverseRemoveNulls, "nihao", "ni\x00ha\x00o")

}

func TestReverseReplaceNulls(t *testing.T) {
	transformationTestHelper(t, reverseReplaceNulls, "ni hao", "ni\x00hao")

}

func TestReverseTrim(t *testing.T) {
	transformationTestHelper(t, reverseTrim, "nihao", "\v\n\f\r\n\r\n\t\f\vnihao\f\n\r\f\r\f\n\t\r\r")
	transformationTestHelper(t, reverseTrim, "", "\v\n\f\r\n\r\n\t\f\v\f\n\r\f\r\f\n\t\r\r")
}

func TestReverseTrimLeft(t *testing.T) {
	transformationTestHelper(t, reverseTrimLeft, "nihao", "\f\n\r\f\r\f\n\t\r\rnihao")

}

func TestReverseTrimRight(t *testing.T) {
	transformationTestHelper(t, reverseTrimRight, "nihao", "nihao\f\n\r\f\r\f\n\t\r\r")

}

func TestReverseUrlDecode(t *testing.T) {
	transformationTestHelper(t, reverseUrlDecode, "ni hao\n", "ni+hao%0A")
	transformationTestHelper(t, reverseUrlDecode, "你好", "%E4%BD%A0%E5%A5%BD")
}

func transformationTestHelper(t *testing.T, f transformReverser, argument, expected string) {
	t.Helper()

	utils.SetRandomSeed(42)
	res := f(argument)
	utils.Assert(t, res, expected)
}
