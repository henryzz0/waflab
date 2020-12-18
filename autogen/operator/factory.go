package operator

import "github.com/hsluoyz/modsecurity-go/seclang/parser"

type operationReverser func(argument string, not bool) (string, error)

var reverserFactory = map[int]operationReverser{
	// string matching operator
	parser.TkOpRx:         reverseRx,
	parser.TkOpBeginsWith: reverseBeginsWith,
	parser.TkOpContains:   reverseContains,
	parser.TkOpEndsWith:   reverseEndsWith,
	parser.TkOpPm:         reversePm,
	parser.TkOpPmFromFile: reversePmFromFile,
	parser.TkOpStrEq:      reverseStrEq,
	parser.TkOpWithin:     reverseWithin,
	// numerical operator
	parser.TkOpEq: reverseEq,
	parser.TkOpGe: reverseGe,
	parser.TkOpGt: reverseGt,
	parser.TkOpLe: reverseLe,
	parser.TkOpLt: reverseLt,
	// validation operator
	parser.TkOpValidateByteRange:    reverseValidateByteRange,
	parser.TkOpValidateUtf8Encoding: reverseValidateUtf8Encoding,
	// miscellaneous operator
	parser.TkOpIpMatch:         reverseIPMatch,
	parser.TkOpIpMatchFromFile: reverseIPMatchFromFile,
}

func ReverseOperator(operator *parser.Operator) (string, error) {
	if operator.Not {
		panic("negative match not supported yet!")
	}
	if f, ok := reverserFactory[operator.Tk]; ok {
		return f(operator.Argument, operator.Not)
	}
	panic("not supported operator!")
}
