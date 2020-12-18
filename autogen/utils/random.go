package utils

import "math/rand"

type random struct {
	rand *rand.Rand
}

var randomGenerator random

func init() {
	randomGenerator = random{
		rand: rand.New(rand.NewSource(42)),
	}
}

// SetRandomSeed set seed for the Random-related func in package utils
func SetRandomSeed(num int64) {
	randomGenerator.rand.Seed(num)
}
