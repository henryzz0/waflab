package operator

import (
	"fmt"
	"regexp"
)

func reverseRx(argument string, not bool) (string, error) {
	generator, err := newGenerator(argument)
	if err != nil {
		return "", err
	}
	return generator.Generate(10), nil
}

func reverseBeginsWith(argument string, not bool) (string, error) {
	return reverseRx(fmt.Sprintf("^%s.*", regexp.QuoteMeta(argument)), not)
}

func reverseContains(argument string, not bool) (string, error) {
	return reverseRx(regexp.QuoteMeta(argument), not)
}

func reverseEndWith(argument string, not bool) (string, error) {
	return reverseRx(fmt.Sprintf(".*%s$", regexp.QuoteMeta(argument)), not)
}