package operator

import (
	"math"
	"strconv"

	"github.com/waflab/waflab/autogen/utils"
)

//TODO: confirm the range, should be [min, max)?
func randomIntWithRange(min int, max int) int {
	return utils.RandomGenerator.Intn(max-min) + min
}

func reverseEq(argument string, not bool) (string, error) {
	// check if valid string
	num, err := strconv.Atoi(argument)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(num), nil
}

func reverseGe(argument string, not bool) (string, error) {
	num, err := strconv.Atoi(argument)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(randomIntWithRange(num, math.MaxInt64)), nil
}

//TODO: potential overflow
func reverseGt(argument string, not bool) (string, error) {
	num, err := strconv.Atoi(argument)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(randomIntWithRange(num+1, math.MaxInt64)), nil
}

func reverseLe(argument string, not bool) (string, error) {
	num, err := strconv.Atoi(argument)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(randomIntWithRange(math.MinInt64, num+1)), nil
}

func reverseLt(argument string, not bool) (string, error) {
	num, err := strconv.Atoi(argument)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(randomIntWithRange(math.MinInt64, num)), nil
}
