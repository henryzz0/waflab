// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package utils

import (
	"fmt"
	"testing"
)

// RandomBiasedBool return a boolean value with probability%
// chance of being true and (1-probability)% of being false
// Notice that probability should in [0, 1]
func RandomBiasedBool(probability float64) bool {
	if probability > 1 || probability < 0 {
		panic(fmt.Sprintf("Invalid probability value %f: It should fall between 0 and 1.", probability))
	}
	return randomGenerator.rand.Float64() < probability
}

// Assert throw an error if the input value is inconsistent with
// expect value. Assert is reserved for testing.
func Assert(t *testing.T, result, expect string) {
	t.Helper()

	if result != expect {
		t.Errorf("Expect %s but get %s", expect, result)
	}
}
