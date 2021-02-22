// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package operator

import (
	"testing"
)

func TestRx(t *testing.T) {
	testHelper(t, reverseRx, ".*", `K\q\^Ol~;`)
	testHelper(t, reverseRx, "(abc|ab)+", "ababababcabcabcabcab")

	// from golang/src/regexp/find_test.go
	testHelper(t, reverseRx, "", "")
	testHelper(t, reverseRx, "^abcdefg", "abcdefg")
	testHelper(t, reverseRx, "a+", "aaaaaaaa")
	testHelper(t, reverseRx, "abcd..", `abcdK\`)
	testHelper(t, reverseRx, "a", "a")
	testHelper(t, reverseRx, ".", `"`)
	testHelper(t, reverseRx, "^", "")
	testHelper(t, reverseRx, "$", "")
	testHelper(t, reverseRx, "^abcd$", "abcd")
	testHelper(t, reverseRx, "^bcd", "bcd")
	testHelper(t, reverseRx, "a+", "aaaaaaaa")
	testHelper(t, reverseRx, "a*", "aaaaaaaaa")
	testHelper(t, reverseRx, "[a-z]+", "ktuztevn")
	testHelper(t, reverseRx, `[a\-\]z]+`, "a]-zz-]]")
	testHelper(t, reverseRx, `[^\n]+`, `K\q\^Ol~`)
	testHelper(t, reverseRx, "()", "")
	testHelper(t, reverseRx, "(a)", "a")
	testHelper(t, reverseRx, "(.)(.)", "K*")
	testHelper(t, reverseRx, "(.*)", "b")
	testHelper(t, reverseRx, "(..)(..)", "b*X*")
	testHelper(t, reverseRx, "(([^xyz]*)(d))", "d")
	testHelper(t, reverseRx, "((a|b|c)*(d))", "abaacd")
	testHelper(t, reverseRx, "(((a|b|c)*)(d))", "d")
}

func TestBeginsWith(t *testing.T) {
	// from CRS v3.3
	testHelper(t, reverseBeginsWith, "/index.php", "/index.php")
	testHelper(t, reverseBeginsWith, "/admin", "/admin")
}

func TestContains(t *testing.T) {
	// from CRS v3.3
	testHelper(t, reverseContains, "/admin/config/", `b/admin/config/\^O`)
	testHelper(t, reverseContains, "/wp-admin/", `b/wp-admin/\^O`)
}

func TestEndsWith(t *testing.T) {
	// from CRS v3.3
	testHelper(t, reverseEndsWith, "/index.php", "b/index.php")
}

func TestPm(t *testing.T) {
	testHelper(t, reversePm, "nihao a|42|b", "aBb")
}

func TestPmFromFile(t *testing.T) {
	testHelper(t, reversePmFromFile, "data/crawlers-user-agents.data", "grapeFX")
}

func TestStrEq(t *testing.T) {
	testHelper(t, reverseStrEq, "abc", "abc")
}

func TestWithin(t *testing.T) {
	testHelper(t, reverseWithin, "abc", `babc\^O`)
}
