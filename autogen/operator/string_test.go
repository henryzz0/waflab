package operator_test

import (
	"testing"

	"github.com/hsluoyz/modsecurity-go/seclang/parser"
	"github.com/waflab/waflab/autogen/operator"
	"github.com/waflab/waflab/autogen/utils"
)

func TestRx(t *testing.T) {
	stringTestHelper(t, parser.TkOpRx, false, ".*", `K\q\^Ol~;`)
	stringTestHelper(t, parser.TkOpRx, false, "(abc|ab)+", "ababababcabcabcabcab")

	// from golang/src/regexp/find_test.go
	stringTestHelper(t, parser.TkOpRx, false, "", "")
	stringTestHelper(t, parser.TkOpRx, false, "^abcdefg", "abcdefg")
	stringTestHelper(t, parser.TkOpRx, false, "a+", "aaaaaaaa")
	stringTestHelper(t, parser.TkOpRx, false, "abcd..", `abcdK\`)
	stringTestHelper(t, parser.TkOpRx, false, "a", "a")
	stringTestHelper(t, parser.TkOpRx, false, ".", `"`)
	stringTestHelper(t, parser.TkOpRx, false, "^", "")
	stringTestHelper(t, parser.TkOpRx, false, "$", "")
	stringTestHelper(t, parser.TkOpRx, false, "^abcd$", "abcd")
	stringTestHelper(t, parser.TkOpRx, false, "^bcd", "bcd")
	stringTestHelper(t, parser.TkOpRx, false, "a+", "aaaaaaaa")
	stringTestHelper(t, parser.TkOpRx, false, "a*", "aaaaaaaaa")
	stringTestHelper(t, parser.TkOpRx, false, "[a-z]+", "ktuztevn")
	stringTestHelper(t, parser.TkOpRx, false, `[a\-\]z]+`, "a]-zz-]]")
	stringTestHelper(t, parser.TkOpRx, false, `[^\n]+`, `K\q\^Ol~`)
	stringTestHelper(t, parser.TkOpRx, false, "()", "")
	stringTestHelper(t, parser.TkOpRx, false, "(a)", "a")
	stringTestHelper(t, parser.TkOpRx, false, "(.)(.)", "K*")
	stringTestHelper(t, parser.TkOpRx, false, "(.*)", "b")
	stringTestHelper(t, parser.TkOpRx, false, "(..)(..)", "b*X*")
	stringTestHelper(t, parser.TkOpRx, false, "(([^xyz]*)(d))", "d")
	stringTestHelper(t, parser.TkOpRx, false, "((a|b|c)*(d))", "abaacd")
	stringTestHelper(t, parser.TkOpRx, false, "(((a|b|c)*)(d))", "d")
}

func TestBeginsWith(t *testing.T) {
	// from CRS v3.3
	stringTestHelper(t, parser.TkOpBeginsWith, false, "/index.php", "/index.php")
	stringTestHelper(t, parser.TkOpBeginsWith, false, "/admin", "/admin")
}

func TestContains(t *testing.T) {
	// from CRS v3.3
	stringTestHelper(t, parser.TkOpContains, false, "/admin/config/", `b/admin/config/\^O`)
	stringTestHelper(t, parser.TkOpContains, false, "/wp-admin/", `b/wp-admin/\^O`)
}

func TestEndsWith(t *testing.T) {
	// from CRS v3.3
	stringTestHelper(t, parser.TkOpEndsWith, false, "/index.php", "b/index.php")
}

func TestPm(t *testing.T) {
	stringTestHelper(t, parser.TkOpPm, false, "nihao a|42|b", "aBb")
}

func TestPmFromFile(t *testing.T) {
	stringTestHelper(t, parser.TkOpPmFromFile, false, "data/crawlers-user-agents.data", "grapeFX")
}

func TestStrEq(t *testing.T) {
	stringTestHelper(t, parser.TkOpStrEq, false, "abc", "abc")
}

func TestWithin(t *testing.T) {
	stringTestHelper(t, parser.TkOpWithin, false, "abc", `babc\^O`)
}

// stringTestHelper takes two string value: argument and output, where
// argument is the argument of operator and output is the expected
// output for executing the operator.
func stringTestHelper(t *testing.T, opTk int, not bool, argument, output string) {
	t.Helper()

	utils.SetRandomSeed(42)
	op := &parser.Operator{
		Tk:       opTk,
		Argument: argument,
		Not:      not,
	}
	res, err := operator.ReverseOperator(op)
	if err != nil {
		t.Errorf("%s: encounter err when calling ReverseOperator %v",
			parser.OperatorNameMap[opTk],
			err)
	}
	if res != output {
		t.Errorf("%s: expect %s but get %s",
			parser.OperatorNameMap[opTk],
			output, res)
	}
}
