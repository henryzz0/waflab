package transformer_test

import (
	"strconv"
	"testing"

	"github.com/hsluoyz/modsecurity-go/seclang/parser"
	"github.com/waflab/waflab/autogen/transformer"
	"github.com/waflab/waflab/autogen/utils"
)

func TestReverseBase64Decode(t *testing.T) {
	transformationTestHelper(t, parser.TkTransBase64Decode, "nihao", "bmloYW8=")
}

func TestReverseCompressWhiteSpace(t *testing.T) {
	transformationTestHelper(t, parser.TkTransCompressWhitespace, "cat /etc/passwd", "cat\f\r\r\n\r /etc/passwd")
	transformationTestHelper(t, parser.TkTransCompressWhitespace, " ", "\f\r\r\n\r ")
	transformationTestHelper(t, parser.TkTransCompressWhitespace, "web application firewall", "web\f\r\r\n\r application\v\f\n\n firewall")
}

func TestReverseHexDecode(t *testing.T) {
	transformationTestHelper(t, parser.TkTransHexDecode, "nihao", "6e6968616f")
}

func TestReverseLength(t *testing.T) {
	transformationTestHelper(t, parser.TkTransLength, "10", "ASG5Ohg1l0")
	transformationTestHelper(t, parser.TkTransLength, "15", "ASG5Ohg1l0CMEef")
}

func TestReverseNormalizePath(t *testing.T) {
	transformationTestHelper(t, parser.TkTransNormalizePath, "/dev/stdout", "/foo/../dev/foo/../stdout")
	transformationTestHelper(t, parser.TkTransNormalizePath, "/", "/foo/../")
	transformationTestHelper(t, parser.TkTransNormalizePath, "", "")
}

func TestReverseNormalizePathWin(t *testing.T) {
	transformationTestHelper(t, parser.TkTransNormalizePathWin, "/dev/stdout", `\foo\..\dev\foo\..\stdout`)
	transformationTestHelper(t, parser.TkTransNormalizePathWin, "/", `\foo\..\`)
	transformationTestHelper(t, parser.TkTransNormalizePathWin, "", "")
}

func TestReverseLowercase(t *testing.T) {
	transformationTestHelper(t, parser.TkTransLowercase, "nihao", "NIhAO")
	transformationTestHelper(t, parser.TkTransLowercase, "Web Application Firewall", "WEb APpLIcaTIOn FireWALL")
}

func TestReverseRemoveComments(t *testing.T) {
	transformationTestHelper(t, parser.TkTransRemoveComments, "function hello() {\n alert(\"hello\")\n}",
		"function hello()/*ASG5Ohg1l0*/ {\n aler/*ASG5Ohg1l0*/t(\"hello/*ASG5Ohg1l0*/\")\n}#QzOrW_BTgE")
	transformationTestHelper(t, parser.TkTransRemoveComments, "", "#CMEefBrPrV")
}

func TestReverseRemoveCommentsChar(t *testing.T) {
	transformationTestHelper(t, parser.TkTransRemoveCommentsChar, "strconv.Quote()", "st#rc--onv.Quote()")
	transformationTestHelper(t, parser.TkTransRemoveCommentsChar, "", "")
}

func TestReverseReplaceComments(t *testing.T) {
	transformationTestHelper(t, parser.TkTransReplaceComments, "Transcript show: 100 factorial.",
		"Transcript/*SG5Ohg1l0C*/show:/*EefBrPrV9Q*/100/*azJoL6uaxF*/factorial.")
	transformationTestHelper(t, parser.TkTransReplaceComments, " ", "/*SG5Ohg1l0C*/")
}

func TestReverseRemoveNulls(t *testing.T) {
	transformationTestHelper(t, parser.TkTransRemoveNulls, "nihao", "ni\x00ha\x00o")

}

func TestReverseReplaceNulls(t *testing.T) {
	transformationTestHelper(t, parser.TkTransReplaceNulls, "ni hao", "ni\x00hao")

}

func TestReverseTrim(t *testing.T) {
	transformationTestHelper(t, parser.TkTransTrim, "nihao", "\v\n\f\r\n\r\n\t\f\vnihao\f\n\r\f\r\f\n\t\r\r")
	transformationTestHelper(t, parser.TkTransTrim, "", "\v\n\f\r\n\r\n\t\f\v\f\n\r\f\r\f\n\t\r\r")
}

func TestReverseTrimLeft(t *testing.T) {
	transformationTestHelper(t, parser.TkTransTrimLeft, "nihao", "\f\n\r\f\r\f\n\t\r\rnihao")

}

func TestReverseTrimRight(t *testing.T) {
	transformationTestHelper(t, parser.TkTransTrimRight, "nihao", "nihao\f\n\r\f\r\f\n\t\r\r")

}

func TestReverseUrlDecode(t *testing.T) {
	transformationTestHelper(t, parser.TkTransUrlDecode, "ni hao\n", "ni+hao%0A")
	transformationTestHelper(t, parser.TkTransUrlDecode, "你好", "%E4%BD%A0%E5%A5%BD")
}

func transformationTestHelper(t *testing.T, opTk int, argument, expected string) {
	t.Helper()

	utils.SetRandomSeed(42)
	trans := []*parser.Trans{
		{
			Tk: opTk,
		},
	}
	res := transformer.ReverseTransform(trans, argument)
	if res != expected {
		t.Errorf("%s: expect %s but get %s",
			parser.TransformationNameMap[opTk],
			expected, strconv.Quote(res))
	}
}
