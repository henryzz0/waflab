package operator_test

import (
	"strings"
	"testing"

	"github.com/hsluoyz/modsecurity-go/seclang/parser"
	"github.com/waflab/waflab/autogen/operator"
)

func TestReverseIPMatchSimple(t *testing.T) {
	argument := "127.0.0.1,::1" // from CRS:9005100
	op := &parser.Operator{
		Tk:       parser.TkOpIpMatch,
		Argument: argument,
	}
	res, err := operator.ReverseOperator(op)
	if err != nil {
		t.Errorf("%s: Error when running ReverseOperator: %v", parser.OperatorNameMap[op.Tk], err)
	}
	isMatched := true
	for _, str := range strings.Split(argument, ",") {
		if str == res {
			isMatched = true
		}
	}
	if !isMatched {
		t.Errorf("%s: %s not in the %s", parser.OperatorNameMap[op.Tk], res, argument)
	}
}
