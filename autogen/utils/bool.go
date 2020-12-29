package utils

import "fmt"

// RandomBiasedBool return a boolean value with probability%
// chance of being true and (1-probability)% of being false
// Notice that probability should in [0, 1]
func RandomBiasedBool(probability float64) bool {
	if probability > 1 || probability < 0 {
		panic(fmt.Sprintf("Invalid probability value %f: It should fall between 0 and 1.", probability))
	}
	return randomGenerator.rand.Float64() < probability
}
