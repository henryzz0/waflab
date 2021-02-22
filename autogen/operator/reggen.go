// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
package operator

import (
	"math"
	"regexp/syntax"
	"strings"

	"github.com/waflab/waflab/autogen/utils"
)

const (
	randomStringLimit    = 5
	repeatedstringLimit  = 10
	negatedProb          = 0.1
	printableChars       = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
	printableCharsLower  = 33  // Lowest byte value of printable chars
	printableCharsUppper = 126 // higher byte value of printable chars
)

func randomString(length int) string {
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteByte(printableChars[utils.RandomIntWithRange(0, len(printableChars))])
	}
	return builder.String()
}

func randomStringWithRange(ranges []rune) string {
	sum := 0
	for i := 0; i < len(ranges); i += 2 {
		sum += int((ranges[i+1] - ranges[i] + 1))
	}
	randomIndex := utils.RandomIntWithRange(0, sum)
	for i := 0; i < len(ranges); i += 2 {
		diff := int(ranges[i+1]-ranges[i]) + 1
		if randomIndex < diff {
			return string(ranges[i] + int32(randomIndex))
		}
		randomIndex -= diff
	}
	panic("randomIndex out of range")
}

func repeatSubexpression(re *syntax.Regexp, times int, not bool) (string, bool) {
	var builder strings.Builder
	isNegated := false
	for i := 0; i < times; i++ {
		for _, r := range re.Sub {
			// execute sub-expression
			res, isChildNegated := generate(r, not)
			builder.WriteString(res)
			isNegated = isNegated || isChildNegated
		}
	}
	return builder.String(), isNegated
}

func generate(re *syntax.Regexp, not bool) (string, bool) {
	isNegated := utils.RandomBiasedBool(negatedProb) && not // only for non-meta operation

	switch re.Op {
	case syntax.OpNoMatch, syntax.OpEmptyMatch:
		// Return a random string for negated operation
		// Return an empty string otherwise
		if isNegated {
			return randomString(randomStringLimit), isNegated
		}
		return "", isNegated
	case syntax.OpLiteral:
		// Return a string with one Rune randomly skipped for negated operation
		// Return the original literal otherwise
		var builder strings.Builder
		if isNegated {
			skippedRuneIndex := utils.RandomIntWithRange(0, len(re.Rune))
			for index, r := range re.Rune {
				if index != skippedRuneIndex {
					builder.WriteRune(r)
				}
			}
		} else {
			for _, r := range re.Rune {
				builder.WriteRune(r)
			}
		}
		return builder.String(), isNegated
	case syntax.OpCharClass:
		validRunes := make([]rune, 0)
		for i := 0; i < len(re.Rune); i += 2 {
			if re.Rune[i] > re.Rune[i+1] { // sanity check
				continue
			}
			if printableCharsLower <= re.Rune[i] && re.Rune[i+1] <= printableCharsUppper {
				validRunes = append(validRunes, re.Rune[i], re.Rune[i+1])
			} else {
				var lb, hb int32 // rune
				lb = int32(math.Max(float64(re.Rune[i]), float64(printableCharsLower)))
				hb = int32(math.Min(float64(re.Rune[i+1]), float64(printableCharsUppper)))
				if lb > hb { // when range [re.Runep[i], re.Rune[i+1]] does not overlap w/ [printableL, printableU]
					continue
				}
				validRunes = append(validRunes, lb, hb)
			}
		}

		if len(validRunes) == 0 { // TODO: better way of handling this
			return "", false
		}

		if isNegated {
			// invert(negate) the range first
			invertedRunes := []int32{}
			if re.Rune[0] > printableCharsLower {
				invertedRunes = append(invertedRunes, printableCharsLower, int32(re.Rune[0]))
			}
			for i := 1; i < len(re.Rune)-1; i += 2 {
				invertedRunes = append(invertedRunes, re.Rune[i]+1, re.Rune[i+1]-1)
			}
			if re.Rune[len(re.Rune)-1] < printableCharsUppper {
				invertedRunes = append(invertedRunes, int32(re.Rune[len(re.Rune)-1]), printableCharsUppper)
			}
			return randomStringWithRange(invertedRunes), isNegated
		}
		return randomStringWithRange(validRunes), isNegated // pick a random rune from range
	case syntax.OpAnyChar, syntax.OpAnyCharNotNL:
		// you cannot negate an AnyChar operator
		return string([]byte{printableChars[utils.RandomIntWithRange(0, len(printableChars))]}), false
	case syntax.OpBeginLine, syntax.OpEndLine, syntax.OpBeginText, syntax.OpEndText, syntax.OpWordBoundary, syntax.OpNoWordBoundary:
		return "", false
	case syntax.OpCapture:
		return generate(re.Sub0[0], not)
	case syntax.OpStar: // Repeat zero or more times
		limit := utils.RandomIntWithRange(0, repeatedstringLimit+1)
		return repeatSubexpression(re, limit, not)
	case syntax.OpPlus: // Repeat one or more times
		limit := utils.RandomIntWithRange(0, repeatedstringLimit) + 1
		isNegated = false
		return repeatSubexpression(re, limit, not)
	case syntax.OpQuest: // Zero or one instances
		limit := utils.RandomIntWithRange(0, 2) // repeat zero or one time
		return repeatSubexpression(re, limit, not)
	case syntax.OpRepeat: // Repeat re.Min to min(re.Max, repeatedstringLimit) times
		max := 0
		if re.Max < 0 { // for re.Max == -1
			max = repeatedstringLimit
		} else {
			max = int(math.Min(float64(re.Max), float64(repeatedstringLimit)))
		}
		count := utils.RandomIntWithRange(int(math.Min(float64(re.Min), float64(max))), max+1)
		return repeatSubexpression(re, count, not)
	case syntax.OpConcat:
		return repeatSubexpression(re, 1, not)
	case syntax.OpAlternate:
		randomIndex := utils.RandomIntWithRange(0, len(re.Sub))
		return generate(re.Sub[randomIndex], not)
	default:
		panic("Unhandled Operation")
	}
}

// Generate a negated string from something
func GenerateStringFromRegex(expression string, not bool) (res string, err error) {
	re, err := syntax.Parse(expression, syntax.Perl)
	if err != nil {
		return "", err
	}
	if not {
		isNegated := false
		for !isNegated {
			res, isNegated = generate(re, not)
		}
	} else {
		res, _ = generate(re, not)
	}

	return res, nil
}
