package utils

//TODO: confirm the range, should be [min, max)?
/*
	Generate an integer from range [min, max)
*/
func randomIntWithRange(min int, max int) int {
	return utils.RandomGenerator.Intn(max-min) + min
}
