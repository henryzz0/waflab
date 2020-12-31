package operator_test

import (
	"log"
	"strconv"
	"testing"

	"github.com/hsluoyz/modsecurity-go/seclang/parser"
	"github.com/pkg/errors"
	"github.com/waflab/waflab/autogen/operator"
	"github.com/waflab/waflab/autogen/utils"
)

func TestReverseEq(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := testNumericalOperator(parser.TkOpEq, func(input int, output int) bool {
			return input == output
		})
		if err != nil {
			t.Error(err)
		}
	}
}

func TestReverseGe(t *testing.T) {
	// @Ge Performs numerical comparison and returns true if the input value is greater than or equal
	// to the provided parameter. Notice that output of our reversing func actually correspond to @Ge's input.
	for i := 0; i < 10; i++ {
		err := testNumericalOperator(parser.TkOpGe, func(input int, output int) bool {
			return output >= input
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestReverseGt(t *testing.T) {
	// @Gt Performs a numerical comparison between a variable and parameter and returns true if the
	// input value is greater than the operator parameter.
	for i := 0; i < 10; i++ {
		err := testNumericalOperator(parser.TkOpGt, func(input int, output int) bool {
			return output > input
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestReverseLe(t *testing.T) {
	// @Le Performs numerical comparison and returns true if the input value is less than or equal to
	// the operator parameter
	for i := 0; i < 10; i++ {
		err := testNumericalOperator(parser.TkOpLe, func(input int, output int) bool {
			return output <= input
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestReverseLt(t *testing.T) {
	// @Lt Performs numerical comparison and returns true if the input value is less than
	// the operator parameter
	for i := 0; i < 10; i++ {
		err := testNumericalOperator(parser.TkOpLt, func(input int, output int) bool {
			return output < input
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

// testNumericalOperator determine if the numerical operator reverser works as we expect.
// The compare is a function with two parameters: output and input. The output is the output
// from specific numerical operator reverser and the input is the randomly generated number
// passed to the reverser. We expect compare return true if the numerical relationship between
// input and output is correct. Otherwise compare should return false
func testNumericalOperator(opTk int, compare func(intput int, output int) bool) error {
	randomNum := utils.RandomNonNegativeInt()
	op := &parser.Operator{
		Tk:       opTk,
		Argument: strconv.Itoa(randomNum),
	}
	str, err := operator.ReverseOperator(op)
	if err != nil {
		return errors.Errorf("%s: Error when running ReverseOperator: %v", parser.OperatorNameMap[opTk], err)
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		return errors.Errorf("%s: Error when converting output string: %v", parser.OperatorNameMap[opTk], err)
	}
	if !compare(randomNum, num) {
		return errors.Errorf("%s: Invalid output %d with input %d", parser.OperatorNameMap[opTk], num, randomNum)
	}
	return nil
}
