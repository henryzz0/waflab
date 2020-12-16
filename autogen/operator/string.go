package operator

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"

	"github.com/waflab/waflab/autogen/utils"
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

func reversePm(argument string, not bool) (string, error) {
	phrases := strings.Split(argument, " ")
	phrase := phrases[utils.RandomGenerator.Intn(len(phrases))] // pick phrase from pm's parameters randomly

	// convert snort style binary data, if there is any
	var builder strings.Builder
	/*
		Since snort content style use | to note the entrance and exit of embedded binary data, we can split phrase
		by separator | and the non-binary and binary will appear in alternating pattern
		Ex:
		"A|41|A" -> ["A", "41", "A"]
		"|41|A" -> ["", "41", "A"]
	*/
	isBinary := false
	for _, part := range strings.Split(phrase, "|") {
		if isBinary {
			decoded, err := hex.DecodeString(part)
			if err != nil {
				return "", err
			}
			builder.Write(decoded)
		} else {
			builder.WriteString(part)
		}
		isBinary = !isBinary
	}
	return builder.String(), nil
}

func reverseStreq(argument string, not bool) (string, error) {
	return argument, nil
}

func reverseWithin(argument string, not bool) (string, error) {
	return reverseRx(fmt.Sprintf(".*%s.*", regexp.QuoteMeta(argument)), not)
}
