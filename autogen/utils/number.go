// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
package utils

// RandomIntWithRange generate an integer by range [min, max)
func RandomIntWithRange(min int, max int) int {
	return randomGenerator.rand.Intn(max-min) + min
}

// RandomNonNegativeInt randomly generate an non-negative int
func RandomNonNegativeInt() int {
	return randomGenerator.rand.Int()
}
