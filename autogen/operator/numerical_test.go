// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package operator

import (
	"testing"
)

func TestReverseEq(t *testing.T) {
	testHelper(t, reverseEq, "100", "100")
	testHelper(t, reverseEq, "1", "1")
	testHelper(t, reverseEq, "0", "0")
	testHelper(t, reverseEq, "1000000000", "1000000000")
	testHelper(t, reverseEq, "-1", "-1")
}

func TestReverseGe(t *testing.T) {
	// @Ge Performs numerical comparison and returns true if the input value is greater than or equal
	// to the provided parameter. Notice that output of our reversing func actually correspond to @Ge's input.
	testHelper(t, reverseGe, "100", "105")
	testHelper(t, reverseGe, "1", "6")
	testHelper(t, reverseGe, "0", "5")
	testHelper(t, reverseGe, "1000000000", "1000000005")
	testHelper(t, reverseGe, "-1", "4")
}

func TestReverseGt(t *testing.T) {
	// @Gt Performs a numerical comparison between a variable and parameter and returns true if the
	// input value is greater than the operator parameter.
	testHelper(t, reverseGt, "100", "145")
	testHelper(t, reverseGt, "1", "46")
	testHelper(t, reverseGt, "0", "45")
	testHelper(t, reverseGt, "1000000000", "1000000045")
	testHelper(t, reverseGt, "-1", "44")
}

func TestReverseLe(t *testing.T) {
	// @Le Performs numerical comparison and returns true if the input value is less than or equal to
	// the operator parameter
	testHelper(t, reverseLe, "100", "97")
	testHelper(t, reverseLe, "1", "-2")
	testHelper(t, reverseLe, "0", "-3")
	testHelper(t, reverseLe, "1000000000", "999999997")
	testHelper(t, reverseLe, "-1", "-4")
}

func TestReverseLt(t *testing.T) {
	// @Lt Performs numerical comparison and returns true if the input value is less than
	// the operator parameter
	testHelper(t, reverseLt, "100", "5")
	testHelper(t, reverseLt, "1", "-94")
	testHelper(t, reverseLt, "0", "-95")
	testHelper(t, reverseLt, "1000000000", "999999905")
	testHelper(t, reverseLt, "-1", "-96")
}
