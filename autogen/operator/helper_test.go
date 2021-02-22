// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
package operator

import (
	"testing"

	"github.com/waflab/waflab/autogen/utils"
)

// testHelper takes two string value: argument and output, where
// argument is the argument of operator and output is the expected
// output for executing the operator.
// stringTestHelper also set not equals to false by default
func testHelper(t *testing.T, f operationReverser, argument, output string, not ...bool) {
	t.Helper()

	utils.SetRandomSeed(42)
	isNot := false
	if len(not) > 0 {
		isNot = true
	}
	res, err := f(argument, isNot)
	if err != nil {
		panic(err)
	}
	utils.Assert(t, res, output)
}
