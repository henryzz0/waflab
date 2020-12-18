package utils

// RandomIntWithRange generate an integer by range [min, max)
func RandomIntWithRange(min int, max int) int {
	return randomGenerator.rand.Intn(max-min) + min
}

// RandomNonNegativeInt randomly generate an non-negative int
func RandomNonNegativeInt() int {
	return randomGenerator.rand.Int()
}

// RandomFloat32 is a wrapper of rand.Float32()
func RandomFloat32() float32 {
	return randomGenerator.rand.Float32()
}
