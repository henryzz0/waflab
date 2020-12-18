package operator

import (
	"strconv"

	"github.com/waflab/waflab/autogen/utils"
)

const (
	randomReverseBound = 100
)

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
	return strconv.Itoa(utils.RandomIntWithRange(num, num+randomReverseBound)), nil
}

//TODO: potential overflow
func reverseGt(argument string, not bool) (string, error) {
	num, err := strconv.Atoi(argument)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(utils.RandomIntWithRange(num+1, num+randomReverseBound)), nil
}

func reverseLe(argument string, not bool) (string, error) {
	num, err := strconv.Atoi(argument)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(utils.RandomIntWithRange(num-randomReverseBound, num+1)), nil
}

func reverseLt(argument string, not bool) (string, error) {
	num, err := strconv.Atoi(argument)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(utils.RandomIntWithRange(num-randomReverseBound, num)), nil
}
